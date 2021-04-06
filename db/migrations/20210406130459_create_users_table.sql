-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE users (
 id         BINARY(16) NOT NULL PRIMARY KEY,
 username   VARCHAR(32),
 password   VARCHAR(64),
 name       VARCHAR(1478),
 surname    VARCHAR(700),
 age        SMALLINT,
 sex        TINYINT(1),
 interests  JSON NOT NULL,
 city       VARCHAR(330)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
