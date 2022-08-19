CREATE DATABASE BOX_CATALOG;

CREATE TABLE IF NOT EXISTS boxes (
    box_id serial PRIMARY KEY,
    label VARCHAR NOT NULL,
    position VARCHAR NOT NULL,
    user_id UUID references users(user_id)
);

CREATE TABLE IF NOT EXISTS categories (
    category_id serial PRIMARY KEY,
    title VARCHAR NOT NULL,
    user_id UUID references users(user_id)
);

CREATE TABLE IF NOT EXISTS items (
    item_id serial PRIMARY KEY,
    item VARCHAR NOT NULL,
    box_id INTEGER REFERENCES boxes(box_id),
    category_id INTEGER REFERENCES categories(category_id),
    user_id UUID references users(user_id)
);

CREATE TABLE IF NOT EXISTS users (
	user_id UUID PRIMARY KEY,
    email VARCHAR UNIQUE,
    password VARCHAR NOT NULL
);
