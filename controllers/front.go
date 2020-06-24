package controllers

import (
	"encoding/json"
	"net/http"
	"io"
)

func RegisterControllers()  {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewDecoder(w)
	enc.Encode(data)
}