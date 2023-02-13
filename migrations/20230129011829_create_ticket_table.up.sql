CREATE TABLE IF NOT EXISTS "hsrloto"."tickets"
(
    id bigserial PRIMARY KEY,
    prize_id bigint default null,
    name varchar(100) not null,
    surname varchar(100) not null,
    middle_name varchar(100) not null,
    phone varchar(100) not null,
    spi varchar(30) not null
)