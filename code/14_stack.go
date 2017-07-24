package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// full spec at https://api.stackexchange.com/docs/info
type Item struct {
	Users     int `json:"total_users"`
	Questions int `json:"total_questions"`
	Answers   int `json:"total_answers"`
	Comments  int `json:"total_comments"`
}

type Info struct {
	Site  string
	Items []Item
}

func FetchInfo(site string, infos chan *Info, errs chan error) {
	url := "https://api.stackexchange.com/2.2/info?site=" + site
	response, err := http.Get(url)
	if err != nil {
		errs <- err
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		errs <- errors.New(response.Status)
		return
	}
	var info Info
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&info); err != nil {
		errs <- err
		return
	}
	info.Site = site
	infos <- &info
}

func main() {
	sites := [...]string{
		"stackoverflow",
		"serverfault",
		"superuser"}

	infos := make(chan *Info)
	errs := make(chan error)

	for _, site := range sites {
		go FetchInfo(site, infos, errs)
	}

	for range sites {
		select {
		case info := <-infos:
			fmt.Printf("%#v\n\n", info)
		case err := <-errs:
			fmt.Printf("%s\n\n", err)
		}
	}
}
