package configs

import (
	// need for go embed feature
	_ "embed"
)

var (
	// LanguageConfigs is a json with necessary mapping from languages to extensions
	//go:embed language_extensions.json
	LanguageConfigs []byte
)
