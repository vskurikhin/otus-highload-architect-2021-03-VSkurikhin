-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE dialog_message ADD INDEX (shard_id);
ALTER TABLE sharding_map ADD INDEX (shard_id);

ALTER TABLE dialog_message
    ADD CONSTRAINT fk_shard_id_9163
    FOREIGN KEY fk_shard_id_9163 (shard_id)
    REFERENCES sharding_map (shard_id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT;

ALTER TABLE dialog_message
    ADD CONSTRAINT fk_from_user_9249
    FOREIGN KEY idx_from_user_9249 (from_user)
    REFERENCES `user` (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT;

ALTER TABLE dialog_message
    ADD CONSTRAINT fk_to_user_9310
    FOREIGN KEY idx_to_user_9310 (to_user)
    REFERENCES `user` (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
