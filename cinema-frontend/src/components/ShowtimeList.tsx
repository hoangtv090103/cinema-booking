'use client';

import { Showtime } from '@/types';
import { useRouter } from 'next/navigation';
import { format } from 'date-fns';

interface ShowtimeListProps {
  showtimes: Showtime[];
  movieId: number;
}

export default function ShowtimeList({ showtimes, movieId }: ShowtimeListProps) {
  const router = useRouter();
  console.log(showtimes);

  const handleShowtimeSelect = (showtimeId: number) => {
    router.push(`/movies/${movieId}/showtimes/${showtimeId}/seats`);
  };

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      {showtimes.map((showtime, index) => {
        const startTime = new Date(showtime.start_time);
        if (isNaN(startTime.getTime())) {
          return <div key={index} className="text-red-500">Invalid time</div>;
        }

        return (
          <button
            key={index}
            onClick={() => handleShowtimeSelect(showtime.id)}
            className="bg-white p-4 rounded-lg shadow hover:shadow-md transition text-left"
          >
            <div className="text-lg font-semibold">
              {format(startTime, 'h:mm a')}
            </div>
            <div className="text-sm text-gray-500">
              Screen {showtime.screenId}
            </div>
          </button>
        );
      })}
    </div>
  );
} 