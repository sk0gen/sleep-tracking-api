package api

type idUriRequest struct {
	ID string `uri:"id" binding:"uuid"`
}
