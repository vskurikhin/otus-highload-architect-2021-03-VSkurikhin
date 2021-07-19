-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE `session` (
 id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
 session_id BIGINT UNSIGNED NOT NULL
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
