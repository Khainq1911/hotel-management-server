package handler

import (
	"booking-website-be/model"
	"booking-website-be/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	RoomRepo repository.RoomRepo
}

func (u *RoomHandler) AddRoom(ctx echo.Context) error {
	req := model.AddRoom{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Println("Binding error:", err)
		return ctx.JSON(http.StatusUnprocessableEntity, model.ResWithOutData{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Invalid request data. Please check your input.",
		})
	}
	cookie, err := ctx.Cookie("employee_id")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := u.RoomRepo.AddRoomRepo(ctx.Request().Context(), req, cookie.Value); err != nil {
		fmt.Println("Repository error:", err)
		return ctx.JSON(http.StatusInternalServerError, model.ResWithOutData{
			StatusCode: http.StatusInternalServerError,
			Message:    "An error occurred while saving the room data.",
		})
	}

	return ctx.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Room added successfully.",
		Data:       req,
	})
}

// view list room
func (u *RoomHandler) ViewListRoom(ctx echo.Context) error {

	response, err := u.RoomRepo.ViewListRoomRepo(ctx.Request().Context())
	if err != nil {

		fmt.Println("Error fetching room list:", err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: 400,
			Message:    "Failed to retrieve room data",
		})
	}
	data := []model.ListRoom{}
	for _, v := range response {
		typeroom := model.SelectTypeRoom{
			TypeID:          v.TypeID,
			TypeName:        v.TypeName,
			TypeDescription: v.TypeDescription,
			MaxOccupancy:    v.MaxOccupancy,
			RoomSize:        v.RoomSize,
		}
		view := model.View{
			ViewId:          v.ViewId,
			ViewName:        v.ViewName,
			ViewDescription: v.ViewDescription,
		}
		price := model.Price{
			PricingId:    v.PricingId,
			PricePerDay:  v.PricePerDay,
			PricePerHour: v.PricePerHour,
			Discount:     v.Discount,
		}
		value := model.ListRoom{
			RoomID:         v.RoomID,
			RoomName:       v.RoomName,
			Floor:          v.Floor,
			TempPrice:      v.TempPrice,
			BookingStatus:  v.BookingStatus,
			CleaningStatus: v.CleaningStatus,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			TypeRoom:       typeroom,
			View:           view,
			Price:          price,
		}
		data = append(data, value)
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: 200,
		Message:    "Data retrieved successfully",
		Data:       data,
	})
}

// update room
func (u *RoomHandler) UpdateRoom(ctx echo.Context) error {
	roomID := ctx.Param("room_id")

	req := model.UpdateRoom{}
	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusUnprocessableEntity, model.ResWithOutData{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable to process input data",
		})
	}
	cookie, err := ctx.Cookie("employee_id")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := u.RoomRepo.UpdateRoomRepo(ctx.Request().Context(), roomID, req, cookie.Value); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, model.ResWithOutData{
			StatusCode: http.StatusInternalServerError,
			Message:    "An error occurred while updating the room",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Room updated successfully",
		Data:       req,
	})
}

// delete room
func (u *RoomHandler) DeleteRoom(ctx echo.Context) error {

	roomID := ctx.Param("room_id")
	if roomID == "" {
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "Room ID is required.",
		})
	}

	cookie, err := ctx.Cookie("employee_id")
	if err != nil {
		fmt.Println("Error retrieving employee ID cookie:", err)
		return ctx.JSON(http.StatusUnauthorized, model.ResWithOutData{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized access. Employee ID not found.",
		})
	}

	if err := u.RoomRepo.DeleteRoomRepo(ctx.Request().Context(), roomID, cookie.Value); err != nil {
		fmt.Println("Error deleting room:", err)
		return ctx.JSON(http.StatusInternalServerError, model.ResWithOutData{
			StatusCode: http.StatusInternalServerError,
			Message:    "An error occurred while deleting the room.",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Room deleted successfully.",
	})
}
