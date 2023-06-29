CREATE TABLE post_likes (
    post_id varchar(40) NOT NULL,
    user_id varchar(40) NOT NULL,
    created_at timestamp DEFAULT now()
);