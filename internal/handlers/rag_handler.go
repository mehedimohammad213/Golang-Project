package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/car-project/internal/dto"
	"github.com/user/car-project/internal/service"
	"github.com/user/car-project/internal/utils"
)

type RAGHandler struct {
	Service service.RAGService
}

func NewRAGHandler(svc service.RAGService) *RAGHandler {
	return &RAGHandler{Service: svc}
}

// Ask godoc
// @Summary      Ask a question (RAG)
// @Description  Ask a natural language question; answers are generated using retrieval-augmented generation over indexed car/inventory data.
// @Tags         rag
// @Accept       json
// @Produce      json
// @Param        body  body      dto.RAGAskRequest  true  "Query"
// @Success      200  {object}  dto.RAGAskResponse
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /rag/ask [post]
// @Security     BearerAuth
func (h *RAGHandler) Ask(c *gin.Context) {
	var req dto.RAGAskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	result, err := h.Service.Ask(c.Request.Context(), req.Query)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "RAG request failed", err.Error())
		return
	}

	sources := make([]dto.RAGSourceRef, len(result.Sources))
	for i, s := range result.Sources {
		sources[i] = dto.RAGSourceRef{
			SourceType: s.SourceType,
			SourceID:   s.SourceID,
			Content:    s.Content,
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "Answer generated", dto.RAGAskResponse{
		Answer:  result.Answer,
		Sources: sources,
	})
}

// IndexCars godoc
// @Summary      Index cars for RAG
// @Description  Re-index all cars (make, model, details, descriptions) into the RAG vector store. Call this after bulk updates or to refresh the knowledge base.
// @Tags         rag
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.RAGIndexResponse
// @Failure      500  {object}  utils.Response
// @Router       /rag/index/cars [post]
// @Security     BearerAuth
func (h *RAGHandler) IndexCars(c *gin.Context) {
	count, err := h.Service.IndexCars(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Indexing failed", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Cars indexed successfully", dto.RAGIndexResponse{
		IndexedCars: count,
	})
}
