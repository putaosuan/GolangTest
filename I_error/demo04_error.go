package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	flies,err:=filepath.Glob("[")
	if err!=nil&&err==filepath.ErrBadPattern {
		fmt.Println(err)
		return
	}
	fmt.Println("flies:",flies)
}
