package search

import (
	"encoding/json"
	"fmt"
	"game-app/param"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Service struct {
	client  *http.Client
	baseUrl string
}

func New(url string) *Service {
	return &Service{
		client: &http.Client{
			Timeout: time.Second * 10, // Timeout after 10 seconds
		},
		baseUrl: url,
	}
}

func (s *Service) SearchRequest(inf param.SearchInfo) (param.SearchResponse, error) {
	fmt.Println("in search service")
	u, err := url.Parse(s.baseUrl)
	if err != nil {
		fmt.Printf("Error parsing URL: %s\n", err)
		return param.SearchResponse{}, err
	}

	params := url.Values{}
	params.Add("cpu", inf.CPU)
	params.Add("ram", strconv.Itoa(int(inf.RAM)))
	params.Add("ssd", strconv.Itoa(int(inf.SSD)))
	params.Add("graphic_ram", strconv.Itoa(int(inf.Graphic)))
	params.Add("screen_size", inf.ScreenSize)
	params.Add("hdd", strconv.Itoa(int(inf.HDD)))
	params.Add("company", inf.Company)

	// Encode the query parameters and set them in the URL
	u.RawQuery = params.Encode()

	// Make the HTTP GET request
	response, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("Error making GET request: %s\n", err)
		return param.SearchResponse{}, err
	}
	defer response.Body.Close()

	var searchRes param.SearchResponse
	err = json.NewDecoder(response.Body).Decode(&searchRes)
	if err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
		return param.SearchResponse{}, err
	}
	return searchRes, nil
}
