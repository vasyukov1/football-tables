package entity

type Match struct {
	ID          int
	Team1ID     int
	Team2ID     int
	Score1      int
	Score2      int
	Stage       string
	GroupID     int
	PlayoffID   int
	IsCompleted bool
}
