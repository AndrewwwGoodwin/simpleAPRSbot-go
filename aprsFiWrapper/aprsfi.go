package aprsFiWrapper

import (
	"encoding/json"
	"io"
	"net/http"
)

var endpoint string = "https://api.aprs.fi/api"

type AprsFiWrapper struct {
	apiKey string
}

func NewAprsFiWrapper(apiKey string) *AprsFiWrapper {
	return &AprsFiWrapper{apiKey: apiKey}
}

func (wrapper AprsFiWrapper) GetLocation(callAndSSID string) *AprsFiLocationStruct {
	resp, err := http.Get(endpoint + "/get?name=" + callAndSSID + "&apikey=" + wrapper.apiKey + "&format=json&what=loc")
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	aprsfiLocation := AprsFiLocationStruct{}
	_ = json.Unmarshal(body, &aprsfiLocation)
	return &aprsfiLocation
}

type AprsFiLocationStruct struct {
	Command string `json:"command"`
	Result  string `json:"result"`
	What    string `json:"what"`
	Found   int    `json:"found"`
	Entries []struct {
		Name     string `json:"name"`
		Type     string `json:"type"`
		Time     string `json:"time"`
		Lasttime string `json:"lasttime"`
		Lat      string `json:"lat"`
		Lng      string `json:"lng"`
		Symbol   string `json:"symbol"`
		Srccall  string `json:"srccall"`
		Dstcall  string `json:"dstcall"`
		Phg      string `json:"phg"`
		Comment  string `json:"comment"`
		Path     string `json:"path"`
	} `json:"entries"`
}
