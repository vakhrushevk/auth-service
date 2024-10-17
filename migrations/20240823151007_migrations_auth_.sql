-- +goose Up
create table role(
                     id serial primary key,
                     name text not null
);

create table auth (
                      id serial primary key,
                      name text not null,
                      email text not null unique,
                      role int references role(id),
                      password text not null,
                      created_at timestamp not null default now(),
                      updated_at timestamp
);

-- +goose Down
drop table auth;
drop table role;
