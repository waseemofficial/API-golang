package entity

type Post struct {
	Id    int64  `json:"id"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

// Write implements io.Writer.
func (Post) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}
