package http

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/mehmetaligok/venom-example-project/src/model"

	"github.com/gin-gonic/gin"
)

type UserServer struct {
	userRepo UserRepo
}

type UserRepo interface {
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
}

// NewUserServer returns new user server instance
func NewUserServer(repo UserRepo) *UserServer {
	return &UserServer{userRepo: repo}
}

func (server *UserServer) GetUserHandler(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Id format is matched, error: %v", err))
		return
	}

	user, err := server.userRepo.GetUser(c, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.String(http.StatusNotFound, "User not found.")
			return
		}

		log.Printf("Error while fetching user: %v \n", err)
		c.String(http.StatusInternalServerError, "Error while fetching user.")
		return
	}

	c.JSON(http.StatusOK, user)
}
