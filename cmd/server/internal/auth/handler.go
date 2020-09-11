package auth

import (
	"net/http"
	"time"

	"github.com/Mockturnal/voting-app-backend/cmd/server/internal/user"
	"github.com/Mockturnal/voting-app-backend/pkg/bcrypt"
	"github.com/Mockturnal/voting-app-backend/pkg/cache"
	"github.com/Mockturnal/voting-app-backend/pkg/jwt"
	gojwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Verify(c *gin.Context)
}

type Services struct {
	Db         *pg.DB
	Cache      cache.Redis
	JWTService jwt.JWT
}

func NewAuthService(db *pg.DB, cache cache.Redis, jwtService jwt.JWT) AuthController {
	return &Services{
		Db:         db,
		Cache:      cache,
		JWTService: jwtService,
	}
}

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type RegisterRequest struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

// @Description Login endpoint that returns a new access token
// @Accept json
// @Produce json
// @Success 200 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /auth/login	[post]
func (c *Services) Login(ctx *gin.Context) {
	var body LoginRequest
	_ = ctx.Bind(&body)

	user := new(user.User)

	if err := c.Db.Model(user).Where("email = ?", body.Email).First(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email / password. Please try again",
		})
		return
	}

	matched := bcrypt.CompareHash(user.Password, body.Password)
	if !matched {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email / password. Please try again",
		})
		return
	}

	token, _ := c.JWTService.GenerateToken(user.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully Logged in!",
		"token":   token,
	})
}

// @Description Register endpoint that makes a new user
// @Accept json
// @Produce json
// @Success 200 {object} gin.H
// @Failure 409 {object} gin.H
// @Router /auth/register [post]
func (c *Services) Register(ctx *gin.Context) {
	var body RegisterRequest
	_ = ctx.Bind(&body)

	existingUser := new(user.User)

	if err := c.Db.Model(&existingUser).Where("email = ?", body.Email).Select(); err == nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": "User already exists with that email",
		})
		return
	}

	hash, _ := bcrypt.GenerateHash(body.Password)

	newUser := &user.User{
		Username:  body.Username,
		Email:     body.Email,
		Password:  hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := c.Db.Model(newUser).Insert()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "User already exists with that email",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered a new account!",
	})
}

// @Description Verify's the user auth token and returns the user's data
// @Accept json
// @Produce json
// @Success 200 {object} user.User
// @Failure 409 {object} gin.H
// @Router /auth/verify [get]
func (c *Services) Verify(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")

	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, Please login and try again",
		})
		return
	}
	bearerToken := header[len("Bearer "):]
	token, err := c.JWTService.Validate(bearerToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, Please login and try again",
		})
		return
	}

	if !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, Please login and try again",
		})
	}

	user := new(user.User)

	if err := c.Db.Model(user).Where("id = ?", token.Claims.(gojwt.MapClaims)["UserId"]).Select(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, Please login and try again",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "You are authorized",
		"user":    user,
	})
}
