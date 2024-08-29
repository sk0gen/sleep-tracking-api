package pagination

type PaginatedResponse[T any] struct {
	Results    []T   `json:"items"`
	PageNumber int32 `json:"pageNumber" example:"1"`
	PageSize   int32 `json:"pageSize" example:"10"`
	TotalItems int64 `json:"totalItems" example:"100"`
}

type PaginatedRequest struct {
	PageNumber int32 `json:"pageNumber" example:"1"`
	PageSize   int32 `json:"pageSize" binding:"max=100" example:"10"`
}

func (req *PaginatedRequest) CheckDefaults() {
	if req.PageNumber <= 0 {
		req.PageNumber = 1
	}
	if req.PageSize <= 10 {
		req.PageSize = 10
	}
}
