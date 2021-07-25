-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE dialog_message
    ADD CONSTRAINT fk_shard_id_9463
        FOREIGN KEY idx_shard_id_9463 (id)
    REFERENCES sharding_map (shard_id)
    ON UPDATE CASCADE
       ON DELETE RESTRICT;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
