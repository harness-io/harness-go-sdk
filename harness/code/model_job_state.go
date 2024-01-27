/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type JobState string

// List of JobState
const (
	CANCELED_JobState JobState = "canceled"
	FAILED_JobState JobState = "failed"
	FINISHED_JobState JobState = "finished"
	RUNNING_JobState JobState = "running"
	SCHEDULED_JobState JobState = "scheduled"
)
