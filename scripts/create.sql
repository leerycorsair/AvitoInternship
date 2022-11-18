
create table if not exists users
( 
    id serial,
    balance float,
);

create table if not exists orders
(
    id serial,
    user_id int,
    service_id int,
    order_id int,
    price float,
    created_at datetime,
    comments varchar(255),
    status int
);

create table if not exists transactions
(
    id serial,
    transaction_id int,
    user_id int,
    transaction_type int,
    value float,
    created_at datetime,
    action_comments varchar(255),
    add_comments varchar(255)
)