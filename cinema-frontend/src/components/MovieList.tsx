'use client';

import { Movie } from '@/types';
import Link from 'next/link';

interface MovieListProps {
  movies: Movie[];
}

export default function MovieList({ movies }: MovieListProps) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {movies.map((movie) => (
        <div
          key={movie.id}
          className="bg-white rounded-lg shadow overflow-hidden"
        >
          <div className="p-6">
            <h2 className="text-xl font-semibold mb-2">{movie.title}</h2>
            <p className="text-gray-600 mb-4">{movie.description}</p>
            <div className="flex justify-between items-center">
              <span className="text-sm text-gray-500">
                {new Date(movie.releaseDate).toLocaleDateString()}
              </span>
              <Link
                href={`/movies/${movie.id}`}
                className="bg-indigo-600 text-white px-4 py-2 rounded-md hover:bg-indigo-700 transition"
              >
                Book Now
              </Link>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}