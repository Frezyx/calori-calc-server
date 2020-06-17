BEGIN;

CREATE TABLE user_products(
    id bigserial not null primary key,
    productId bigserial not null,
    name varchar(100),
    category varchar(100),
    calory real,
    squi real,
    fat real,
    carboh real,
    grams real,
    date_created int
);

COMMIT;