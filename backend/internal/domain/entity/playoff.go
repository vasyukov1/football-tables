package entity

type Playoff struct {
    ID     int     `json:"id"`
    Rounds []Stage `json:"rounds,omitempty"`
}
