package mqs

import "github.com/treeder/functions/api/models"

func validate(job *models.Task) error {
	if job == nil {
		return models.ErrMQMissingTask
	}
	if job.ID == "" {
		return models.ErrMQEmptyTaskID
	}
	if job.Priority == nil {
		return models.ErrMQMissingTaskPriority
	}
	if *job.Priority < 0 || *job.Priority > 2 {
		return models.NewErrMQInvalidTaskPriority(job)
	}
	return nil
}
