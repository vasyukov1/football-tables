package request

type CreateGroupRequest struct {
    Name    string `json:"name" validate:"required"`
    TeamIDs []int  `json:"teamIds"`
}

type UpdateGroupRequest struct {
    Name    *string `json:"name,omitempty"`
    TeamIDs *[]int  `json:"teamIds,omitempty"`
}
