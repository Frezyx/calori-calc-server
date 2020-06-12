CREATE TABLE users(
    id bigserial not null primary key,
    email varchar not null unique,
    encripted_password varchar not null
);