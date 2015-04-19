package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/catkins/heartbeat/config"
	"gopkg.in/redis.v2"
)

var appConfig config.Configuration

func init() {
	appConfig = config.Load()
}

func main() {
	client := connectToRedis()
	defer client.Close()

	startInterval(func(tick time.Time) {
		message, err := buildMessage(tick)
		if err == nil {
			publish(client, message)
		} else {
			fmt.Println(time.Now().String(), err.Error())
		}
	})
}

func connectToRedis() *redis.Client {
	fmt.Printf("Connecting to redis at redis://%s/%d\n",
		appConfig.RedisAddress,
		appConfig.RedisDatabase)

	options := appConfig.RedisOptions()
	return redis.NewTCPClient(&options)
}

func startInterval(callback func(time.Time)) {
	fmt.Printf("Starting heartbeat on channel \"%s\" every %d seconds\n",
		appConfig.HeartbeatChannel,
		appConfig.HeartbeatInterval)

	interval := time.Duration(appConfig.HeartbeatInterval) * time.Second
	ticker := time.NewTicker(interval)

	for {
		tick := <-ticker.C
		go callback(tick)
	}
}

func publish(client *redis.Client, message string) {
	_, err := client.Publish(appConfig.HeartbeatChannel, message).Result()

	if err != nil {
		fmt.Println(time.Now().String(), err.Error())
	}
}

func buildMessage(tick time.Time) (string, error) {
	var context struct {
		Timestamp string
	}

	context.Timestamp = fmt.Sprintf("%d", tick.Unix())
	var buffer bytes.Buffer
	err := appConfig.HeartbeatTemplate.Execute(&buffer, context)

	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
