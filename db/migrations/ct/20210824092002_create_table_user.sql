-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE `user` (
 id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
 username   VARCHAR(32) NOT NULL
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
