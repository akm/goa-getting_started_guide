// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cellar": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=goa-getting_started_guide/design
// --out=$(GOPATH)/src/goa-getting_started_guide
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// BottleHref returns the resource href.
func BottleHref(bottleID interface{}) string {
	parambottleID := strings.TrimLeftFunc(fmt.Sprintf("%v", bottleID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/bottles/%v", parambottleID)
}
