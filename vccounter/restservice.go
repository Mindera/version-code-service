package vccounter

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"net/http"
	"strconv"
)

const (
	AppIdPathParameter       = "app-id"
	VersionCodePathParameter = "code"
)

type AppVersionCodeDataStore interface {
	UpdateAppVersionCode(ac *AppCode) error
	CurrentAppVersionCode(appId string) (*AppCode, error)
	DeleteApp(appId string) error
}

type VersionCodeService struct {
	VcStorage AppVersionCodeDataStore
}

func (vc VersionCodeService) Register() {

	ws := new(restful.WebService)

	ws.
		Path("/versionCode").
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET(fmt.Sprintf("/{%s}", AppIdPathParameter)).To(vc.currentVersionCode))
	ws.Route(ws.DELETE(fmt.Sprintf("/{%s}", AppIdPathParameter)).To(vc.removeApp))
	ws.Route(ws.PUT(fmt.Sprintf("/{%s}/{%s}", AppIdPathParameter, VersionCodePathParameter)).To(vc.updateVersionCode))
	ws.Route(ws.GET(fmt.Sprintf("/{%s}/next", AppIdPathParameter)).To(vc.nextVersionCode))

	restful.Add(ws)
}

func (vc VersionCodeService) currentVersionCode(request *restful.Request, response *restful.Response) {

	appId := request.PathParameter(AppIdPathParameter)
	ac, err := vc.VcStorage.CurrentAppVersionCode(appId)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteAsJson(ac)
}

func (vc VersionCodeService) nextVersionCode(request *restful.Request, response *restful.Response) {

	appId := request.PathParameter(AppIdPathParameter)
	ac, err := vc.VcStorage.CurrentAppVersionCode(appId)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	ac.NextVersionCode()
	err = vc.VcStorage.UpdateAppVersionCode(ac)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteAsJson(ac)
}

func (vc VersionCodeService) updateVersionCode(request *restful.Request, response *restful.Response) {

	appId := request.PathParameter(AppIdPathParameter)
	code, err := strconv.Atoi(request.PathParameter(VersionCodePathParameter))

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	ac := AppCode{appId, code}
	err = vc.VcStorage.UpdateAppVersionCode(&ac)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteAsJson(ac)
}

func (vc VersionCodeService) removeApp(request *restful.Request, response *restful.Response) {

	appId := request.PathParameter(AppIdPathParameter)

	err := vc.VcStorage.DeleteApp(appId)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}
