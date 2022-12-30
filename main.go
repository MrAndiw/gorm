package main

import (
	"fmt"
	"net/http"
	"time"
	"tutor/gorm/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// connection to db
	db.Connect()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/create-order", CreateOrder)
	e.GET("/get-order", GetOrder)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

type Order struct {
	ID          uint
	IDUser      string
	Code        string
	Notes       string
	Status      int
	TotalAmount uint
	CreatedBy   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Handler
func CreateOrder(c echo.Context) error {
	con := db.GetConnection()

	order := Order{
		IDUser:      "1",
		Code:        "ORDER001",
		Notes:       "Hellow World",
		Status:      1,
		TotalAmount: 10000,
		CreatedBy:   1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result := con.Create(&order) // pass pointer of data to Create

	return c.String(http.StatusOK, fmt.Sprintf("Created Succesfuly %d ", result.RowsAffected))
}

func GetOrder(c echo.Context) error {
	id := c.QueryParam("id")

	con := db.GetConnection()

	var order Order

	// Get all matched records
	con.Where("id = ?", id).Find(&order)

	if order.ID == 0 {
		return c.JSON(http.StatusBadRequest, "Record not found")
	}

	return c.JSON(http.StatusOK, order)
}
