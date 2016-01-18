package iblclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const HOST = "http://ibl.api.bbci.co.uk/ibl/v1/"

type RequestParams struct {
	lang         string
	availability string
	rights       string
}

func paramsToString(params RequestParams) (out string) {
	values := make(url.Values)

	if params.lang == "" {
		params.lang = "en"
	}
	values.Add("lang", params.lang)

	if params.availability == "" {
		params.availability = "available"
	}
	values.Add("availability", params.availability)

	if params.rights == "" {
		params.rights = "web"
	}
	values.Add("rights", params.rights)

	return values.Encode()
}

func Request(route string, params RequestParams) chan string {
	out := make(chan string)
	go func() {
		fullUrl := HOST + route + "?" + paramsToString(params)
		fmt.Println(fullUrl)
		response, _ := http.Get(fullUrl)
		defer response.Body.Close()
		contents, _ := ioutil.ReadAll(response.Body)

		out <- string(contents)
	}()
	return out

}
