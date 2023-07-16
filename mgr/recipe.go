package mgr

import (
	"fmt"

	"github.com/dattaray-basab/template-lib/common"
)

func NewRecipe(absPathToRecipe string, absPathToRecipeContainer string) {
	fmt.Println("NewRecipe")
	RecipeDirectory := common.GetRecipePath(absPathToRecipe)
	fmt.Println("RecipeDirectory::", RecipeDirectory)
}
