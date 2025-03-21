package service

import (
	appexception "project-common/exception"
	"project-main/domain"
	"project-main/model"
)

type CustomerService interface {
	GetAllCustomer(string) ([]model.CustomerResponse, *appexception.AppError)
	GetCustomer(string) (*model.CustomerResponse, *appexception.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]model.CustomerResponse, *appexception.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]model.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*model.CustomerResponse, *appexception.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
