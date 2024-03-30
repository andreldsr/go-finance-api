-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE transaction
(
    id          bigserial primary key,
    description varchar(255),
    color       varchar(7),
    user_id     bigint not null,
    category_id bigint not null,
    account_id  bigint not null,
    date        timestamp default now(),
    foreign key (user_id) references "user" (id),
    foreign key (category_id) references category (id),
    foreign key (account_id) references account (id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table transaction;