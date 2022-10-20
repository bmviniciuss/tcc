CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS cardpaymentms;

CREATE TABLE IF NOT EXISTS cardpaymentms.payments (
  id uuid not null default uuid_generate_v4(),
  authorization_id uuid not null,
  client_id uuid not null,
  payment_type varchar(30) not null,
  amount float not null,
  cardholder_name varchar(256) not null,
  card_token varchar(300) not null,
  masked_number varchar(20) not null,
  payment_date TIMESTAMPTZ not null,
  created_at TIMESTAMPTZ default now(),
  constraint cards_pkey primary key (id)
);

create table if not exists cardpaymentms.payables (
  id uuid not null default uuid_generate_v4(),
  client_id uuid not null,
  payment_id uuid not null,
  amount float not null,
  payment_date TIMESTAMPTZ not null,
  created_at TIMESTAMPTZ default now(),
  constraint payables_pkey primary key (id),
  constraint payables_payment_id_fkey foreign key (payment_id) references cardpaymentms.payments(id)
);
