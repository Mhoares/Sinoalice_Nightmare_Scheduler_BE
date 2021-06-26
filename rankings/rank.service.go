package rankings

import (
	"github.com/Mhoares/Sinoalice_Nightmare_Scheduler_BE/auth"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

var Origin = "*"
type Service struct {
	Auth *auth.Service
	Repo *RankRepository
}
func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil{
		println(err.Error())
	}
	Origin = viper.Get("ORIGIN").(string)
	println(Origin)
}

type UpdateRequest struct {
	Password string `json:"password"`
	Ranks *Ranks	`json:"ranks"`

}
func (sr *Service) UpdateRanks() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var (
			ur UpdateRequest
			updated  = true
		)
		httpStatus := http.StatusOK
		err := ctx.BindJSON(&ur)
		if  err != nil{
			httpStatus = http.StatusBadRequest
			ctx.JSON(httpStatus, err.Error())
			return
		}
		if ok, err := sr.Auth.CheckPassword(ur.Password); err != nil || !ok{
			httpStatus = http.StatusUnauthorized
			ctx.JSON(httpStatus, struct{}{})
			return
		}
		updated, err = sr.Repo.UpdateOrSave(ur.Ranks,ur.Ranks.Day,ur.Ranks.GC)
		if err != nil  {
			httpStatus = http.StatusInternalServerError
			ctx.JSON(httpStatus, err.Error())
			return
		}
		if !updated {
			httpStatus = http.StatusInternalServerError
			ctx.JSON(httpStatus,struct{}{})
			return
		}
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", Origin)
		ctx.JSON(httpStatus, struct{}{})

	}
}
type GetRequest struct {
	Day int `form:"day"`
	GC  int `form:"gc"`
}
func (sr *Service) Get() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var gr GetRequest
		httpStatus := http.StatusOK
		err := ctx.Bind(&gr)
		if  err != nil{
			httpStatus = http.StatusBadRequest
			ctx.JSON(httpStatus, err.Error())
			return
		}
		r, err := sr.Repo.Get(gr.Day,gr.GC)
		if err != nil {
			httpStatus = http.StatusInternalServerError
			ctx.JSON(httpStatus, err.Error())
			return
		}
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", Origin)
		ctx.JSON(httpStatus,r)
	}
}
