package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type producto struct {
    Nombre     string  `json:"nombre"`
    Descripcion  string  `json:"descripcion"`
    Valor float64  `json:"valor"`
    Fecha_de_Expiracion  string `json:"fecha_de_expiracion"`
}

var productos = []producto{
    {Nombre: "Computador", Descripcion: "Computador Lenovo 4 gb ram", Valor: 600000, Fecha_de_Expiracion: "03/03/2023"},
	{Nombre: "Audifonos", Descripcion: "Audifonos beats", Valor: 200000, Fecha_de_Expiracion: "03/03/2023"},
	{Nombre: "Anteojos", Descripcion: "Anteojos RayBan", Valor: 300000, Fecha_de_Expiracion: "03/03/2023"},
}

func getProductos(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, productos)
}

func postProductos(c *gin.Context) {
    var nuevoProducto producto

    if err := c.BindJSON(&nuevoProducto); err != nil {
        return
    }

    productos = append(productos, nuevoProducto)
    c.IndentedJSON(http.StatusCreated, nuevoProducto)
}

func getProductoPorNombre(c *gin.Context) {
    nombre := c.Param("nombre")

    for _, a := range productos {
        if a.Nombre == nombre {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "producto no encontrado"})
}

func eliminarProductoPorNombre(c *gin.Context) {
    nombre := c.Param("nombre")
	

    for index, a := range productos {
        if a.Nombre == nombre {
			productos = append(productos[:index], productos[index+1:]...)
            c.IndentedJSON(http.StatusOK, productos)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "producto no encontrado"})
}

func actualizarProductoPorNombre(c *gin.Context) {
    nombre := c.Param("nombre")
	var nuevoProducto producto

    if err := c.BindJSON(&nuevoProducto); err != nil {
        return
    }

    for index, a := range productos {
        if a.Nombre == nombre {
			productos[index] = nuevoProducto
            c.IndentedJSON(http.StatusOK, productos)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "producto no encontrado"})
}


func main() {
    router := gin.Default()
    router.GET("/productos", getProductos)
	router.POST("/productos", postProductos)
	router.GET("/productos/:nombre", getProductoPorNombre)
	router.PATCH("/productos/:nombre", actualizarProductoPorNombre)
	router.DELETE("/productos/:nombre", eliminarProductoPorNombre)
    router.Run("localhost:3000")
}