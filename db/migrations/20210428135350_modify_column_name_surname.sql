-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- workaround ERROR 1071 (42000): Specified key was too long; max key length is 3072 bytes

ALTER TABLE `user` MODIFY COLUMN name VARCHAR(512);
ALTER TABLE `user` MODIFY COLUMN surname VARCHAR(256);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
