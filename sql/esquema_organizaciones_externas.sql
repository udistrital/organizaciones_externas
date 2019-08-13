CREATE SCHEMA organizaciones_externas;

-- Creación de secuencia y tabla arl
CREATE SEQUENCE organizaciones_externas.arl_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE organizaciones_externas.arl(
	id integer NOT NULL DEFAULT nextval('organizaciones_externas.arl_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_arl PRIMARY KEY (id)

);

-- Creación de secuencia y tabla banco
CREATE SEQUENCE organizaciones_externas.banco_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE organizaciones_externas.banco(
	id integer NOT NULL DEFAULT nextval('organizaciones_externas.banco_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	denominacion_banco character varying(200),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_banco PRIMARY KEY (id)

);

-- Creación de secuencia y tabla caja_compensacion
CREATE SEQUENCE organizaciones_externas.caja_compensacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE organizaciones_externas.caja_compensacion(
	id integer NOT NULL DEFAULT nextval('organizaciones_externas.caja_compensacion_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),	
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_caja_compensacion PRIMARY KEY (id)

);

-- Creación de secuencia y tabla entidad_aseguradora
CREATE SEQUENCE organizaciones_externas.entidad_aseguradora_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE organizaciones_externas.entidad_aseguradora(
	id integer NOT NULL DEFAULT nextval('organizaciones_externas.entidad_aseguradora_id_seq'::regclass),
	codigo numeric(4,0),
 	nombre character varying(100) NOT NULL,
	descripcion character varying(100),	
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_entidad_aseguradora PRIMARY KEY (id)

);

-- Creación de secuencia y tabla eps
CREATE SEQUENCE organizaciones_externas.eps_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE organizaciones_externas.eps(
	id integer NOT NULL DEFAULT nextval('organizaciones_externas.eps_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_eps PRIMARY KEY (id)

);

-- Creación de secuencia y tabla fondo_pension
CREATE SEQUENCE organizaciones_externas.fondo_pension_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE organizaciones_externas.fondo_pension(
	id integer NOT NULL DEFAULT nextval('organizaciones_externas.fondo_pension_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_fondo_pension PRIMARY KEY (id)

);

-- Creación de secuencia y tabla sucursal
CREATE SEQUENCE organizaciones_externas.sucursal_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE organizaciones_externas.sucursal(
	id integer NOT NULL DEFAULT nextval('organizaciones_externas.sucursal_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	encargado character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_sucursal PRIMARY KEY (id)

);
