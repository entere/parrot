package rest

import (
	"encoding/json"
	"net/http"
)

type Rest struct {
}

type rest struct {
	Msg  string      `json:"msg"`
	Code int32       `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, msg string, code int32, data interface{}) {
	response := &rest{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func Success(w http.ResponseWriter, data interface{}) {
	response := &rest{
		Code: 200,
		Msg:  "Success",
		Data: data,
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

func Error(w http.ResponseWriter, msg string, code int32) {
	response := &rest{
		Code: code,
		Msg:  msg,
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}
