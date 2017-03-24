package avutil

//#include <libavutil/mathematics.h>
import "C"

func AvRescaleQ(a int64, bq Rational, cq Rational) int64 {
	return (int64)(C.av_rescale_q(C.int64_t(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq)))
}
