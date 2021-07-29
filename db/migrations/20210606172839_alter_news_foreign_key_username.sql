-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE news
    ADD CONSTRAINT fk_news_7833
        FOREIGN KEY idx_user_5283 (username)
    REFERENCES login (username)
    ON UPDATE CASCADE
       ON DELETE RESTRICT;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
