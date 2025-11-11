package bookapp

type Chapter struct {
	// Chapter title
	Title string `json:"title"`
	// The amount of words in this chapter
	NumberOfWords uint `json:"number_of_words"`
}
