-- +goose Up
create table posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT NULL,
    published_at TIMESTAMP NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE    
);

-- +goose Down
drop table posts;