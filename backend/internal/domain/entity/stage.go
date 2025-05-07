package entity

type Stage struct {
    ID        int      `json:"id"`
    PlayoffID int      `json:"playoffId"`
    Name      string   `json:"name"`
    Matches   []*Match `json:"matches,omitempty"`
}
