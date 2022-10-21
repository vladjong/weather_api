CREATE TABLE town
(
    id serial not null unique,
    name varchar(255) not null,
    lon numeric not null,
    lat numeric not null
);