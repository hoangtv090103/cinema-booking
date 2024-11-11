'use client';

import { useQuery } from 'react-query';
import { getMovies, getTheaters } from '@/services/api';
import MovieList from './MovieList';
import TheaterList from './TheaterList';

export default function MainContent() {
  const { 
    data: movies, 
    isLoading: isLoadingMovies, 
    error: movieError 
  } = useQuery('movies', getMovies);

  const { 
    data: theaters, 
    isLoading: isLoadingTheaters, 
    error: theaterError 
  } = useQuery('theaters', getTheaters);

  if (isLoadingMovies || isLoadingTheaters) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-500"></div>
      </div>    
    );
  }

  if (movieError || theaterError) {
    return (
      <div className="text-center text-red-500">
        Error loading content. Please try again later.
      </div>
    );
  }

  return (
    <div className="space-y-12">
      <div className="space-y-6">
        <h1 className="text-3xl font-bold">Now Showing</h1>
        <MovieList movies={movies || []} />
      </div>
      
      <div className="space-y-6">
        <h1 className="text-3xl font-bold">Our Theaters</h1>
        <TheaterList theaters={theaters || []} />
      </div>
    </div>
  );
}