CREATE TABLE IF NOT EXISTS public.user (
   id int8 not null primary key,
   username varchar(255) not null,
   parent int8 default 0
);
