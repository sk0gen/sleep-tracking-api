package pagination

type PaginatedResponse[T any] struct {
	Results    []T   `json:"items"`
	PageNumber int32 `json:"pageNumber"`
	PageSize   int32 `json:"pageSize"`
	TotalItems int64 `json:"totalItems"`
}

type PaginatedRequest struct {
	PageNumber int32 `json:"pageNumber""`
	PageSize   int32 `json:"pageSize" binding:"max=100"`
}

func (req *PaginatedRequest) CheckDefaults() {
	if req.PageNumber <= 0 {
		req.PageNumber = 1
	}
	if req.PageSize <= 10 {
		req.PageSize = 10
	}
}
