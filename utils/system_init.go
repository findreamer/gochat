package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
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

const (
	PublishKey = "webscket"
)

// 发布消息到redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error

	err = Red.Publish(ctx, channel, msg).Err()
	return err
}

// 订阅 redis 消息
func Subscribe(ctx context.Context, channel string) (string, error) {

	sub := Red.Subscribe(ctx, channel)

	fmt.Println("Subscribe ==> ", sub)

	msg, err := sub.ReceiveMessage(ctx)

	return msg.Payload, err

}
