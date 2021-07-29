-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE user_has_friends
    ADD CONSTRAINT fk_user_has_friends_0242
        FOREIGN KEY key_user_has_friends_0242 (user_id)
    REFERENCES `user` (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT;

ALTER TABLE user_has_friends
    ADD CONSTRAINT fk_user_has_friends_7122
        FOREIGN KEY key_user_has_friends_7122 (friend_id)
    REFERENCES `user` (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
