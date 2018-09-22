package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/dict.h>
import "C"
import (
	"unsafe"
)

type (
	Dictionary      C.struct_AVDictionary
	DictionaryEntry C.struct_AVDictionaryEntry
)

func AvDictSet(d **Dictionary, key, value string, flags int) int {
	return int(C.av_dict_set((**C.struct_AVDictionary)(unsafe.Pointer(d)), C.CString(key), C.CString(value), C.int(flags)))
}

func AvDictParseString(d **Dictionary, i, keyValSep, pairsSep string, flags int) int {
	return int(C.av_dict_parse_string((**C.struct_AVDictionary)(unsafe.Pointer(d)), C.CString(i), C.CString(keyValSep), C.CString(pairsSep), C.int(flags)))
}

func AvDictGet(d *Dictionary, key string, prev *DictionaryEntry, flags int) *DictionaryEntry {
	return (*DictionaryEntry)(C.av_dict_get((*C.struct_AVDictionary)(d), C.CString(key), (*C.struct_AVDictionaryEntry)(prev), C.int(flags)))
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
