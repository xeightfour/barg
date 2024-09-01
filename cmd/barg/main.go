package main

import (
	"fmt"

	"github.com/xeightfour/barg/internal/pkg/prober"
	"github.com/xeightfour/barg/internal/pkg/scrman"
)

func main() {
	fmt.Println("Hello Barg!")
	var pb prober.Prober
	if err := pb.Init("assets/prober.json"); err != nil {
		fmt.Println(err)
		return
	}
	scrman.Test()
	if err := scrman.DrawProber(&pb); err != nil {
		fmt.Println(err)
		return
	}
}
