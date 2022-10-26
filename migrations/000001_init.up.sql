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

CREATE TABLE User_lists
(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Cities (id) ON DELETE CASCADE NOT NULL,
    like_list_id INT REFERENCES LikeLists (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE LikeLists
(
    id SERIAL PRIMARY KEY,
    city_id INT REFERENCES Cities (id) NOT NULL
);