package models

type Blog struct {
	ID        string
	Title     string
	Body      string
	CreatedAt string
}

// func (p *Blog) Save() error {
// 	filename := p.Title + ".txt"
// 	return os.WriteFile(filename, p.Body, 0600)
// }
