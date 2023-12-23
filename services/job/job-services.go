package jobservices

import jobrepo "finder-job/repository/job"

type JobServices struct {
	jobRepo *jobrepo.JobRepo
}

func NewJobServices(jobRepo *jobrepo.JobRepo) *JobServices {
     return &JobServices{jobRepo: jobRepo}
}

func (s *JobServices) GetJobs() {}

func (s *JobServices) GetJobById() {}

func (s *JobServices) CreateJob() {}

func (s *JobServices) UpdateJob() {}

func (s *JobServices) DeleteJob() {}

func (s *JobServices) ApplyJob() {}

func (s *JobServices) GetAppliedJob() {}

