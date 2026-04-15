package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)


func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())


	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})


	e.GET("/products", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})


	// get dla jednego produktu
	e.GET("/products/:id", func(c *echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		if id < 0 || id >= len(products) {
			return c.JSON(http.StatusNotFound, "Product not found")
		}

		return c.JSON(http.StatusOK, products[id])
	})



	e.POST("/products", func(c *echo.Context) error {
		var newProduct Product

		// parsowanie JSON 
		if err := c.Bind(&newProduct); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		newProduct.ID = len(products) + 1

		products = append(products, newProduct)
		return c.JSON(http.StatusCreated, newProduct)
	})



	e.PUT("/products/:id", func(c *echo.Context) error {
		idParam := c.Param("id")

		// tresc body jsona
		var updated Product
		if err := c.Bind(&updated); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		// szukanie produktu
		for i, p := range products {
			if p.ID == id {
				products[i].Name = updated.Name
				products[i].Price = updated.Price
				return c.JSON(http.StatusOK, products[i])
			}
		}

		return c.JSON(http.StatusNotFound, "Product not found")
	})




	e.DELETE("/products/:id", func(c *echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		for i, p := range products {
			if p.ID == id {
				// usuwamy poprzez dodanie dwoch wyciknow
				products = append(products[:i], products[i+1:]...)
				return c.NoContent(http.StatusNoContent)
			}
		}

		return c.JSON(http.StatusNotFound, "Product not found")
	})




	
	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}


}
