package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/transfashion/olahproductimage"
)

func main() {
	fmt.Println("Olah Product Image")

	// ambil lokasi file ini
	currdir, _ := os.Getwd()

	// ambil parameter
	dirSource := filepath.Join(currdir, "data", "source")
	dirTarget := filepath.Join(currdir, "data", "target")

	flag.StringVar(&dirSource, "source", dirSource, "Direktori dari sumber data image yang akan diolah")
	flag.StringVar(&dirTarget, "target", dirTarget, "Direktori target untuk menyimpan image hasil olah")

	fmt.Printf("Source: %s\r\n", dirSource)
	fmt.Printf("Target: %s\r\n", dirTarget)

	ok := YesNoPrompt("Apakah akan dilanjutkan ?", false)
	if ok {
		fmt.Println("Program dimulai...")
		pi := olahproductimage.New(dirSource, dirTarget)
		err := pi.Execute()
		if err != nil {
			fmt.Println("Program Error, tidak bisa dilanjutkan")
			os.Exit(1)
		}
	} else {
		fmt.Println("Eksekusi dibatalkan")
	}

}
