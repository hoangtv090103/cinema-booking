-- Create Roles table
CREATE TABLE roles
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    active     BOOLEAN   DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255)        NOT NULL,
    email      VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id    INT                 REFERENCES roles (id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    active     BOOLEAN   DEFAULT TRUE
);


-- Index on email for fast user lookup by email (used in login).
CREATE UNIQUE INDEX idx_users_email ON users (email);


-- Permissions table with CRUD-based permissions
CREATE TABLE permissions
(
    id          SERIAL PRIMARY KEY,
    create_perm BOOLEAN   DEFAULT FALSE,
    read_perm   BOOLEAN   DEFAULT FALSE,
    update_perm BOOLEAN   DEFAULT FALSE,
    delete_perm BOOLEAN   DEFAULT FALSE,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    active      BOOLEAN   DEFAULT TRUE
);

-- Role-Permission join table to associate roles with permissions
CREATE TABLE IF NOT EXISTS user_roles
(
    user_id INT REFERENCES users (id) ON DELETE CASCADE,
    role_id INT REFERENCES roles (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

CREATE TABLE role_permissions
(
    role_id       INT REFERENCES roles (id) ON DELETE CASCADE,
    permission_id INT REFERENCES permissions (id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);




CREATE TABLE IF NOT EXISTS movies
(
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    description  TEXT,
    release_date DATE,
    duration     INT          NOT NULL, -- Duration in minutes
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW(),
    active       BOOLEAN   DEFAULT TRUE
);

-- Index on title for fast lookups when searching by title.
CREATE INDEX idx_movies_title ON movies (title);

-- Index on release_date for filtering movies by release dates.
CREATE INDEX idx_movies_release_date ON movies (release_date);

-- Theaters
CREATE TABLE IF NOT EXISTS theaters
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    location   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    active     BOOLEAN   DEFAULT TRUE
);
-- Index on location to improve performance for location-based searches.
CREATE INDEX idx_theaters_location ON theaters (location);

-- Screens
CREATE TABLE IF NOT EXISTS screens
(
    id            SERIAL PRIMARY KEY,
    theater_id    INT REFERENCES theaters (id),
    name          VARCHAR(50) NOT NULL, -- e.g., "Screen 1", "VIP Screen"
    seat_capacity INT         NOT NULL, -- Total number of seats
    created_at    TIMESTAMP DEFAULT NOW(),
    updated_at    TIMESTAMP DEFAULT NOW(),
    active        BOOLEAN   DEFAULT TRUE
);
-- Index on theater_id for quick lookup of screens by theater.
CREATE INDEX idx_screens_theater_id ON screens (theater_id);


CREATE TABLE IF NOT EXISTS showtimes
(
    id         SERIAL PRIMARY KEY,
    movie_id   INT REFERENCES movies (id),
    screen_id  INT REFERENCES screens (id),
    start_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    active     BOOLEAN   DEFAULT TRUE
);
-- Index on movie_id for faster lookup by movie.
CREATE INDEX idx_showtimes_movie_id ON showtimes (movie_id);

-- Index on screen_id to support filtering by screen.
CREATE INDEX idx_showtimes_screen_id ON showtimes (screen_id);

-- Index on start_time to improve performance for filtering by showtime.
CREATE INDEX idx_showtimes_start_time ON showtimes (start_time);


CREATE TABLE IF NOT EXISTS seats
(
    id         SERIAL PRIMARY KEY,
    screen_id  INT REFERENCES screens (id),
    row        VARCHAR(5)  NOT NULL, -- Seat row, e.g., "A", "B"
    number     INT         NOT NULL, -- Seat number, e.g., "1", "2"
    seat_type  VARCHAR(50) NOT NULL, -- e.g., "Regular", "VIP"
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    active     BOOLEAN   DEFAULT TRUE,
    UNIQUE (screen_id, row, number)  -- Unique seat position per screen
);
-- Index on screen_id for retrieving all seats in a particular screen.
CREATE INDEX idx_seats_screen_id ON seats (screen_id);

-- Composite index on (screen_id, row, number) for unique seat positions in a screen.
CREATE UNIQUE INDEX idx_seats_screen_row_number ON seats (screen_id, row, number);


CREATE TABLE IF NOT EXISTS bookings
(
    id           SERIAL PRIMARY KEY,
    user_id      INT REFERENCES users (id),
    showtime_id  INT REFERENCES showtimes (id),
    booking_time TIMESTAMP DEFAULT NOW(),
    total_price  DECIMAL(10, 2) NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW(),
    active       BOOLEAN   DEFAULT TRUE
);
-- Index on user_id for fast retrieval of a user's booking history.
CREATE INDEX idx_bookings_user_id ON bookings (user_id);

-- Index on showtime_id for efficient lookup of bookings by showtime.
CREATE INDEX idx_bookings_showtime_id ON bookings (showtime_id);

-- Booking Seats
CREATE TABLE IF NOT EXISTS booking_seats
(
    id         SERIAL PRIMARY KEY,
    booking_id INT REFERENCES bookings (id) ON DELETE CASCADE,
    seat_id    INT REFERENCES seats (id),
    price      DECIMAL(10, 2) NOT NULL, -- Seat-specific price (e.g., VIP pricing)
    UNIQUE (booking_id, seat_id)        -- A specific seat can be booked only once per booking
);
-- Composite index on (booking_id, seat_id) for efficient joins and lookups.
CREATE UNIQUE INDEX idx_booking_seats_booking_seat_id ON booking_seats (booking_id, seat_id);


CREATE TABLE IF NOT EXISTS payments
(
    id             SERIAL PRIMARY KEY,
    booking_id     INT REFERENCES bookings (id),
    amount         DECIMAL(10, 2) NOT NULL,
    payment_time   TIMESTAMP DEFAULT NOW(),
    payment_status VARCHAR(50)    NOT NULL, -- e.g., "Completed", "Pending", "Failed"
    created_at     TIMESTAMP DEFAULT NOW(),
    updated_at     TIMESTAMP DEFAULT NOW()
);
-- Index on booking_id for fast lookup of payments by booking.
CREATE INDEX idx_payments_booking_id ON payments (booking_id);

-- Index on payment_status to improve filtering performance by status.
CREATE INDEX idx_payments_status ON payments (payment_status);


CREATE TABLE IF NOT EXISTS ticket_pricing
(
    id          SERIAL PRIMARY KEY,
    seat_type   VARCHAR(50)    NOT NULL, -- e.g., "Regular", "VIP"
    day_of_week INT            NOT NULL, -- 0=Sunday, 1=Monday, etc.
    price       DECIMAL(10, 2) NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    active      BOOLEAN   DEFAULT TRUE,
    UNIQUE (seat_type, day_of_week)      -- Each seat type has unique pricing per day of the week
);
-- Composite index on (seat_type, day_of_week) for fast lookups of pricing rules.
CREATE UNIQUE INDEX idx_ticket_pricing_seat_type_day ON ticket_pricing (seat_type, day_of_week);
