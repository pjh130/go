package str

import (
	"log"
	"reflect"
	"testing"
)

func TestIsLittleEndian(t *testing.T) {
	log.Println("IsLittleEndian", IsLittleEndian())
}

func Test_NumAndBytes(t *testing.T) {
	if true {
		var data int = -20
		var result int = 0

		tt := reflect.TypeOf(data)
		log.Println("data type is:", tt)
		b, err := NumToBytes(data, true)
		if nil != err {
			log.Println("NumToBytes err: ", err)
			t.Fail()
			return
		}

		err = BytesToNum(b, &result, true)
		if nil != err {
			log.Println("BytesToNum err: ", err)
			t.Fail()
			return
		}

		log.Println("data:", data, "====== result:", result)
		if data != result {
			t.Fail()
			return
		}
	}

	if true {
		var data uint = 20
		var result uint = 0

		tt := reflect.TypeOf(data)
		log.Println("data type is:", tt)
		b, err := NumToBytes(data, true)
		if nil != err {
			log.Println("NumToBytes err: ", err)
			t.Fail()
			return
		}

		err = BytesToNum(b, &result, true)
		if nil != err {
			log.Println("BytesToNum err: ", err)
			t.Fail()
			return
		}

		log.Println("data:", data, "====== result:", result)
		if data != result {
			t.Fail()
			return
		}
	}

	if true {
		var data float32 = -32.2
		var result float32 = 0.0

		tt := reflect.TypeOf(data)
		log.Println("data type is:", tt)
		b, err := NumToBytes(data, true)
		if nil != err {
			log.Println("NumToBytes err: ", err)
			t.Fail()
			return
		}

		err = BytesToNum(b, &result, true)
		if nil != err {
			log.Println("BytesToNum err: ", err)
			t.Fail()
			return
		}

		log.Println("data:", data, "====== result:", result)
		if data != result {
			t.Fail()
			return
		}
	}

	if true {
		var data float64 = -64.4
		var result float64 = 0

		tt := reflect.TypeOf(data)
		log.Println("data type is:", tt)
		b, err := NumToBytes(data, true)
		if nil != err {
			log.Println("NumToBytes err: ", err)
			t.Fail()
			return
		}

		err = BytesToNum(b, &result, true)
		if nil != err {
			log.Println("BytesToNum err: ", err)
			t.Fail()
			return
		}

		log.Println("data:", data, "====== result:", result)
		if data != result {
			t.Fail()
			return
		}
	}

	if true {
		var data int64 = -64
		var result int64 = 0

		tt := reflect.TypeOf(data)
		log.Println("data type is:", tt)
		b, err := NumToBytes(data, true)
		if nil != err {
			log.Println("NumToBytes err: ", err)
			t.Fail()
			return
		}

		err = BytesToNum(b, &result, true)
		if nil != err {
			log.Println("BytesToNum err: ", err)
			t.Fail()
			return
		}

		log.Println("data:", data, "====== result:", result)
		if data != result {
			t.Fail()
			return
		}
	}

	if true {
		var data uint64 = 64
		var result uint64 = 0

		tt := reflect.TypeOf(data)
		log.Println("data type is:", tt)
		b, err := NumToBytes(data, true)
		if nil != err {
			log.Println("NumToBytes err: ", err)
			t.Fail()
			return
		}

		err = BytesToNum(b, &result, true)
		if nil != err {
			log.Println("BytesToNum err: ", err)
			t.Fail()
			return
		}

		log.Println("data:", data, "====== result:", result)
		if data != result {
			t.Fail()
			return
		}
	}

	if true {
		var data int32 = -32
		var result int32 = 0

		tt := reflect.TypeOf(data)
		log.Println("data type is:", tt)
		b, err := NumToBytes(data, true)
		if nil != err {
			log.Println("NumToBytes err: ", err)
			t.Fail()
			return
		}

		err = BytesToNum(b, &result, true)
		if nil != err {
			log.Println("BytesToNum err: ", err)
			t.Fail()
			return
		}

		log.Println("data:", data, "====== result:", result)
		if data != result {
			t.Fail()
			return
		}
	}

	if true {
		var data uint32 = 32
		var result uint32 = 0

		tt := reflect.TypeOf(data)
		log.Println("data type is:", tt)
		b, err := NumToBytes(data, true)
		if nil != err {
			log.Println("NumToBytes err: ", err)
			t.Fail()
			return
		}

		err = BytesToNum(b, &result, true)
		if nil != err {
			log.Println("BytesToNum err: ", err)
			t.Fail()
			return
		}

		log.Println("data:", data, "====== result:", result)
		if data != result {
			t.Fail()
			return
		}
	}
}
