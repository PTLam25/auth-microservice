package polo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	polo = "polo"
)

func Marco(c *gin.Context) {
	// энд поинт для клауд провайдера, чтобы проверять, что наш сервер работает. Если микросервис не работает, но клауд остановит его и запустить новый
	c.String(http.StatusOK, polo)
}
