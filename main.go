package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string 		`json:"id"`
	Item string		`json:"item"`
	Completed bool	`json:"completed"`
}


var todos = []todo{
	{"1","work",true},
	{"5","sleep",true},
	{"9","play",false},
}

//get data from the server have the data
func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK,todos)
}

func addTodo(context *gin.Context) {
	var new todo
	if err := context.BindJSON(&new); err!=nil {
		return
	}

	todos = append(todos, new)
	context.IndentedJSON(http.StatusCreated,new)
}


//extracting path variable 
func getTodo(context *gin.Context) {
	id:=context.Param("id")
	todo,err:= findTodo(id)
	if err!=nil {
		context.IndentedJSON(http.StatusNotFound,gin.H{"not found":"todo not foind"})
		return
	}
	context.IndentedJSON(http.StatusOK,todo)
}

//func to iterate the struct and find the id
func findTodo(id string)(*todo,error) {
	for i,t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil,errors.New("todos not found")
}

//update the todo
func updateTodo(context *gin.Context) {
	id:=context.Param("id")
	todo,err:= findTodo(id)
	if err!=nil {
		context.IndentedJSON(http.StatusNotFound,gin.H{"not found":"todo not foind"})
		return
	}
	todo.Completed = ! todo.Completed //flip the boolean values
	context.IndentedJSON(http.StatusOK,todo)
}

func removeItem(context *gin.Context){
	id:=context.Param("id")
	todo,err:= findTodo(id)
	if err!=nil {
		context.IndentedJSON(http.StatusNotFound,gin.H{"not found":"todo not foind"})
		return
	}
	todo.Item =""
	context.IndentedJSON(http.StatusOK,todo)

}

func main () {

	//roter is the server 
	router := gin.Default()
	router.GET("/todos",getTodos)
	router.POST("/todos",addTodo)
	router.GET("/todos/:id",getTodo)
	router.PATCH("/todos/:id",updateTodo)
	router.DELETE("/todo/:id",removeItem)
	router.Run("localhost:8080")

}



