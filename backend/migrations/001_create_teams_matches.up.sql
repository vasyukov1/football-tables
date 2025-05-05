CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    team1_id INTEGER REFERENCES teams(id),
    team2_id INTEGER REFERENCES teams(id),
    score1 INTEGER CHECK (score1 >= 0),
    score2 INTEGER CHECK (score2 >= 0),
    group_id INTEGER,
    playoff_id INTEGER,
    stage VARCHAR(50),
    is_completed BOOLEAN NOT NULL DEFAULT FALSE
);