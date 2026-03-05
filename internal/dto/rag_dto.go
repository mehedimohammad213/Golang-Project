package dto

// RAGAskRequest is the request body for POST /rag/ask.
type RAGAskRequest struct {
	Query string `json:"query" binding:"required"`
}

// RAGAskResponse is the response for POST /rag/ask.
type RAGAskResponse struct {
	Answer  string         `json:"answer"`
	Sources []RAGSourceRef `json:"sources,omitempty"`
}

// RAGSourceRef references a retrieved chunk.
type RAGSourceRef struct {
	SourceType string `json:"source_type"`
	SourceID   string `json:"source_id"`
	Content    string `json:"content"`
}

// RAGIndexResponse is the response for POST /rag/index/cars.
type RAGIndexResponse struct {
	IndexedCars int `json:"indexed_cars"`
}
