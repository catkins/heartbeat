package config_test

import (
	"bytes"
	"fmt"
	"github.com/catkins/heartbeat/config"
	"github.com/joho/godotenv"
	"testing"
	"text/template"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	msg := fmt.Sprintf("%s %v != %v", message, a, b)
	t.Fatal(msg)
}

// Ensure the default values work
func TestDefaultValues(t *testing.T) {
	e_redisUrl := "localhost:6379"
	e_redisPassword := ""
	e_redisDatabase := int64(0)
	e_hbChannel := "heartbeat"
	e_hbTemplate, _ := template.New("testhb-message").Parse("{{.Timestamp}}")
	e_hbInterval := int64(1)

	var e_buffer bytes.Buffer
	var a_buffer bytes.Buffer

	var e_context struct {
		Timestamp string
	}
	e_context.Timestamp = "1505556760"
	e_hbTemplate.Execute(&e_buffer, e_context)
	e_msg := e_buffer.String()

	var a_context struct {
		Timestamp string
	}
	a_context.Timestamp = "1505556760"

	c := config.Load()
	c.HeartbeatTemplate.Execute(&a_buffer, a_context)
	a_msg := a_buffer.String()

	assertEqual(t, e_redisUrl, c.RedisAddress, "Default RedisURL mismatch")
	assertEqual(t, e_redisPassword, c.RedisPassword, "Default RedisPassword mismatch")
	assertEqual(t, e_redisDatabase, c.RedisDatabase, "Default RedisDb mismatch")
	assertEqual(t, e_hbChannel, c.HeartbeatChannel, "Default Heartbeat channel mismatch")
	assertEqual(t, e_hbInterval, c.HeartbeatInterval, "Default Heartbeat interval mismatch")
	assertEqual(t, e_msg, a_msg, "Default Heartbeat Template mismatch")
}

// Ensure the default values can be overriden
func TestOverriddenValues(t *testing.T) {
	godotenv.Load(".env.test")
	c := config.Load()
	assertEqual(t, "localhost:7000", c.RedisAddress, "Overriden RedisURL mismatch")
	assertEqual(t, "foo", c.RedisPassword, "Overriden RedisPassword mismatch")
	assertEqual(t, int64(1), c.RedisDatabase, "Overriden RedisDb mismatch")
	assertEqual(t, "mayday", c.HeartbeatChannel, "Overriden Heartbeat channel mismatch")
	assertEqual(t, int64(10), c.HeartbeatInterval, "Overriden Heartbeat interval mismatch")
}
