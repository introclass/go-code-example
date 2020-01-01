package unmarshal

import "encoding/json"

type Obj struct {
	Value int
	Array [100]int
}

type ObjP struct {
	Value int
	Array []int
}

func UnmarshalObj(str []byte, v interface{}) {
	json.Unmarshal(str, v)
}

func UnmarshalObjP(str []byte, v interface{}) {
	json.Unmarshal(str, v)
}
