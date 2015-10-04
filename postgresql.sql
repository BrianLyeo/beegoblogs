-- Table: users

DROP TABLE users;

CREATE TABLE users
(
  id serial NOT NULL, -- user's id
  email character varying(80) NOT NULL, -- user's E-mail address
  "name" character varying(40) NOT NULL, -- user's display name.
  "password" character varying(40) NOT NULL, -- user's password.
  "create" time with time zone NOT NULL, -- the user created time
  lastlogin time with time zone, -- the time when last time login for the user
  showemail boolean NOT NULL DEFAULT false, -- show email to other users
  "level" integer NOT NULL DEFAULT 0, -- 0: guest, lowest level; 1: user, common; 2: manager; 3:admin
  CONSTRAINT pk PRIMARY KEY (id),
  CONSTRAINT "only" UNIQUE (email)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE users OWNER TO postgres;
COMMENT ON COLUMN users."name" IS 'user''s display name.';
COMMENT ON COLUMN users.email IS 'user''s E-mail address';
COMMENT ON COLUMN users.id IS 'user''s id';
COMMENT ON COLUMN users."create" IS 'the user created time';
COMMENT ON COLUMN users.lastlogin IS 'the time when last time login for the user';
COMMENT ON COLUMN users.showemail IS 'show email to other users';
COMMENT ON COLUMN users."level" IS '0: guest, lowest level; 1: user, common; 2: manager; 3:admin';


-- Index: qemail

-- DROP INDEX qemail;

CREATE INDEX qemail
  ON users
  USING btree
  (email);

