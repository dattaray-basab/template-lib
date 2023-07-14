package misc

import (
	"fmt"
	"strings"

	"github.com/dattaray-basab/template-lib/mgr/globals"
)

func SetRecipePath(dirPath string) {
	fmt.Println("misc/envmgr.go::SetRecipePath: path::", dirPath)
	fmt.Println("misc/envmgr.go::SetRecipePath: globals.RecipePathKey::", globals.RecipePathKey)
	foundPath, err := GetNamedPath(dirPath, "___DEV_3")
	if err != nil {
		fmt.Println("misc/envmgr.go::SetRecipePath: err::", err)
		return 
	}
	fmt.Println("misc/envmgr.go::SetRecipePath: foundPath::", foundPath)
}

func GetNamedPath(dirPath string, searchPathFragment string) (string, error)  {
	dirPathParts := strings.Split(dirPath, "/")
	newPath := ""
	for i, dirPathPart := range dirPathParts {
		newPath = newPath + dirPathPart + "/"
		if dirPathPart == searchPathFragment {
			fmt.Println("misc/envmgr.go::SetRecipePath: index, dirPathPart::", i, dirPathPart)
			return newPath, nil
		}
	}
	return "", fmt.Errorf("misc/envmgr.go::SetRecipePath: %s not found in %s", searchPathFragment, dirPath)
}