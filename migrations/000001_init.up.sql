CREATE TABLE Cities
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL UNIQUE,
    lat NUMERIC NOT NULL,
    lon NUMERIC NOT NULL,
    country VARCHAR(30) NOT NULL
);

CREATE TABLE Weathers
(
    id SERIAL PRIMARY KEY,
    city_id INT REFERENCES Cities (id) NOT NULL,
    temp NUMERIC NOT NULL,
    date TIMESTAMP NOT NULL,
    info JSON,
    CONSTRAINT unique_weather UNIQUE(city_id, date)
);

CREATE TABLE Users
(
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE User_Lists
(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users (id) ON DELETE CASCADE NOT NULL,
    title TEXT
);

CREATE TABLE List_Items
(
    id SERIAL PRIMARY KEY,
    city_id INT REFERENCES Cities (id) NOT NULL,
    list_id INT REFERENCES User_Lists (id) NOT NULL
);

