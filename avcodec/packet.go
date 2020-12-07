// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"

	"github.com/ymosleh/goav/avutil"
)

func AvPacketAlloc() *Packet {
	return (*Packet)(C.av_packet_alloc())
}

//Initialize optional fields of a packet with default values.
func (p *Packet) AvInitPacket() {
	C.av_init_packet((*C.struct_AVPacket)(p))
}

//Allocate the payload of a packet and initialize its fields with default values.
func (p *Packet) AvNewPacket(s int) int {
	return int(C.av_new_packet((*C.struct_AVPacket)(p), C.int(s)))
}

//Reduce packet size, correctly zeroing padding.
func (p *Packet) AvShrinkPacket(s int) {
	C.av_shrink_packet((*C.struct_AVPacket)(p), C.int(s))
}

//Increase packet size, correctly zeroing padding.
func (p *Packet) AvGrowPacket(s int) int {
	return int(C.av_grow_packet((*C.struct_AVPacket)(p), C.int(s)))
}

//Initialize a reference-counted packet from av_malloc()ed data.
func (p *Packet) AvPacketFromData(d *uint8, s int) int {
	return int(C.av_packet_from_data((*C.struct_AVPacket)(p), (*C.uint8_t)(d), C.int(s)))
}

// By definition this needs to perform a byte copy.   libav takes ownership
// of the buf and frees it when the AvPacket is freed
func (p *Packet) AvPacketFromByteSlice(buf []byte) int {
	ptr := C.malloc(C.size_t(len(buf)))
	cBuf := (*[1 << 30]byte)(ptr)
	copy(cBuf[:], buf)
	return p.AvPacketFromData((*uint8)(ptr), len(buf))
}

//Copy packet, including contents.
func (p *Packet) AvCopyPacket(r *Packet) int {
	panic("deprecated")
	return 0
	//return int(C.av_copy_packet((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))

}

//Copy packet side data.
func (p *Packet) AvCopyPacketSideData(r *Packet) int {
	panic("deprecated")
	return 0
	//return int(C.av_copy_packet_side_data((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))

}

//Free a packet.
func AvPacketFree(p *Packet) {
	var ptr *C.struct_AVPacket = (*C.struct_AVPacket)(p)
	C.av_packet_free(&ptr)
}

//Allocate new information of a packet.
func (p *Packet) AvPacketNewSideData(t AvPacketSideDataType, s int) *uint8 {
	return (*uint8)(C.av_packet_new_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

//Shrink the already allocated side data buffer.
func (p *Packet) AvPacketShrinkSideData(t AvPacketSideDataType, s int) int {
	return int(C.av_packet_shrink_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

//Get side information from packet.
func (p *Packet) AvPacketGetSideData(t AvPacketSideDataType, s *int) *uint8 {
	return (*uint8)(C.av_packet_get_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(s))))
}

//int 	av_packet_merge_side_data (Packet *pkt)
// deprecated
// func (p *Packet) AvPacketMergeSideData() int {
// 	return int(C.av_packet_merge_side_data((*C.struct_AVPacket)(p)))
// }

//int 	av_packet_split_side_data (Packet *pkt)
// deprecated
// func (p *Packet) AvPacketSplitSideData() int {
// 	return int(C.av_packet_split_side_data((*C.struct_AVPacket)(p)))
// }

//Convenience function to free all the side data stored.
func (p *Packet) AvPacketFreeSideData() {
	C.av_packet_free_side_data((*C.struct_AVPacket)(p))
}

//Setup a new reference to the data described by a given packet.
func (p *Packet) AvPacketRef(s *Packet) int {
	return int(C.av_packet_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

//Wipe the packet.
func (p *Packet) AvPacketUnref() {
	C.av_packet_unref((*C.struct_AVPacket)(p))
}

//Move every field in src to dst and reset src.
func (p *Packet) AvPacketMoveRef(s *Packet) {
	C.av_packet_move_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s))
}

//Copy only "properties" fields from src to dst.
func (p *Packet) AvPacketCopyProps(s *Packet) int {
	return int(C.av_packet_copy_props((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

//Convert valid timing fields (timestamps / durations) in a packet from one timebase to another.
func (p *Packet) AvPacketRescaleTs(r, r2 avutil.Rational) {
	C.av_packet_rescale_ts((*C.struct_AVPacket)(p), *(*C.struct_AVRational)(unsafe.Pointer(&r)), *(*C.struct_AVRational)(unsafe.Pointer(&r2)))
}
