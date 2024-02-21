package controller

import (
	"net/http"
	"strconv"

	"github.com/be/perpustakaan/exception"
	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
	"github.com/be/perpustakaan/service"
	"github.com/julienschmidt/httprouter"
)

type PenaltiesControllerImpl struct {
	PenaltiesService service.PenaltiesService
}

func NewPenaltiesController(penaltiesService service.PenaltiesService) PenaltiesController {
	return &PenaltiesControllerImpl{
		PenaltiesService: penaltiesService,
	}
}

func (c *PenaltiesControllerImpl) PayPenalties(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		panic(exception.CustomEror{Code: 400, Error: "id must be number"})
	}
	updatePenaltiesRequest := webrequest.UpdatePenaltiesRequest{}
	helper.ReadFromRequestBody(request, &updatePenaltiesRequest)
	update := c.PenaltiesService.PayPenalties(request.Context(), updatePenaltiesRequest, id)

	webResponse := webresponse.ResponseApi{
		Code:   200,
		Status: "OK",
		Data:   update,
	}
	helper.WriteToResponseBody(writer, webResponse)

}
