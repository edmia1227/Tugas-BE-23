create table "users"(
	id int,
	name varchar(10),
	email varchar(10),
	pass varchar(10),
	address varchar(10),
	phone varchar(13),
	birth_date varchar(10),
	primary key(id)
);

create table "admins"(
	id int,
	name varchar(10),
	email varchar(10),
	pass varchar(10),
	phone varchar(13),
	primary key(id)
);

create table "genres"(
	id int,
	genre_name varchar(10),
	primary key(id)
);

create table "books_request"(
	id int,
	user_id int,
	admin_id int,
	tittle varchar(10),
	author varchar(10),
	publisher varchar(10),
	publish_year varchar(10),
	status_req varchar(10),
	foreign key(user_id) references users(id),
	foreign key(admin_id) references admins(id),
	primary key(id)
	
);


create table "books"(
	id int,
	genre_id int,
	tittle varchar(10),
	author varchar(10),
	publiser varchar(10),
	publish_year varchar(10),
	foreign key(genre_id) references genres(id),
	primary key(id)
	
);


create table "loans"(
	id int,
	user_id int,
	book_id int,
	deadline_date varchar(10),
	date_loan varchar(10),
	date_return varchar(10),
	status_loan varchar(10),
	foreign key(user_id) references users(id),
	foreign key(book_id) references books(id),
	primary key(id)
);