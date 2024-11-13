export interface Movie {
    id: number;
    title: string;
    description: string;
    releaseDate: string;
    duration: number;
}

export interface Theater {
    id: number;
    name: string;
    location: string;
    active: boolean;
}

export interface Showtime {
    id: number;
    movieId: number;
    screenId: number;
    startTime: string;
    active: boolean;
}

export interface Seat {
    id: number;
    row: string;
    number: number;
    isAvailable: boolean;
    screenId: number;
}

export interface Booking {
    id: number;
    user_id: number;
    showtime_id: number;
    seat_ids: number[];
    status: string;
    created_at: string;
}