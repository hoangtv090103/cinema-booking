package main

import (
	authhandler "bookingcinema/pkg/auth/handlers"
	authinfra "bookingcinema/pkg/auth/infrastructe/postgres"
	authusecase "bookingcinema/pkg/auth/usecases"
	authutils "bookingcinema/pkg/auth/utils"
	dbpg "bookingcinema/pkg/database"

	moviehandler "bookingcinema/pkg/movie/handlers"
	"bookingcinema/pkg/movie/infrastructure/movieinfra"
	movieusecases "bookingcinema/pkg/movie/usecases"

	theaterhandler "bookingcinema/pkg/theater/handlers"
	theaterinfra "bookingcinema/pkg/theater/infrastructure/postgres"
	theaterusecase "bookingcinema/pkg/theater/usecases"

	bookinghandler "bookingcinema/pkg/booking/handlers"
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
		c.Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(200)
		}
		return c.Next()
	})

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
	bookingUseCase := bookingUsecase.NewBookingUseCase(bookingRepo, pricingRepo, showtimeRepo, seatRepo)
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
	movies := v1.Group("/movies")
	movies.Get("/", movieHandler.GetAllMovies)
	movies.Get("/search", movieHandler.SearchMovies)
	movies.Post("/", authutils.AuthMiddleware(userRepo), movieHandler.CreateMovie)
	movies.Delete("/:id", authutils.AuthMiddleware(userRepo), movieHandler.DeleteMovie)

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

	showtimes := v1.Group("/showtimes")
	showtimes.Get("/movie/:movie_id", showtimeHandler.GetShowtimesByMovie)
	showtimes.Post("/", authutils.AuthMiddleware(userRepo), showtimeHandler.CreateShowtime)

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
