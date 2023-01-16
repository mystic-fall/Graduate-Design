package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.RecoveryMiddleware())
	v1 := ginRouter.Group("/api/v1")

	v1.POST("/user/login", handler.UserLogin)
	v1.POST("/upload", handler.Upload)
	v1.GET("/download", handler.Download)

	v1.Use(middleware.JWT())
	v1.POST("/user/logout", handler.UserLogout)

	v1.GET("/student", handler.StudentList)
	v1.POST("/student", handler.CreateStu)
	v1.PUT("/student", handler.UpdateStu)
	v1.DELETE("/student", handler.DeleteStu)

	v1.GET("/admin", handler.AdminList)
	v1.POST("/admin", handler.CreateAdmin)
	v1.PUT("/admin", handler.UpdateAdmin)
	v1.DELETE("/admin", handler.DeleteAdmin)

	v1.GET("/dm", handler.DMList)
	v1.POST("/dm", handler.CreateDM)
	v1.PUT("/dm", handler.UpdateDM)
	v1.DELETE("/dm", handler.DeleteDM)

	v1.GET("/building", handler.BuildingList)
	v1.POST("/building", handler.CreateBuilding)
	v1.PUT("/building", handler.UpdateBuilding)
	v1.DELETE("/building", handler.DeleteBuilding)

	v1.GET("/dormitory", handler.DormitoryList)
	v1.POST("/dormitory", handler.CreateDormitory)
	v1.PUT("/dormitory", handler.UpdateDormitory)
	v1.DELETE("/dormitory", handler.DeleteDormitory)

	v1.GET("/live", handler.LiveList)
	v1.POST("/live", handler.CreateLive)
	v1.PUT("/live", handler.UpdateLive)
	v1.DELETE("/live", handler.DeleteLive)
	return ginRouter
}
