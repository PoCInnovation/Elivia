package bridge

import (
	"fmt"
	"strings"
)

// Response represent a Json object that will be returned by the Modules
type Response struct {
	Tag      string                 `json:"package"`
	Response string                 `json:"response"`
	Data     map[string]interface{} `json:"data"`
}

// Init is a way to simulate the Response constructor
func (r *Response) Init(Tag, Response string, data ...map[string]interface{}) Response {
	r.Tag = Tag
	r.Response = Response
	r.Data = make(map[string]interface{})
	r.AppendData(data...)
	return *r
}

// AppendData to the original return structure
func (r *Response) AppendData(data ...map[string]interface{}) Response {
	for _, elm := range data {
		for key, value := range elm {
			r.Data[key] = value
		}
	}
	return *r
}

// Format the Response by replacing variadic data
func (r *Response) Format() Response {
OUTER:
	for i := strings.Index(r.Response, "%") + 1; i > 0; i = strings.Index(r.Response, "%") + 1 {
		ii := strings.Index(r.Response[i:], "%") + 1
		if ii == 0 {
			fmt.Println("Response is ill formated: abording format")
			break
		}
		vname := r.Response[i : ii+i-1]
		if vars, exist := r.Data[vname]; exist {
			if value, ok := vars.(string); ok {
				r.Response = r.Response[:i-1] + value + r.Response[ii+i:]
			}
			continue OUTER
		} else {
			fmt.Println(i, ii, "missing data \"", vname, "\", can't format response: abording format")
			return *r
		}
	}
	return *r
}
