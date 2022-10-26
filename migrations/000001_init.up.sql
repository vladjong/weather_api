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
    login VARCHAR(30) NOT NULL UNIQUE,
    password_hash VARCHAR(30) NOT NULL
);

CREATE TABLE Users_lists
(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Cities (id) ON DELETE CASCADE NOT NULL,
    list_id INT REFERENCES Lists (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE Lists
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    city_id INT REFERENCES Cities (id) NOT NULL
);