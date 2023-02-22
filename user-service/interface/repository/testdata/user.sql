CREATE TABLE public.users (
  id bigserial NOT NULL PRIMARY KEY,
  first_name character varying(25),
  last_name character varying(25),
  username character varying(25) NOT NULL UNIQUE,
  password character varying(255),
  email character varying(255)
);