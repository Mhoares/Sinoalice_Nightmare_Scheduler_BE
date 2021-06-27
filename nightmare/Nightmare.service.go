package nightmare

import (
	"github.com/Mhoares/Sinoalice_Nightmare_Scheduler_BE/auth"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

var Origin ="*"
type Service struct {
	Repo 	Repository
	SinoDB *SinoaliceDBService
	Auth   *auth.Service
}

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil{
		println(err.Error())
	}
	Origin = viper.Get("ORIGIN").(string)
	println(Origin)
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
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", Origin)
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
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", Origin)
		ctx.JSON(httpStatus, struct{}{})

	}
}
type GetImageDataURL struct {
	  Icon string `form:"icon"`
}
func (sn *Service) GetImageDataURL() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var (
			gidURL GetImageDataURL
			imgBase64 string
		)
		httpStatus := http.StatusOK
		err := ctx.Bind(&gidURL)
		if  err != nil{
			httpStatus = http.StatusBadRequest
			ctx.JSON(httpStatus, err.Error())
			return
		}
		imgBase64, err = sn.SinoDB.GetImageDataUrl(gidURL.Icon)
		if err != nil {
			httpStatus = http.StatusServiceUnavailable
			ctx.JSON(httpStatus, err.Error())
			return
		}
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", Origin)
		ctx.JSON(httpStatus,imgBase64)
	}
}
