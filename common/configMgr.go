package common

import (
	"github.com/dattaray-basab/template-lib/globals"

	"github.com/spf13/viper"
)

func GetRecipePath(absPathToRecipe string) string {

	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		viper.Set(globals.RecipePathKey, absPathToRecipe)
		viper.WriteConfig()
	}
	return viper.GetString(globals.RecipePathKey)
}
