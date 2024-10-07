package jobapplication

import (
	"context"
	"fmt"

	"github.com/d4499/jager/internal"
	"github.com/d4499/jager/internal/database/db"
)

type JobApplicationService struct {
	db *db.Queries
}

func NewJobApplicationService(db *db.Queries) JobApplicationService {
	return JobApplicationService{
		db: db,
	}
}

func (j *JobApplicationService) CreateJobApplication(params db.CreateJobApplicationParams) (db.JobApplication, error) {
	jobApp, err := j.db.CreateJobApplication(context.Background(), db.CreateJobApplicationParams{
		ID:          internal.NewCUID(),
		Title:       params.Title,
		Company:     params.Company,
		UserID:      params.UserID,
		AppliedDate: params.AppliedDate,
	})
	if err != nil {
		return db.JobApplication{}, fmt.Errorf("unable to create job application: %v", err)
	}

	return jobApp, nil
}

func (j *JobApplicationService) DeleteJobApplication(id string) error {
	err := j.db.DeleteJobApplication(context.Background(), id)
	if err != nil {
		return fmt.Errorf("unable to delete job application: %v", err)
	}
	return nil
}

func (j *JobApplicationService) GetAllJobApplications(userId string) ([]db.JobApplication, error) {
	jobApps, err := j.db.GetAllJobApplications(context.Background(), userId)
	if err != nil {
		return []db.JobApplication{}, fmt.Errorf("unable to get job applications: %v", err)
	}

	return jobApps, nil
}
