-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE `user`
    ADD CONSTRAINT fk_user_9403
        FOREIGN KEY idx_user_9403 (id)
    REFERENCES login (id)
    ON UPDATE CASCADE
       ON DELETE RESTRICT;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
