package filter

import (
	"encoding/json"
	"strings"

	"github.com/dendi239/gitfame/configs"
)

type Language struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Extensions []string `json:"extensions"`
}

var (
	languages = make(map[string]Language)
)

func init() {
	var langs []Language
	err := json.Unmarshal(configs.LanguageConfigs, &langs)
	if err != nil {
		panic(err)
	}

	for _, l := range langs {
		languages[strings.ToLower(l.Name)] = l
	}
}
