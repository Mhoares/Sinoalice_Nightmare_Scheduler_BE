package analyzer

import (
	"github.com/Mhoares/Sinoalice_Nightmare_Scheduler_BE/auth"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type Service struct {
	Auth *auth.Service
	Blue *BlueService
	Repo *WeaponRepository

}
var Origin = "*"
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
}
func (sw *Service) UpdateWeapons() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var (
			ur UpdateRequest
		)
		httpStatus := http.StatusOK
		err := ctx.BindJSON(&ur)
		if  err != nil{
			httpStatus = http.StatusBadRequest
			ctx.JSON(httpStatus, err.Error())
			return
		}
		if ok, err := sw.Auth.CheckPassword(ur.Password); err != nil || !ok{
			httpStatus = http.StatusUnauthorized
			ctx.JSON(httpStatus, struct{}{})
			return
		}

		weapons, err := sw.Blue.GetWeapons()
		if err != nil {
			httpStatus = http.StatusServiceUnavailable
			ctx.JSON(httpStatus, err.Error())
			return

		}
		err = sw.Repo.SaveWeapons(weapons)
		if err != nil  {
			httpStatus = http.StatusInternalServerError
			ctx.JSON(httpStatus, err.Error())
			return
		}

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", Origin)
		ctx.JSON(httpStatus, struct{}{})

	}
}
type UpdateSupportRequest struct {
	Password string `json:"password"`
}
func (sw *Service) UpdateSupportSkills() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var (
			ur UpdateRequest
		)
		httpStatus := http.StatusOK
		err := ctx.BindJSON(&ur)
		if  err != nil{
			httpStatus = http.StatusBadRequest
			ctx.JSON(httpStatus, err.Error())
			return
		}
		if ok, err := sw.Auth.CheckPassword(ur.Password); err != nil || !ok{
			httpStatus = http.StatusUnauthorized
			ctx.JSON(httpStatus, struct{}{})
			return
		}

		supportsRaw, err := sw.Blue.GetRawSupport()
		if err != nil {
			httpStatus = http.StatusServiceUnavailable
			ctx.JSON(httpStatus, err.Error())
			return

		}
		supportBoon, err := sw.Blue.GetBoostSupport("Support Boon", "C4:F24")
		if err != nil {
			httpStatus = http.StatusServiceUnavailable
			ctx.JSON(httpStatus, err.Error())
			return

		}
		dauntlessCourage, err := sw.Blue.GetBoostSupport("Dauntless Courage", "C4:F24")
		if err != nil {
			httpStatus = http.StatusServiceUnavailable
			ctx.JSON(httpStatus, err.Error())
			return

		}
		recoverySupport, err := sw.Blue.GetBoostSupport("Recovery Support", "C4:F24")
		if err != nil {
			httpStatus = http.StatusServiceUnavailable
			ctx.JSON(httpStatus, err.Error())
			return

		}
		replenishMagic, err := sw.Blue.GetBoostSupport("Replenish Magic", "C4:E24")
		if err != nil {
			httpStatus = http.StatusServiceUnavailable
			ctx.JSON(httpStatus, err.Error())
			return

		}
		supportsRaw = append(supportsRaw, supportBoon[0], recoverySupport[0], dauntlessCourage[0],replenishMagic[0])
		err = sw.Repo.SaveSupportSkills(supportsRaw)
		if err != nil  {
			httpStatus = http.StatusInternalServerError
			ctx.JSON(httpStatus, err.Error())
			return
		}

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", Origin)
		ctx.JSON(httpStatus, struct{}{})

	}
}
type GetRequest struct {
	Name string `form:"name"`
}
func (sw *Service) GetSupportSkillByName() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var gr GetRequest
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", Origin)
		httpStatus := http.StatusOK
		err := ctx.Bind(&gr)
		if  err != nil{
			httpStatus = http.StatusBadRequest
			ctx.JSON(httpStatus, err.Error())
			return
		}
		r, err := sw.Repo.GetSupportSkillByName(gr.Name)
		if err != nil {
			httpStatus = http.StatusInternalServerError
			ctx.JSON(httpStatus, err.Error())
			return
		}

		ctx.JSON(httpStatus,r)
	}
}
func (sw *Service) GetWeapons() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		httpStatus := http.StatusOK

		r, err := sw.Repo.GetAllWeapons()
		if err != nil {
			httpStatus = http.StatusInternalServerError
			ctx.JSON(httpStatus, err.Error())
			return
		}
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", Origin)
		ctx.JSON(httpStatus,r)
	}
}

