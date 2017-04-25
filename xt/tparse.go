package xt

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type tParser struct {
	rec     []string
	lastTag string
	t       Text
}

func (t *tParser) parseAll(data interface{}) {
	err := t.pvalue(reflect.Indirect(reflect.ValueOf(data)))
	if err != nil {
		log.Fatal(err)
	}

	if len(t.rec) == 1 && t.rec[0] == "" {
		t.rec = t.rec[1:]
	}
	if len(t.rec) != 0 {
		log.Fatalf("record not fully consumed: %v %v", t.rec, len(t.rec))
	}
}

func (t *tParser) pint(v reflect.Value) error {
	n, err := strconv.Atoi(t.rec[0])
	if err != nil {
		return err
	}
	v.SetInt(int64(n))
	t.rec = t.rec[1:]
	return nil
}

func (t *tParser) pfloat(v reflect.Value) error {
	n, err := strconv.ParseFloat(t.rec[0], 64)
	if err != nil {
		return err
	}
	v.SetFloat(n)
	t.rec = t.rec[1:]
	return nil
}

func tagParse(tag string) map[string]string {
	ret := make(map[string]string)
	for _, t := range strings.Split(tag, ",") {
		ts := strings.SplitN(t, ":", 2)
		if len(ts) == 1 {
			ret[ts[0]] = "true"
		} else {
			ret[ts[0]] = ts[1]
		}
	}
	return ret
}

func (t *tParser) pstring(v reflect.Value) error {
	if t.lastTag != "" {
		tags := tagParse(t.lastTag)
		if tags["page"] != "" {
			pid, err := strconv.Atoi(tags["page"])
			if err != nil || t.t[pid] == nil {
				return fmt.Errorf("Bad page tag: %v", t.lastTag)
			}
			var off int
			if tags["offset"] != "" {
				off, err = strconv.Atoi(tags["offset"])
				if err != nil {
					return fmt.Errorf("Bad offset: %v", tags["offset"])
				}
			}
			tid, err := strconv.Atoi(t.rec[0])
			tid += off
			if err == nil && t.t[pid][tid] != "" {
				v.SetString(t.t[pid][tid])
				t.rec = t.rec[1:]
				return nil
			}
			if t.rec[0] == "0" || t.rec[0] == tags["ignore"] {
				v.SetString("")
				t.rec = t.rec[1:]
				return nil
			}
			log.Printf("bad string ID: %d/%v", pid, t.rec[0])
		}
	}
	v.SetString(t.rec[0])
	t.rec = t.rec[1:]
	return nil
}

func (t *tParser) parray(v reflect.Value) error {
	for i := 0; i < v.Len(); i++ {
		var err error
		err = t.pvalue(v.Index(i))
		if err != nil {
			return fmt.Errorf("Array field (%d): %v", i, err)
		}
	}
	return nil
}

func (t *tParser) pstruct(v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		sf := v.Type().Field(i)
		t.lastTag = sf.Tag.Get("x3t")
		err := t.pvalue(fv)
		if err != nil {
			return fmt.Errorf("Parse Field (%s): %v", v.Type().Field(i).Name, err)
		}
	}
	return nil
}

func (t *tParser) pslice(v reflect.Value) error {
	// Slices are prefixed with a length
	l, err := strconv.Atoi(t.rec[0])
	if err != nil {
		return fmt.Errorf("slice length: %v", err)
	}
	t.rec = t.rec[1:]
	v.Set(reflect.MakeSlice(v.Type(), l, l))
	for i := 0; i < l; i++ {
		err := t.pvalue(v.Index(i))
		if err != nil {
			return fmt.Errorf("slice field (%d): %v", i, err)
		}
	}
	return nil
}

func (t *tParser) pvalue(v reflect.Value) error {
	switch v.Kind() {
	case reflect.Int:
		return t.pint(v)
	case reflect.Float64:
		return t.pfloat(v)
	case reflect.String:
		return t.pstring(v)
	case reflect.Array:
		return t.parray(v)
	case reflect.Struct:
		return t.pstruct(v)
	case reflect.Slice:
		return t.pslice(v)
	default:
		return fmt.Errorf("bad kind: %v", v.Kind())
	}
}