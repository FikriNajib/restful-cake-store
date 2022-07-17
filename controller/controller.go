package controller

import (
	"restful-cake-store/config"
	"restful-cake-store/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func DetailOfCake(c *gin.Context) {
	db := config.ConnectDB()
	var (
		cake model.Cake
		result gin.H
	)
	id := c.Param("id")
	row := db.QueryRow("select * from cake where id = ?", id)
	err := row.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.Created_at, &cake.Updated_at)
	if err == sql.ErrNoRows {
		log.Printf("Record Not Found",err.Error())
		result = gin.H{
			"statusCode": http.StatusInternalServerError,
			"result": "Record Not Found",
			"count":  0,
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else if (err != nil) {
		log.Printf("Response Failed",err.Error())
		result = gin.H{
			"statusCode": http.StatusInternalServerError,
			"result": nil,
			"count":  0,
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}else {
		result = gin.H{
			"statusCode": http.StatusOK,
			"result": cake,
			"count":  1,
		}
		c.JSON(http.StatusOK, result)
		return
	}
	return
}

func ListOfCakes(c *gin.Context) {
	db := config.ConnectDB()
	var (
		cake  model.Cake
		cakes []model.Cake
	)
	rows, err := db.Query("select * from cake order by rating desc, title")
	if err != nil {
		log.Printf("Query error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"statusCode" : 500,
			"message": fmt.Sprintf("Response Error"),
		})
		return
	}
	for rows.Next() {
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.Created_at, &cake.Updated_at)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode" : 500,
				"result" : "Response Failed",
			})
			return
		}
		cakes = append(cakes, cake)
		if err != nil {
			log.Printf("Build response error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode" : 500,
				"message": fmt.Sprintf("Response Error"),
			})
			return
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"result": cakes,
		"count":  len(cakes),
	})
	return
}

func AddCake(c *gin.Context) {
	db := config.ConnectDB()
	var input model.InputCake
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"statusCode": 400,
			"message": fmt.Sprintf("Invalid Request"),
		})
		log.Printf("Invalid Request", err.Error())
		return
	}
	input.Created_at = time.Now()

	_, err = db.Exec("INSERT INTO cake(title, description, rating, image, created_at) VALUES(?, ?, ?, ?, ?)", input.Title, input.Description, input.Rating, input.Image, input.Created_at)
	if err != nil {
		log.Printf("Query error", err.Error())
		c.JSON(http.StatusInternalServerError,gin.H{
			"statusCode": 500,
			"message": fmt.Sprintf("Failed create cake"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"statusCode" : 200,
		"message": fmt.Sprintf(" Successfully created"),
	})
	return
}

func UpdateCake(c *gin.Context) {
	db := config.ConnectDB()
	id := c.Param("id")
	var input model.InputCake
	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.Printf("Invalid Request", err.Error())
		c.JSON(http.StatusBadRequest,gin.H{
			"statusCode": 400,
			"message": fmt.Sprintf("Invalid Request"),
		})
		return
	}

	input.Updated_at = time.Now()
	_, err = db.Exec("update cake set title= ?, description= ? , rating= ?, image= ?, updated_at= ? where id= ?",input.Title, input.Description, input.Rating, input.Image, input.Updated_at, id)
	if err == sql.ErrNoRows {
		log.Printf("Record Not Found",err.Error())
		c.JSON(http.StatusInternalServerError,gin.H{
			"statusCode": 500,
			"message": fmt.Sprintf("Failed updated ,id:%s not found",id),
		})
		return
	} else if err != nil {
		log.Printf("Query error", err.Error())
		c.JSON(http.StatusInternalServerError,gin.H{
			"statusCode": 500,
			"message": fmt.Sprintf("Failed updated id:%s , Internal Error",id),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": 200,
			"message":    fmt.Sprintf("Successfully updated id: %s", id),
		})
	}
	return
}

func DeleteCake(c *gin.Context) {
	db := config.ConnectDB()
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Failed to Parsing", err.Error())
		return
	}
	_, err = db.Exec("delete from cake where id= ?", idInt)
	if err == sql.ErrNoRows {
		log.Printf("Record Not Found", err.Error())
		c.JSON(http.StatusInternalServerError,gin.H{
			"statusCode": 500,
			"message" : fmt.Sprintf("Delete Failed, id : %s not found", id),
		})
		return
	} else if (err != nil) {
		log.Printf("Query error", err.Error())
		c.JSON(http.StatusInternalServerError,gin.H{
			"statusCode": 500,
			"message" : fmt.Sprintf("Delete Failed at id : %s", id),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted id: %s", id),
	})
	return
}