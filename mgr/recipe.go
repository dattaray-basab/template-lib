package mgr

import (
	"fmt"

	"github.com/dattaray-basab/template-lib/common"
	"github.com/dattaray-basab/template-lib/globals"
	"github.com/spf13/viper"
)

func NewRecipe(absPathToRecipe string, absPathToRecipeContainer string) {
	fmt.Println("NewRecipe")
	recipeDirName := initConfig(absPathToRecipe)
	fmt.Println("recipeDirName::", recipeDirName)
}


func initConfig(absPathToRecipe string) string {

	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		// dirPath, _ := os.Getwd()
		foundPath, err := common.GetNamedPath(absPathToRecipe, globals.RecipeDirName)
		if err != nil {
			fmt.Println(err)
		}
		viper.Set(globals.RecipeDirName, foundPath)
		viper.WriteConfig()
	}
	return viper.GetString(globals.RecipeDirName)
}

