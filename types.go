package main

import "fmt"

type database struct {
	Name   string
	Tables []table
}

type table struct {
	Name    string
	Fields  []field
	Keys    []string
	Details string // eg. ENGINE etc.
	Done    bool
	Drop    bool
	Changes []change
}

type field struct {
	Name    string
	Details string
	Done    bool
}

type change struct {
	OriginalField field
	NewField      field
}

func (c *change) SQL(table string) string {
	// Modify
	if c.OriginalField.Name != "" && c.NewField.Name != "" {
		return fmt.Sprintf("ALTER TABLE `%s` MODIFY COLUMN %s %s;",
			table,
			c.OriginalField.Name,
			c.NewField.Details)
	}

	// Create
	if c.OriginalField.Name == "" && c.NewField.Name != "" {
		return fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN %s %s;",
			table,
			c.NewField.Name,
			c.NewField.Details)
	}

	// Remove
	if c.OriginalField.Name != "" && c.NewField.Name == "" {
		return fmt.Sprintf("ALTER TABLE `%s` DROP COLUMN %s;",
			table,
			c.OriginalField.Name,
		)
	}

	return "-- NOT SURE WHAT TO DO"
}
