-- +goose Up
CREATE TABLE history_search_address 
(
    id         serial4 NOT NULL,
    history_id int4 NOT NULL,
    created_at timestamp DEFAULT now() NULL,
    CONSTRAINT fk_history_search_address_1 FOREIGN KEY (history_id) REFERENCES public.search_history(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX pk_history_search_address ON history_search_address USING btree (id);

-- +goose Down
DROP TABLE IF NOT EXISTS history_search_address;
