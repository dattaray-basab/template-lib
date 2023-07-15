package mgr

import (
	"fmt"

	"github.com/dattaray-basab/template-lib/common"
	
)

func NewRecipe(absPathToRecipe string, absPathToRecipeContainer string) {
	fmt.Println("NewRecipe")
	recipeDirName := common.GetRecipePath(absPathToRecipe)
	fmt.Println("recipeDirName::", recipeDirName)
}


