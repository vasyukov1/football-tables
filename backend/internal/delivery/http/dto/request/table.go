package request

type CreateTableRequest struct {
    GroupIDs []int `json:"group_ids" validate:"required"`
}

type UpdateTableRequest struct {
    GroupIDs *[]int `json:"group_ids,omitempty"`
}
