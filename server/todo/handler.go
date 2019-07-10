package todo

import (
	context "context"
	"log"

	"github.com/satori/go.uuid"
)

// Server grpc server
type Server struct {
	Todos []*TodoObject
}

// ListTodo list all
func (s *Server) ListTodo(ctx context.Context, _ *ListTodoParams) (*ListTodoResponse, error) {
	log.Printf("ListTodo")
	return &ListTodoResponse{Todos: s.Todos}, nil
}

// AddTodo add one
func (s *Server) AddTodo(ctx context.Context, todoParams *AddTodoParams) (*AddTodoResponse, error) {
	log.Printf("AddTodo task %s", todoParams.Task)
	uuid, _ := uuid.NewV1()
	todoObject := &TodoObject{
		Id:   uuid.String(),
		Task: todoParams.Task,
	}
	s.Todos = append(s.Todos, todoObject)
	return &AddTodoResponse{Message: "success"}, nil
}

// GetTodo get one
func (s *Server) GetTodo(ctx context.Context, todoParams *GetTodoParams) (*GetTodoResponse, error) {
	log.Printf("GetTodo id %s", todoParams.Id)
	id := todoParams.Id
	var todoObject *TodoObject
	for _, todo := range s.Todos {
		if todo.Id == id {
			todoObject = todo
			break
		}
	}
	return &GetTodoResponse{Todo: todoObject}, nil
}

// DeleteTodo delete one
func (s *Server) DeleteTodo(ctx context.Context, todoParams *DeleteTodoParams) (*DeleteTodoResponse, error) {
	log.Printf("DeleteTodo id %s", todoParams.Id)
	id := todoParams.Id
	var updatedTodos []*TodoObject
	for index, todo := range s.Todos {
		if todo.Id == id {
			updatedTodos = append(s.Todos[:index], s.Todos[index+1:]...)
			break
		}
	}
	s.Todos = updatedTodos
	return &DeleteTodoResponse{Message: "success"}, nil
}

// UpdateTodo update one
func (s *Server) UpdateTodo(ctx context.Context, todoParams *UpdateTodoParams) (*UpdateTodoResponse, error) {
	log.Printf("UpdateTodo id %s", todoParams.Id)
	id := todoParams.Id
	for _, todo := range s.Todos {
		if todo.Id == id {
			todo.Task = todoParams.Task
			break
		}
	}
	return &UpdateTodoResponse{Message: "success"}, nil
}
