package unmarshal

import (
	"testing"
)

var objstr_noarray = []byte(`{"Value":10}`)
var objstr_array = []byte(`{"Value":10, "Array":[1,2,3,4,5,7,8]}`)

func BenchmarkUnmarshalObjNoArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := &Obj{}
		UnmarshalObj(objstr_noarray, obj)
	}
}

func BenchmarkUnmarshalObjPNoArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := &ObjP{}
		UnmarshalObjP(objstr_noarray, obj)
	}
}

func BenchmarkUnmarshalObjArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := &Obj{}
		UnmarshalObj(objstr_array, obj)
	}
}

func BenchmarkUnmarshalObjPArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := &ObjP{}
		UnmarshalObjP(objstr_array, obj)
	}
}
