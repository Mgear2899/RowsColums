Automatically recognize columns as a result of the response
Example:
```
pacage main

import rowscol "github.com/Mgear2899/RowsColums"

func main() {
	row, err := db.Query(`select name, age from table`)
  if err != nil {
    log.Fatal(err)
  }

	for row.Next() {
		cols, err := rows.Columns()
     if err != nil {
        log.Fatal(err)
      }
			res := rowscol.QueryRowsColums(cols, rows)
			fmt.Println(res)
	}
}
```
