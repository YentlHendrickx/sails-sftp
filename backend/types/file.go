package types

type FileInfo struct {
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Mode  string `json:"mode"`  // File permissions
	Mtime int64  `json:"mtime"` // Last modified time in Unix timestamp
	IsDir bool   `json:"is_dir"`
}
