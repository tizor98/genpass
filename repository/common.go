package repository

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

// ALERT!! This is only for fun.
// Do not use [reflect] package to scan a sql result since you should know both, the struct fields and return columns.
// Reflection is useful when use for dynamic or not known information. But commonly is slower than regular code.
func scanOneStruct[T *Q, Q any](rows *sql.Rows, target T) {
	if reflect.Pointer != reflect.TypeOf(target).Kind() {
		log.Fatal(fmt.Errorf("target must be a pointer to a struct"))
	}

	elem := reflect.ValueOf(target).Elem()

	fields := reflect.VisibleFields(reflect.TypeOf(*target))
	pointers := make([]any, len(fields))
	for i := 0; i < len(fields); i++ {
		field := elem.FieldByName(fields[i].Name)
		pointers[i] = field.Addr().Interface()
	}

	if !rows.Next() {
		return
	}

	err := rows.Scan(pointers...)
	if err != nil {
		log.Fatal(err)
	}
}
