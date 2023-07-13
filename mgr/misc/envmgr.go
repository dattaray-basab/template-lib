package misc

import (
	"fmt"
	"github.com/dattaray-basab/template-lib/mgr/globals"
)

func SetRecipePath(path string) {
	fmt.Println("misc/envmgr.go::SetRecipePath: ", path)
	fmt.Println("misc/envmgr.go::SetRecipePath: ", globals.RecipePathKey)
}
