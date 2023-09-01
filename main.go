//
//author:abel
//date:2023/9/1
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Println(idx, arg)
	}
}
