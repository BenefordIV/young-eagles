BEGIN;

create table child
(
    id              uuid primary key     default uuidv7(),
    first_name      text        not null,
    last_name       text        not null,
    date_of_birth   date        not null,
    has_certificate boolean     not null default false,
    created_at      timestamptz not null default now(),
    updated_at      timestamptz not null default now()
);

create table pilot
(
    id         uuid primary key default uuidv7(),
    first_name text not null,
    last_name  text not null,
    email      text not null,
    eaa_chapter integer not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create table plane (
    call_number text primary key not null,
    model       text not null,
    make        text not null,
    create_at    timestamptz not null default now(),
    updated_at   timestamptz not null default now()
);


create table flight (
    id uuid primary key default uuidv7(),
    child_id uuid not null references child(id),
    pilot_id uuid not null references pilot(id),
    plane_call_number text not null references plane(call_number),
    flight_date date not null default now(),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

COMMIT;