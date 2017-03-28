package main

import (
	"log"

	"github.com/amarburg/goav/avcodec"
	"github.com/amarburg/goav/avdevice"
	"github.com/amarburg/goav/avfilter"
	"github.com/amarburg/goav/avformat"
	"github.com/amarburg/goav/avutil"
	"github.com/amarburg/goav/swresample"
	"github.com/amarburg/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
