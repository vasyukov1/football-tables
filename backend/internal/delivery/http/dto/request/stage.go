package request

type CreateStageRequest struct {
    PlayoffID int    `json:"playoff_id" validate:"required"`
    Name      string `json:"name" validate:"required"`
}

type UpdateStageRequest struct {
    Name *string `json:"name,omitempty"`
}
