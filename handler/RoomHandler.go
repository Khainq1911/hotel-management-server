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

	if err := u.RoomRepo.AddRoomRepo(ctx.Request().Context(), req); err != nil {
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
			TypeID:       v.TypeID,
			TypeName:     v.TypeName,
			Description:  v.Description,
			MaxOccupancy: v.MaxOccupancy,
			RoomSize:     v.RoomSize,
		}
		value := model.ListRoom{
			RoomID:         v.RoomID,
			RoomName:       v.RoomName,
			Floor:          v.Floor,
			PricePerDay:    v.PricePerDay,
			PricePerHour:   v.PricePerHour,
			BookingStatus:  v.BookingStatus,
			Discount:       v.Discount,
			CleaningStatus: v.CleaningStatus,
			CheckInTime:    v.CheckInTime,
			CheckOutTime:   v.CheckOutTime,
			CurrentGuest:   v.CurrentGuest,
			Note:           v.Note,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			TypeRoom:       typeroom,
		}
		data = append(data, value)
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: 200,
		Message:    "Data retrieved successfully",
		Data:       data,
	})
}

// view room detail
func (u *RoomHandler) ViewDetailRoom(ctx echo.Context) error {
	room_id := ctx.Param("room_id")

	response, err := u.RoomRepo.ViewDetailRoomRepo(ctx.Request().Context(), room_id)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, model.ResWithOutData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error retrieving room details",
		})
	}

	if len(response) == 0 {
		return ctx.JSON(http.StatusNotFound, model.ResWithOutData{
			StatusCode: http.StatusNotFound,
			Message:    "Room not found",
		})
	}

	data := []model.ListRoom{}
	v := response[0]

	typeroom := model.SelectTypeRoom{
		TypeID:       v.TypeID,
		TypeName:     v.TypeName,
		Description:  v.Description,
		MaxOccupancy: v.MaxOccupancy,
		RoomSize:     v.RoomSize,
	}
	value := model.ListRoom{
		RoomID:         v.RoomID,
		RoomName:       v.RoomName,
		Floor:          v.Floor,
		BookingStatus:  v.BookingStatus,
		CleaningStatus: v.CleaningStatus,
		CheckInTime:    v.CheckInTime,
		CheckOutTime:   v.CheckOutTime,
		CurrentGuest:   v.CurrentGuest,
		Note:           v.Note,
		CreatedAt:      v.CreatedAt,
		UpdatedAt:      v.UpdatedAt,
		TypeRoom:       typeroom,
	}
	data = append(data, value)

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Room details retrieved successfully",
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

	if err := u.RoomRepo.UpdateRoomRepo(ctx.Request().Context(), roomID, req); err != nil {
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
	var req model.DeleteRoom

	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	room_id := ctx.Param("room_id")

	if err := u.RoomRepo.DeleteRoomRepo(ctx.Request().Context(), room_id, req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to update type room",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}
