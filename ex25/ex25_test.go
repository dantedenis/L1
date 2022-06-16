package ex25

import (
	"testing"
	"fmt"
	"time"
)

func TestSleep(t *testing.T) {
	fmt.Println(time.Now())
	SleepWithAfter(time.Second * 2)
	fmt.Println(time.Now())
}
