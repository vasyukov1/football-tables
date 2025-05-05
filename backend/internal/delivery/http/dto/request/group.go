package request

type CreateGroupRequest struct {
    Name    string `json:"name" validate:"required"`
    TeamIDs []int  `json:"team_ids"` // optional для initial assignment
}

type UpdateGroupRequest struct {
    Name    *string `json:"name,omitempty"`
    TeamIDs *[]int  `json:"team_ids,omitempty"`
}
