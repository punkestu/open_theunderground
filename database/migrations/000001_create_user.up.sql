CREATE TABLE users(
    id varchar(40) PRIMARY KEY,
    fullname varchar(255),
    username varchar(255) UNIQUE,
    password varchar(255),
    email varchar(255)
)