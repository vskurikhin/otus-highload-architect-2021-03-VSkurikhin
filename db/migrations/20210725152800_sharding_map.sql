-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE sharding_map (
 id         INT PRIMARY KEY AUTO_INCREMENT,
 city       VARCHAR(128)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
