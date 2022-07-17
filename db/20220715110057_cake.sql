-- +goose Up
CREATE TABLE IF NOT EXISTS cake
(
    id          int(100) auto_increment
        primary key,
    title       varchar(100) not null,
    description varchar(200) not null,
    rating      float(5)     not null,
    image       varchar(500) not null,
    created_at  datetime     not null,
    updated_at  datetime     not null
);


-- +goose Down
DROP TABLE cake;
