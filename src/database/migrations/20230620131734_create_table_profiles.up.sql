create table profiles
(
    id         serial primary key,
    user_id    integer,
    avatar_url varchar,
    nickname   varchar,
    created_at integer,
    updated_at integer
)