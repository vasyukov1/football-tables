package entity

type Group struct {
    ID      int       `json:"id"`
    Name    string    `json:"name"`
    Teams   []*Team   `json:"teams"`
    Matches []*Match  `json:"matches"`
}
