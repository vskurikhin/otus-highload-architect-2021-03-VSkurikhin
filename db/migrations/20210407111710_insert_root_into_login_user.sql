-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT INTO login
    (id, username, password)
  VALUES
    (0x5CB9CAE1977911000000000000000000, 'root', '$2a$10$1rUifnvj61PCVusPXXowJuHePZe8JiisXFmQpytyhQI4Gy0Eq2NPe');

INSERT INTO user
    (id, username, name, surname, age, sex, city)
  VALUES
    (0x5CB9CAE1977911000000000000000000, 'root', 'Charlie', 'Root', 51, 1, 'Murray Hill');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
