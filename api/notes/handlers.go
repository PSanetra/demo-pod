package notes

import "github.com/gin-gonic/gin"

func AddHandlers(
	router *gin.RouterGroup,
	settings *Settings,
) {
	AddGetNotesHandler(router, settings)
	AddPutNotesHandler(router, settings)
}
