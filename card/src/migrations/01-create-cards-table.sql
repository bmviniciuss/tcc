CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS cardms;

CREATE TABLE IF NOT EXISTS cardms.cards (
  id uuid not null default uuid_generate_v4(),
	pan varchar(16) not null,
	masked_pan varchar(16) not null,
	cvv varchar(3) not null,
	cardholder_name varchar(256) not null,
	token varchar(256) not null unique,
	expiration_year smallint not null,
	expiration_month smallint not null,
	active boolean not null,
	is_debit boolean not null,
	is_credit boolean not null,
	constraint cards_pkey primary key (id)
);

CREATE TABLE IF NOT EXISTS cardms.payment_authorization (
  	id uuid not null default uuid_generate_v4(),
	amount double precision not null,
	"status" varchar(16) not null,
	card_id uuid not null,
	transaction_date timestamptz not null,
	created_at timestamptz default now(),
	constraint payment_authorization_pkey primary key (id),
	constraint payment_authorization_card_fkey foreign key (card_id) references cardms.cards(id)
);

