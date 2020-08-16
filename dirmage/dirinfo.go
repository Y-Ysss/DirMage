package dirmage

type dirInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Enabled     bool   `json:"enabled"`
}

func (dir *dirInfo) SetValues(name string, description string, path string, enabled bool) {
	dir.Name = name
	dir.Description = description
	dir.Path = path
	dir.Enabled = enabled
}

func (dir *dirInfo) SetName(name string) {
	dir.Name = name
}

func (dir *dirInfo) SetDesc(description string) {
	dir.Description = description
}

func (dir *dirInfo) SetPath(path string) {
	dir.Path = path
}

func (dir *dirInfo) SetEnabled(enabled bool) {
	dir.Enabled = enabled
}
