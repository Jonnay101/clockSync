package main

import (
	"fmt"

	"github.com/Jonnay101/clockSync/msync"
)

func main() {
	localClock := msync.NewClock()

	fmt.Printf("%+v", localClock)
}
