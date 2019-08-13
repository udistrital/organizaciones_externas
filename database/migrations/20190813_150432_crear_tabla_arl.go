package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaArl_20190813_150432 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaArl_20190813_150432{}
	m.Created = "20190813_150432"

	migration.Register("CrearTablaArl_20190813_150432", m)
}

// Run the migrations
func (m *CrearTablaArl_20190813_150432) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE organizaciones_externas.arl(id integer NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20),activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_arl PRIMARY KEY (id));")
	m.SQL("ALTER TABLE organizaciones_externas.arl OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaArl_20190813_150432) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS organizaciones_externas.arl")
}
