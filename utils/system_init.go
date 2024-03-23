package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("config app: ", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))

}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{

		Addr:         viper.GetString("redis.address"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.pollSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})

	pong, err := Red.Ping().Result()

	if err != nil {
		fmt.Println("Init Redis err: ", err)
	} else {
		fmt.Println("Redis has been inited successful", pong)
	}
}

func InitMysql() {
	// 自定义日志模版，打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
}
