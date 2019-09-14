package parser

import (
	"bytes"
	"html/template"
	"reflect"
	"strings"
	"time"
)

func executeTemplate(text string, state map[string]interface{}) (string, error) {
	t := template.New("microspector")
	_, err := t.Parse(text)

	if err != nil {
		return "", nil
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, state); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// IsTrue reports whether the value is 'true', in the sense of not the zero of its type,
// and whether the value has a meaningful truth value. This is the definition of
// truth used by if and other such actions.
func IsTrue(val interface{}) bool {
	return isTrue(reflect.ValueOf(val))
}

func isTrue(val reflect.Value) (truth bool) {
	if !val.IsValid() {
		// Something like var x interface{}, never set. It's a form of nil.
		return false
	}
	switch val.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		truth = val.Len() > 0
	case reflect.Bool:
		truth = val.Bool()
	case reflect.Complex64, reflect.Complex128:
		truth = val.Complex() != 0
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Interface:
		truth = !val.IsNil()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		truth = val.Int() != 0
	case reflect.Float32, reflect.Float64:
		truth = val.Float() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		truth = val.Uint() != 0
	case reflect.Struct:
		truth = true // Struct values are always true.
	default:
		return false
	}
	return truth
}

var zero reflect.Value

func toVariableName(str string) string {
	segments := strings.Split(str, "-")
	for key, value := range segments {
		segments[key] = strings.Title(value)
	}

	return strings.Join(segments, "")
}
