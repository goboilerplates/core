package main

import (
	"fmt"

	"github.com/goboilerplates/core"
)

func main() {
	var interactor core.GetSamplesInteractor
	interactor = core.CreateDefaultGetSamples("Kaka", "Ronaldo")

	fmt.Println("Get all samples")
	samples, _ := interactor.All()
	for _, sample := range samples {
		fmt.Println(sample)
	}

	fmt.Println("\nGet all samples by name keyword: ka")
	samples, _ = interactor.AllByName("ka")
	for _, sample := range samples {
		fmt.Println(sample)
	}
}
