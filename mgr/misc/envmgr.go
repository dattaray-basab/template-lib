package misc

import (
	"fmt"

	"github.com/dattaray-basab/template-lib/mgr/globals"
		"github.com/dattaray-basab/template-lib/common"
)

func SetRecipePath(dirPath string) {
	fmt.Println("misc/envmgr.go::SetRecipePath: path::", dirPath)
	fmt.Println("misc/envmgr.go::SetRecipePath: globals.RecipePathKey::", globals.RecipePathKey)
	foundPath, err := common.GetNamedPath(dirPath, "__DEV_3")
	if err != nil {
		fmt.Println("misc/envmgr.go::SetRecipePath: err::", err)
		return 
	}
	fmt.Println("misc/envmgr.go::SetRecipePath: foundPath::", foundPath)
}

