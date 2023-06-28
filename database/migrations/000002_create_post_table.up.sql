CREATE TABLE posts (
    id varchar(40) PRIMARY KEY,
    topic TEXT NOT NULL,
    author_id varchar(40),
    created_at timestamp DEFAULT now()
);