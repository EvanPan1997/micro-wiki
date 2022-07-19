package v1

import (
	"micro-wiki/api/v1/example"
	"micro-wiki/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ExampleApiGroup
}

var ApiGroupApp = new(ApiGroup)
