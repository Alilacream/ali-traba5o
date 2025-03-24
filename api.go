package main

import "github.com/gin-gonic/gin"
import "net/http"

// import "os"
import "errors"

// import "fmt"
type info struct {
	Id   string `json:id`
	Name string `json:name`
	Age  uint   `json:age`
}
var not = errors.New("Id does not exit")
var empty = errors.New("Empty JSON")
var wrong = errors.New("Something went Wrong.")

var slc = []info{
	{Id: "2", Name: "Mohammed", Age: 18},
	{Id: "1", Name: "Ali", Age: 19},
}

func Getinfo(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, slc)
}
func Createinfo(ctx *gin.Context) {
	var NewInfo info
	err := ctx.BindJSON(&NewInfo)
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{"error": empty})
		return
	}
	slc = append(slc, NewInfo)
	ctx.IndentedJSON(http.StatusOK, slc)
}
func infobyId(ctx *gin.Context) {
	id := ctx.Param("id")
	info, err := GetinfobyId(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{"error":not })
		return
	}
	ctx.IndentedJSON(http.StatusOK, info)
}
func GetinfobyId(id string) (*info, error) {
	
	for i, ele := range slc {
		if ele.Id == id {
			return &slc[i], nil
		}
	}
	return nil, not
}

func main() {
	server := gin.Default()
	server.GET("/info", Getinfo)
	server.GET("/info/:id", infobyId)
	server.POST("/info", Createinfo)
	server.Run(":8080")
}
