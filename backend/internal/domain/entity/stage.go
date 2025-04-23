package entity

type Stage struct {
    ID        int
    PlayoffID int
    Name      string
    Matches   []*Match
}
