-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE dialog_message (
 id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
 shard_id   INT NOT NULL,
 from_user  BIGINT UNSIGNED NOT NULL,
 to_user    BIGINT UNSIGNED NOT NULL,
 message    VARCHAR(256),
 parent_id  BIGINT UNSIGNED
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
