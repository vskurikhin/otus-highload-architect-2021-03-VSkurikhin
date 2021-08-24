-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE sharding_map (
 id         INT PRIMARY KEY AUTO_INCREMENT,
 field      VARCHAR(64) NOT NULL,
 city       VARCHAR(128),
 shard_id   INT
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
