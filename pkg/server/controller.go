package server

import (
	"MEND/pkg/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//go:generate mockery --name=UserRepository --inpackage --filename=controller.mock.go
//
// UserRepository provides functionality to manager User entity.
//
// NOTE: I had discussion some time ago with my coworkers about place where interfaces should be placed.
// There was 2 approaches: 1. Wherever caller uses interface 2. Wherever some shared funcionalities could be abstracted.
// Here I used 1st approach - we concluded that it's cleaner way of doing so and prevent import cycles.
// Doing so in 2nd way would be defining it in "repository" package and use it here as "type UserController struct {client repository.Repository}".
type UserRepository interface {
	// Create creates new user.
	Create(ctx context.Context, user models.User) error

	// Get gets created user.
	Get(ctx context.Context, id int) (*models.User, error)

	// Update updates existing user.
	// Note: We pass id and whole User entity as there is no constraint that id cannot be updated.
	Update(ctx context.Context, id int, user models.User) error

	// Delete deletes existing user.
	Delete(ctx context.Context, id int) error
}

// UserController provides http handlers for gin router.
type UserController struct {
	// NOTE: Controller should probably contain logger (eg. ZeroLog) etc. omitted for simplicity.
	repo UserRepository
}

// NewUserController returns new UserController instance.
func NewUserController(repo UserRepository) *UserController {
	return &UserController{repo: repo}
}

// Create takes user from request body and inserts it in repository.
func (u UserController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid payload",
			"details": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	if err := u.repo.Create(ctx, user); err != nil {
		/*
			NOTE: Errors from repository can be read here like this:

			if errors.Is(err, repository.Duplicate) {
				c.JSON(http.StatusBadRequest, gin.H{
				"error":   "user with given id already exists",
				"details": err.Error(),
			})
			return
			}

			This is omitted for simplicity.

			Additionally errors from another layers should be mapped for user friendly messages.
		*/

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "unknown error",
			"details": err.Error(),
		})
		return
	}

	// NOTE: Because of sql and no-sql implementation we assumed that ID of user will be passed
	// from user. In real sql scenario it would be auto-incremented or UUID id that should
	// be passed back to the user in a response.
	c.Status(http.StatusNoContent)
}

// Get takes id from request and gets user from repository.
func (u UserController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID passed in request"})
		return
	}

	ctx := c.Request.Context()
	user, err := u.repo.Get(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "unknown error",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update takes id and whole user from request and updates it in repository.
func (u UserController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID passed in request"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid payload",
			"details": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	if err := u.repo.Update(ctx, id, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "unknown error",
			"details": err.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}

// Delete takes id param from request and deletes user from repository.
func (u UserController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID passed in request"})
		return
	}

	ctx := c.Request.Context()
	if err := u.repo.Delete(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "unknown error",
			"details": err.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}
