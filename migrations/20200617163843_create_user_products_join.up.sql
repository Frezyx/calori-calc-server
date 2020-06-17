BEGIN;

CREATE TABLE user_products_join(
    id bigserial not null primary key,
    product_id bigserial not null,
    user_id bigserial not null
);

COMMIT;