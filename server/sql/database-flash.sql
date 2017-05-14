-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.0-beta
-- PostgreSQL version: 9.6
-- Project Site: pgmodeler.com.br
-- Model Author: ---

-- object: nico | type: ROLE --
-- DROP ROLE IF EXISTS nico;
CREATE ROLE nico WITH 
	SUPERUSER
	CREATEDB
	CREATEROLE
	INHERIT
	LOGIN
	REPLICATION
	ENCRYPTED PASSWORD '********';
-- ddl-end --

-- object: admin | type: ROLE --
-- DROP ROLE IF EXISTS admin;
CREATE ROLE admin WITH 
	SUPERUSER
	CREATEDB
	CREATEROLE
	INHERIT
	LOGIN
	REPLICATION
	ENCRYPTED PASSWORD '********';
-- ddl-end --

-- object: root | type: ROLE --
-- DROP ROLE IF EXISTS root;
CREATE ROLE root WITH 
	SUPERUSER
	CREATEDB
	CREATEROLE
	INHERIT
	LOGIN
	REPLICATION
	ENCRYPTED PASSWORD '********';
-- ddl-end --


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: flash | type: DATABASE --
-- -- DROP DATABASE IF EXISTS flash;
-- CREATE DATABASE flash
-- 	ENCODING = 'UTF8'
-- 	LC_COLLATE = 'en_CA.UTF-8'
-- 	LC_CTYPE = 'en_CA.UTF-8'
-- 	TABLESPACE = pg_default
-- 	OWNER = postgres
-- ;
-- -- ddl-end --
-- 

-- object: main_schema | type: SCHEMA --
-- DROP SCHEMA IF EXISTS main_schema CASCADE;
CREATE SCHEMA main_schema;
-- ddl-end --
ALTER SCHEMA main_schema OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,main_schema;
-- ddl-end --

-- object: main_schema.card | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.card CASCADE;
CREATE TABLE main_schema.card(
	id serial NOT NULL,
	contents varchar(10000),
	creation_time timestamp NOT NULL,
	CONSTRAINT card_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE main_schema.card OWNER TO postgres;
-- ddl-end --

-- object: main_schema.collection | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.collection CASCADE;
CREATE TABLE main_schema.collection(
	id serial NOT NULL,
	deck_id varchar(128) NOT NULL,
	creation_time timestamp NOT NULL,
	CONSTRAINT collection_pk PRIMARY KEY (id,deck_id)

);
-- ddl-end --
ALTER TABLE main_schema.collection OWNER TO postgres;
-- ddl-end --

-- object: main_schema.card_diff | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.card_diff CASCADE;
CREATE TABLE main_schema.card_diff(
	id integer NOT NULL,
	card_id integer NOT NULL,
	user_id integer NOT NULL,
	action varchar(32) NOT NULL,
	data varchar(10000),
	creation_time timestamp NOT NULL,
	CONSTRAINT diff_pk PRIMARY KEY (id,card_id)

);
-- ddl-end --
ALTER TABLE main_schema.card_diff OWNER TO postgres;
-- ddl-end --

-- object: main_schema.card_action | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.card_action CASCADE;
CREATE TABLE main_schema.card_action(
	name varchar(32) NOT NULL,
	CONSTRAINT action_pk PRIMARY KEY (name)

);
-- ddl-end --
ALTER TABLE main_schema.card_action OWNER TO postgres;
-- ddl-end --

-- object: main_schema.owner | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.owner CASCADE;
CREATE TABLE main_schema.owner(
	user_id integer NOT NULL,
	object_id integer NOT NULL,
	object_type varchar(32) NOT NULL,
	creation_time timestamp NOT NULL,
	CONSTRAINT owner_pkey PRIMARY KEY (user_id,object_id,object_type)

);
-- ddl-end --
ALTER TABLE main_schema.owner OWNER TO postgres;
-- ddl-end --

-- object: main_schema.object_type | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.object_type CASCADE;
CREATE TABLE main_schema.object_type(
	name varchar(32) NOT NULL,
	CONSTRAINT object_types_pk PRIMARY KEY (name)

);
-- ddl-end --
ALTER TABLE main_schema.object_type OWNER TO postgres;
-- ddl-end --

-- object: main_schema.collection_diff | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.collection_diff CASCADE;
CREATE TABLE main_schema.collection_diff(
	id integer NOT NULL,
	collection_id integer NOT NULL,
	user_id integer NOT NULL,
	action varchar(32) NOT NULL,
	data varchar(100),
	creation_time timestamp NOT NULL,
	CONSTRAINT collection_diff_pk PRIMARY KEY (id,collection_id)

);
-- ddl-end --
ALTER TABLE main_schema.collection_diff OWNER TO postgres;
-- ddl-end --

-- object: main_schema.collection_action | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.collection_action CASCADE;
CREATE TABLE main_schema.collection_action(
	name varchar(32) NOT NULL,
	CONSTRAINT collection_action_pk PRIMARY KEY (name)

);
-- ddl-end --
ALTER TABLE main_schema.collection_action OWNER TO postgres;
-- ddl-end --

-- object: main_schema."user" | type: TABLE --
-- DROP TABLE IF EXISTS main_schema."user" CASCADE;
CREATE TABLE main_schema."user"(
	id serial NOT NULL,
	email varchar(128),
	hash varchar(128),
	token varchar(128),
	creation_time timestamp NOT NULL,
	CONSTRAINT user_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE main_schema."user" OWNER TO postgres;
-- ddl-end --

-- object: main_schema.session | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.session CASCADE;
CREATE TABLE main_schema.session(
	user_id integer NOT NULL,
	login_time timestamp NOT NULL,
	logout_time timestamp,
	active bool NOT NULL,
	CONSTRAINT session_pk PRIMARY KEY (user_id,login_time)

);
-- ddl-end --
ALTER TABLE main_schema.session OWNER TO postgres;
-- ddl-end --

-- object: main_schema.deck | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.deck CASCADE;
CREATE TABLE main_schema.deck(
	id serial NOT NULL,
	card_id integer NOT NULL,
	creation_time timestamp NOT NULL,
	CONSTRAINT deck_pk PRIMARY KEY (id,card_id)

);
-- ddl-end --
ALTER TABLE main_schema.deck OWNER TO postgres;
-- ddl-end --

-- object: main_schema.deck_diff | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.deck_diff CASCADE;
CREATE TABLE main_schema.deck_diff(
	id integer NOT NULL,
	deck_id integer NOT NULL,
	user_id integer NOT NULL,
	action varchar(32) NOT NULL,
	data varchar(100),
	creation_time timestamp NOT NULL,
	CONSTRAINT deck_diff_pk PRIMARY KEY (id,deck_id)

);
-- ddl-end --
ALTER TABLE main_schema.deck_diff OWNER TO postgres;
-- ddl-end --

-- object: main_schema.deck_action | type: TABLE --
-- DROP TABLE IF EXISTS main_schema.deck_action CASCADE;
CREATE TABLE main_schema.deck_action(
	name varchar(32) NOT NULL,
	CONSTRAINT deck_action_pk PRIMARY KEY (name)

);
-- ddl-end --
ALTER TABLE main_schema.deck_action OWNER TO postgres;
-- ddl-end --

-- object: collection_deck_id_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.collection DROP CONSTRAINT IF EXISTS collection_deck_id_fkey CASCADE;
ALTER TABLE main_schema.collection ADD CONSTRAINT collection_deck_id_fkey FOREIGN KEY (deck_id)
REFERENCES main_schema.deck (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: action_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.card_diff DROP CONSTRAINT IF EXISTS action_fkey CASCADE;
ALTER TABLE main_schema.card_diff ADD CONSTRAINT action_fkey FOREIGN KEY (action)
REFERENCES main_schema.card_action (name) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: card_id_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.card_diff DROP CONSTRAINT IF EXISTS card_id_fkey CASCADE;
ALTER TABLE main_schema.card_diff ADD CONSTRAINT card_id_fkey FOREIGN KEY (card_id)
REFERENCES main_schema.card (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: user_id_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.card_diff DROP CONSTRAINT IF EXISTS user_id_fkey CASCADE;
ALTER TABLE main_schema.card_diff ADD CONSTRAINT user_id_fkey FOREIGN KEY (user_id)
REFERENCES main_schema."user" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: object_name_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.owner DROP CONSTRAINT IF EXISTS object_name_fkey CASCADE;
ALTER TABLE main_schema.owner ADD CONSTRAINT object_name_fkey FOREIGN KEY (object_type)
REFERENCES main_schema.object_type (name) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: object_id_collection_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.owner DROP CONSTRAINT IF EXISTS object_id_collection_fkey CASCADE;
ALTER TABLE main_schema.owner ADD CONSTRAINT object_id_collection_fkey FOREIGN KEY (object_id)
REFERENCES main_schema.collection (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;
-- ddl-end --

-- object: object_id_card_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.owner DROP CONSTRAINT IF EXISTS object_id_card_fkey CASCADE;
ALTER TABLE main_schema.owner ADD CONSTRAINT object_id_card_fkey FOREIGN KEY (object_id)
REFERENCES main_schema.card (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;
-- ddl-end --

-- object: object_id_deck_id_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.owner DROP CONSTRAINT IF EXISTS object_id_deck_id_fkey CASCADE;
ALTER TABLE main_schema.owner ADD CONSTRAINT object_id_deck_id_fkey FOREIGN KEY (object_id)
REFERENCES main_schema.deck (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;
-- ddl-end --

-- object: collection_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.collection_diff DROP CONSTRAINT IF EXISTS collection_fkey CASCADE;
ALTER TABLE main_schema.collection_diff ADD CONSTRAINT collection_fkey FOREIGN KEY (collection_id)
REFERENCES main_schema.collection (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: action_collection_name_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.collection_diff DROP CONSTRAINT IF EXISTS action_collection_name_fkey CASCADE;
ALTER TABLE main_schema.collection_diff ADD CONSTRAINT action_collection_name_fkey FOREIGN KEY (action)
REFERENCES main_schema.collection_action (name) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: user_id_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.collection_diff DROP CONSTRAINT IF EXISTS user_id_fkey CASCADE;
ALTER TABLE main_schema.collection_diff ADD CONSTRAINT user_id_fkey FOREIGN KEY (user_id)
REFERENCES main_schema."user" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: session_user_id_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.session DROP CONSTRAINT IF EXISTS session_user_id_fkey CASCADE;
ALTER TABLE main_schema.session ADD CONSTRAINT session_user_id_fkey FOREIGN KEY (user_id)
REFERENCES main_schema."user" (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;
-- ddl-end --

-- object: deck_card_if_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.deck DROP CONSTRAINT IF EXISTS deck_card_if_fkey CASCADE;
ALTER TABLE main_schema.deck ADD CONSTRAINT deck_card_if_fkey FOREIGN KEY (card_id)
REFERENCES main_schema.card (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: deck_id_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.deck_diff DROP CONSTRAINT IF EXISTS deck_id_fkey CASCADE;
ALTER TABLE main_schema.deck_diff ADD CONSTRAINT deck_id_fkey FOREIGN KEY (deck_id)
REFERENCES main_schema.deck (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: user_id_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.deck_diff DROP CONSTRAINT IF EXISTS user_id_fkey CASCADE;
ALTER TABLE main_schema.deck_diff ADD CONSTRAINT user_id_fkey FOREIGN KEY (user_id)
REFERENCES main_schema."user" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: action_name_fkey | type: CONSTRAINT --
-- ALTER TABLE main_schema.deck_diff DROP CONSTRAINT IF EXISTS action_name_fkey CASCADE;
ALTER TABLE main_schema.deck_diff ADD CONSTRAINT action_name_fkey FOREIGN KEY (action)
REFERENCES main_schema.deck_action (name) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


