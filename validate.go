package openf1go

import (
	"reflect"
	"strconv"
	"time"
)

func buildArgs(i interface{}) []Arg {
	st := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	args := []Arg{}

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		if !v.FieldByName(field.Name).IsZero() {
			r := reflect.Indirect(v).FieldByName(field.Name)
			fieldvalue := r.Interface()

			switch v := fieldvalue.(type) {
			case int:
				args = append(args, Arg{Key: field.Tag.Get("json"), Value: strconv.Itoa(v)})
			case time.Time:
				args = append(args, Arg{Key: field.Tag.Get("json"), Value: v.Format(time.RFC3339)})
			case string:
				args = append(args, Arg{Key: field.Tag.Get("json"), Value: v})
			}

		}
	}

	return args
}
