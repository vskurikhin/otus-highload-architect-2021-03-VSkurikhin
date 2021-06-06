-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE news (
 id         BINARY(16) NOT NULL PRIMARY KEY,
 title      VARCHAR(256) NOT NULL,
 content    VARCHAR(1024),
 public_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 enabled    BOOL DEFAULT true
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
