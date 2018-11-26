package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var statusCode = 200

// SetStatusCode -- Set status code
func SetStatusCode(statCode int) int {
	statusCode := statCode
	return statusCode
}

// RespondJSON -- set response to json format
func RespondJSON(c *gin.Context, data interface{}) {
	c.JSON(statusCode, data)
	return
}

// RespondCreated -- Set response for create process
func RespondCreated(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Created"
	}
	statusCode := SetStatusCode(201)
	c.JSON(statusCode, gin.H{"message": message})
	return
}

// RespondUpdated -- Set response for update process
func RespondUpdated(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Updated"
	}
	statusCode := SetStatusCode(200)
	c.JSON(statusCode, gin.H{"message": message})
	return
}

// RespondDeleted -- Set response for delete process
func RespondDeleted(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Deleted"
	}
	statusCode := SetStatusCode(200)
	c.JSON(statusCode, gin.H{"message": message})
	return
}

// RespondError -- Set response for error
func RespondError(c *gin.Context, message interface{}, statusCode int) {
	c.JSON(statusCode, gin.H{"errors": message})
	return
}

// RespondFailValidation -- Set response for fail validation
func RespondFailValidation(c *gin.Context, message interface{}) {
	RespondError(c, message, 422)
	return
}

// RespondUnauthorized -- Set response not authorized
func RespondUnauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "Unauthorized"
	}
	statusCode := SetStatusCode(401)
	c.JSON(statusCode, gin.H{"errors": message})
	return
}

// RespondNotFound -- Set response not found
func RespondNotFound(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Not Found"
	}
	statusCode := SetStatusCode(404)
	c.JSON(statusCode, gin.H{"errors": message})
	return
}

// RespondMethodNotAllowed -- Set response method not allowed
func RespondMethodNotAllowed(c *gin.Context, message string) {
	if message == "" {
		message = "Method Not Allowed"
	}
	statusCode := SetStatusCode(405)
	c.JSON(statusCode, gin.H{"errors": message})
	return
}

func RespondRedirect(c *gin.Context, url string) {
	if url == "" {
		return
	}

	c.Redirect(http.StatusFound, url)
	return
}
