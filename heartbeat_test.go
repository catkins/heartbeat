package main

import (
	"fmt"
	"regexp"
	"testing"
	"time"
)

func TestBuildMessage(t *testing.T) {
	str, _ := buildMessage(time.Now())
	r := regexp.MustCompile("^[0-9]+$")
	if len(r.FindString(str)) == 0 {
		msg := fmt.Sprintf("Expected template to produce an epoch, received: %s", str)
		t.Fatal(msg)
	}
}
