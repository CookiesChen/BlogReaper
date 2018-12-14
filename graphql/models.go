// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type Article struct {
	URL        string   `json:"url"`
	Title      string   `json:"title"`
	Published  string   `json:"published"`
	Updated    string   `json:"updated"`
	Content    string   `json:"content"`
	Summary    string   `json:"summary"`
	Categories []string `json:"categories"`
	Read       bool     `json:"read"`
	Later      bool     `json:"later"`
	FeedID     string   `json:"feedId"`
}

type Category struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Feeds []Feed `json:"feeds"`
}

type Feed struct {
	ID       string    `json:"id"`
	URL      string    `json:"url"`
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Follow   int       `json:"follow"`
	Articles []Article `json:"articles"`
}

type User struct {
	Email string   `json:"email"`
	Info  UserInfo `json:"info"`
}

type UserInfo struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
	Gender int    `json:"gender"`
}
