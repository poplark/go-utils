package common

import (
	"bytes"
	"encoding/json"
	"github.com/derekstavis/go-qs"
	"io/ioutil"
	"net/http"
	"strings"
	"uframework/log"
)

func InvokeService(url string, method string, data map[string]interface{}) ([]byte, error) {
	m := strings.ToUpper(method)

	if m == "POST" {
		b, err := json.Marshal(data)
		if err != nil {
			uflog.ERROR("[Invoke Service] - "+m+" ("+url+") - json error: ", err)
			return nil, err
		}
		body := bytes.NewBuffer([]byte(b))

		uflog.INFO("[Invoke Service] - POST "+url+": ", body)

		res, err := http.Post(url, "application/json", body)

		if err != nil {
			uflog.ERROR("[Invoke Service] - POST ("+url+"): ", data, ` Error: `, err)
			return nil, err
		}
		defer res.Body.Close()
		result, err := ioutil.ReadAll(res.Body)
		return result, err
	} else {
		querystring, err := qs.Marshal(data)

		if err != nil {
			uflog.ERROR("[Invoke Service] - "+m+" ("+url+") - querystring error: ", err)
			return nil, err
		}

		uflog.INFO("[Invoke Service] - GET " + url + "?" + querystring)

		res, err := http.Get(url + "?" + querystring)
		if err != nil {
			uflog.ERROR("[Invoke Service] - GET ("+url+"): ", data, ` Error: `, err)
			return nil, err
		}
		defer res.Body.Close()
		return ioutil.ReadAll(res.Body)
	}
}
