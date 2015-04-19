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

# add some more info to the heartbeat message with go templates
HEARTBEAT_MESSAGE="$(hostname) - {{.Timestamp}}" ./heartbeat
```

### Debugging messages

Start the app, then in another shell watch whats happening in your redis server

```bash
redis-cli MONITOR
```

### Environment Variables

This app uses environment variables to override configuration through [godotenv](github.com/joho/godotenv).

#### Redis Connection

- `REDIS_URL` (defaults to "localhost:6379")
- `REDIS_PASSWORD` (defaults to "")
- `REDIS_DATABASE` which redis database to connect to (defaults to "0")

#### App settings

- `HEARTBEAT_CHANNEL` channel to publish heartbeats to (defaults to "heartbeat")
- `HEARTBEAT_MESSAGE` message to publish as a heartbeats as a Go Template (defaults to just sending the current unix timestamp)
- `HEARTBEAT_INTERVAL` in seconds (defaults to "1")

### TODO

- [ ] Write tests
- [ ] Support dynamic messages using Go Templates
- [ ] Add debug flag
