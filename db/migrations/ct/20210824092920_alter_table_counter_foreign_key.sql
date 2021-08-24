-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE counter
    ADD CONSTRAINT fk_session_1609
    FOREIGN KEY idx_session_1609 (username)
    REFERENCES `user` (username)
    ON UPDATE CASCADE
    ON DELETE RESTRICT;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
