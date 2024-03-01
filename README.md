Automatically recognize columns as a result of the response
Example:
```
package main

import (
	"fmt"
	"log"

	rowscol "github.com/Mgear2899/RowsColums"
)

func main() {
	db := connDB()
	rows, err := db.Query(`select name, age from table`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		cols, err := rows.Columns()
		if err != nil {
			log.Fatal(err)
		}
		res := rowscol.QueryRowsColums(cols, rows)
		fmt.Println(res)
	}
}
```
