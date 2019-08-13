package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaBanco_20190813_151710 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaBanco_20190813_151710{}
	m.Created = "20190813_151710"

	migration.Register("CrearTablaBanco_20190813_151710", m)
}

// Run the migrations
func (m *CrearTablaBanco_20190813_151710) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE organizaciones_externas.banco(id integer NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), denominacion_banco character varying(200), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_banco PRIMARY KEY (id));")
	m.SQL("ALTER TABLE organizaciones_externas.banco OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaBanco_20190813_151710) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS organizaciones_externas.banco")

}
