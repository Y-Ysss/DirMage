package dirmage

type Config struct {
	Data     tomlData
	Selector tomlSelector
	Prompt   tomlPrompt
}

type tomlData struct {
	DirsFile string `toml:"dirs_file"`
}

type tomlSelector struct {
	PageSize    int       `toml:"page_size"`
	Text        string    `toml:"selector_text"`
	EditText    string    `toml:"edit_selector_text"`
	EnabledText [2]string `toml:"enabled_text"`
}

type tomlPrompt struct {
	Text string `toml:"prompt_text"`
}
