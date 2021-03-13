package reflect

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

var (
	DefaultTimeLoc = time.Local
)

const (
	format_Date     = "2006-01-02"
	format_DateTime = "2006-01-02 15:04:05"
)

//通过反射调用函数(加保护)
func ProtectCall(fn interface{}, args ...interface{}) (result []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 65535)
			l := runtime.Stack(buf, false)
			err = fmt.Errorf(fmt.Sprintf("%v: %s", r, buf[:l]))
		}
	}()
	result = Call(fn, args...)
	return
}

//通过反射调用函数
func Call(fn interface{}, args ...interface{}) (result []interface{}) {
	fnType := reflect.TypeOf(fn)
	fnValue := reflect.ValueOf(fn)
	numIn := fnType.NumIn()

	var out []reflect.Value
	if numIn == 0 {
		out = fnValue.Call(nil)
	} else {
		argsLength := len(args)
		argumentIn := numIn
		if fnType.IsVariadic() {
			argumentIn--
		}

		if argsLength < argumentIn {
			panic("with too few input arguments")
		}

		/*if !fnType.IsVariadic() && argsLength > argumentIn {
			panic("ProtectCall with too many input arguments")
		}*/

		in := make([]reflect.Value, numIn)
		for i := 0; i < argumentIn; i++ {
			if args[i] == nil {
				in[i] = reflect.Zero(fnType.In(i))
			} else {
				in[i] = reflect.ValueOf(args[i])
			}
		}

		if fnType.IsVariadic() {
			m := argsLength - argumentIn
			slice := reflect.MakeSlice(fnType.In(numIn-1), m, m)
			in[numIn-1] = slice
			for i := 0; i < m; i++ {
				x := args[argumentIn+i]
				if x != nil {
					slice.Index(i).Set(reflect.ValueOf(x))
				}
			}
			out = fnValue.CallSlice(in)
		} else {
			out = fnValue.Call(in)
		}
	}

	if out != nil && len(out) > 0 {
		result = make([]interface{}, len(out))
		for i, v := range out {
			result[i] = v.Interface()
		}
	}
	return
}

// set field value to row container
func setFieldValue(ind reflect.Value, value interface{}) {
	switch ind.Kind() {
	case reflect.Bool:
		if value == nil {
			ind.SetBool(false)
		} else if v, ok := value.(bool); ok {
			ind.SetBool(v)
		} else {
			v, _ := StrTo(ToStr(value)).Bool()
			ind.SetBool(v)
		}

	case reflect.String:
		if value == nil {
			ind.SetString("")
		} else {
			ind.SetString(ToStr(value))
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if value == nil {
			ind.SetInt(0)
		} else {
			val := reflect.ValueOf(value)
			switch val.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				ind.SetInt(val.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				ind.SetInt(int64(val.Uint()))
			default:
				v, _ := StrTo(ToStr(value)).Int64()
				ind.SetInt(v)
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if value == nil {
			ind.SetUint(0)
		} else {
			val := reflect.ValueOf(value)
			switch val.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				ind.SetUint(uint64(val.Int()))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				ind.SetUint(val.Uint())
			default:
				v, _ := StrTo(ToStr(value)).Uint64()
				ind.SetUint(v)
			}
		}
	case reflect.Float64, reflect.Float32:
		if value == nil {
			ind.SetFloat(0)
		} else {
			val := reflect.ValueOf(value)
			switch val.Kind() {
			case reflect.Float64:
				ind.SetFloat(val.Float())
			default:
				v, _ := StrTo(ToStr(value)).Float64()
				ind.SetFloat(v)
			}
		}

	case reflect.Struct:
		if value == nil {
			ind.Set(reflect.Zero(ind.Type()))

		} else if _, ok := ind.Interface().(time.Time); ok {
			var str string
			switch d := value.(type) {
			case time.Time:
				TimeFromDB(&d, time.Now().Location())
				ind.Set(reflect.ValueOf(d))
			case []byte:
				str = string(d)
			case string:
				str = d
			}
			if str != "" {
				if len(str) >= 19 {
					str = str[:19]
					t, err := time.ParseInLocation(format_DateTime, str, time.Now().Location())
					if err == nil {
						t = t.In(DefaultTimeLoc)
						ind.Set(reflect.ValueOf(t))
					}
				} else if len(str) >= 10 {
					str = str[:10]
					t, err := time.ParseInLocation(format_Date, str, DefaultTimeLoc)
					if err == nil {
						ind.Set(reflect.ValueOf(t))
					}
				}
			}
		}
	}
}

// convert time from db.
func TimeFromDB(t *time.Time, tz *time.Location) {
	*t = t.In(tz)
}

type argInt []int

// get int by index from int slice
func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}

// interface to string
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

type StrTo string

// set string
func (f *StrTo) Set(v string) {
	if v != "" {
		*f = StrTo(v)
	} else {
		f.Clear()
	}
}

// clean string
func (f *StrTo) Clear() {
	*f = StrTo(0x1E)
}

// check string exist
func (f StrTo) Exist() bool {
	return string(f) != string(0x1E)
}

// string to bool
func (f StrTo) Bool() (bool, error) {
	return strconv.ParseBool(f.String())
}

// string to float32
func (f StrTo) Float32() (float32, error) {
	v, err := strconv.ParseFloat(f.String(), 32)
	return float32(v), err
}

// string to float64
func (f StrTo) Float64() (float64, error) {
	return strconv.ParseFloat(f.String(), 64)
}

// string to int
func (f StrTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int(v), err
}

// string to int8
func (f StrTo) Int8() (int8, error) {
	v, err := strconv.ParseInt(f.String(), 10, 8)
	return int8(v), err
}

// string to int16
func (f StrTo) Int16() (int16, error) {
	v, err := strconv.ParseInt(f.String(), 10, 16)
	return int16(v), err
}

// string to int32
func (f StrTo) Int32() (int32, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int32(v), err
}

// string to int64
func (f StrTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	return int64(v), err
}

// string to uint
func (f StrTo) Uint() (uint, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint(v), err
}

// string to uint8
func (f StrTo) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

// string to uint16
func (f StrTo) Uint16() (uint16, error) {
	v, err := strconv.ParseUint(f.String(), 10, 16)
	return uint16(v), err
}

// string to uint31
func (f StrTo) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint32(v), err
}

// string to uint64
func (f StrTo) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(f.String(), 10, 64)
	return uint64(v), err
}

// string to string
func (f StrTo) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}
