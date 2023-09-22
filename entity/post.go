package entity

type Post struct {
	ID    int64  `json:"ID"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

// Write implements io.Writer.
func (Post) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}
