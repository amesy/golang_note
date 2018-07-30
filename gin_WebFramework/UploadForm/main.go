package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"static": "Ok!",
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request.")
			return
		}

		filename := header.Filename

		// 将文件上传到指定的upload目录下.
		out, err := os.Create("upload/" + filename)
		if err != nil {
			log.Fatal(err)
		}

		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusOK, "Upload successful.")

	})

	r.POST("/multi/upload", func(c *gin.Context) {
		err := c.Request.ParseMultipartForm(1000000)
		if err != nil {
			log.Fatal(err)
		}

		formdata := c.Request.MultipartForm
		files := formdata.File["upload_file"]
		fmt.Println("files: ", files)
		for i, _ := range files {
			file, err := files[i].Open()
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			out, err := os.Create("upload/" + files[i].Filename)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
			}
			c.String(http.StatusCreated, "Upload all successful.")
		}

	})

	r.Run(":8088")
}
