package types

type User struct {
	Name string
	RegisterDate string
	IsCompany bool
	// Image
}

type Item struct {
	Title string
	Description string
	PhoneNumber string
	User User
	Address string
}

// Parser parse items from a given url
type Parser interface {
	// Parse - download page from url and parse an item from it
	Parse(itemUrl string) (*Item, error)
}