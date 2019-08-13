package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchema_20190813_150001 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchema_20190813_150001{}
	m.Created = "20190813_150001"

	migration.Register("CrearSchema_20190813_150001", m)
}

// Run the migrations
func (m *CrearSchema_20190813_150001) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA organizaciones_externas;")
	m.SQL("ALTER SCHEMA organizaciones_externas OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,organizaciones_externas;")

}

// Reverse the migrations
func (m *CrearSchema_20190813_150001) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS organizaciones_externas");
}
