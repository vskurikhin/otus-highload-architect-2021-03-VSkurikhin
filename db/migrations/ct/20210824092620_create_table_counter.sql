-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE counter (
 username   VARCHAR(32) NOT NULL PRIMARY KEY,
 total      BIGINT UNSIGNED NOT NULL,
 unread     BIGINT NOT NULL
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
