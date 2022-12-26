package middle

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UUID(c *gin.Context) {
	rid := uuid.NewString()
	cid := c.GetHeader("client_id")
	if cid == "" {
		cid, _ = c.Cookie("client_id")
	}
	if cid == "" {
		cid = uuid.NewString()
		c.SetCookie("client_id", cid, 86400*100, "/", "", false, true)
	}
	c.Set("rid", rid)
	c.Set("cid", cid)
}
