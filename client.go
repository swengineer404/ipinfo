package ipinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	client *http.Client
)

func init() {
	client = &http.Client{
		Timeout: time.Second * 30,
	}
}

func Get(addr string) (*IPInfo, error) {
	res, err := client.Get("https://api.incolumitas.com/?q=" + addr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var info IPInfo
	if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("invalid status code: %d", res.StatusCode)
	}

	return &info, nil
}
