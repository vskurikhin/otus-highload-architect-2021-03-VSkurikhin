-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT INTO interest (id, interests) VALUES (unhex(replace(uuid(),'-','')), 'Computer science');
INSERT INTO interest (id, interests) VALUES (unhex(replace(uuid(),'-','')), 'Software engineer');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
