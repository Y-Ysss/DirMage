package dirmage

type DirInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Enabled bool `json:enabled`
}

type Config struct {
	Selector tomlSelector
	Prompt tomlPrompt
}

type tomlSelector struct {
	DirsListPath string `toml:"dirs_list_path"`
	PageSize int `toml:"page_size"`
}

type tomlPrompt struct {
	String string `toml:"prompt_string"`
}

var PromptString string = "\n\x1b[32m[{dirName}] \x1b[0m\x1b[36m{workingDir} {git}\x1b[37m\n$ "
var DirectoriesList string = "directories.json"
// type Dirs struct {
// 	DirInfo []DirInfo `json:"directories"`
// }

