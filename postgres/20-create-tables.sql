\connect persons

create table if not exists persons (
    id serial primary key,
    name text not null,
    age     integer,
    address text,
    work    text
);

grant all privileges on all tables in schema public to program;
grant all privileges on all sequences IN schema public to program;