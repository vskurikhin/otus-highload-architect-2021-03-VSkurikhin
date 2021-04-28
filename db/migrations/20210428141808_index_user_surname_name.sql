-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE user ADD INDEX idx_user_surname_name_1342 USING BTREE (surname, name) COMMENT 'with index selectivity';

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
