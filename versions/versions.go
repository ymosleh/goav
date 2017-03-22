package main

import (
	"log"

	"github.com/ioblank/goav/avcodec"
	"github.com/ioblank/goav/avdevice"
	"github.com/ioblank/goav/avfilter"
	"github.com/ioblank/goav/avformat"
	"github.com/ioblank/goav/avutil"
	"github.com/ioblank/goav/swresample"
	"github.com/ioblank/goav/swscale"
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
