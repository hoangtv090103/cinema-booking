'use client';

import { useQuery, QueryClient, QueryClientProvider, useMutation } from 'react-query';
import { useParams, useRouter } from 'next/navigation';
import { getSeats, createBooking } from '@/services/api';
import Layout from '@/components/Layout';
import { useState } from 'react';

const queryClient = new QueryClient();

export default function SeatSelectionPage() {
  return (
    <QueryClientProvider client={queryClient}>
      <SeatSelectionContent />
    </QueryClientProvider>
  );
}

function SeatSelectionContent() {
  const params = useParams();
  const router = useRouter();
  const [selectedSeats, setSelectedSeats] = useState<number[]>([]);

  const { data: seats, isLoading } = useQuery(['seats', params.showtimeId], () =>
    getSeats(Number(params.showtimeId))
  );

  const bookingMutation = useMutation(createBooking, {
    onSuccess: () => {
      router.push('/booking/confirmation');
    },
  });

  const handleSeatSelect = (seatId: number) => {
    setSelectedSeats((prev) =>
      prev.includes(seatId)
        ? prev.filter((id) => id !== seatId)
        : [...prev, seatId]
    );
  };

  const handleBooking = () => {
    bookingMutation.mutate({
      showtime_id: Number(params.showtimeId),
      seat_ids: selectedSeats,
    });
  };

  if (isLoading) {
    return (
      <Layout>
        <div className="flex justify-center items-center h-64">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-500"></div>
        </div>
      </Layout>
    );
  }

  return (
    <Layout>
      <div className="space-y-6">
        <h1 className="text-3xl font-bold">Select Your Seats</h1>
        <div className="grid grid-cols-8 gap-2">
          {seats?.map((seat) => (
            console.log(seat),
            <button
              key={seat.id}
              onClick={() => handleSeatSelect(seat.id)}
              disabled={!seat.available}
              className={`p-2 text-center rounded ${
                selectedSeats.includes(seat.id)
                  ? 'bg-indigo-600 text-white'
                  : seat.available
                  ? 'bg-white hover:bg-gray-100'
                  : 'bg-gray-200 cursor-not-allowed'
              }`}
            >
              {seat.row}{seat.seat_number}
            </button>
          ))}
        </div>
        <button
          onClick={handleBooking}
          disabled={selectedSeats.length === 0}
          className="w-full bg-indigo-600 text-white py-2 px-4 rounded-md hover:bg-indigo-700 disabled:bg-gray-300"
        >
          Book Selected Seats
        </button>
      </div>
    </Layout>
  );
} 