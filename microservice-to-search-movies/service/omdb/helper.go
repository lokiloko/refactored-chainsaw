package omdb

import (
	"bytes"
	"encoding/json"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

func Request(params *dto.WebServiceRequestParams, out interface{}) error {
	var payload io.Reader = nil
	if params.Payload != nil {
		payloadByte, _ := json.Marshal(&params.Payload)
		payload = bytes.NewBuffer([]byte(payloadByte))
	}

	req, err := http.NewRequest(params.Method, params.Url, payload)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("WebService request statusCode is not ok")
	}

	if out != nil {
		json.Unmarshal(body, &out)
	}

	return nil
}

