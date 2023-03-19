package repository

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type IAuthorization interface {
	CreateUser(user cargodelivery.User) (int, error)
	GetUser(username string) (cargodelivery.User, error)
}

type IProject interface {
	GetUserProject(userId, projectId int) (cargodelivery.Project, error)
	CreateProject(userId int, project cargodelivery.Project, access string) (int, error)
	GetAllProjects(userId int) ([]cargodelivery.Project, error)
	GetProjectById(userId, projectId int) (cargodelivery.Project, error)
	DeleteProject(userId, projectId int) error
	UpdateProject(userId, projectId int, input cargodelivery.UpdateProject) error
}

type ICar interface {
	CreateCar(projectId int, car cargodelivery.Car) (int, error)
	GetAllCars(projectId int) ([]cargodelivery.Car, error)
	GetCarById(projectId, carId int) (cargodelivery.Car, error)
	DeleteCar(projectId, carId int) error
	UpdateCar(projectId, carId int, input cargodelivery.UpdateCar) error
}

type ICargo interface {
	CreateCargo(projectId int, cargo cargodelivery.Cargo) (int, error)
	GetAllCargos(projectId int) ([]cargodelivery.Cargo, error)
	GetCargoById(projectId, cargoId int) (cargodelivery.Cargo, error)
	DeleteCargo(projectId, cargoId int) error
	UpdateCargo(projectId, cargoId int, input cargodelivery.UpdateCargo) error
}

type IClient interface {
	CreateClient(projectId int, client cargodelivery.Client) (int, error)
	GetAllClients(projectId int) ([]cargodelivery.Client, error)
	GetClientById(projectId, clientId int) (cargodelivery.Client, error)
	DeleteClient(projectId, clientId int) error
	UpdateClient(projectId, clientId int, input cargodelivery.UpdateClient) error
}

type IOrder interface {
	CreateOrder(clientId int, order cargodelivery.Order) (int, error)
}

type Repository struct {
	IAuthorization
	IProject
	ICar
	ICargo
	IClient
	IOrder
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IAuthorization: NewAuthPostgres(db),
		IProject:       NewProjectPostgres(db),
		ICar:           NewCarPostgres(db),
		ICargo:         NewCargoPostgres(db),
		IClient:        NewClientPostgres(db),
		IOrder:         NewOrderPostgres(db),
	}
}