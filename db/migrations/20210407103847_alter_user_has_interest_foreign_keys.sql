-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE user_has_interest
    ADD CONSTRAINT fk_user_has_interest_9774
        FOREIGN KEY idx_user_has_interest_9774 (user_id)
        REFERENCES `user` (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT;

ALTER TABLE user_has_interest
    ADD CONSTRAINT fk_user_has_interest_6459
        FOREIGN KEY idx_user_has_interest_6459 (interest_id)
        REFERENCES interest (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
