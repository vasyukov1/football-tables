package entity

type Match struct {
    ID           int    `json:"id"`
    Team1ID      int    `json:"team1Id"`
    Team2ID      int    `json:"team2Id"`
    Score1       int    `json:"score1"`
    Score2       int    `json:"score2"`
    Stage        string `json:"stage"`
    GroupID      *int   `json:"groupId,omitempty"`
    PlayoffID    *int   `json:"playoffId,omitempty"`
    IsCompleted  bool   `json:"isCompleted"`
    NextMatchID  *int   `json:"nextMatchId,omitempty"`
}
