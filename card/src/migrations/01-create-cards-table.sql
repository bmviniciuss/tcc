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

INSERT INTO cardms.cards
(id, pan, masked_pan, cvv, cardholder_name, "token", expiration_year, expiration_month, active, is_debit, is_credit)
VALUES('5e92ac92-9327-4aee-b6a0-430d6fc59a50', '123412341231234', '123********1234', '123', 'Vinicius Barboasa', '3a15b6f4-6c3a-46a6-bf1c-c6874e0fbf3d8e8160f8-cccf-4db4-b0a5-3a9ce30a179e', 2022,11,true, true, true);
