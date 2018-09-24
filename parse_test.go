package main

import (
	"testing"
)

const (
	addressSQLProd = "./sql/address_prod.sql"
	addressSQLDev  = "./sql/address_dev.sql"
	prod           = "./sql/prod.sql"
	dev            = "./sql/dev.sql"
)

func TestAddress(t *testing.T) {

	diff(addressSQLProd, addressSQLDev)
	diff(prod, dev)

	t.Error("ha")
}
