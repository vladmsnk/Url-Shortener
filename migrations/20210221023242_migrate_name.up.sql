CREATE TABLE IF NOT EXISTS activity(
    id uuid PRIMARY KEY,
    s_title VARCHAR(25) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price       NUMERIC(5, 2),
    available_from   timestamptz,
    available_to     timestamptz,
    created_at       timestamp
);

CREATE TABLE IF NOT EXISTS selection(
    id uuid PRIMARY KEY,
    user_id uuid,
    title VARCHAR(25) NOT NULL,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS activities_for_selection (
    selection_id uuid,
    activity_id uuid,
    CONSTRAINT activities_for_selection_pk
        PRIMARY KEY (selection_id, activity_id)
);