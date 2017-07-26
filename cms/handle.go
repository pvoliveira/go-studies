package cms

import (
	"net/http"
	"time"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Go Projects CMS",
		Content: "Welcome to our home page",
		Posts: []*Post{
			&Post{
				Title:         "Hello, gophers' world!",
				Content:       "Hello world! Thanks for coming to the site.",
				DatePublished: time.Now(),
			},
			&Post{
				Title:         "A post with comments!",
				Content:       "Hello world! Thanks for coming to the site.",
				DatePublished: time.Now(),
				Comments: []*Comment{
					&Comment{
						Author:        "Ben Tranter",
						Comment:       "Nermind, I guess I just commented on my own post...",
						DatePublished: time.Now().Add(-time.Hour / 2),
					},
				},
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}
