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
	Note float64  `json:note`
}
 var not = errors.New("Id does not exit")
// var empty = errors.New("Empty JSON")
// var wrong = errors.New("Something went Wrong.")

var slc = []info{
	{Id: "2", Name: "Mohammed", Age: 18, Note:13.5},
	{Id: "1", Name: "Ali", Age: 19, Note:14.5},
}

func Getinfo(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, slc)
}
func Createinfo(ctx *gin.Context) {
	var NewInfo info
	err := ctx.BindJSON(&NewInfo)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error":"wrong"})
		return
	}	

	slc = append(slc, NewInfo)
	ctx.IndentedJSON(http.StatusOK, slc)
}
func infobyId(ctx *gin.Context) {
	id := ctx.Param("id")
	info, err := GetinfobyId(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error":"not" })
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
func checkoutInfo(ctx *gin.Context){
	id, err := ctx.GetQuery("id")
	if !err  {
		ctx.IndentedJSON(http.StatusBadRequest,gin.H{"message": "cannot check this info"} )
		return
	}
	info, er := GetinfobyId(id)
	if er != nil {

		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}
	 if info.Note <= 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Inlogic"})
		
info.Note *= -1
		return	
	}

	ctx.IndentedJSON(http.StatusOK,info)


}

func main() {
	server := gin.Default()
	server.GET("/info", Getinfo)
	server.GET("/info/:id", infobyId)
	server.POST("/info", Createinfo)
	server.PATCH("/checkout", checkoutInfo)
	server.Run(":8080")
}
