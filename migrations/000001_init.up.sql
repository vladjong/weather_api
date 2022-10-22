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