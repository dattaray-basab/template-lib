package mgr

import (
	"fmt"

	"github.com/dattaray-basab/template-lib/common"
)

func NewPhase(
	absPathToRecipe string,
	phaseName string,
	codeBlock string,
	dependentPhaseNames string) {
		
		fmt.Println("NewRecipe")
		recipeDirName := common.GetRecipePath(absPathToRecipe)
		fmt.Println("recipeDirName::", recipeDirName)
}
