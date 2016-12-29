# goav
Golang binding for FFmpeg.   Forked from [here](https://github.com/giorgisio/goav)

A binding to the ffmpeg video/audio manipulation library.   It's incomplete but is enough to decode frames of ProRes right now.

[![GoDoc](https://godoc.org/github.com/amarburg/goav?status.svg)](https://godoc.org/github.com/amarburg/goav)

[![Travis CI](https://travis-ci.org/amarburg/goav.svg?branch=master)](https://travis-ci.org/amarburg/goav)

## Usage

See [go-prores-ffmpeg](https://github.com/amarburg/go-prores-ffmpeg/blob/master/prores.go) for an example.

I found the original [examples](example/) weren't complete.

## Libraries

* `avcodec` corresponds to the ffmpeg library: libavcodec [provides implementation of a wider range of codecs]
* `avformat` corresponds to the ffmpeg library: libavformat [implements streaming protocols, container formats and basic I/O access]
* `avutil` corresponds to the ffmpeg library: libavutil [includes hashers, decompressors and miscellaneous utility functions]
* `avfilter` corresponds to the ffmpeg library: libavfilter [provides a mean to alter decoded Audio and Video through chain of filters]
* `avdevice` corresponds to the ffmpeg library: libavdevice [provides an abstraction to access capture and playback devices]
* `swresample` corresponds to the ffmpeg library: libswresample [implements audio mixing and resampling routines]
* `swscale` corresponds to the ffmpeg library: libswscale [implements color conversion and scaling routines]


## Installation

Depends on FFMpeg: [install instructions here](https://github.com/FFmpeg/FFmpeg/blob/master/INSTALL.md)

``` sh
sudo apt-get -y --force-yes install autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev

sudo apt-get install yasm

export FFMPEG_ROOT=$HOME/ffmpeg
export CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"
export CGO_CFLAGS="-I$FFMPEG_ROOT/include"
export LD_LIBRARY_PATH=$HOME/ffmpeg/lib
```

```
go get github.com/giorgisio/goav

```

## More Examples

Coding examples are available in the examples/ directory.

## Note
- Function names in Go are consistent with that of the libraries to help with easy search
- [cgo: Extending Go with C](http://blog.giorgis.io/cgo-examples)
- goav comes with absolutely no warranty.

## Contribute
- Fork this repo and create your own feature branch.
- Follow standard Go conventions
- Test your code.
- Create pull request

## TODO

- [ ] Returning Errors
- [ ] Garbage Collection
- [X] Review included/excluded functions from each library
- [ ] Go Tests
- [ ] Possible restructuring packages
- [x] Tutorial01.c
- [ ] More Tutorial

## License
This library is under the [MIT License](http://opensource.org/licenses/MIT)
