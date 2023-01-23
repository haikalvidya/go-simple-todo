package usecase

import (
	"github.com/haikalvidya/go-simple-todo/internal/delivery/payload"
	"github.com/haikalvidya/go-simple-todo/internal/models"
	"gorm.io/gorm"
)

type ITodoUsecase interface {
	GetAllTodos(query *int64) ([]*payload.TodoInfo, error)
	SelectTodoByID(id int64) (*payload.TodoInfo, error)
	CreateTodo(todo *payload.CreateTodoRequest) (*payload.TodoInfo, error)
	UpdateTodo(id int64, todo *payload.UpdateTodoRequest) (*payload.TodoInfo, error)
	DeleteTodo(id int64) error
}

type todoUsecase usecaseType

func (u *todoUsecase) GetAllTodos(query *int64) ([]*payload.TodoInfo, error) {
	var todos []*models.Todo
	var err error
	if query != nil {
		todos, err = u.Repo.Todo.SelectByQueryActivityGroupID(*query)
		if err != nil {
			return nil, err
		}
	} else {
		todos, err = u.Repo.Todo.GetAll()
		if err != nil {
			return nil, err
		}

	}

	todoInfos := []*payload.TodoInfo{}
	for _, todo := range todos {
		todoInfos = append(todoInfos, todo.PublicInfo())
	}

	return todoInfos, nil
}

func (u *todoUsecase) SelectTodoByID(id int64) (*payload.TodoInfo, error) {
	todo, err := u.Repo.Todo.SelectByID(id)
	if err != nil {
		return nil, err
	}

	return todo.PublicInfo(), nil
}

func (u *todoUsecase) CreateTodo(todo *payload.CreateTodoRequest) (*payload.TodoInfo, error) {
	todoModel := &models.Todo{
		Title: todo.Title,
	}

	var err error

	err = u.Repo.Tx.DoInTransaction(func(tx *gorm.DB) error {
		todoModel, err = u.Repo.Todo.CreateTx(tx, todoModel)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return todoModel.PublicInfo(), nil
}

func (u *todoUsecase) UpdateTodo(id int64, todo *payload.UpdateTodoRequest) (*payload.TodoInfo, error) {
	todoModel, err := u.Repo.Todo.SelectByID(id)
	if err != nil {
		return nil, err
	}

	todoModel.Title = todo.Title

	var err2 error

	err2 = u.Repo.Tx.DoInTransaction(func(tx *gorm.DB) error {
		err2 = u.Repo.Todo.UpdateTx(tx, todoModel)
		if err2 != nil {
			return err2
		}
		return nil
	})
	if err2 != nil {
		return nil, err2
	}

	return todoModel.PublicInfo(), nil
}

func (u *todoUsecase) DeleteTodo(id int64) error {
	todoModel, err := u.Repo.Todo.SelectByID(id)
	if err != nil {
		return err
	}

	var err2 error

	err2 = u.Repo.Tx.DoInTransaction(func(tx *gorm.DB) error {
		err2 = u.Repo.Todo.DeleteTx(tx, todoModel)
		if err2 != nil {
			return err2
		}
		return nil
	})
	if err2 != nil {
		return err2
	}

	return nil
}
