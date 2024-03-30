-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE account
(
    id      bigserial primary key,
    name    varchar(255),
    color   varchar(7),
    user_id bigint not null,
    foreign key (user_id) references "user" (id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table account;