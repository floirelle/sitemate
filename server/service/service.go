package service

import (
	"server/models"
	"server/repository"
)

type Service interface {
	GetIssues() ([]models.Issue, error)
	GetIssue(id string) (models.Issue, error)
	AddIssue(issue models.Issue) (models.Issue, error)
	UpdateIssue(id string, newIssue models.Issue) (models.Issue, error)
	DeleteIssue(id string) (models.Issue, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetIssues() ([]models.Issue, error) {
	return s.repository.GetIssues()
}

func (s *service) GetIssue(id string) (models.Issue, error) {
	return s.repository.GetIssue(id)
}

func (s *service) AddIssue(issue models.Issue) (models.Issue, error) {
	return s.repository.AddIssue(issue)
}

func (s *service) UpdateIssue(id string, newIssue models.Issue) (models.Issue, error) {
	return s.repository.UpdateIssue(id, newIssue)
}

func (s *service) DeleteIssue(id string) (models.Issue, error) {
	return s.repository.DeleteIssue(id)
}
