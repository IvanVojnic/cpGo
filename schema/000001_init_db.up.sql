CREATE TABLE users
(
    id       serial       not null unique,
    email    varchar(255) not null,
    name     varchar(255) not null,
    password varchar(255) not null
);

CREATE TABLE friends
(
    user_sender   int references users (id) not null,
    user_receiver int references users (id) not null,
    status       varchar(255)              not null
);

CREATE TABLE rooms
(
    id              serial                    not null unique,
    id_user_creator int references users (id) not null,
    date_event      date,
    id_place           int
);

CREATE TABLE statuses
(
    id     serial       not null unique,
    status varchar(255) not null
);

CREATE TABLE invites
(
    id serial                          not null unique,
    user_id int references users (id)      not null,
    room_id int references rooms (id)      not null,
    status_id int references statuses (id) not null
);