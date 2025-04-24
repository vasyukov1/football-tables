package request

type CreateMatchRequest struct {
	Team1ID int    `json:"team1_id" binding:"required,min=1"`
	Team2ID int    `json:"team2_id" binding:"required,min=1"`
	Stage   string `json:"stage" binding:"required,oneof=group playoff"`
}
