export interface Team {
    id: number;
    name: string;
    played: number;
    wins: number;
    draws: number;
    losses: number;
    goalsFor: number;
    goalsAgainst: number;
    points: number;
}

export interface Match {
    id: number;
    team1: Team;
    team2: Team;
    score1: number;
    score2: number;
    stage: 'semifinal' | 'final';
}

export interface Group {
    id: number;
    name: string;
    teams: Team[];
}

export interface ApiError {
    message: string;
    statusCode?: number;
}