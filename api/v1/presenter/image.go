package presenter

type Image struct {
	Path     string `json:"path"`
	Distance int    `json:"distance"`
	Hash     string `json:"hash"`
}
