create table if not exists snippets
(
    id serial not null
        constraint snippets_pk
            primary key,
    title text not null,
    content text not null,
    created timestamp not null,
    expires timestamp
);

create unique index if not exists snippets_title_uindex
    on snippets (title);

create table if not exists users
(
    id serial not null
        constraint users_pk
            primary key,
    name text not null,
    email text not null,
    password text not null,
    created timestamp not null
);

create unique index if not exists users_email_uindex
    on users (email);

