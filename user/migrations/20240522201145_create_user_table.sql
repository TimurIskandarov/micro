-- Active: 1715467903239@@127.0.0.1@54321@geo
-- +goose Up
CREATE TABLE user (
    id       serial4 NOT NULL
    name     text NOT NULL,
    email    text NOT NULL,
    password text NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS user; 
