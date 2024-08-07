package repository

import (
	"database/sql"
	"fmt"
	"reflect"
)

func scanResult[T any](rows *sql.Rows, target T) (T, error) {
	if reflect.Pointer != reflect.TypeOf(target).Kind() {
		return target, fmt.Errorf("target must be a pointer to a struct")
	}

	fields := reflect.VisibleFields(reflect.TypeOf(target))

	for i := 0; i < len(fields); i++ {
		//reflect.PointerTo(fields[i].Name)
	}

	for rows.Next() {
		rows.Scan()
	}

	return target, nil
}
