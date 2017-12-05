CREATE TABLE items (
    id serial UNIQUE,
    name TEXT not null
);

CREATE TABLE stocks (
    id serial UNIQUE,
    item_id integer not null REFERENCES items(id),
    price integer,
    sold boolean default false not null
);
