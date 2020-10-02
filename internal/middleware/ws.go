package middleware
//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//	"net/http"
//	"server_simple/internal/models"
//	"server_simple/internal/models/protoCompiles"
//	"time"
//)
//
//var Ws wser
//
//type ws struct {
//}
//
//func init() {
//	Ws = &ws{}
//}
//
//type wser interface {
//	HttpUpgrader() gin.HandlerFunc
//}
//
//func (this *ws) HttpUpgrader() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		es := models.SetGinContext(c)
//		wsu := websocket.Upgrader{
//			HandshakeTimeout: time.Duration(time.Second * 30),
//			CheckOrigin: func(r *http.Request) bool {
//				return true
//			},
//		}
//		wsconn, err := wsu.Upgrade(c.Writer, c.Request, nil)
//		if err != nil {
//			es.JsonError(protoCompiles.ErrorCodes_ERROR, err.Error())
//			c.Abort()
//			return
//		}
//		c.Set("wsconn", wsconn)
//		c.Next()
//	}
//}
