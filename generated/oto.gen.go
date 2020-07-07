// Code generated by oto; DO NOT EDIT.

package generated

import (
	"context"
	"log"
	"net/http"

	"github.com/pacedotdev/oto/otohttp"

	uuid "github.com/google/uuid"
)

type CardService interface {
	Card(context.Context, CardRequest) (*CardResponse, error)
	Cards(context.Context, CardsRequest) (*CardsResponse, error)
	New(context.Context, CreateRequest) (*CardResponse, error)
}

type cardServiceServer struct {
	server      *otohttp.Server
	cardService CardService
}

func RegisterCardService(server *otohttp.Server, cardService CardService) {
	handler := &cardServiceServer{
		server:      server,
		cardService: cardService,
	}
	server.Register("CardService", "Card", handler.handleCard)
	server.Register("CardService", "Cards", handler.handleCards)
	server.Register("CardService", "New", handler.handleNew)
}

func (s *cardServiceServer) handleCard(w http.ResponseWriter, r *http.Request) {
	var request CardRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.cardService.Card(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *cardServiceServer) handleCards(w http.ResponseWriter, r *http.Request) {
	var request CardsRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.cardService.Cards(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *cardServiceServer) handleNew(w http.ResponseWriter, r *http.Request) {
	var request CreateRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.cardService.New(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

type CardRequest struct {
	ID string `json:"id"`
}

type CardResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    int       `json:"status"`
	Creator   uuid.UUID `json:"creator"`
	Volunteer uuid.UUID `json:"volunteer"`
	CreatedAt string    `json:"createdAt"`
	Error     string    `json:"error,omitempty"`
}

type CardsRequest struct {
}

type CardsResponse struct {
	Cards []CardResponse `json:"cards"`
	Error string         `json:"error,omitempty"`
}

type CreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Creator string `json:"creator"`
}
