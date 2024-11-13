import axios from 'axios';
import { Movie, Theater, Showtime, Seat, Booking } from '@/types';

const api = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
  headers: {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${localStorage.getItem('token')}`
  }
});


export const getMovies = async (): Promise<Movie[]> => {
  const { data } = await api.get('/movies');
  return data;
};

export const getTheaters = async (): Promise<Theater[]> => {
  const { data } = await api.get('/theaters');
  return data;
};

export const getShowtimes = async (movieId: number): Promise<Showtime[]> => {
  const { data } = await api.get(`/movies/${movieId}/showtimes`);
  return data;
};

export const getMovie = async (id: number): Promise<Movie> => {
  const { data } = await api.get(`/movies/${id}`);
  return data;
};

export const getSeats = async (showtimeId: number): Promise<Seat[]> => {
  const { data } = await api.get(`/showtimes/${showtimeId}/seats`);
  return data;
};

export const createBooking = async (booking: {
  showtime_id: number;
  seat_ids: number[];
}): Promise<Booking> => {
  console.log(booking);
  const { data } = await api.post(
    '/bookings',
    booking
  );
  return data;
};