-- +goose Up
create table posts (
    id UUID primary key ,
    created_at TIMESTAMP not null ,
    updated_at timestamp not null ,
    title TEXT NOT NULL ,
    description TEXT,
    published_at timestamp ,
    url TEXT not null unique ,
    feed_id UUID not null references feeds(id) on delete cascade
);

-- +goose Down
drop table posts;