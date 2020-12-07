package main

import (
	"log"

	"github.com/ymosleh/goav/avcodec"
	"github.com/ymosleh/goav/avdevice"
	"github.com/ymosleh/goav/avfilter"
	"github.com/ymosleh/goav/avutil"
	"github.com/ymosleh/goav/swresample"
	"github.com/ymosleh/goav/swscale"
)

func main() {
	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())
}
