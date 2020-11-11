package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int //单位cm
	Weight     int //单位kg
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hukou      string
	Xingzuo    string
	House      string
	Car        string
	Workplace  string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
