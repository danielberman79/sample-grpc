CREATE TABLE IF NOT EXISTS comments(
                                       id uuid PRIMARY KEY,
                                       name text,
                                       comment text,
                                       created_at timestamp
);
