create table "users"(
	id serial primary key,
	name varchar(200),
	email varchar(200),
	pass varchar(100),
	address varchar(200),
	phone varchar(13),
	birth_date int
);

create table "admins"(
	id serial primary key,
	name varchar(200),
	email varchar(200),
	pass varchar(100),
	phone varchar(13)
);

create table "genres"(
	id serial primary key,
	genre_name varchar(200)
);

create table "books_request"(
	id serial primary key,
	user_id int,
	admin_id int,
	tittle varchar(200),
	author varchar(200),
	publisher varchar(200),
	publish_year varchar(200),
	status_req varchar(200),
	foreign key(user_id) references users(id),
	foreign key(admin_id) references admins(id)
	
);


create table "books"(
	id serial primary key,
	genre_id int,
	tittle varchar(200),
	author varchar(200),
	publiser varchar(100),
	publish_year varchar(100),
	foreign key(genre_id) references genres(id)
);


create table "loans"(
	id serial primary key,
	user_id int,
	book_id int,
	deadline_date varchar(100),
	date_loan varchar(100),
	date_return varchar(100),
	status_loan varchar(100),
	foreign key(user_id) references users(id),
	foreign key(book_id) references books(id)
);


insert into loans (status_loan) values
('dipinjam'), ('dikembalikan');

insert into genres (genre_name) values
('drama'), ('thriller'), ('Sci-fi'), ('adventure'), ('comedy');

insert into books (tittle, genre_id) values
('gumiho', 1), ('7 escape', 1), ('wrong turn', 2), ('scream', 2), ('planet of appes', 3), ('transformers', 3), ('king scorpion', 4), ('rambo', 4), ('tetangga sebelah', 5), ('cinta brontosaurus', 5);


insert into users (id,"name") values
(1, 'edi'), (2, 'edward'), (3, 'mia'), (4, 'zelen'), (5, 'anos');

select * from books b where tittle = 'gumiho';

select * from books b where genre_id = 4;

insert into loans (user_id, book_id, status_loan) values
(1,1, 'dipinjam'), (1,2, 'dipinjam'), (1, (select id from books where genre_id = 2 limit 1), 'dipinjam'),
(2,3, 'dipinjam'), (3,1, 'dipinjam');

update loans set status_loan = 'dikembalikan' where user_id = 1 and book_id = 1;

select users.name as user_name, books.tittle as book_tittle
from users 
join loans on users.id = loans.user_id 
join books  on loans.book_id = books.id 
where loans.status_loan = 'dipinjam';


select books.tittle as book_tittle
from users
join loans on users.id = loans.user_id 
join books on loans.book_id = books.id
where loans.status_loan = 'dikembalikan' and users.name = 'edi'