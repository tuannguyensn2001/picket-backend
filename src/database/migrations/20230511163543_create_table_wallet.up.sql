create table wallets
(
    id         serial primary key,
    user_id    int not null,
    balance    int not null default 0,
    created_at integer,
    updated_at integer
)