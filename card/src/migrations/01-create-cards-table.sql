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
)
