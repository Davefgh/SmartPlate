package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"vehicle-api/internal/models"
	"vehicle-api/internal/repository"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo *repository.UserRepository
}
func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	rand.Seed(time.Now().UnixNano()) // Initialize random seed
	return &UserHandler{repo: repo}
}

//create handles POST /users
// func (h *UserHandler) CreateUser(c echo.Context) error{
// 	var user models.User
// 	if err := c.Bind(&user); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string] string{"error": "invalid request body"})

// 	}
// 	if user.LAST_NAME == "" || user.FIRST_NAME == "" || user.MIDDLE_NAME == "" || user.EMAIL == "" || user.PASSWORD == "" || user.ROLE == "" || user.STATUS == "" || user.LTO_CLIENT_ID == "" {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
// 	}
// 	if err := h.repo.Create(&user); err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
// 	}
// 	return c.JSON(http.StatusOK, user)
// }
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		log.Printf("CreateUser bind error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// Validate required fields
	if user.LAST_NAME == "" || user.FIRST_NAME == "" || user.EMAIL == "" || user.PASSWORD == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing required fields"})
	}

	// Generate LTO_CLIENT_ID if not provided
	if user.LTO_CLIENT_ID == "" {
		ltoID, err := h.generateUniqueLTOID()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate LTO ID"})
		}
		user.LTO_CLIENT_ID = ltoID
	}

	// Create user with transaction
	if err := h.repo.Create(&user); err != nil {
		log.Printf("CreateUser error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
	}

	return c.JSON(http.StatusCreated, user)
}


// GetAllUsers handles GET /users
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.repo.GetAll()
	if err != nil {
		log.Printf("GetAllUsers error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch users"})
	}
	return c.JSON(http.StatusOK, users)
}

//GetUserByID handles GET /users/:id
func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	user, err := h.repo.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

//GetUserByEmail handles GET /users/email/:email
func (h *UserHandler) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")

	user, err := h.repo.GetByEmail(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

//UpdateUser handles PUT /users/:id
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	user.USER_ID = id
	if err := h.repo.Update(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}

	// Fetch the updated user to return it with updated timestamp
	updatedUser, err := h.repo.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch updated user"})
	}

	return c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles DELETE /users/:id
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
	}

	if err := h.repo.Delete(id); err != nil {
		log.Printf("DeleteUser error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete user"})
	}

	return c.NoContent(http.StatusNoContent)
}


// UpdateUserByLTO handles PUT /users/by-lto/:lto_client_id
func (h *UserHandler) UpdateUserByLTO(c echo.Context) error {
    ltoID := c.Param("lto_client_id")
    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
    }
    
    existingUser, err := h.repo.GetByLTOClientID(ltoID)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }
    
    user.USER_ID = existingUser.USER_ID
    if err := h.repo.Update(&user); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
    }
    
    return c.JSON(http.StatusOK, user)
}

// DeleteUserByLTO handles DELETE /users/by-lto/:lto_client_id
func (h *UserHandler) DeleteUserByLTO(c echo.Context) error {
    ltoID := c.Param("lto_client_id")
    user, err := h.repo.GetByLTOClientID(ltoID)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }
    
    if err := h.repo.Delete(user.USER_ID); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
    }
    
    return c.NoContent(http.StatusNoContent)
}
//get user by lto client id
func (h *UserHandler) GetUserByLTOID(c echo.Context) error {
    ltoID := c.Param("lto_client_id")
    user, err := h.repo.GetByLTOClientID(ltoID)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }
    return c.JSON(http.StatusOK, user)
}

// GenerateLTOID handles GET /generate-lto-id
//23-041120-3925000
func (h *UserHandler) GenerateLTOID(c echo.Context) error {
	ltoID, err := h.generateUniqueLTOID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"lto_client_id": ltoID,
	})
}

// Helper function with proper 15-digit generation
func (h *UserHandler) generateUniqueLTOID() (string, error) {
	const (
		prefix      = "25" // 2-digit prefix 25 for 2025
		totalLength = 15
		maxAttempts = 10
	)
	remainingDigits := totalLength - len(prefix)

	for i := 0; i < maxAttempts; i++ {
		// Generate random numbers with leading zeros
		randomPart := fmt.Sprintf("%0*d", remainingDigits, rand.Intn(1e13))
		generatedID := prefix + randomPart

		// Check uniqueness
		_, err := h.repo.GetByLTOClientID(generatedID)
		if err != nil {
			// If not found, return the unique ID
			return generatedID, nil
		}
	}
	return "", fmt.Errorf("failed to generate unique LTO ID after %d attempts", maxAttempts)
}