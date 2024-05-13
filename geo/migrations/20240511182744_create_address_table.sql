-- +goose Up
CREATE TABLE address 
(
    id         serial4 NOT NULL,
    history_id int4 NOT NULL,
    address    varchar NULL,
    latitude   float8 NULL,
    longitude  float8 NULL,
    created_at timestamp DEFAULT now() NULL,
    CONSTRAINT fk_address_1 FOREIGN KEY (history_id) REFERENCES public.history_search_address(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose Down
DROP TABLE IF NOT EXISTS address;
