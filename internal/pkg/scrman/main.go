package scrman

import (
	"fmt"
)

func Test() {
	var sc Screen
	sc.Init()
	if err := sc.Box(0, 5, 0, 10); err != nil {
		fmt.Println(err)
		return
	}
	if err := sc.Box(5, 10, 0, 10); err != nil {
		fmt.Println(err)
		return
	}
	if err := sc.VLine(2, 10, 20); err != nil {
		fmt.Println(err)
		return
	}
	if err := sc.TexBox(20, 2, "TexBox!"); err != nil {
		fmt.Println(err)
		return
	}
	if err := sc.Text(6, 18, "Hi There!"); err != nil {
		fmt.Println(err)
		return
	}
	sc.Show()
}
