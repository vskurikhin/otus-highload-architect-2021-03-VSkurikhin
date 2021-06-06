-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE news ADD COLUMN username VARCHAR(32);

UPDATE news SET username = (SELECT username FROM login WHERE id = 0x5cb9cae1977911000000000000000000);

ALTER TABLE news MODIFY username VARCHAR(32) NOT NULL;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
