package main

type KitMeta struct {
	Title    string            `json:"title"`
	Summary  string            `json:"summary"`
	Category string            `json:"category"`
	Env      map[string]string `json:"env,omitempty"`
	Inputs   []KitInput        `json:"inputs"`
}

type KitInput struct {
	Path string `json:"path"`
	Role string `json:"role"`
}

type KitIndex struct {
	Title      string      `json:"title"`
	Summary    string      `json:"summary"`
	Category   string      `json:"category"`
	Files      []FileEntry `json:"files"`
	TotalBytes int64       `json:"total_bytes"`
}

type FileEntry struct {
	Path   string `json:"path"`
	Sha256 string `json:"sha256"`
	Bytes  int64  `json:"bytes"`
}

type Outputs struct {
	Index    KitIndex
	Manifest string // manifest.sha256 contents
}

type VerifyReport struct {
	OK         bool   `json:"ok"`
	KitTitle   string `json:"kit_title"`
	FileCount  int    `json:"file_count"`
	TotalBytes int64  `json:"total_bytes"`
}
