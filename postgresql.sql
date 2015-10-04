-- Table: users

DROP TABLE users;

CREATE TABLE users
(
  id serial NOT NULL, -- user's id
  email character varying(80) NOT NULL, -- user's E-mail address
  dispalyname character varying(40) NOT NULL, -- user's display name.
  userpassword character varying(40) NOT NULL, -- user's password.
  createtime timestamp NOT NULL, -- the user created time
  lastlogin timestamp, -- the time when last time login for the user
  showemail boolean NOT NULL DEFAULT false, -- show email to other users
  userlevel integer NOT NULL DEFAULT 0, -- 0: guest, lowest level; 1: user, common; 2: manager; 3:admin
  CONSTRAINT pk PRIMARY KEY (id),
  CONSTRAINT "only" UNIQUE (email)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE users OWNER TO postgres;
COMMENT ON COLUMN users.dispalyname IS 'user''s display name.';
COMMENT ON COLUMN users.email IS 'user''s E-mail address';
COMMENT ON COLUMN users.id IS 'user''s id';
COMMENT ON COLUMN users.createtime IS 'the user created time';
COMMENT ON COLUMN users.lastlogin IS 'the time when last time login for the user';
COMMENT ON COLUMN users.showemail IS 'show email to other users';
COMMENT ON COLUMN users.userlevel IS '0: guest, lowest level; 1: user, common; 2: manager; 3:admin';


-- Index: qemail

-- DROP INDEX qemail;

CREATE INDEX qemail
  ON users
  USING btree
  (email);

