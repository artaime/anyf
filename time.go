package anyf

import (
	"fmt"
	"time"
)


func PrintTime() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}