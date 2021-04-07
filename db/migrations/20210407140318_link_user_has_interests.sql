-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT user_has_interests
    (id, user_id, interest_id)
  VALUES
    (unhex(replace(uuid(),'-','')), 0x5CB9CAE1977911000000000000000000,
     (SELECT id FROM interest WHERE interests = 'Computer science'));

INSERT user_has_interests
    (id, user_id, interest_id)
  VALUES
    (unhex(replace(uuid(),'-','')), 0x5CB9CAE1977911000000000000000000,
     (SELECT id FROM interest WHERE interests = 'Software engineer'));

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
