BEGIN;

CREATE TABLE dates(
    id bigserial not null primary key,
    products_ids varchar(255),
    date_created int
);

COMMIT;