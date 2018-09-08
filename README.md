# goav
Golang binding for __FFmpeg 4.0.2__.   Forked from [here](https://github.com/amarburg/goav)

## Libraries

* `avcodec` corresponds to the ffmpeg library: libavcodec [provides implementation of a wider range of codecs]
* `avformat` corresponds to the ffmpeg library: libavformat [implements streaming protocols, container formats and basic I/O access]
* `avutil` corresponds to the ffmpeg library: libavutil [includes hashers, decompressors and miscellaneous utility functions]
* `avfilter` corresponds to the ffmpeg library: libavfilter [provides a mean to alter decoded Audio and Video through chain of filters]
* `avdevice` corresponds to the ffmpeg library: libavdevice [provides an abstraction to access capture and playback devices]
* `swresample` corresponds to the ffmpeg library: libswresample [implements audio mixing and resampling routines]
* `swscale` corresponds to the ffmpeg library: libswscale [implements color conversion and scaling routines]


## Installation

Depends on FFMpeg __3.2.xx__: [install instructions here](https://github.com/FFmpeg/FFmpeg/blob/master/INSTALL.md).   __Note__:  This version is included in Ubuntu 16.10 (Yakkety), but Ubuntu 16.04 (Xenial) is still on 2.7.xx.   You will need to install an [updated version](https://launchpad.net/~jonathonf/+archive/ubuntu/ffmpeg-3)

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

- [ ] __Write some Go Tests__
- [ ] Get working in Wercker -- requires tests

## License
This library is under the [MIT License](http://opensource.org/licenses/MIT)
