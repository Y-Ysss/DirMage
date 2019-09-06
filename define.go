package main

type DirInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Enabled bool `json:enabled`
}

var PromptString string = "\n\x1b[32m[{dirName}] \x1b[0m\x1b[36m{workingDir} {git}\x1b[37m\n$ "

// type Dirs struct {
// 	DirInfo []DirInfo `json:"directories"`
// }

