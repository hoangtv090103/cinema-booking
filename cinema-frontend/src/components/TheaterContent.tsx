'use client';

import { useQuery } from 'react-query';
import { getTheaters } from '@/services/api';
import TheaterList from './TheaterList';

export default function TheaterContent() {
  const { data: theaters, isLoading, error } = useQuery('theaters', getTheaters);

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
        Error loading theaters. Please try again later.
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <h1 className="text-3xl font-bold">Our Theaters</h1>
      <TheaterList theaters={theaters || []} />
    </div>
  );
}