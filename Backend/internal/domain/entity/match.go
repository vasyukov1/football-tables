package entity

type Match struct {
    ID          int
    GroupID     *int
    PlayoffID   *int
    Team1ID     int
    Team2ID     int
    Score1      int
    Score2      int
    Stage       string
    NextMatchID *int
    IsCompleted bool
}
