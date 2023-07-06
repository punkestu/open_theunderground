CREATE TABLE post_comments
(
    id         varchar(40) PRIMARY KEY,
    post_id    varchar(40) NOT NULL,
    user_id    varchar(40) NOT NULL,
    comment    text        NOT NULL,
    created_at timestamp DEFAULT now()
);