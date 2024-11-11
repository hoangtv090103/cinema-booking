import axios from 'axios';
import { Movie, Theater, Showtime } from '@/types';

const api = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
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