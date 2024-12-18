package handler

import (
	"booking-website-be/model"
	"booking-website-be/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TypeRoomHandler struct {
	TypeRoomRepo repository.TypeRoomRepo
}

func (u *TypeRoomHandler) ListPrice(ctx echo.Context) error {

	res, err := u.TypeRoomRepo.ListPriceRepo(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ResWithOutData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve price list. Please try again later.",
		})
	}

	finalData := []model.ListRoomPricing{}
	for _, v := range res {
		typeRoom := model.SelectTypeRoom{
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
		data := model.ListRoomPricing{
			PricingId:    v.PricingId,
			PricePerDay:  v.PricePerDay,
			PricePerHour: v.PricePerHour,
			Discount:     v.Discount,
			TypeRoom:     typeRoom,
			View:         view,
		}
		finalData = append(finalData, data)
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Price list retrieved successfully.",
		Data:       finalData,
	})
}

func (u *TypeRoomHandler) UpdatePrice(ctx echo.Context) error {

	id := ctx.Param("id")

	bodyReq := model.UpdatePrice{}
	if err := ctx.Bind(&bodyReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body. Please check the data format.",
		})
	}

	cookie, err := ctx.Cookie("employee_id")
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, model.ResWithOutData{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized. Employee ID is missing.",
		})
	}

	if err := u.TypeRoomRepo.UpdatePriceRepo(ctx.Request().Context(), bodyReq, id, cookie.Value); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, model.ResWithOutData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update price. Please try again later.",
		})
	}

	return ctx.JSON(http.StatusOK, model.ResWithOutData{
		StatusCode: http.StatusOK,
		Message:    "Price updated successfully.",
	})
}
