package mgr

import (
	"fmt"
	"path/filepath"

	"github.com/dattaray-basab/template-lib/common"
)

func NewRecipe(absPathToSource string, absPathToRecipeParent string) {
	fmt.Println()
	fmt.Println("<<< NewRecipe/n")
	fmt.Println("absPathToRecipeParent::", absPathToRecipeParent)
	fmt.Println("absPathToSource::", absPathToSource)
	absPathToRecipe := filepath.Join(absPathToRecipeParent, "__recipe")
	pathToRecipe := common.GetRecipePath(absPathToRecipe)
	fmt.Println("pathToRecipe::", pathToRecipe)
	fmt.Println(">>> NewRecipe")
	fmt.Println()
}
