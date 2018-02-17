/**
* Author : Siddhant Goenka
*/
package main

import (
	"encoding/json"
)

/* Capitalization of the fields is mandatory for json handler to serialize it */
/* Interface generalizing both JSON structures */
type jsonObject interface {
	marshal() string
	unmarshal(str string)
}


/* JSON structure to send information to workers */
type dataInfo struct	{
	Filename string `json:"datafile"`
	Start uint64 `json:"start"`
	End uint64 `json:"end"`
}

/* JSON structure to send back partial sum to manager */
type result struct {
	Prefix string `json:"prefix"`
	Value uint64 `json:"value"`
	Suffix string `json:"suffix"`
	Chunk uint64 `json:"chunk"`
}


func (di *dataInfo) marshal() string {
	bytarr, err := json.Marshal(*di)			// returns []byte
	if check(err) {
		return string(bytarr)
	}else	{
		panic(err)
	}
}
func (di *dataInfo) unmarshal(jsonStr string)	{
	json.Unmarshal([]byte(jsonStr), di)
}


func (r *result) marshal() string {
	bytarr, err := json.Marshal(*r)
	if check(err) {
		return string(bytarr)
	}else	{
		panic(err)
	}
}
func (r *result) unmarshal(jsonStr string)	{
	json.Unmarshal([]byte(jsonStr), r)
}