package delivery

import (
	"net/http"
	"strconv"

	"github.com/haikalvidya/go-simple-todo/internal/delivery/payload"
	"github.com/haikalvidya/go-simple-todo/pkg/common"
	"github.com/haikalvidya/go-simple-todo/pkg/utils"
	"github.com/labstack/echo/v4"
)

type activityDelivery deliveryType

func (d *activityDelivery) CreateActivity(c echo.Context) error {
	res := common.Response{}
	req := &payload.CreateActivityRequest{}

	c.Bind(req)

	if err := c.Validate(req); err != nil {
		res.Error = utils.GetErrorValidation(err)
		res.Status = false
		res.Message = "Failed Create Activity"
		return c.JSON(http.StatusBadRequest, res)
	}

	activityRes, err := d.Usecase.Activity.CreateActivity(req)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Create Activity"
	res.Data = activityRes
	res.Status = true

	return c.JSON(http.StatusOK, res)
}

func (d *activityDelivery) GetAllActivity(c echo.Context) error {
	res := common.Response{}

	activityRes, err := d.Usecase.Activity.GetAllActivities()
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Get All Activity"
	res.Data = activityRes
	res.Status = true

	return c.JSON(http.StatusOK, res)
}

func (d *activityDelivery) GetActivityById(c echo.Context) error {
	res := common.Response{}

	activityId := c.Param("id")

	// convert to int64
	activityIdInt, err := strconv.ParseInt(activityId, 10, 64)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	activityRes, err := d.Usecase.Activity.SelectActivityByID(activityIdInt)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Get Activity By Id"
	res.Data = activityRes
	res.Status = true

	return c.JSON(http.StatusOK, res)
}

func (d *activityDelivery) UpdateActivity(c echo.Context) error {
	res := common.Response{}
	req := &payload.UpdateActivityRequest{}

	activityId := c.Param("id")

	// convert to int64
	activityIdInt, err := strconv.ParseInt(activityId, 10, 64)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	c.Bind(req)

	if err := c.Validate(req); err != nil {
		res.Error = utils.GetErrorValidation(err)
		res.Status = false
		res.Message = "Failed Update Activity"
		return c.JSON(http.StatusBadRequest, res)
	}

	activityRes, err := d.Usecase.Activity.UpdateActivity(activityIdInt, req)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Update Activity"
	res.Data = activityRes
	res.Status = true

	return c.JSON(http.StatusOK, res)
}

func (d *activityDelivery) DeleteActivity(c echo.Context) error {
	res := common.Response{}

	activityId := c.Param("id")

	// convert to int64
	activityIdInt, err := strconv.ParseInt(activityId, 10, 64)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	err = d.Usecase.Activity.DeleteActivity(activityIdInt)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Delete Activity"
	res.Status = true

	return c.JSON(http.StatusOK, res)
}
