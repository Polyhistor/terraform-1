package types

import (
	"reflect"

	"github.com/zclconf/go-cty/cty"
)

// TypeType is a capsule type used to represent a cty.Type as a cty.Value. This
// is used by the `type()` console function to smuggle cty.Type values to the
// REPL session, where it can be displayed to the user directly.
//
// We implement the raw equality operation to allow comparison of two types. No
// other operation is permitted.
var TypeType = cty.CapsuleWithOps(
	"type",
	reflect.TypeOf(cty.Type{}),
	&cty.CapsuleOps{
		RawEquals: func(a, b interface{}) bool {
			aType := *(a.(*cty.Type))
			bType := *(b.(*cty.Type))
			return aType.Equals(bType)
		},
	},
)
