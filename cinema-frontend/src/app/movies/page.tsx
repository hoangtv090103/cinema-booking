'use client';

import { useQuery, QueryClient, QueryClientProvider } from 'react-query';
import { getMovies } from '@/services/api';
import MovieList from '@/components/MovieList';

const queryClient = new QueryClient();

export default function MoviesPage() {
  return (
    <QueryClientProvider client={queryClient}>
      <MoviesContent />
    </QueryClientProvider>
  );
}

function MoviesContent() {
  const { data: movies, isLoading, error } = useQuery('movies', getMovies);

  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-500"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="text-center text-red-500">
        Error loading movies. Please try again later.
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <h1 className="text-3xl font-bold">Now Showing</h1>
      <MovieList movies={movies || []} />
    </div>
  );
}