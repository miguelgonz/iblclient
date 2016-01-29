package iblclient

import (
	"encoding/json"
)

func FetchChannelProgrammes(channel string) ChannelProgrammes {
	var parsed ChannelProgrammesResponse
	response := <-request("channels/"+channel+"/programmes", requestParams{})
	json.Unmarshal([]byte(response), &parsed)
	return parsed.ChannelProgrammes
}

func FetchChannels() []Channel {
	var parsed ChannelsResponse
	response := <-request("channels", requestParams{})
	json.Unmarshal([]byte(response), &parsed)
	return parsed.Channels
}

func FetchSearch(query string) []Item {
	var parsed SearchResponse
	response := <-request("search", requestParams{query: query})
	json.Unmarshal([]byte(response), &parsed)
	return parsed.Search.Items

}
