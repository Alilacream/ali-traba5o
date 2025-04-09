package main

import "github.com/gin-gonic/gin"
import "todo-list/helpF"
import "net/http"
import "fmt"

var not =helpF.Error{Msg: "Stats weren't found"} 

var todos = []helpF.Todo{
	{Stat: "1", Todo: "Clean room", Done: true},
	{Stat: "2", Todo: "Salat", Done: true},
	{Stat: "3", Todo: "Learning", Done: false},
}

func Gettodo(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, todos)
}
func Newtodo(ctx *gin.Context) {
	var newtoe helpF.Todo
	err := ctx.BindJSON(&newtoe)
	for _, ele := range todos {
		if newtoe.Stat == ele.Stat {
			err = fmt.Errorf("Cannot give same stat to diffrents todo lists")
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	if err != nil {
		err = fmt.Errorf("Something went wrong")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos = append(todos, newtoe)
	ctx.IndentedJSON(http.StatusOK, todos)

}

func updateTodo(ctx *gin.Context) {
	stats := ctx.Param("stat")
	newstats, err := helpF.Helpertodo(stats,todos)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	for i, ele := range todos {
		if newstats.Stat == ele.Stat  {
			// todos[i].Done = !(todos[i].Done)	
			helpF.Reverse(&todos[i])
			ctx.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": not})
}

func GetbyidTodo(ctx *gin.Context) {
	stats := ctx.Param("stat")
	newstats, err := helpF.Helpertodo(stats,todos)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, newstats)
}

func deleteallTodo(ctx *gin.Context) {
	// var point *helpF.Todo = todos
	stats := ctx.Param("stat")
	for i, ele := range todos {
		if ele.Stat == stats {
			helpF.Free(&todos[i])
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Deleting Succesfully"})
			return
		}
	}

	var txt = fmt.Sprintf("Something went wrong")
	ctx.IndentedJSON(http.StatusOK, gin.H{"error": txt})
}

/*
func deleteTodo(ctx *gin.Context) {

}
*/
func main() {
	server := gin.Default()
	server.GET("/todo", Gettodo)
	server.POST("/todo", Newtodo)
	server.GET("/todo/:stat", GetbyidTodo)
	server.PATCH("/todo/:stat", updateTodo)
	server.DELETE("/todo/:stat", deleteallTodo)
	server.Run()
}
