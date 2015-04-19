package main

import (
	"fmt"
	"time"

	"github.com/catkins/heartbeat/config"
	"gopkg.in/redis.v2"
)

var appConfig config.Configuration

func main() {
	appConfig = config.Load()

	interval := time.Duration(appConfig.HeartbeatInterval) * time.Second
	ticker := time.NewTicker(interval)

	fmt.Printf("Connecting to redis at redis://%s/%d\n",
		appConfig.RedisAddress, appConfig.RedisDatabase)

	options := appConfig.RedisOptions()
	client := redis.NewTCPClient(&options)
	defer client.Close()

	fmt.Printf("Starting heartbeat on channel \"%s\" every %d seconds\n", appConfig.HeartbeatChannel, appConfig.HeartbeatInterval)

	for {
		tick := <-ticker.C

		go func() {
			var message string
			if len(appConfig.HeartbeatMessage) > 0 {
				message = appConfig.HeartbeatMessage
			} else {
				message = fmt.Sprintf("%d", tick.Unix())
			}

			_, err := client.Publish(appConfig.HeartbeatChannel, message).Result()

			if err != nil {
				fmt.Println(time.Now().String(), err.Error())
			}
		}()

	}
}
