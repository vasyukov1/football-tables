package entity

type Table struct {
    ID       int   `json:"id"`
    GroupIDs []int `json:"groupIds"`
}
