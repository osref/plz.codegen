package cpSimpleValueToJson

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpSimpleValueToJson"] = F
}

// F the function definition
var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		//"cpSimpleValue": F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	dst.Write{{ .ST|opFuncName }}(src)
}
`,
	FuncMap: map[string]interface{}{
		"opFuncName": GenOpFuncName,
	},
}

// GenOpFuncName get corresponding read/write operation name for this type
func GenOpFuncName(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Int:
		return "Int"
	case reflect.Int8:
		return "Int8"
	case reflect.String:
		return "String"
	}
	return ""
}