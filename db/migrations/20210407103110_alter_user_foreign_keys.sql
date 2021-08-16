-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE `session`
    ADD CONSTRAINT fk_session_2250
        FOREIGN KEY idx_session_2250 (id)
        REFERENCES login (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
