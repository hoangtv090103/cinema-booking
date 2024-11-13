import { ReactNode } from 'react';
import Link from 'next/link';

interface LayoutProps {
  children: ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  return (
    <div className="min-h-screen bg-gray-100">
      <nav className="bg-white shadow-lg">
        <div className="max-w-7xl mx-auto px-4">
          <div className="flex justify-between h-16">
            <div className="flex">
              <Link href="/" className="flex items-center">
                <span className="text-xl font-bold">Cinema Booking</span>
              </Link>
              <div className="ml-10 flex items-center space-x-4">
                <Link href="/movies" className="text-gray-700 hover:text-gray-900">
                  Movies
              </Link>
                <Link href="/theaters" className="text-gray-700 hover:text-gray-900">
                  Theaters
                </Link>
              </div>
            </div>
          </div>
        </div>
      </nav>
      <main className="max-w-7xl mx-auto py-6 px-4">{children}</main>
    </div>
  );
}