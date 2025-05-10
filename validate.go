package openf1go

import (
	"reflect"
	"strconv"
	"time"
)

// buildArgs takes a struct as input and converts its fields into a slice of Arg objects.
// Each Arg object contains a key (from the struct's `json` tag) and a value (converted to a string).
func buildArgs(i interface{}) []Arg {
	// Get the type and value of the input struct using reflection.
	st := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	args := []Arg{}

	// Iterate over all fields of the struct.
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		// Check if the field is not zero (i.e., it has a value).
		if !v.FieldByName(field.Name).IsZero() {
			// Get the field's value using reflection.
			r := reflect.Indirect(v).FieldByName(field.Name)
			fieldvalue := r.Interface()

			// Handle the field's value based on its type.
			switch v := fieldvalue.(type) {
			case int:
				// Convert int to string and append to args with the key from the `json` tag.
				args = append(args, Arg{Key: field.Tag.Get("json"), Value: strconv.Itoa(v)})
			case time.Time:
				// Format time.Time to RFC3339 string and append to args with the key from the `json` tag.
				args = append(args, Arg{Key: field.Tag.Get("json"), Value: v.Format(time.RFC3339)})
			case string:
				// Append string value to args with the key from the `json` tag.
				args = append(args, Arg{Key: field.Tag.Get("json"), Value: v})
			}
		}
	}

	// Return the constructed slice of Arg objects.
	return args
}
