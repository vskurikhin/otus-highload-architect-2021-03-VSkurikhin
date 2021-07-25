-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE dialog_message (
 id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
 shard_id   INT NOT NULL,
 message    VARCHAR(256)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
