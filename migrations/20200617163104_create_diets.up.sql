BEGIN;

CREATE TABLE diets(
    id bigserial not null primary key,
    user_id bigserial not null,
    name varchar(100),
    calory real,
    squi real,
    fat real,
    carboh real
);

COMMIT;