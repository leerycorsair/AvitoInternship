alter table "users" add constraint "user_id" primary key ("id");
alter table "orders" add constraint "order_id" primary key ("id");
alter table "transactions" add constraint "transaction_id" primary key ("id");

alter table "orders" add foreign key ("order_user_id") references "users" ("id") on delete cascade;
alter table "transaction" add foreign key ("transaction_user_id") references "users" ("id") on delete cascade;

alter table "users" alter column "available_balance" set not null;
alter table "users" alter column "reserved_balance" set not null;
alter table "users" add constraint "ab_check" check (available_balance >= 0);
alter table "users" add constraint "rb_check" check (reserved_balance >= 0);

alter table "orders" alter column "user_id" set not null;
alter table "orders" alter column "service_id" set not null;
alter table "orders" alter column "order_id" set not null;
alter table "orders" alter column "price" set not null;
alter table "orders" alter column "created_at" set not null;
alter table "orders" alter column "status" set not null;
alter table "orders" add constraint "price_check" check (price >= 0);

alter table "transactions" alter column "transaction_id" set not null;
alter table "transactions" alter column "user_id" set not null;
alter table "transactions" alter column "value" set not null;
alter table "transactions" alter column "created_at" set not null;
alter table "transactions" add constraint "value_check" check (value >= 0);
