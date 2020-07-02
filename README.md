# goshowexif
 Simple script in Golang to show EXIF data from Image


Awailable modes:
* lensmodel  ( Show lense model)
* lensmake ( Show lense make)
* Model (Show Camera model)
* iso (Shows ISO setting)
* date ( Shows Date when picture was taken )
* fstop (Shows Aperture Values used for image)



# usage 

* Install all needed dependecies 

For reading exif : 

`go get github.com/rwcarlsen/goexif/exif`
`go get 	"github.com/rwcarlsen/goexif/mknote"`

For arg parsing : 

`go get 	"github.com/akamensky/argparse"`



* run the script 


`go run main.go -m <mode here>`



##  Future to do 

* Add that can provide modes in bulk
* More dynamic approach for functions, So that have 1 function and methon for asking data gets build based on input 
* option to provide image via input param. Currently it is hardcoded (Very very bad approach, but for simple test it works) 