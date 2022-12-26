CREATE TABLE IF NOT EXISTS activities(
    id uuid PRIMARY KEY,
    s_title VARCHAR(25) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price       NUMERIC(5, 2),
    available_from   date,
    available_to     date,
    created_at       timestamp
);