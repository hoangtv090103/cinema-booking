import { Theater } from '@/types';
import Link from 'next/link';

interface TheaterListProps {
  theaters: Theater[];
}

export default function TheaterList({ theaters }: TheaterListProps) {
  return (
    <div className="space-y-4">
      {theaters.map((theater) => (
        <div
          key={theater.id}
          className="bg-white p-4 rounded-lg shadow hover:shadow-md transition"
        >
          <div className="flex justify-between items-center">
            <div>
              <h2 className="text-xl font-semibold text-gray-900">{theater.name}</h2>
              <p className="text-gray-600">{theater.location}</p>
            </div>
            <Link
              href={`/theaters/${theater.id}`}
              className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              View Screens
            </Link>
          </div>
        </div>
      ))}
    </div>
  );
}
