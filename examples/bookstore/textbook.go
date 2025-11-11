package bookapp

type TextBook struct {
	Book  `tstype:",extends"`
	Pages int `json:"pages"`
}
