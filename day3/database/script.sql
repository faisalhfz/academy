create table if not exists customers (
	id int primary key,
	customer_name char(50) not null
);

create table if not exists products (
	id int primary key,
	name char(50) not null
);

create table if not exists orders (
	order_id int primary key,
	customer_id int references customers (id),
	product_id int references products (id),
	order_date date not null,
	total float(8) not null
);

insert into customers (id, customer_name) values (1, 'faisal');
insert into customers (id, customer_name) values (2, 'hafizh');

insert into products (id, name) values (1, 'apple');
insert into products (id, name) values (2, 'banana');
insert into products (id, name) values (3, 'cheese');

insert into orders (order_id, customer_id, product_id, order_date, total) 
	values (1, 1, 1, '2022/10/26', 1);
insert into orders (order_id, customer_id, product_id, order_date, total) 
	values (2, 1, 2, '2022/10/27', 2);
insert into orders (order_id, customer_id, product_id, order_date, total) 
	values (3, 2, 1, '2022/10/27', 2);
insert into orders (order_id, customer_id, product_id, order_date, total) 
	values (4, 2, 3, '2022/10/28', 1);

update customers set customer_name = 'hilman' where id = 2;

update products set name = 'orange' where id = 2;

update orders set total = 3 where order_id = 1;

delete from orders where order_date = '2022/10/27';

select o.order_id, c.customer_name, p.name, o.order_date, o.total
	from orders as o
	inner join customers as c
	on o.customer_id = c.id
	inner join products as p 
	on o.product_id = p.id;


