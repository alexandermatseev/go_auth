-- +goose Up
CREATE TABLE users(
    id serial primary key,
    name varchar not null,
    role int not null ,
    email varchar,
    password varchar not null ,
    confirm_password varchar not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
DROP table users;
