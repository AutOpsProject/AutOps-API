package workflow

import (
	"strings"

	"github.com/AutOpsProject/AutOps-API/internal/domain/common"
)

// Workflow represents a workflow entity composed of inputs, outputs, and steps.
// It embeds StatefulNamedEntity and VersionedSource for identification, naming, status, and versioning.
type Workflow struct {
	common.StatefulNamedEntity
	common.VersionedSource
	inputs  *common.List[*WorkflowAttribute]
	outputs *common.List[*WorkflowAttribute]
	steps   *common.List[*WorkflowStep]
	runs    *common.List[*WorkflowRun]
}

type WorkflowComparator struct{}

func (WorkflowComparator) Compare(t1, t2 *Workflow) int {
	return strings.Compare(t1.GetIdentifier().ToString(), t2.GetIdentifier().ToString())
}

// ExistingWorkflow creates a Workflow instance using an existing identifier, status, version, and provided lists of inputs, outputs, and steps.
// Returns an error if the name, description, or source path are invalid.
func ExistingWorkflow(workflowIdentifier string, name string, description string, status common.Status, sourcePath string, version int, inputs []*WorkflowAttribute, outputs []*WorkflowAttribute, steps []*WorkflowStep, runs []*WorkflowRun) (*Workflow, error) {
	statefulEntity, err := common.NewStatefulNamedEntity(workflowIdentifier, name, description, status)
	if err != nil {
		return nil, err
	}
	versionedEntity, err := common.NewVersionedSource(sourcePath, version)
	if err != nil {
		return nil, err
	}
	workflowEntity := Workflow{
		StatefulNamedEntity: *statefulEntity,
		VersionedSource:     *versionedEntity,
		inputs:              common.NewList(common.Comparator[*WorkflowAttribute](WorkflowAttributeComparator{}), inputs),
		outputs:             common.NewList(common.Comparator[*WorkflowAttribute](WorkflowAttributeComparator{}), outputs),
		steps:               common.NewList(common.Comparator[*WorkflowStep](WorkflowStepComparator{}), steps),
		runs:                common.NewList(common.Comparator[*WorkflowRun](WorkflowRunComparator{}), runs),
	}
	return &workflowEntity, nil
}

// NewWorkflow creates a new Workflow instance with a freshly generated identifier, default status (PENDING), version 1, and empty lists.
func NewWorkflow(projectIdentifier string, name string, description string, sourcePath string) (*Workflow, error) {
	workflowIdentifier, err := common.BuildWorkflowIdentifier(projectIdentifier)
	if err != nil {
		return nil, err
	}
	return ExistingWorkflow(workflowIdentifier.ToString(), name, description, common.PENDING, sourcePath, 1, []*WorkflowAttribute{}, []*WorkflowAttribute{}, []*WorkflowStep{}, []*WorkflowRun{})
}

// ListInputs returns the list of WorkflowAttributes defined as inputs.
func (w *Workflow) ListInputs() []*WorkflowAttribute {
	return w.inputs.Items()
}

// ListOutputs returns the list of WorkflowAttributes defined as outputs.
func (w *Workflow) ListOutputs() []*WorkflowAttribute {
	return w.outputs.Items()
}

// ListSteps returns the list of WorkflowSteps defined in the workflow.
func (w *Workflow) ListSteps() []*WorkflowStep {
	return w.steps.Items()
}

// ListRuns returns the list of WorkflowRuns of the workflow.
func (w *Workflow) ListRuns() []*WorkflowRun {
	return w.runs.Items()
}

// AddInput adds a new WorkflowAttribute as input to the workflow.
// Returns an error if the input is already present.
func (w *Workflow) AddInput(input *WorkflowAttribute) error {
	if w.inputs.Contains(input) {
		return ErrWorkflowInputAlreadyPresent
	}
	w.inputs.Append(input)
	return nil
}

// AddOutput adds a new WorkflowAttribute as output to the workflow.
// Returns an error if the output is already present.
func (w *Workflow) AddOutput(output *WorkflowAttribute) error {
	if w.outputs.Contains(output) {
		return ErrWorkflowOutputAlreadyPresent
	}
	w.outputs.Append(output)
	return nil
}

// AddStep inserts a WorkflowStep into the workflow and shifts subsequent step numbers.
func (w *Workflow) AddStep(step *WorkflowStep) {
	w.shiftStepsFrom(step.GetStepNumber(), 1)
	w.steps.Append(step)
}

func (w *Workflow) AddRun(run *WorkflowRun) error {
	if w.runs.Contains(run) {
		return ErrWorkflowRunAlreadyPresent
	}
	w.runs.Append(run)
	return nil
}

// RemoveInput removes an input by its identifier from the workflow.
// Returns an error if the input is not found.
func (w *Workflow) RemoveInput(inputIdentifier string) error {
	attribute, found := w.inputs.SelectOne(func(t *WorkflowAttribute) bool {
		return t.GetIdentifier().ToString() == inputIdentifier
	})
	if !found {
		return ErrWorkflowInputNotFound
	}
	w.inputs.Remove(attribute)
	return nil
}

// RemoveOutput removes an output by its identifier from the workflow.
// Returns an error if the output is not found.
func (w *Workflow) RemoveOutput(outputIdentifier string) error {
	attribute, found := w.outputs.SelectOne(func(t *WorkflowAttribute) bool {
		return t.GetIdentifier().ToString() == outputIdentifier
	})
	if !found {
		return ErrWorkflowOutputNotFound
	}
	w.outputs.Remove(attribute)
	return nil
}

// RemoveStep removes a step by its step number from the workflow and shifts subsequent step numbers.
// Returns an error if the step is not found.
func (w *Workflow) RemoveStep(stepNumber int) error {
	attribute, found := w.steps.SelectOne(func(s *WorkflowStep) bool {
		return s.GetStepNumber() == stepNumber
	})
	if !found {
		return ErrWorkflowStepNotFound
	}
	w.steps.Remove(attribute)
	w.shiftStepsFrom(stepNumber, -1)
	return nil
}

// shiftStepsFrom shifts the step numbers of all steps with a step number >= stepNumber by the provided shiftValue.
func (w *Workflow) shiftStepsFrom(stepNumber int, shiftValue int) {
	steps, _ := w.steps.SelectAll(func(s *WorkflowStep) bool {
		return s.GetStepNumber() >= stepNumber
	})
	for _, step := range steps {
		step.stepNumber += shiftValue
	}
}
