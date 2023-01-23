package delivery

import (
	"github.com/haikalvidya/go-simple-todo/internal/middlewares"
	"github.com/haikalvidya/go-simple-todo/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	Activity *activityDelivery
	Todo     *todoDelivery
}

type deliveryType struct {
	Usecase    *usecase.Usecase
	Middleware *middlewares.CustomMiddleware
}

func NewDelivery(e *echo.Echo, usecase *usecase.Usecase, mid *middlewares.CustomMiddleware) *Delivery {
	deliveryType := &deliveryType{
		Usecase:    usecase,
		Middleware: mid,
	}
	delivery := &Delivery{
		Activity: (*activityDelivery)(deliveryType),
		Todo:     (*todoDelivery)(deliveryType),
	}

	Route(e, delivery, mid)

	return delivery
}

func Route(e *echo.Echo, delivery *Delivery, mid *middlewares.CustomMiddleware) {
	// activity-groups
	activityGroups := e.Group("/activity-groups")
	{
		activityGroups.GET("", delivery.Activity.GetAllActivity)
		activityGroups.GET("/:id", delivery.Activity.GetActivityById)
		activityGroups.POST("", delivery.Activity.CreateActivity)
		activityGroups.PATCH("/:id", delivery.Activity.UpdateActivity)
		activityGroups.DELETE("/:id", delivery.Activity.DeleteActivity)
	}

	// todo-items
	todoItems := e.Group("/todo-items")
	{
		todoItems.GET("", delivery.Todo.GetAllTodos)
		todoItems.GET("/:id", delivery.Todo.GetTodoByID)
		todoItems.POST("", delivery.Todo.CreateTodo)
		todoItems.PATCH("/:id", delivery.Todo.UpdateTodoByID)
		todoItems.DELETE("/:id", delivery.Todo.DeleteTodoByID)
	}
}
