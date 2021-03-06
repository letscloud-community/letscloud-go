package letscloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func processResponse(b []byte, target interface{}) error {
	return json.Unmarshal(b, &target)
}

func splitText(s string) []string {
	return strings.Split(s, ",")
}

func contains(arr []string, v string) bool {
	for _, s := range arr {
		if s == v {
			return true
		}
	}

	return false
}

func validateStruct(val interface{}) error {
	v := reflect.ValueOf(val)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		splittedTags := splitText(jsonTag)

		if v.Field(i).IsZero() && !contains(splittedTags, "omitempty") {
			return errors.New(fmt.Sprintf("field %s [type %s with tag %s] contains empty value",
				t.Field(i).Name, v.Field(i).Kind(), t.Field(i).Tag))
		}
	}

	return nil
}
