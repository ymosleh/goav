package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/log.h>
import "C"

// Logging constants
const (
	AV_LOG_QUIET   = C.AV_LOG_QUIET
	AV_LOG_PANIC   = C.AV_LOG_PANIC
	AV_LOG_FATAL   = C.AV_LOG_FATAL
	AV_LOG_ERROR   = C.AV_LOG_ERROR
	AV_LOG_WARNING = C.AV_LOG_WARNING
	AV_LOG_INFO    = C.AV_LOG_INFO
	AV_LOG_VERBOSE = C.AV_LOG_VERBOSE
	AV_LOG_DEBUG   = C.AV_LOG_DEBUG
)

// AvLogGetLevel returns the current log level.
func AvLogGetLevel() int {
	return int(C.av_log_get_level())
}

// AvLogSetLevel sets the log level.
func AvLogSetLevel(level int) {
	C.av_log_set_level(C.int(level))
}
