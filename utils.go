package main

import (
	"io/ioutil"
	"strings"
)

const (
	typeUnknown = iota
	typeCommentContained
	typeCommentOpen
	typeCommentClose
	typeCreateTable
	typeTableColumn
	typePrimaryKey
	typeKey
	typeCloseTable
)

// getEntity gets the string inbetween ``
// eg. "`abc` something else" returns "abc"
func getEntity(line string) string {
	c := strings.Count(line, "`")
	if c != 2 {
		return ""
	}

	s := strings.Index(line, "`")
	e := strings.LastIndex(line, "`")
	return line[s+1 : e]
}

// getColumnDetails returns the details for the column (without the name)
func getColumnDetails(line string) string {
	c := strings.Count(line, "`")
	if c != 2 {
		return ""
	}

	e := strings.LastIndex(line, "`")
	return strings.TrimSpace(line[e+1:])
}

// getType returns the lines type
func getType(line string) int {
	line = strings.TrimSpace(line)

	// Self enclosed comments
	if len(line) > 2 && line[0:2] == "/*" && strings.Contains(line, "*/") {
		if strings.Count(line, "/*") == strings.Count(line, "*/") {
			return typeCommentContained
		}
		return typeCommentOpen
	}

	// Comment open
	if len(line) > 2 && line[0:2] == "/*" {
		return typeCommentOpen
	}

	// Comment close
	if strings.Contains(line, "*/") {
		return typeCommentClose
	}

	// Create table
	if strings.Contains(line, "CREATE TABLE") {
		return typeCreateTable
	}

	// Going to assume that a table column will begin with "`"
	if len(line) > 4 && line[0:1] == "`" {
		return typeTableColumn
	}

	// Primary key
	if len(line) > 13 && strings.Index(line, "PRIMARY KEY") == 0 {
		return typePrimaryKey
	}

	// Key
	if len(line) > 4 && strings.Index(line, "KEY") == 0 {
		return typeKey
	}

	// Close table
	if strings.Index(line, ")") == 0 && strings.LastIndex(line, ";") == (len(line)-1) {
		return typeCloseTable
	}

	// Unknown
	// fmt.Println("Not sure what this line type is: ", line)
	return typeUnknown
}

// cleanLine trims the line, removes commas
func cleanLine(line string) string {
	line = strings.TrimSpace(line)
	line = strings.TrimRight(line, ",")
	return strings.TrimSpace(line)
}

func loadSQL(originalSQLFile, newSQLFile string) (string, string) {
	original, err := ioutil.ReadFile(originalSQLFile)
	if err != nil {
		panic(err)
	}
	new, err := ioutil.ReadFile(newSQLFile)
	if err != nil {
		panic(err)
	}

	return string(original), string(new)
}
