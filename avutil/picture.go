package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/imgutils.h>
import "C"
import "unsafe"

func AvAllocateImageBuffer(size int) []uint8 {
	p := AvMalloc(uintptr(size))
	if p == nil {
		return nil
	}
	return (*[1 << 30]uint8)(p)[:size]
}

func AvFreeImageBuffer(buffer []uint8) {
	AvFree(unsafe.Pointer(&buffer[0]))
}

//Return the size in bytes of the amount of data required to store an image with the given parameters.
func AvImageGetBufferSize(pixFmt PixelFormat, width, height, align int) int {
	return int(C.av_image_get_buffer_size((C.enum_AVPixelFormat)(pixFmt), C.int(width), C.int(height), C.int(align)))
}

//Setup the data pointers and linesizes based on the specified image parameters and the provided array.
func AvImageFillArrays(dstData [8]*uint8, dstLinesize [8]int32, src []uint8,
	pixFmt PixelFormat, width, height, align int) int {
	cData := (**C.uint8_t)(unsafe.Pointer(&dstData[0]))
	cLinesize := (*C.int)(unsafe.Pointer(&dstLinesize))
	cSrc := (*C.uint8_t)(unsafe.Pointer(&src))
	cPixFmt := (C.enum_AVPixelFormat)(pixFmt)

	return int(C.av_image_fill_arrays(cData, cLinesize, cSrc, cPixFmt, C.int(width), C.int(height), C.int(align)))
}
