package tc

// GenerateTokensRequest request contract
type GenerateTokensRequest struct {
	UserID int
	Email  string
	Count  int
}

// GenerateTokensResponse response contract
type GenerateTokensResponse struct {
	Success bool
}
