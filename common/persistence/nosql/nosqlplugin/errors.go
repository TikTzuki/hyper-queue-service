package nosqlplugin

import "fmt"

// Condition Errors for NoSQL interfaces
type (
	// Only one of the fields must be non-nil
	WorkflowOperationConditionFailure struct {
		UnknownConditionFailureDetails   *string // return some info for logging
		ShardRangeIDNotMatch             *int64  // return the previous shardRangeID
		WorkflowExecutionAlreadyExists   *WorkflowExecutionAlreadyExists
		CurrentWorkflowConditionFailInfo *string // return the logging info if fail on condition of CurrentWorkflow
	}

	WorkflowExecutionAlreadyExists struct {
		RunID            string
		CreateRequestID  string
		State            int
		CloseStatus      int
		LastWriteVersion int64
		OtherInfo        string
	}

	TaskOperationConditionFailure struct {
		RangeID int64
		Details string // detail info for logging
	}

	ShardOperationConditionFailure struct {
		RangeID int64
		Details string // detail info for logging
	}

	ConditionFailure struct {
		componentName string
	}
)

var _ error = (*WorkflowOperationConditionFailure)(nil)

func (e *WorkflowOperationConditionFailure) Error() string {
	return "workflow operation condition failure"
}

var _ error = (*TaskOperationConditionFailure)(nil)

func (e *TaskOperationConditionFailure) Error() string {
	return "task operation condition failure"
}

var _ error = (*ShardOperationConditionFailure)(nil)

func (e *ShardOperationConditionFailure) Error() string {
	return "shard operation condition failure"
}

var _ error = (*ConditionFailure)(nil)

func NewConditionFailure(name string) error {
	return &ConditionFailure{
		componentName: name,
	}
}
func (e *ConditionFailure) Error() string {
	return fmt.Sprintf("%s operation condition failure", e.componentName)
}
