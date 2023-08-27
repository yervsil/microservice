package handler

import (
	"context"
	"encoding/json"

	"io"

	"net/http"

	"time"

	"github.com/yervsil/api_gateway/internal/dto"
	"github.com/yervsil/api_gateway/pkg/utils"
	"errors"
)

func (h *Handler) userSignUp(w http.ResponseWriter, r *http.Request) {
	var inp dto.SignIn

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.AppLogger.Info("error in reading request body: %v", err)
		utils.HandleError(w, err, 400)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(bytes, &inp); err != nil {
		h.AppLogger.Info("error in unmarshaling: %v", err)
		utils.HandleError(w, err, 400)
		return 
	}

	if inp.Email == "" || inp.Password == ""{
		h.AppLogger.Info("not all fields are filled")
		utils.HandleError(w, errors.New("not all fields are filled"), 400)
		return 
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	id, err := h.Clients.Login(ctx, inp)
	if err != nil {
		h.AppLogger.Info("error in sending request to user service: %v", err)
		utils.HandleError(w, err, 400)
		return 
	}

	utils.SendJSON(w, map[string]string{"user_id": id}, 200)
}

func (h *Handler) userSignIn(w http.ResponseWriter, r *http.Request) {
	var inp dto.SignIn

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.AppLogger.Info("error in reading request body: %v", err)
		utils.HandleError(w, err, 400)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(bytes, &inp); err != nil {
		h.AppLogger.Info("error in unmarshaling: %v", err)
		utils.HandleError(w, err, 400)
		return 
	}

	if inp.Email == "" || inp.Password == "" {
		h.AppLogger.Info("not all fields are filled")
		utils.HandleError(w, errors.New("not all fields are filled"), 400)
		return 
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	id, err := h.Clients.Login(ctx, inp)
	if err != nil {
		h.AppLogger.Info("error in sending request to user service: %v", err)
		utils.HandleError(w, err, 400)
		return 
	}

	utils.SendJSON(w, map[string]string{"user_id": id}, 200)
}