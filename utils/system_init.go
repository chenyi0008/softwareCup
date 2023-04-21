package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
)

var Db *xorm.Engine

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))
	fmt.Println("config path:", viper.Get("file"))

}

func InitMySQL() *xorm.Engine {
	fmt.Println(viper.GetString("mysql.dns"))
	var err error
	Db, err = xorm.NewEngine("mysql", viper.GetString("mysql.dns"))
	if err != nil {
		fmt.Println(err)
	}
	if Db != nil {
		fmt.Println("MySQL inited ~~~~~~~~~~~~")
	}
	return Db
}
