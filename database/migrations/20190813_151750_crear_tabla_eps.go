package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaEps_20190813_151750 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaEps_20190813_151750{}
	m.Created = "20190813_151750"

	migration.Register("CrearTablaEps_20190813_151750", m)
}

// Run the migrations
func (m *CrearTablaEps_20190813_151750) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE organizaciones_externas.eps(id integer NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_eps PRIMARY KEY (id));")
	m.SQL("ALTER TABLE organizaciones_externas.eps OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaEps_20190813_151750) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS organizaciones_externas.eps")

}
