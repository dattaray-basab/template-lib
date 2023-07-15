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
		recipeDirName := common.GetRecipePath(templaterDirpath)
		if templaterDirpath != recipeDirName || !strings.HasPrefix(templaterDirpath, recipeDirName) {
			_ = fmt.Errorf("Recipe path %s is not a subdirectory of %s", templaterDirpath, recipeDirName)
		}
		fmt.Println("recipeDirName::", recipeDirName)
}
