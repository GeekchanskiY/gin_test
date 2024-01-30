package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize the Gin router
    r := gin.Default()
    r.LoadHTMLGlob("templates/*")
    // Define a route for serving HTML
    r.GET("/", func(c *gin.Context) {
        // Return "Hello, World!" HTML page
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Hello, World!",
        })

//c.String(http.StatusOK, "Hello, World!")
    })

    // Run the server on port 8080
    r.Run(":8080")
}

