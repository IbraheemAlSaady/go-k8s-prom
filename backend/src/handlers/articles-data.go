package handlers

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func GetArticlesData() []Article {
	return []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 3", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 4", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 5", Desc: "Article Description", Content: "Article Content"},
	}
}
