package controllers

import (
    "gin-crud-project/db"
    "gin-crud-project/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetUsers(c *gin.Context) {
    rows, err := db.DB.Query("SELECT id, name, email FROM users")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        users = append(users, user)
    }

    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := db.DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")

    // buscar el usuario actual
    var currentUser models.User
    err := db.DB.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&currentUser.ID, &currentUser.Name, &currentUser.Email)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }

    var updates map[string]interface{}
    if err := c.ShouldBindJSON(&updates); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // aplicar solo las actualizaciones que se pasen
    if name, ok := updates["name"].(string); ok {
        currentUser.Name = name
    }
    if email, ok := updates["email"].(string); ok {
        currentUser.Email = email
    }

    _, err = db.DB.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", currentUser.Name, currentUser.Email, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, currentUser)
}


func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    _, err := db.DB.Exec("DELETE FROM users WHERE id=$1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
