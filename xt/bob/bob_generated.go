package bob

// Automatically generated by bobgen, do not edit.

func (x *partX3) Decode(r *bobReader) error {
	var buf []byte
	var err error
	x.FacesX3, err = dec_sl_faceListX3(r)
	if err != nil {
		return err
	}
	buf, err = r.data(40, true)
	if err != nil {
		return err
	}
	for i_arr10_int32 := range x.X3Vals {
		x.X3Vals[i_arr10_int32] = dec32(buf[0:][i_arr10_int32*4:])
	}
	return nil
}

func (x *Mat6Pair) Decode(r *bobReader) error {
	var buf []byte
	var err error
	x.Name, err = r.decodeString()
	if err != nil {
		return err
	}
	buf, err = r.data(2, true)
	if err != nil {
		return err
	}
	x.Value = dec16(buf[0:])
	return nil
}

func (x *mat6big) Decode(r *bobReader) error {
	var buf []byte
	var err error
	buf, err = r.data(2, true)
	if err != nil {
		return err
	}
	x.Technique = dec16(buf[0:])
	x.Effect, err = r.decodeString()
	if err != nil {
		return err
	}
	x.Value, err = dec_sl_mat6Value(r)
	if err != nil {
		return err
	}
	return nil
}

func (x *weight) Decode(r *bobReader) error {
	var err error
	x.Weights, err = dec_sl_wgt(r)
	if err != nil {
		return err
	}
	return nil
}

func (x *uv) decodeBuf(b []byte) {
	var err error
	_ = err
	x.Idx = dec32(b[0:])
	for i_arr6_float32 := range x.Values {
		x.Values[i_arr6_float32] = decf32(b[4:][i_arr6_float32*4:])
	}
}

func dec_sl_mat6Value(r *bobReader) ([]mat6Value, error) {
	var err error
	l16, err := r.decode16()
	if err != nil {
		return nil, err
	}
	l := int(l16)
	ret := make([]mat6Value, l)
	for i_sl_mat6Value := range ret {
		err = ret[i_sl_mat6Value].Decode(r)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func dec_sl_wgt(r *bobReader) ([]wgt, error) {
	var err error
	l16, err := r.decode16()
	if err != nil {
		return nil, err
	}
	l := int(l16)
	ret := make([]wgt, l)
	b, err := r.data(l*6, true)
	if err != nil {
		return nil, err
	}
	for i_sl_wgt := range ret {
		ret[i_sl_wgt].decodeBuf(b[i_sl_wgt*6:])
	}
	return ret, nil
}

func (x *faceListX3) Decode(r *bobReader) error {
	var buf []byte
	var err error
	buf, err = r.data(4, true)
	if err != nil {
		return err
	}
	x.MaterialIndex = dec32(buf[0:])
	x.Faces, err = dec_sl_arr4_int32(r)
	if err != nil {
		return err
	}
	x.UVList, err = dec_sl_uv(r)
	if err != nil {
		return err
	}
	return nil
}

func dec_sl_faceListX3(r *bobReader) ([]faceListX3, error) {
	var err error
	l16, err := r.decode16()
	if err != nil {
		return nil, err
	}
	l := int(l16)
	ret := make([]faceListX3, l)
	for i_sl_faceListX3 := range ret {
		err = ret[i_sl_faceListX3].Decode(r)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (x *wgt) decodeBuf(b []byte) {
	var err error
	_ = err
	x.Idx = dec16(b[0:])
	x.Coeff = dec32(b[2:])
}

func dec_sl_arr4_int32(r *bobReader) ([][4]int32, error) {
	var err error
	l32, err := r.decode32()
	if err != nil {
		return nil, err
	}
	l := int(l32)
	ret := make([][4]int32, l)
	b, err := r.data(l*16, true)
	if err != nil {
		return nil, err
	}
	for i_sl_arr4_int32 := range ret {
		for i_arr4_int32 := range ret[i_sl_arr4_int32] {
			ret[i_sl_arr4_int32][i_arr4_int32] = dec32(b[i_sl_arr4_int32*16:][i_arr4_int32*4:])
		}
	}
	return ret, nil
}

func dec_sl_uv(r *bobReader) ([]uv, error) {
	var err error
	l32, err := r.decode32()
	if err != nil {
		return nil, err
	}
	l := int(l32)
	ret := make([]uv, l)
	b, err := r.data(l*28, true)
	if err != nil {
		return nil, err
	}
	for i_sl_uv := range ret {
		ret[i_sl_uv].decodeBuf(b[i_sl_uv*28:])
	}
	return ret, nil
}
