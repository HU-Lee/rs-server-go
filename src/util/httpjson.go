package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetReqJson(req *http.Request, json_data interface{}) error {
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &json_data); err != nil {
		return err
	}

	return nil
}

func GetResJson(res *http.Response, json_data interface{}) error {
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &json_data); err != nil {
		return err
	}

	return nil
}
