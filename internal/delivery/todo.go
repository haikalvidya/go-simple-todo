package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/haikalvidya/go-simple-todo/internal/delivery/payload"
	"github.com/haikalvidya/go-simple-todo/pkg/common"
	"github.com/haikalvidya/go-simple-todo/pkg/utils"
	"github.com/labstack/echo/v4"
)

type todoDelivery deliveryType

func (d *todoDelivery) CreateTodo(c echo.Context) error {
	res := common.Response{}
	req := &payload.CreateTodoRequest{}

	c.Bind(req)

	if err := c.Validate(req); err != nil {
		res.Error = utils.GetErrorValidation(err)
		res.Status = false
		res.Message = "Failed Create Todo"
		return c.JSON(http.StatusBadRequest, res)
	}

	todoRes, err := d.Usecase.Todo.CreateTodo(req)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Create Todo"
	res.Data = todoRes
	res.Status = true

	return c.JSON(http.StatusCreated, res)
}

func (d *todoDelivery) GetAllTodos(c echo.Context) error {
	res := common.Response{}
	query := c.QueryParam("activity_group_id")
	var activity_group_id *int64

	if query == "" {
		activity_group_id = nil
	} else {
		activity_group_id_int, err := strconv.ParseInt(query, 10, 64)
		if err != nil {
			activity_group_id = nil
		} else {
			activity_group_id = &activity_group_id_int
		}
	}

	todoRes, err := d.Usecase.Todo.GetAllTodos(activity_group_id)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Get All Todo"
	res.Data = todoRes
	res.Status = true

	return c.JSON(http.StatusOK, res)
}

func (d *todoDelivery) GetTodoByID(c echo.Context) error {
	res := common.Response{}

	todoID := c.Param("id")

	// convert string to int64
	todoIDInt, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil {
		res.Status = false
		res.Message = fmt.Sprintf(payload.TODO_ID_NOT_FOUND, todoID)
		return c.JSON(http.StatusBadRequest, res)
	}

	todoRes, err := d.Usecase.Todo.SelectTodoByID(todoIDInt)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Get Todo By ID"
	res.Data = todoRes
	res.Status = true

	return c.JSON(http.StatusOK, res)
}

func (d *todoDelivery) UpdateTodoByID(c echo.Context) error {
	res := common.Response{}
	req := &payload.UpdateTodoRequest{}

	c.Bind(req)

	if err := c.Validate(req); err != nil {
		res.Error = utils.GetErrorValidation(err)
		res.Status = false
		res.Message = "Failed Update Todo"
		return c.JSON(http.StatusBadRequest, res)
	}

	todoID := c.Param("id")

	// convert string to int64
	todoIDInt, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil {
		res.Status = false
		res.Message = fmt.Sprintf(payload.TODO_ID_NOT_FOUND, todoID)
		return c.JSON(http.StatusBadRequest, res)
	}

	todoRes, err := d.Usecase.Todo.UpdateTodo(todoIDInt, req)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Update Todo By ID"
	res.Data = todoRes
	res.Status = true

	return c.JSON(http.StatusOK, res)
}

func (d *todoDelivery) DeleteTodoByID(c echo.Context) error {
	res := common.Response{}

	todoID := c.Param("id")

	// convert string to int64
	todoIDInt, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil {
		res.Status = false
		res.Message = fmt.Sprintf(payload.TODO_ID_NOT_FOUND, todoID)
		return c.JSON(http.StatusBadRequest, res)
	}

	err = d.Usecase.Todo.DeleteTodo(todoIDInt)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Delete Todo By ID"
	res.Status = true

	return c.JSON(http.StatusOK, res)
}
