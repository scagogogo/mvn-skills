package main

import (
	"fmt"
	"github.com/scagogogo/mvn-skills/pkg/finder"
)

func main() {

	maven, err := finder.FindMaven()
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Maven可执行路径： %s", maven))
	// Output:
	// Maven可执行路径： mvn

}
