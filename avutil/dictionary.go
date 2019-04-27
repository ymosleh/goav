package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/dict.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

type (
	Dictionary      C.struct_AVDictionary
	DictionaryEntry C.struct_AVDictionaryEntry
)

func AvDictSet(d **Dictionary, key, value string, flags int) int {
	ck := C.CString(key)
	defer C.free(unsafe.Pointer(ck))
	cv := C.CString(value)
	defer C.free(unsafe.Pointer(cv))
	return int(C.av_dict_set((**C.struct_AVDictionary)(unsafe.Pointer(d)), ck, cv, C.int(flags)))
}

func AvDictParseString(d **Dictionary, i, keyValSep, pairsSep string, flags int) int {
	ci := C.CString(i)
	defer C.free(unsafe.Pointer(ci))
	ck := C.CString(keyValSep)
	defer C.free(unsafe.Pointer(ck))
	cp := C.CString(pairsSep)
	defer C.free(unsafe.Pointer(cp))
	return int(C.av_dict_parse_string((**C.struct_AVDictionary)(unsafe.Pointer(d)), ci, ck, cp, C.int(flags)))
}

func AvDictGet(d *Dictionary, key string, prev *DictionaryEntry, flags int) *DictionaryEntry {
	ck := C.CString(key)
	defer C.free(unsafe.Pointer(ck))
	return (*DictionaryEntry)(C.av_dict_get((*C.struct_AVDictionary)(d), ck, (*C.struct_AVDictionaryEntry)(prev), C.int(flags)))
}

func AvDictFree(d **Dictionary) {
	C.av_dict_free((**C.struct_AVDictionary)(unsafe.Pointer(d)))
}

func (e *DictionaryEntry) Key() string {
	return C.GoString(e.key)
}

func (e *DictionaryEntry) Value() string {
	return C.GoString(e.value)
}
