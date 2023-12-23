package jobcontroller

import (
	jobservices "finder-job/services/job"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	JobService *jobservices.JobServices
}

func NewJobController(JobServices *jobservices.JobServices) *JobController {
   return &JobController{JobService: JobServices}
}


func (c *JobController) GetJobs(ctx *gin.Context) {}

func (c *JobController) GetJobById(ctx *gin.Context){}

func (c *JobController) CreateJob(ctx *gin.Context) {}

func (c *JobController) UpdateJob(ctx *gin.Context) {}

func (c *JobController) DeleteJob(ctx *gin.Context) {}

func (c *JobController) GetAppliedJob(ctx *gin.Context) {}

func (c *JobController) GetApplierInJob(ctx *gin.Context) {}

func (c *JobController) ApplyJob(ctx *gin.Context) {}