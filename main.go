package main

import (
    "context"
    "github.com/Mhoares/Sinoalice_Nightmare_Scheduler_BE/auth"
    "github.com/Mhoares/Sinoalice_Nightmare_Scheduler_BE/nightmare"
    "github.com/Mhoares/Sinoalice_Nightmare_Scheduler_BE/rankings"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "net/http"

    "time"
)
func preflight(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", nightmare.Origin)
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
    c.JSON(http.StatusOK, struct{}{})
}
func main() {

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer nightmare.Mongo.Client.Disconnect(ctx)
    defer rankings.RankRepo.Client.Disconnect(ctx)
    defer cancel()
    auth := new(auth.Service)
    if  err := auth.Init(); err != nil{
        println(err.Error())
        return
    }
    sinoDB := new(nightmare.SinoaliceDBService)
    ns := new(nightmare.Service)
    ns.Init(nightmare.Mongo,sinoDB,auth)
    rs := rankings.Service{Repo: rankings.RankRepo, Auth: auth}
    //gin.SetMode(gin.ReleaseMode)
    r := gin.New()
    nightmares := r.Group("/nightmares")
    {
        nightmares.GET("", ns.GetNightmares())
        nightmares.GET("image",ns.GetImageDataURL())
        nightmares.POST("update", ns.UpdateNightmares())
        nightmares.OPTIONS("", preflight)
    }
    rankings := r.Group("/rank")
    {
        rankings.GET("",rs.Get())
        rankings.POST("update", rs.UpdateRanks())
        rankings.OPTIONS("", preflight)
    }

    viper.SetConfigFile("config.json")
    if err := viper.ReadInConfig(); err != nil{
        println(err.Error())
    }
    if err := r.Run(":"+viper.Get("PORT").(string)); err != nil {
        println(err)
    }

}
