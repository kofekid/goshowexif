package main

import (
	"fmt"
	"log"
	"os"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/akamensky/argparse"
	"github.com/rwcarlsen/goexif/mknote"
)


func main() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	module := parser.String("m", "module", &argparse.Options{Required: true, Help: "Specify what module you want to use. like Mode, FocalLength etc "})
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	fname := "sample.jpg"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	exif.RegisterParsers(mknote.All...)
	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	switch {
    case *module == "Model" :
        camModel, _ := x.Get(exif.Model)
        fmt.Println(camModel)
    case *module == "iso":
        iso, _ := x.Get(exif.FocalLength)
        fmt.Println(iso)
    case *module == "fstop":
        fstop, _ := x.Get(exif.ApertureValue)
        fmt.Println(fstop)   
    case *module == "lensmake":
        lenseMake, _ := x.Get(exif.LensMake)
        fmt.Println(lenseMake)
    case *module == "lensmodel":
        lenseModel, _ := x.Get(exif.LensModel)
        fmt.Println(lenseModel)     
    case *module == "date":
        	tm, _ := x.DateTime()
			fmt.Println("Taken: ", tm)            
    default:
        fmt.Println("Invalid Request")
    }
    
	

}




