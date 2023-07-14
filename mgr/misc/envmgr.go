package misc

import (
	"fmt"
	"strings"

	"github.com/dattaray-basab/template-lib/mgr/globals"
)

func SetRecipePath(dirPath string) {
	fmt.Println("misc/envmgr.go::SetRecipePath: path::", dirPath)
	fmt.Println("misc/envmgr.go::SetRecipePath: globals.RecipePathKey::", globals.RecipePathKey)

	// split a string into a slice of strings
	dirPathParts := strings.Split(dirPath, "/")
	for i, dirPathPart := range dirPathParts {
		if dirPathPart == globals.RecipeDirName {
			fmt.Println("misc/envmgr.go::SetRecipePath: index, dirPathPart::", i, dirPathPart)
		}
	}

}
