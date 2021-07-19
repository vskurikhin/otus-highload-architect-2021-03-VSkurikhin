-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE `user` (
 id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
 username   VARCHAR(32) NOT NULL,
 name       VARCHAR(128),
 surname    VARCHAR(128),
 age        SMALLINT,
 sex        TINYINT(1),
 city       VARCHAR(128)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
