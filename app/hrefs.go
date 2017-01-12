//************************************************************************//
// API "mVisa": Application Resource Href Factories
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=ChamaconektVisa/design
// --out=$(GOPATH)/src/ChamaconektVisa
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strings"
)

// DepositsHref returns the resource href.
func DepositsHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/deposits/%v", paramid)
}
