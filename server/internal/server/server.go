package server

import (
	"strings"

	"github.com/Sergio-Na/argus/server/config"
	"github.com/Sergio-Na/argus/server/internal/auth"
	"github.com/Sergio-Na/argus/server/internal/database"
	"github.com/Sergio-Na/argus/server/internal/handler"
	"github.com/Sergio-Na/argus/server/internal/supabase"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	config     *config.Config
	supaClient *supabase.SupabaseClient
	auth       *auth.Service
	db         *database.Service
	handler    *handler.Handler
}

func New(cfg *config.Config, supaClient *supabase.SupabaseClient) (*Server, error) {
	authService := auth.NewService(supaClient)
	dbService := database.NewService(supaClient)
	handlerService := handler.NewHandler(authService, dbService)

	s := &Server{
		router:     gin.Default(),
		config:     cfg,
		supaClient: supaClient,
		auth:       authService,
		db:         dbService,
		handler:    handlerService,
	}

	s.registerRoutes()
	return s, nil
}

func (s *Server) Run() error {
	return s.router.Run(s.config.ServerAddress)
}

func (s *Server) registerRoutes() {
	s.router.GET("/", s.handler.Home)
	s.router.POST("/signup", s.handler.SignUp)
	s.router.POST("/signin", s.handler.SignIn)

	authorized := s.router.Group("/")
	authorized.Use(s.authMiddleware())
	{
		authorized.GET("/user", s.handler.GetUser)
	}
}

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		token := bearerToken[1]

		user, err := s.auth.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
