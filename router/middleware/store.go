package middleware

import (
	"github.com/TeaMeow/KitSvc/store"
	"github.com/TeaMeow/KitSvc/store/datastore"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

func Store(c *cli.Context) gin.HandlerFunc {
	v := datastore.Open(
		c.String("database-user"),
		c.String("database-password"),
		c.String("database-host"),
		c.String("database-name"),
		c.String("database-charset"),
		c.Bool("database-parseTime"),
		c.String("database-loc"),
	)

	return func(c *gin.Context) {
		store.ToContext(c, v)
		c.Next()
	}
}