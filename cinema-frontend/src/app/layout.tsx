import './globals.css';
import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Cinema Booking',
  description: 'Book your favorite movies',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}