create table users
(
	id varchar(36) not null,
	name varchar(36) not null,
	cpf varchar(15) not null,
	cnpj varchar(25) not null,
	email varchar(200) not null,
	balance float default 0 null,
	is_seller boolean not null,
	password text not null,
	created_at datetime default current_timestamp() not null,
	updated_at datetime default null on update current_timestamp,
	primary key (id)
);

create unique index users_email_uindex
	on users (email);

create unique index users_id_uindex
	on users (id);

-- -------------------------------------------------------------------------------------------

create table transactions
(
	id varchar(36) not null,
	id_payer varchar(36) not null,
	id_payee varchar(36) not null,
	value float default 0 null,
	created_at datetime default current_timestamp() not null,
	updated_at datetime default null on update current_timestamp,
	primary key (id)
);
