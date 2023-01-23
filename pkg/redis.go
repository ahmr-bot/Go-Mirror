package pkg

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

type Config struct {
	RedisAddr string `json:"redis_addr"`
	RedisPass string `json:"redis_pass"`
	RedisDB   int    `json:"redis_db"`
}

func (c Config) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

var config Config
var client *redis.Client

func init() {
	log.Println("初始化中...")
	//  读取配置文件
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("配置文件读取失败", err)
	}
	defer file.Close()
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal(err)
	}

	client = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPass,
		DB:       config.RedisDB,
	})
}
func RedisClient() *redis.Client {
	return client
}
