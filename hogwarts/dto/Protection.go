package dto

// Protection is the action of
// the Ministry to help Hogwarts
// quick/strong determine how efficient
// the protection will be
type Protection struct {
	Quick  bool `json:"quick"`
	Strong bool `json:"strong"`
}