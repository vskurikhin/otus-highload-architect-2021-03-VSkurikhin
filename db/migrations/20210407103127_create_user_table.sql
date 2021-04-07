-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE `user` (
 id         BINARY(16) NOT NULL PRIMARY KEY,
 username   VARCHAR(32) NOT NULL,
 name       VARCHAR(1478),
 surname    VARCHAR(700),
 age        SMALLINT,
 sex        TINYINT(1),
 interests  JSON NOT NULL,
 city       VARCHAR(330),
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 enabled    BOOL DEFAULT true
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
