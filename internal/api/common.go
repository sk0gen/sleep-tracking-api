package api

type idUriRequest struct {
	ID string `uri:"id" binding:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
}
