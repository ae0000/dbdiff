package main

import (
	"strings"
)

// parseSQL converts the raw sql into a database type
func parseSQL(sql, name string) database {
	var t table
	d := database{Name: name}
	commenting := false

	// Split lines
	s := strings.Split(sql, "\n")

	for _, s := range s {
		s = cleanLine(s)

		lineType := getType(s)
		if lineType == typeCommentContained {
			commenting = false
		}

		// If we are in commenting mode... dont do anything else until we get a
		// close comment back
		if commenting && lineType != typeCloseTable {
			continue
		}

		// What is it
		switch lineType {
		case typeCommentOpen:
			commenting = true
		case typeCommentClose:
			commenting = false
		case typeUnknown, typeCommentContained:
			// Do nothing
		case typeCreateTable:
			t = table{
				Name: getEntity(s),
			}
		case typeTableColumn:
			f := field{
				Name:    getEntity(s),
				Details: getColumnDetails(s),
			}
			t.Fields = append(t.Fields, f)
		case typePrimaryKey, typeKey:
			t.Keys = append(t.Keys, s)
		case typeCloseTable:
			t.Details = s
			d.Tables = append(d.Tables, t)
		}
	}

	return d
}
