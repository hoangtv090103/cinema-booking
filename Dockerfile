FROM golang:1.22.3-alpine

WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o bookingCinema ./cmd/web

EXPOSE 8080
CMD ["./bookingCinema"]