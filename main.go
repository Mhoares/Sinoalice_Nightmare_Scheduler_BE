package main

import (
    "context"
    "github.com/Mhoares/Sinoalice_Nightmare_Scheduler_BE/auth"
    "github.com/Mhoares/Sinoalice_Nightmare_Scheduler_BE/nightmare"
    "github.com/gin-gonic/gin"
    "net/http"

    "time"
)
func preflight(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
    c.JSON(http.StatusOK, struct{}{})
}
func main() {

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer nightmare.Mongo.Client.Disconnect(ctx)
    defer cancel()
    auth := new(auth.Service)
    if  err := auth.Init(); err != nil{
        println(err.Error())
        return
    }
    sinoDB := new(nightmare.SinoaliceDBService)
    ns := new(nightmare.Service)
    ns.Init(nightmare.Mongo,sinoDB,auth)
    r := gin.Default()
    nightmares := r.Group("/nightmares")
    {
        nightmares.GET("", ns.GetNightmares())
        nightmares.POST("update", ns.UpdateNightmares())
        nightmares.OPTIONS("", preflight)
    }
    if err := r.Run(":8081"); err != nil {
        println(err)
    }

}
