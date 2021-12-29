package main

import (
	"/kapal"
	"fmt"
)

func main() {
	kapal := kapal.NewKapal("Kapal Pesiar", "Harmony of the Seas", 50, 20)
	fmt.Println(kapal.GetKapalJenis())
}
