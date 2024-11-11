'use client';

import { useState } from 'react';
import { QueryClient, QueryClientProvider } from 'react-query';
import MainContent from '@/components/MainContent';
import Navbar from '@/components/Navbar';

export default function Home() {
  const [queryClient] = useState(() => new QueryClient());

  return (
    <QueryClientProvider client={queryClient}>
      <div className="min-h-screen bg-gray-100">
        <Navbar />
        <main className="container mx-auto px-4 py-8">
          <MainContent />
        </main>
      </div>
    </QueryClientProvider>
  );
}