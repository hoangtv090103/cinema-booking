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