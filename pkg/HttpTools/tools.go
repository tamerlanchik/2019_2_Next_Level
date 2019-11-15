package HttpTools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Inflate struct from json written in body
func StructFromBody(r http.Request, s interface{}) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, s)
}

// Inflatr body with struct content
func BodyFromStruct(w http.ResponseWriter, s interface{}) error {
	js, err := json.Marshal(s)
	if err != nil {
		return err
	}
	w.Write(js)	//sets inchangeble status 200 if not set already
	return nil
}



type Answer interface {
	SetStatus(string)
}

const (
	defaultStatus = "OK"
)

type Response struct {
	Body  interface{}
	status int
	writer http.ResponseWriter
}

func (r *Response) SetWriter(writer http.ResponseWriter) *Response {
	r.writer = writer
	return r
}

func (r *Response) SetStatus(status int) *Response {
	r.status = status
	return r
}

func (r *Response) SetError(err interface{}) *Response {
	r.Body = err
	return r
}

func (r *Response) Copy() Response {
	return *r
}

func (r *Response) Send() {
	if r.Body == nil {
		fmt.Println("Nil error")
	}
	body, err := json.Marshal(r.Body)
	if err != nil {
		fmt.Println("Cannot encode json")
		return
	}
	r.writer.Header().Set("Content-Type", "application/json")
	if r.status == 0 {
		r.status = http.StatusOK
	}
	r.writer.WriteHeader(r.status)
	r.writer.Write(body)
}

func (r *Response) String() string {
	body, _ := json.Marshal(r.Body)
	return string(body)
}
