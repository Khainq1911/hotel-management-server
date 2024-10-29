package main

import (
	"booking-website-be/database"
	"booking-website-be/handler"
	"booking-website-be/repository"
	"log"
	"net/http"
	"os"
	"strconv"

	"booking-website-be/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Error loading Port")
	}

	sql := &database.Sql{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     DB_PORT,
		Dbname:   os.Getenv("DB_NAME"),
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.CORS())

	TypeRoomDb := handler.TypeRoomHandler{
		TypeRoomRepo: repository.NewTypeRoomRepo(sql),
	}

	AccountHandler := handler.AccountHandler{
		Repo: repository.NewAccountRepo(sql),
	}

	RoomDb := handler.RoomHandler{
		RoomRepo: repository.NewRoomRepo(sql),
	}

	bookingDb := handler.BookingHandler{
		BookingRepo: repository.NewBookingRepo(sql),
	}

	salaryDb := handler.SalaryHandler{
		Repo: repository.NewSalaryRepo(sql),
	}

	employeeDb := handler.EmployeeHandler{
		EmployeeRepo: repository.NewEmployeeRepo(sql),
	}

	paymentDb := handler.PaymentHandler{
		Repo: repository.NewPaymentRepo(sql),
	}
	api := router.Api{
		Echo:            e,
		AccountHandler:  AccountHandler,
		TypeRoomHandler: TypeRoomDb,
		RoomHandler:     RoomDb,
		BookingHandler:  bookingDb,
		SalaryHandler:   salaryDb,
		EmployeeHandler: employeeDb,
		PaymentHandler:  paymentDb,
	}

	api.SetupRouter()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
