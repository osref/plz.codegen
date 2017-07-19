package maxStructByField

import (
	"github.com/v2pro/plz/util"
	"github.com/v2pro/wombat/fp/cmpStructByField"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	util.GenMaxStructByField = genF
}

// F the function definition
var F = &gen.FuncTemplate{
	TemplateName: "maxStructByField",
	Dependencies: []*gen.FuncTemplate{cmpStructByField.F},
	TemplateParams: map[string]string{
		"T": "the struct type to max",
		"F": "the field name of T",
	},
	FuncName: `Max_{{ .T|name }}_by_{{ .F }}`,
	Source: `
{{ $compare := gen "cmpStructByField" "T" .T "F" .F }}
func Exported_{{ .funcName }}(objs []interface{}) interface{} {
	currentMaxObj := objs[0]
	for i := 1; i < len(objs); i++ {
		currentMax := {{ cast "currentMaxObj" .T }}
		elem := {{ cast "objs[i]" .T }}
		if {{ $compare }}(elem, currentMax) > 0 {
			currentMaxObj = objs[i]
		}
	}
	return currentMaxObj
}`,
}

func genF(typ reflect.Type, fieldName string) func([]interface{}) interface{} {
	funcObj := gen.Expand(F, "T", typ, "F", fieldName)
	return funcObj.(func([]interface{}) interface{})
}
