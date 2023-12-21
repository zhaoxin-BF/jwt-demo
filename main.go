package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhaoxin-BF/jwt-demo/pkg"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/set", setting)
	r.GET("/get", getting)

	r.Run(":8080")
}

// set
func setting(ctx *gin.Context) {
	token, err := pkg.GenToken(100100100, "beifeng")
	if err != nil {

	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// get
func getting(ctx *gin.Context) {
	tokenstr := ctx.GetHeader("Authorization")

	if tokenstr == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "no auth"})
		return
	}

	claims, err := pkg.ParseToken(tokenstr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "no auth"})
		return
	}

	fmt.Println(claims)
	fmt.Println(claims.UserID)
	fmt.Println(claims.UserName)
}
