create table if not exists brand(
  brand serial primary key,
  name varchar(80) not null
);

create table if not exists model(
  model serial primary key,
  brand int references brand(brand),
  name varchar(80) not null
);

create table if not exists condition(
  condition serial primary key,
  name varchar(24) not null
);

create table if not exists product(
  product serial primary key,
  name varchar(80) not null,
  code varchar(80) not null,
  model int references model(model),
  descriptions jsonb,
  price int,
  image text,
  amount int,
  deleted boolean
);