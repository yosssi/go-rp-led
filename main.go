package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString("GPIO No should be specified.\n")
		os.Exit(1)
	}

	gpio := os.Args[1]

	f, err := os.Create("/sys/class/gpio/export")
	if err != nil {
		panic(err)
	}

	f.WriteString(gpio)

	f, err = os.Create(fmt.Sprintf("/sys/class/gpio/gpio%s/direction", gpio))
	if err != nil {
		panic(err)
	}

	f.WriteString("out")

	f, err = os.Create(fmt.Sprintf("/sys/class/gpio/gpio%s/value", gpio))
	if err != nil {
		panic(err)
	}

	f.WriteString("1")

	time.Sleep(3 * time.Second)

	f.WriteString("0")
}
