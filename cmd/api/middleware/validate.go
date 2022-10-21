package middleware

import "github.com/gin-gonic/gin"

func TokenValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.GetHeader("TOKEN")
		// if token == "" {
		// 	web.Failure(c, http.StatusUnauthorized, errors.New("token not found"))
		// 	c.Abort()
		// 	return
		// }
		// if token != os.Getenv("TOKEN") {
		// 	web.Failure(c, http.StatusUnauthorized, errors.New("invalid token"))
		// 	c.Abort()
		// 	return
		// }
		c.Next()
	}
}
