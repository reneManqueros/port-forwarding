package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Settings struct {
	WebAddress   string        `json:"web_address"`
	Redirections []Redirection `json:"redirections"`
}

func (s *Settings) AddRedirection(redirection Redirection) {
	// ToDo: make threadsafe
	s.Redirections = append(s.Redirections, redirection)

	// ToDo: Add marker for threadsafe killswitch
	s.Save()
}

func (s *Settings) Load() {
	data, err := ioutil.ReadFile("settings.json")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(data, &s)
	if err != nil {
		log.Fatalln(err)
	}
}

func (s Settings) Save() {
	data, err := json.Marshal(s)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile("settings.json", data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
