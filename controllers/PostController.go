package controllers

import (
	"GO_CRUD_API/initializers"
	"GO_CRUD_API/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostsCreate(c *gin.Context) {

	var X struct {
		Title string
		Body  string
	}

	c.Bind(&X)

	post := models.Post{Title: X.Title, Body: X.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {

	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func GetPost(c *gin.Context) {
	// Récupérer l'ID du paramètre de la requête
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}

func UpdatePost(c *gin.Context) {
	// Récupérer l'ID du paramètre de la requête
	id := c.Param("id")

	var X struct {
		Title string
		Body  string
	}

	c.Bind(&X)

	var post models.Post

	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: X.Title, Body: X.Body})

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {
	// Récupérer l'ID du paramètre de la requête
	id := c.Param("id")

	// Rechercher l'enregistrement pour vérifier s'il existe
	var post models.Post
	result := initializers.DB.Unscoped().First(&post, id)

	// Si l'enregistrement n'existe pas, renvoyer une erreur
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "Post not found"})
		} else {
			c.JSON(500, gin.H{"error": "Failed to query the database"})
		}
		return
	}

	// pour l'archiver ctt(remplir deleteAt dans bdd)
	// initializers.DB.Delete(&models.Post{}, id)

	// Suppression définitive de l'enregistrement
	result_delete := initializers.DB.Unscoped().Delete(&models.Post{}, id)

	// Vérifier si l'opération s'est bien déroulée
	if result_delete.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete post"})
		return
	}

	// Retourner le statut de réussite
	c.Status(200)
}
