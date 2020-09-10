package poll


import (
	"github.com/labstack/echo"
)

func NewPollRoutes(g *echo.Group) {
	// g.GET("", GetPolls, middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey:    []byte(os.Getenv("ACCESS_TOKEN_SECRET")),
	// 	SigningMethod: jwt.SigningMethodHS256.Name,
	// }))

	g.GET("", GetPolls)
}
