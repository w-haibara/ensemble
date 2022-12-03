package docker

import (
	"net/url"
	"reflect"
	"strconv"
)

type TypeCode int

const (
	BoolType TypeCode = iota
	Int64Type
	Float64Type
	StringType
)

type Request map[string]*string

func NewRequest(m map[string]string) Request {
	r := Request{}
	for k := range m {
		v := m[k]
		r[k] = &v
	}
	return r
}

func NewRequestFromURLValues(values url.Values) Request {
	r := Request{}
	for k, v := range values {
		if len(v) == 0 {
			continue
		}
		str := v[0]
		r[k] = &str
	}
	return r
}

func (r Request) Value(key string) *string {
	v, ok := r[key]
	if !ok {
		return nil
	}

	return v
}

func (r Request) ParsedValue(key string, typ TypeCode) (res any, ok bool, err error) {
	v := r.Value(key)
	if v == nil {
		return nil, false, nil
	}

	switch typ {
	case BoolType:
		v, err := strconv.ParseBool(*v)
		if err != nil {
			return nil, false, err
		}

		return v, true, nil
	case Int64Type:
		v, err := strconv.Atoi(*v)
		if err != nil {
			return nil, false, err
		}

		return int64(v), true, nil
	case Float64Type:
		v, err := strconv.ParseFloat(*v, 64)
		if err != nil {
			return nil, false, err
		}

		return v, true, nil
	case StringType:
		return *v, true, nil
	default:
		return nil, false, nil
	}
}

func (r Request) Unmarshal(v any) error {
	if v == nil {
		return nil
	}

	val := reflect.ValueOf(v).Elem()

	fields := reflect.VisibleFields(val.Type())

	for _, fv := range fields {
		f := val.FieldByName(fv.Name)
		if !f.IsValid() || !f.CanSet() {
			continue
		}

		switch f.Type().Kind() {
		case reflect.Bool:
			if err := r.setBool(f, fv.Name); err != nil {
				return err
			}
		case reflect.Int64:
			if err := r.setInt64(f, fv.Name); err != nil {
				return err
			}
		case reflect.Float64:
			if err := r.setFloat64(f, fv.Name); err != nil {
				return err
			}
		case reflect.String:
			if err := r.setString(f, fv.Name); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r Request) setBool(fv reflect.Value, name string) error {
	v, ok, err := r.ParsedValue(name, BoolType)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	fv.SetBool(v.(bool))
	return nil
}

func (r Request) setInt64(fv reflect.Value, name string) error {
	v, ok, err := r.ParsedValue(name, Int64Type)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	fv.SetInt(v.(int64))
	return nil
}

func (r Request) setFloat64(fv reflect.Value, name string) error {
	v, ok, err := r.ParsedValue(name, Float64Type)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	fv.SetFloat(v.(float64))
	return nil
}

func (r Request) setString(fv reflect.Value, name string) error {
	v, ok, err := r.ParsedValue(name, StringType)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	fv.SetString(v.(string))
	return nil
}
