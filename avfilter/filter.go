// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import "unsafe"

//Get a filter definition matching the given name.
func AvfilterGetByName(n string) *Filter {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return (*Filter)(C.avfilter_get_by_name(cn))
}

//Register a filter.
func (f *Filter) AvfilterRegister() int {
	panic("deprecated")
	return 0
	//return int(C.avfilter_register((*C.struct_AVFilter)(f)))
}

//Iterate over all registered filters.
func (f *Filter) AvfilterNext() *Filter {
	panic("deprecated")
	return nil
	//return (*Filter)(C.avfilter_next((*C.struct_AVFilter)(f)))
}
