-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE user_has_interests
    ADD CONSTRAINT fk_user_has_interests_9774
        FOREIGN KEY key_user_has_interests_9774 (user_id)
        REFERENCES `user` (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT;

ALTER TABLE user_has_interests
    ADD CONSTRAINT fk_user_has_interests_6459
        FOREIGN KEY key_user_has_interests_6459 (interest_id)
        REFERENCES interest (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
