package main

import "fmt"

func diff(original, new string) {
	o, n := loadSQL(original, new)

	do := parseSQL(o, "original")
	dn := parseSQL(n, "new")

	// Start with the original, look for changed fields and keys
	for h, t := range do.Tables {

		// Go through the fields
		for i, f := range t.Fields {
			// Find in the new db
			foundTable := false
			for j, tn := range dn.Tables {
				if tn.Name == t.Name {
					// Find the field
					// fmt.Println("FOund ", t.Name)
					foundTable = true
					for k, fn := range tn.Fields {
						if fn.Name == f.Name {
							do.Tables[h].Done = true
							do.Tables[h].Fields[i].Done = true
							dn.Tables[j].Done = true
							dn.Tables[j].Fields[k].Done = true

							if fn.Details != f.Details {
								c := change{
									OriginalField: f,
									NewField:      fn,
								}

								do.Tables[h].Changes = append(do.Tables[h].Changes, c)
							}
						}
					}

				}
			}
			if !foundTable {
				do.Tables[h].Drop = true
			}
		}

		// We have gone through the fields once.. now we go through again
		// looking for fields that were not in the new db
		if !do.Tables[h].Drop {
			for i, f := range t.Fields {
				if !f.Done {
					c := change{
						OriginalField: f,
					}

					do.Tables[h].Changes = append(do.Tables[h].Changes, c)
					do.Tables[h].Fields[i].Done = true
				}
			}
		}

		// Now check the new db for fields that have not been taken care of
		for j, tn := range dn.Tables {
			if tn.Name == t.Name {
				// Find the field
				for k, fn := range tn.Fields {
					if !fn.Done {
						c := change{
							NewField: fn,
						}
						do.Tables[h].Changes = append(do.Tables[h].Changes, c)
						dn.Tables[j].Fields[k].Done = true
					}
				}
			}
		}

	}

	// Show changes
	if true {
		for _, t := range do.Tables {
			if t.Drop {
				fmt.Println("\n-- :::::::::: Changes for ", t.Name)
				fmt.Println("-- .....................................................")
				fmt.Println("-- Change   : DROP ", t.Name)
				fmt.Printf("DROP TABLE %s;\n", t.Name)

			} else {

				if len(t.Changes) > 0 {
					fmt.Println("\n-- :::::::::: Changes for ", t.Name)
					for _, c := range t.Changes {
						fmt.Println("-- .....................................................")
						if len(c.OriginalField.Name) > 0 {
							fmt.Println("-- Column   : ", c.OriginalField.Name)
						} else {
							fmt.Println("-- Column   : ", c.NewField.Name)
						}
						fmt.Println("-- Original : ", c.OriginalField.Details)
						fmt.Println("-- New      : ", c.NewField.Details)
						fmt.Println(c.SQL(t.Name))
					}
				}
			}
		}
	}
}
