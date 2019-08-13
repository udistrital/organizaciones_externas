package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaEntidadAseguradora_20190813_151743 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaEntidadAseguradora_20190813_151743{}
	m.Created = "20190813_151743"

	migration.Register("CrearTablaEntidadAseguradora_20190813_151743", m)
}

// Run the migrations
func (m *CrearTablaEntidadAseguradora_20190813_151743) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE organizaciones_externas.entidad_aseguradora(id integer NOT NULL, codigo numeric(4,0), nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_entidad_aseguradora PRIMARY KEY (id));")
	m.SQL("ALTER TABLE organizaciones_externas.entidad_aseguradora OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaEntidadAseguradora_20190813_151743) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS organizaciones_externas.entidad_aseguradora")
	
}
