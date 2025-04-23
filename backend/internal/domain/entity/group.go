package entity

type Group struct {
    ID      int
    Name    string
    Teams   []*Team
    Matches []*Match
}
