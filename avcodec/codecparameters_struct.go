package avcodec

func (cp *CodecParameters) CodecType() MediaType {
	return MediaType(cp.codec_type)
}

func (cp *CodecParameters) CodecId() CodecId {
	return CodecId(cp.codec_id)
}
