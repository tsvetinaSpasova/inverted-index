package document

import (
	"net/http"

	"xedni/pkg/domain/document"
	"xedni/pkg/service"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// FetchResponse is the shape of data for a loaded document record
type FetchResponse struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

// Render satisfies the chi interface
func (fr *FetchResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}

// CreateResponse contains the ID post document creation
type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

// Render setups up the correct http status code.
func (cr *CreateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}

// NewFetchResponse instantiate a new response post load
func NewFetchResponse(d document.Document, _ *service.DocumentService) *FetchResponse {
	return &FetchResponse{
		ID:   d.ID,
		Text: d.Text,
	}
}

// NewCreateResponse instantiates a new response when document is created
func NewCreateResponse(ID uuid.UUID, _ *service.DocumentService) *CreateResponse {
	return &CreateResponse{ID: ID}
}