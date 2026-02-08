package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/car-project/internal/db"
	"github.com/user/car-project/internal/handlers"
	"github.com/user/car-project/internal/middleware"
	"github.com/user/car-project/internal/repository"
	"github.com/user/car-project/internal/service"
	"github.com/user/car-project/internal/utils"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/user/car-project/docs"
)

func SetupRouter(jwtSecret string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	// In a real professional app, we'd add structured logging middleware here
	r.Use(gin.Logger())

	// Initialize Repositories
	carRepo := repository.NewCarRepository(db.DB)
	userRepo := repository.NewUserRepository(db.DB)
	roleRepo := repository.NewRoleRepository(db.DB)
	permRepo := repository.NewPermissionRepository(db.DB)

	// Initialize Services
	carService := service.NewCarService(carRepo)
	userService := service.NewUserService(userRepo)
	roleService := service.NewRoleService(roleRepo)
	permService := service.NewPermissionService(permRepo)

	// Initialize Handlers
	carHandler := handlers.NewCarHandler(carService)
	userHandler := handlers.NewUserHandler(userService)
	roleHandler := handlers.NewRoleHandler(roleService)
	permHandler := handlers.NewPermissionHandler(permService)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "up",
		})
	})

	// Login route for demo (should be improved with real auth)
	r.POST("/login", func(c *gin.Context) {
		// In a real app, you'd verify credentials here
		token, err := utils.GenerateToken(1, jwtSecret, 24)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	api := r.Group("/api/v1")
	api.Use(middleware.AuthMiddleware(jwtSecret))
	{
		// User CRUD routes
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// Role CRUD routes
		roles := api.Group("/roles")
		{
			roles.POST("", roleHandler.CreateRole)
			roles.GET("", roleHandler.GetRoles)
			roles.GET("/:id", roleHandler.GetRoleByID)
			roles.PUT("/:id", roleHandler.UpdateRole)
			roles.DELETE("/:id", roleHandler.DeleteRole)
			roles.POST("/assign", roleHandler.AssignRole)
		}

		// Permission CRUD routes
		perms := api.Group("/permissions")
		{
			perms.POST("", permHandler.CreatePermission)
			perms.GET("", permHandler.GetPermissions)
			perms.GET("/:id", permHandler.GetPermissionByID)
			perms.PUT("/:id", permHandler.UpdatePermission)
			perms.DELETE("/:id", permHandler.DeletePermission)
		}

		// Car CRUD routes
		cars := api.Group("/cars")
		{
			cars.POST("", middleware.RequirePermission(permService, "car-create"), carHandler.CreateCar)
			cars.GET("", middleware.RequirePermission(permService, "car-read"), carHandler.GetCars)
			cars.GET("/:id", middleware.RequirePermission(permService, "car-read"), carHandler.GetCarByID)
			cars.PUT("/:id", middleware.RequirePermission(permService, "car-update"), carHandler.UpdateCar)
			cars.DELETE("/:id", middleware.RequirePermission(permService, "car-delete"), carHandler.DeleteCar)
		}
	}

	return r
}
