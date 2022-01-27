package main

import (
	"fmt"

	"github.com/Lennern/gotesting/myquote"
)

func main() {
	fmt.Println(myquote.GetGo())
	fmt.Println(myquote.GetGlass())
	fmt.Println(myquote.GetOpt())
	fmt.Println(myquote.GetHello())
}
