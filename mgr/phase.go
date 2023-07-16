package mgr

import (
	"fmt"
	"strings"

	"github.com/dattaray-basab/template-lib/common"
)

func NewPhase(
	templaterDirpath string,
	phaseName string,
	codeBlock string,
	dependentPhaseNames string) {
	fmt.Println("NewPhase")
	recipePath := common.GetRecipePath(templaterDirpath)
	fmt.Println("templaterDirpath::", templaterDirpath)
	fmt.Println("recipePath::", recipePath)
	if templaterDirpath != recipePath && !strings.HasPrefix(templaterDirpath, recipePath) {
		_ = fmt.Errorf("Recipe path %s is not a subdirectory of %s", templaterDirpath, recipePath)
	}
}
