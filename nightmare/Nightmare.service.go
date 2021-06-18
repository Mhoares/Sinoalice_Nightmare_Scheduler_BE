package nightmare

import (
	"github.com/Mhoares/Sinoalice_Nightmare_Scheduler_BE/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	Repo 	Repository
	SinoDB *SinoaliceDBService
	Auth   *auth.Service
}

func (sn *Service) Init( r Repository, sdb *SinoaliceDBService, a *auth.Service)  {
	sn.Auth = a
	sn.SinoDB  = sdb
	sn.Repo = r
}
func (sn *Service) GetNightmares() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		httpStatus := http.StatusOK
		nms, err := sn.Repo.GetAll()
		if err != nil {
			httpStatus = http.StatusInternalServerError
		}
		ctx.JSON(httpStatus,nms)
	}
}
type UpdateRequest struct {
	Version string `json:"version"`
	Password string `json:"password"`
}

func (sn *Service) UpdateNightmares() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var (
			ur UpdateRequest
			nms []*Nightmare
		)
		httpStatus := http.StatusOK
		err := ctx.BindJSON(&ur)
		if  err != nil{
			httpStatus = http.StatusBadRequest
			ctx.JSON(httpStatus, err.Error())
			return
		}
		if ok, err := sn.Auth.CheckPassword(ur.Password); err != nil || !ok{
			httpStatus = http.StatusUnauthorized
			ctx.JSON(httpStatus, struct{}{})
			return
		}
		nms, err = sn.SinoDB.GetNightmares(ur.Version)
		if err != nil {
			httpStatus = http.StatusServiceUnavailable
			ctx.JSON(httpStatus, err.Error())
			return
		}
		err = sn.Repo.SaveNightmares(nms)
		if err != nil {
			httpStatus = http.StatusInternalServerError
			ctx.JSON(httpStatus, err.Error())
			return
		}
		ctx.JSON(httpStatus, struct{}{})

	}
}
