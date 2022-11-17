
create table if not exists users
( 
    id serial,
    available_balance float,
    reserved_balance float 
);

create type order_status as enum ('pending', 'released', 'cancelled');

create table if not exists orders
(
    id serial,
    user_id int,
    service_id int,
    order_id int,
    price float,
    created_at datetime,
    comments varchar(255),
    status order_status
);