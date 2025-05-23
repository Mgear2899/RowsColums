Automatically recognize columns as a result of the response
Example:
```
package main

import (
	"fmt"
	"log"

	rowscol "github.com/mgear2899/rowscolums"
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

Example output:
```
map[age:26 name:Bob]
map[age:18 name:John]
map[age:43 name:Ivan]
```
