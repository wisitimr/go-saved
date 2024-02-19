package handler

import (
	"encoding/json"
	"net/http"
	mDaybookDetail "saved/http/rest/internal/model/daybook_detail"
	mHandler "saved/http/rest/internal/model/handler"
	mRes "saved/http/rest/internal/model/response"
	mService "saved/http/rest/internal/model/service"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type DaybookDetailHandler struct {
	DaybookDetailService mService.DaybookDetailService
	logger               *logrus.Logger
	mRes.ResponseDto
}

func InitDaybookDetailHandler(DaybookDetailService mService.DaybookDetailService, logger *logrus.Logger) mHandler.DaybookDetailHandler {
	return DaybookDetailHandler{
		DaybookDetailService: DaybookDetailService,
		logger:               logger,
	}
}

func (h DaybookDetailHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	DaybookDetails, err := h.DaybookDetailService.FindAll(r.Context(), r.URL.Query())
	if err != nil {
		h.Respond(w, r, err, 0)
		return
	}
	h.Respond(w, r, DaybookDetails, http.StatusOK)
}

func (h DaybookDetailHandler) FindById(w http.ResponseWriter, r *http.Request) {
	DaybookDetail, err := h.DaybookDetailService.FindById(r.Context(), chi.URLParam(r, "id"))
	if err != nil {
		h.Respond(w, r, err, 0)
		return
	}
	h.Respond(w, r, DaybookDetail, http.StatusOK)
}

func (h DaybookDetailHandler) Create(w http.ResponseWriter, r *http.Request) {
	DaybookDetailPayload := mDaybookDetail.DaybookDetail{}
	err := json.NewDecoder(r.Body).Decode(&DaybookDetailPayload)
	if err != nil {
		h.Respond(w, r, err, 0)
		return
	}
	res, err := h.DaybookDetailService.Create(r.Context(), DaybookDetailPayload)
	if err != nil {
		h.Respond(w, r, err, 0)
		return
	}
	h.Respond(w, r, res, http.StatusCreated)
}

func (h DaybookDetailHandler) Update(w http.ResponseWriter, r *http.Request) {
	DaybookDetailPayload := mDaybookDetail.DaybookDetail{}
	err := json.NewDecoder(r.Body).Decode(&DaybookDetailPayload)
	if err != nil {
		h.Respond(w, r, err, 0)
		return
	}
	res, err := h.DaybookDetailService.Update(r.Context(), chi.URLParam(r, "id"), DaybookDetailPayload)
	if err != nil {
		h.Respond(w, r, err, 0)
		return
	}
	h.Respond(w, r, res, http.StatusOK)
}
