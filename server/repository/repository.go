package repository

import (
	"fmt"
	"server/custom_errors"
	"server/models"

	"github.com/google/uuid"
)

type Repository interface {
	GetIssues() ([]models.Issue, error)
	GetIssue(id string) (models.Issue, error)
	AddIssue(issue models.Issue) (models.Issue, error)
	UpdateIssue(id string, newIssue models.Issue) (models.Issue, error)
	DeleteIssue(id string) (models.Issue, error)
}

type repository struct {
	issues map[string]models.Issue
}

func NewRepository() Repository {
	return &repository{
		issues: map[string]models.Issue{
			"1": {
				Id:          "1",
				Title:       "FIRST TITLE",
				Description: "FIRST DESCRIPTION",
				Reporter:    "ADINATA",
				Watchers:    []string{"ADINATA"},
			},
		},
	}
}

func (r *repository) GetIssues() ([]models.Issue, error) {
	issues := []models.Issue{}
	fmt.Println(r.issues)
	for _, v := range r.issues {
		issues = append(issues, v)
	}

	return issues, nil
}

func (r *repository) GetIssue(id string) (models.Issue, error) {
	issue, exist := r.issues[id]
	if !exist {
		return models.Issue{}, custom_errors.NotFoundError()
	}

	return issue, nil
}

func (r *repository) AddIssue(issue models.Issue) (models.Issue, error) {
	newIssue := models.Issue{
		Id:          uuid.NewString(),
		Title:       issue.Title,
		Description: issue.Description,
		Reporter:    issue.Reporter,
		Watchers:    issue.Watchers,
	}

	r.issues[newIssue.Id] = newIssue
	return newIssue, nil
}

func (r *repository) UpdateIssue(id string, newIssue models.Issue) (models.Issue, error) {
	issue, exist := r.issues[id]
	if !exist {
		return models.Issue{}, custom_errors.NotFoundError()
	}

	issue.Title = newIssue.Title
	issue.Description = newIssue.Description
	issue.Watchers = newIssue.Watchers

	r.issues[issue.Id] = issue
	return issue, nil
}

func (r *repository) DeleteIssue(id string) (models.Issue, error) {
	issue, exist := r.issues[id]
	if !exist {
		return models.Issue{}, custom_errors.NotFoundError()
	}

	delete(r.issues, issue.Id)

	return issue, nil
}
