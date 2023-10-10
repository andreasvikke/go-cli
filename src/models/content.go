package models

type ContentResult struct {
	Data   []Content      `json:"data"`
	Errors []ContentError `json:"errors"`
}

type ContentError struct {
	Id           string `json:"id"`
	Version      int    `json:"version"`
	ErrorMessage string `json:"errorMessage"`
}

type Content struct {
	Id     string `json:"id,omitempty"`
	Type   string `json:"type"`
	Status string `json:"status,omitempty"`
	Title  string `json:"title"`
	// Ancestors []Ancestor `json:"ancestors,omitempty"`
	Body    Body     `json:"body"`
	Version *Version `json:"version,omitempty"`
	// Space     *Space     `json:"space"`
	// History   *History   `json:"history,omitempty"`
	// Links     *Links     `json:"_links,omitempty"`
	// Metadata  *Metadata  `json:"metadata"`
}

type Version struct {
	By        *User  `json:"by,omitempty"`
	When      string `json:"when,omitempty"`
	Message   string `json:"message,omitempty"`
	Number    int    `json:"number"`
	MinorEdit bool   `json:"minorEdit"`
}

type Body struct {
	View    *Storage `json:"view,omitempty"`
	Storage Storage  `json:"storage"`
}

type Storage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}
