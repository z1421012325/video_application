package tools

import "encoding/json"

func ReqMapStruct(req ,to interface{}){
	b,_ := json.Marshal(req)
	_  = json.Unmarshal(b,to)
}
