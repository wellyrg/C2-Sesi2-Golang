package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type Book struct{
	ID int			        `json:"id"`	
	Judul string		    `json:"tittle"`	
	Pengarang string		`json:"author"`	
	Deskripsi string 		`json:"desc"`
}

var mapBooks = make(map[int]Book,0) 
var index int = 1

func init(){
	mapBooks[1] = Book{ID: 1, Judul :"Golang", Pengarang: "Gopher", Deskripsi: "A book for GO"}
}

func main() {
	
	g:= gin.Default()
	g.GET("/book",getAllBook)
	g.GET("/book/:id",getBookById)
	g.POST("/book",addBook)
	g.DELETE("/book/:id", deleteBook)
	g.PUT("/book/:id",updateBook)
	g.Run(":8080")

}

func getAllBook(ctx *gin.Context) {
	
   books := make([]Book,0)
   for _, v :=range mapBooks{
	books = append(books, v)
   }



	ctx.JSON(http.StatusOK, books)


}
func getBookById(ctx *gin.Context){
	idString := ctx.Param("id")

	id,err:= strconv.Atoi(idString)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error" : err,
		})
		return
	}

	foundBook, found := mapBooks[id]
	if !found{
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error" : "not found",
		})
		return
	}
	ctx.JSON(http.StatusOK,foundBook)


}
func addBook(ctx *gin.Context){
	var newBook Book

	err :=ctx.ShouldBindJSON(&newBook)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err)
		return

	}else{
		ctx.Writer.Write([]byte("Created"))
	}
	newBook.ID = index
	mapBooks[index] = newBook
	index++
}
func deleteBook(ctx *gin.Context){
	idString := ctx.Param("id")

	id,err:= strconv.Atoi(idString)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error" : err,
		})
		return
	}

	_, found := mapBooks[id]
	if !found{
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error" : "not found",
		})
		return
	}else{
		delete(mapBooks,id)
		ctx.Writer.Write([]byte("Deleted"))

	}

	

}

func updateBook(ctx *gin.Context){
	idString := ctx.Param("id")

	id,err:= strconv.Atoi(idString)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error" : err,
		})
		return
	}

	_, found := mapBooks[id]
	if !found{
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error" : "not found",
		})
		return
	}

	var newUpdateBook Book
	err = ctx.ShouldBindJSON(&newUpdateBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err)
		return
	}else{
		ctx.Writer.Write([]byte("Updated"))
	}
	newUpdateBook.ID = id

	mapBooks[id] = newUpdateBook

	

}