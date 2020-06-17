BEGIN;

CREATE TABLE products(
    id bigserial not null primary key,
    name varchar(100),
    category varchar(100),
    calory real,
    squi real,
    fat real,
    carboh real
);

COMMIT;