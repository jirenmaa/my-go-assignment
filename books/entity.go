package books

type Books struct {
	Books []Book `json:"books"`
}

type Book struct {
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Author      string `json:"author"`
	Published   string `json:"published"`
	Publisher   string `json:"publisher"`
	Pages       int    `json:"pages"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Borrower    string `json:"borrower"`
}
