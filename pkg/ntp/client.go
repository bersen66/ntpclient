package ntp

import (
	"fmt"
	"os"
	"time"

	ntplib "github.com/beevik/ntp"
)

func RunClient(server string) {
	currTime, err := GetCurrentTime(server)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Current time: %v\n", currTime)
}

func GetCurrentTime(server string) (time.Time, error) {
	return ntplib.Time(server)
}
