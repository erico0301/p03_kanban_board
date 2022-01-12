package server

import (
	"p03_kanban_board/controller/controllertask"
	"p03_kanban_board/controller/controlleruser"
	_ "p03_kanban_board/docs"
	"p03_kanban_board/middleware"
	"p03_kanban_board/repository/repositorytask"
	"p03_kanban_board/repository/repositoryuser"
	"p03_kanban_board/service/servicetask"
	"p03_kanban_board/service/serviceuser"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {

	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	repoTask := repositorytask.New(db)
	srvTask := servicetask.New(repoTask)
	ctrlTask := controllertask.New(srvTask)

	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("/update-account", middleware.Authorization, ctrlUser.Update)
	routeUser.DELETE("/delete-account", middleware.Authorization, ctrlUser.DeleteByID)

	// route task
	r.POST("tasks", middleware.Authorization, ctrlTask.Create)
	r.GET("tasks", middleware.Authorization, ctrlTask.Gets)
	r.PUT("tasks/:taskID", middleware.Authorization, ctrlTask.Update)
	r.PATCH("tasks/update-status/:taskID", middleware.Authorization, ctrlTask.UpdateStatus)
	r.PATCH("tasks/update-category/:taskID", middleware.Authorization, ctrlTask.UpdateCategory)
	r.DELETE("tasks/:taskID", middleware.Authorization, ctrlTask.Delete)

	// routing docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
