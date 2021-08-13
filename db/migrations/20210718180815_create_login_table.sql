-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE login (
 id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
 username   VARCHAR(32) NOT NULL,
 password   VARCHAR(64) NOT NULL
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
