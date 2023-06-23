create table profiles
(
    id         serial primary key,
    user_id    integer,
    avatar_url string,
    nickname   string,
    created_at integer,
    updated_at integer
)