-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE role
(
    id   bigserial primary key,
    name varchar(255)
);

CREATE TABLE "user"
(
    id       bigserial primary key,
    name     varchar(255),
    email    varchar(255) unique,
    password text,
    active   bool default true
);

CREATE TABLE user_roles
(
    user_id bigint,
    role_id bigint
);

create unique index user_email_idx on "user" (email);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table user_roles;
drop table "user";
drop table role;
