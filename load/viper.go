package load

import (
	"fisco/global"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	VP = &viper.Viper{}
)

func init() {
	VP = viper.New()
	VP.SetConfigName("configuration")
	VP.AddConfigPath(".")
	VP.ReadInConfig()

	//watch viper
	VP.WatchConfig()

	//reload viper
	VP.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := VP.Unmarshal(&global.GlobalConfig); err != nil {
			fmt.Println(err)
		}
	})

	//global load viper
	if err := VP.Unmarshal(&global.GlobalConfig); err != nil {
		fmt.Println(err)
	}
}
