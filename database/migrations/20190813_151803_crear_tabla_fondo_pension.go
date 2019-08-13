package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaFondoPension_20190813_151803 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaFondoPension_20190813_151803{}
	m.Created = "20190813_151803"

	migration.Register("CrearTablaFondoPension_20190813_151803", m)
}

// Run the migrations
func (m *CrearTablaFondoPension_20190813_151803) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE organizaciones_externas.fondo_pension(id integer NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_fondo_pension PRIMARY KEY (id));")
	m.SQL("ALTER TABLE organizaciones_externas.fondo_pension OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaFondoPension_20190813_151803) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS organizaciones_externas.fondo_pension")
}
