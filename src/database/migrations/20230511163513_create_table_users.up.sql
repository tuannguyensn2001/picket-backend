create table users
(
    id                serial primary key,
    username          varchar,
    email             varchar,
    password          varchar,
    email_verified_at integer,
    created_at        integer,
    updated_at        integer
)