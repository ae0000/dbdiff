package main

import "testing"

func TestGetEntity(t *testing.T) {
	e := getEntity("CREATE TABLE `Addresses` (")
	if e != "Addresses" {
		t.Errorf("Got %s, wanted Addresses", e)
	}

	e = getEntity("CREATE TABLE Addresses` (")
	if e != "" {
		t.Errorf("Got %s, wanted <empty>", e)
	}

	e = getEntity("")
	if e != "" {
		t.Errorf("Got %s, wanted <empty>", e)
	}

	e = getEntity("``````")
	if e != "" {
		t.Errorf("Got %s, wanted <empty>", e)
	}

	e = getEntity("``")
	if e != "" {
		t.Errorf("Got %s, wanted <empty>", e)
	}

	e = getEntity("`heading` varchar(250) NOT NULL,")
	if e != "heading" {
		t.Errorf("Got %s, wanted heading", e)
	}
}

func TestGetType(t *testing.T) {
	e := getType("CREATE TABLE `Addresses` (")
	if e != typeCreateTable {
		t.Error("Wrong type")
	}

	e = getType("TABLE `Addresses` (")
	if e != typeUnknown {
		t.Error("Wrong type")
	}

	e = getType("/*!40101 SET @saved_cs_client     = @@character_set_client */;")
	if e != typeCommentContained {
		t.Error("Wrong type")
	}

	e = getType("`heading` varchar(250) NOT NULL,")
	if e != typeTableColumn {
		t.Error("Wrong type")
	}

	e = getType("PRIMARY KEY (`a_id`)")
	if e != typePrimaryKey {
		t.Error("Wrong type")
	}

	e = getType("KEY `ID` (`id`),")
	if e != typeKey {
		t.Error("Wrong type")
	}

	e = getType(") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='Hold a list of all the addresses.';")
	if e != typeCloseTable {
		t.Error("Wrong type")
	}
}

func TestCleanLine(t *testing.T) {
	e := cleanLine("  abc   ,")
	if e != "abc" {
		t.Error("cleanLine did not clean")
	}

	e = cleanLine("  abc   , ")
	if e != "abc" {
		t.Error("cleanLine did not clean")
	}

	e = cleanLine("     , ")
	if e != "" {
		t.Error("cleanLine did not clean")
	}

	e = cleanLine(" abc,,,, ")
	if e != "abc" {
		t.Error("cleanLine did not clean: ", e)
	}
}

func TestGetColumnDetails(t *testing.T) {
	e := getColumnDetails("`heading` varchar(250) NOT NULL")
	if e != "varchar(250) NOT NULL" {
		t.Error("Wrong type")
	}

	e = getColumnDetails("varchar(250) NOT NULL")
	if e != "" {
		t.Error("Wrong type")
	}
}
