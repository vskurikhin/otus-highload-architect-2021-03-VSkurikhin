-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE user_has_friends (
 id          BINARY(16) NOT NULL PRIMARY KEY,
 user_id     BINARY(16) NOT NULL,
 friend_id   BINARY(16) NOT NULL,
 created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 enabled     BOOL DEFAULT true
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
