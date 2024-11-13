package main

import (
	authhandler "bookingcinema/pkg/auth/handlers"
	authinfra "bookingcinema/pkg/auth/infrastructe/postgres"
	authusecase "bookingcinema/pkg/auth/usecases"
	authutils "bookingcinema/pkg/auth/utils"
	dbpg "bookingcinema/pkg/database"
	"context"

	moviehandler "bookingcinema/pkg/movie/handlers"
	"bookingcinema/pkg/movie/infrastructure/movieinfra"
	movieusecases "bookingcinema/pkg/movie/usecases"

	theaterhandler "bookingcinema/pkg/theater/handlers"
	theaterinfra "bookingcinema/pkg/theater/infrastructure/postgres"
	theaterusecase "bookingcinema/pkg/theater/usecases"

	bookinghandler "bookingcinema/pkg/booking/handlers"
	bookingkafka "bookingcinema/pkg/booking/infrastructure/kafka"
	bookinginfra "bookingcinema/pkg/booking/infrastructure/postgres"
	bookingUsecase "bookingcinema/pkg/booking/usecases"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Cannot load .env file")
	}

	conn, err := dbpg.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer conn.Close()

	// Initialize Fiber
	app := fiber.New()

	// Add CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(200)
		}
		return c.Next()
	})

	// Initialize Kafka Producer
	producer := bookingkafka.NewProducer("kafka:9092") // Adjust the broker address as needed

	// Initialize Kafka Consumer
	consumer := bookingkafka.NewConsumer("kafka:9092") // Adjust the broker address as needed
	go consumer.ReadMessages(context.Background())

	// Repositories and Services
	userRepo := authinfra.NewUserRepository(conn)
	authUsecase := authusecase.NewAuthenticationUseCase(userRepo)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	// Movie
	movieRepo := movieinfra.NewMovieRepository(conn)
	movieUseCase := movieusecases.NewMovieUseCase(movieRepo)
	movieHandler := moviehandler.NewMovieHandler(movieUseCase)

	// Theater Repositories and Services
	theaterRepo := theaterinfra.NewTheaterRepository(conn)
	screenRepo := theaterinfra.NewScreenRepository(conn)
	seatRepo := theaterinfra.NewSeatRepository(conn)
	showtimeRepo := theaterinfra.NewShowtimeRepository(conn)

	theaterUseCase := theaterusecase.NewTheaterUseCase(theaterRepo)
	screenUseCase := theaterusecase.NewScreenUseCase(screenRepo)
	seatUseCase := theaterusecase.NewSeatUseCase(seatRepo)
	showtimeUseCase := theaterusecase.NewShowtimeUseCase(showtimeRepo)

	theaterHandler := theaterhandler.NewTheaterHandler(theaterUseCase)
	screenHandler := theaterhandler.NewScreenHandler(screenUseCase)
	seatHandler := theaterhandler.NewSeatHandler(seatUseCase)
	showtimeHandler := theaterhandler.NewShowtimeHandler(showtimeUseCase)

	// Booking Repositories and Services
	pricingRepo := bookinginfra.NewTicketPricingRepository(conn)

	bookingRepo := bookinginfra.NewBookingRepository(conn)
	bookingUseCase := bookingUsecase.NewBookingUseCase(bookingRepo, pricingRepo, showtimeRepo, seatRepo, producer)
	bookingHandler := bookinghandler.NewBookingHandler(bookingUseCase)

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to the Booking Cinema API"})
	})

	// Auth Routes
	v1.Post("/auth/register", authHandler.RegisterHandler)
	v1.Post("/auth/login", authHandler.LoginHandler)
	v1.Get("/auth/user", authutils.AuthMiddleware(userRepo), authHandler.UserHandler)

	// Movie Routes
	movieRoute := v1.Group("/movies")
	movieRoute.Get(":id", movieHandler.GetMovieByID)
	movieRoute.Get("/", movieHandler.GetAllMovies)
	movieRoute.Get("/search", movieHandler.SearchMovies)
	movieRoute.Post("/", authutils.AuthMiddleware(userRepo), movieHandler.CreateMovie)
	movieRoute.Delete("/:id", authutils.AuthMiddleware(userRepo), movieHandler.DeleteMovie)

	// Showtimes
	movieShowtimesRoute := movieRoute.Group("/:movie_id/showtimes")
	movieShowtimesRoute.Get("/:id", showtimeHandler.GetShowtimeByID)
	movieShowtimesRoute.Get("/", showtimeHandler.GetShowtimesByMovie)
	movieShowtimesRoute.Get("/:id/seats", seatHandler.GetSeatsByShowtime)
	movieShowtimesRoute.Post("/", authutils.AuthMiddleware(userRepo), showtimeHandler.CreateShowtime)

	showtimesRoute := v1.Group("/showtimes")
	showtimesRoute.Get("/:id", showtimeHandler.GetShowtimeByID)
	showtimesRoute.Get("/:id/seats", seatHandler.GetSeatsByShowtime)
	showtimesRoute.Get("/", showtimeHandler.GetShowtimesByMovie)
	showtimesRoute.Post("/", authutils.AuthMiddleware(userRepo), showtimeHandler.CreateShowtime)

	// Theater Routes
	theaters := v1.Group("/theaters")
	theaters.Get("/", theaterHandler.GetAllTheaters)
	theaters.Get("/:id", theaterHandler.GetTheaterByID)
	theaters.Post("/", authutils.AuthMiddleware(userRepo), theaterHandler.CreateTheater)
	theaters.Put("/:id", authutils.AuthMiddleware(userRepo), theaterHandler.UpdateTheater)
	theaters.Delete("/:id", authutils.AuthMiddleware(userRepo), theaterHandler.DeleteTheater)

	screens := theaters.Group("/:theater_id/screens")
	screens.Get("/", screenHandler.GetScreensByTheater)
	screens.Post("/", authutils.AuthMiddleware(userRepo), screenHandler.CreateScreen)

	seats := screens.Group("/:screen_id/seats")
	seats.Get("/", seatHandler.GetSeatsByScreen)
	seats.Post("/", authutils.AuthMiddleware(userRepo), seatHandler.CreateSeat)

	// // Booking Routes
	bookings := v1.Group("/bookings")
	bookings.Post("/", authutils.AuthMiddleware(userRepo), bookingHandler.CreateBooking)
	bookings.Get("/", authutils.AuthMiddleware(userRepo), bookingHandler.GetUserBookings)

	// Start the Fiber server
	err = app.Listen(":8080")
	if err != nil {
		return
	}
}
