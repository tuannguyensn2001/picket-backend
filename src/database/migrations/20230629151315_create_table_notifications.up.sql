CREATE TABLE notifications
(
    id         serial primary key,
    "from"     integer,
    "to"       integer,
    type       integer,
    payload    jsonb,
    template   varchar,
    read_at    integer,
    created_at integer,
    updated_at integer
);