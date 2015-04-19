## Heartbeat

Dead simple Redis based heartbeat service written in Go which uses less than 1MB of RAM.

By default every second, heartbeat will send a Redis [PUBLISH](http://redis.io/commands/publish) command to a specified channel.

### Development

```bash
git clone git://github.com/catkins/heartbeat $GOPATH/src/catkins/heartbeat
cd $GOPATH/src/catkins/heartbeat

# fetch dependencies
go get github.com/joho/godotenv
go get gopkg.in/redis.v2

# start the service
go run heartbeat.go

# or build it and you're away
go build
./heartbeat
```

### Environment Variables

This app uses environment variables to override configuration through [godotenv](github.com/joho/godotenv).

#### Redis Connection

- `REDIS_URL` (defaults to "localhost:6379")
- `REDIS_PASSWORD` (defaults to "")
- `REDIS_DATABASE` which redis database to connect to (defaults to "0")

#### App settings

- `HEARTBEAT_CHANNEL` channel to publish heartbeats to (defaults to "heartbeat")
- `HEARTBEAT_MESSAGE` message to publish as a heartbeats (if blank it sends the current time as a unix timestamp)
- `HEARTBEAT_INTERVAL` in seconds (defaults to "1")

### TODO

- [ ] Write tests
- [ ] Support dynamic messages using Go Templates
