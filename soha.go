package soha

import (
	"html/template"

	"github.com/flysnow-org/soha/deps"

	"github.com/flysnow-org/soha/internal"

	// Init the namespaces
	_ "github.com/flysnow-org/soha/cast"
	_ "github.com/flysnow-org/soha/collections"
	_ "github.com/flysnow-org/soha/crypto"
	_ "github.com/flysnow-org/soha/encoding"
	_ "github.com/flysnow-org/soha/lang"
	_ "github.com/flysnow-org/soha/math"
	_ "github.com/flysnow-org/soha/path"
	_ "github.com/flysnow-org/soha/reflect"
	_ "github.com/flysnow-org/soha/safe"
	_ "github.com/flysnow-org/soha/strings"
	_ "github.com/flysnow-org/soha/time"
	_ "github.com/flysnow-org/soha/transform"
)

func CreateFuncMap() map[string]interface{} {
	funcMap := template.FuncMap{}
	d:=&deps.Deps{}

	tfns := internal.TemplateFuncsNamespaces{}
	// Merge the namespace funcs
	for _, nsf := range internal.TemplateFuncsNamespaceRegistry {
		ns := nsf(d)
		tfns = append(tfns,ns)
		if _, exists := funcMap[ns.Name]; exists {
			panic(ns.Name + " is a duplicate template func")
		}
		funcMap[ns.Name] = ns.Context
		for _, mm := range ns.MethodMappings {
			for _, alias := range mm.Aliases {
				if _, exists := funcMap[alias]; exists {
					panic(alias + " is a duplicate template func")
				}
				funcMap[alias] = mm.Method
			}

		}

	}
	return funcMap
}
