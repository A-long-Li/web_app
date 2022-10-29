/**
 *@filename       settings.go
 *@Description
 *@author          liyajun
 *@create          2022-10-28 23:07
 */

package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(ConfigUnit)

type ConfigUnit struct {
	*AppConfig      `mapstructure:"app"`
	*LogConfig      `mapstructure:"log"`
	*DatabaseConfig `mapstructure:"datasource"`
	*RedisConfig    `mapstructure:"redis"`
}
type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	Port      int    `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID string `mapstructure:"Machine_id"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"username"`
	CharSet      string `mapstructure:"charset"`
	PassWord     string `mapstructure:"password"`
	Location     string `mapstructure:"loc"`
	DriverName   string `mapstructure:"driver_name"`
	DatabaseName string `mapstructure:"database"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	PassWord string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"database_name"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml") // 指定配置文件路径
	err = viper.ReadInConfig()                  // 查找并读取配置文件
	if err != nil {                             // 处理读取配置文件的错误
		panic(fmt.Errorf("viper.ReadInConfig() failed: %s \n", err))
	}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("viper.Unmarshal() failed: %s \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//配置文件修改之后自动修改Conf
		fmt.Printf("config file changed %s \n:", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("viper.Unmarshal() failed: %s \n", err))
		}
	})
	return
}
