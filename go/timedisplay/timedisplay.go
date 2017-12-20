package timedisplay

import (
	"fmt"
	"time"
)

func displayDate() {
	fmt.Printf("%s", time.Now().Format("2006/01/02 15:04"))
}
