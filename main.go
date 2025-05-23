package rowscolums

import (
	"database/sql"
)

// automatically recognize columns as a result of the response
func QueryRowsColums(cols []string, rows *sql.Rows) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	colums := make([]interface{}, len(cols))
	columsPointer := make([]interface{}, len(cols))

	for i := range colums {
		columsPointer[i] = &colums[i]
	}

	if err := rows.Scan(columsPointer...); err != nil {
		return nil, err
	}

	for i, colName := range cols {
		val := columsPointer[i].(*interface{})
		m[colName] = *val
	}
	return m, nil
}
