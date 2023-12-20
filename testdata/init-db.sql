CREATE TABLE if not exists movies (
  id int4 NOT NULL,
  "name" varchar NULL,
  CONSTRAINT movies_pk PRIMARY KEY (id)
);

insert into movies (id, name) values (99, 'test movie');