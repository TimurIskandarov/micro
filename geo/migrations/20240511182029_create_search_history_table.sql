-- +goose Up
CREATE TABLE search_history 
(
	id         serial4 NOT NULL,
	history    varchar,
	created_at timestamp DEFAULT now() NULL
);
CREATE UNIQUE INDEX pk_search_history ON search_history USING btree (id);

-- +goose Down
DROP TABLE IF EXISTS search_history; 
