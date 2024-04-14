// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package transport

import (
	"context"
	"sync"

	"github.com/kubefirst/tink/internal/agent/event"
	"github.com/kubefirst/tink/internal/agent/workflow"
)

// Ensure, that WorkflowHandlerMock does implement WorkflowHandler.
// If this is not the case, regenerate this file with moq.
var _ WorkflowHandler = &WorkflowHandlerMock{}

// WorkflowHandlerMock is a mock implementation of WorkflowHandler.
//
//	func TestSomethingThatUsesWorkflowHandler(t *testing.T) {
//
//		// make and configure a mocked WorkflowHandler
//		mockedWorkflowHandler := &WorkflowHandlerMock{
//			CancelWorkflowFunc: func(workflowID string)  {
//				panic("mock out the CancelWorkflow method")
//			},
//			HandleWorkflowFunc: func(contextMoqParam context.Context, workflowMoqParam workflow.Workflow, recorder event.Recorder)  {
//				panic("mock out the HandleWorkflow method")
//			},
//		}
//
//		// use mockedWorkflowHandler in code that requires WorkflowHandler
//		// and then make assertions.
//
//	}
type WorkflowHandlerMock struct {
	// CancelWorkflowFunc mocks the CancelWorkflow method.
	CancelWorkflowFunc func(workflowID string)

	// HandleWorkflowFunc mocks the HandleWorkflow method.
	HandleWorkflowFunc func(contextMoqParam context.Context, workflowMoqParam workflow.Workflow, recorder event.Recorder)

	// calls tracks calls to the methods.
	calls struct {
		// CancelWorkflow holds details about calls to the CancelWorkflow method.
		CancelWorkflow []struct {
			// WorkflowID is the workflowID argument value.
			WorkflowID string
		}
		// HandleWorkflow holds details about calls to the HandleWorkflow method.
		HandleWorkflow []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// WorkflowMoqParam is the workflowMoqParam argument value.
			WorkflowMoqParam workflow.Workflow
			// Recorder is the recorder argument value.
			Recorder event.Recorder
		}
	}
	lockCancelWorkflow sync.RWMutex
	lockHandleWorkflow sync.RWMutex
}

// CancelWorkflow calls CancelWorkflowFunc.
func (mock *WorkflowHandlerMock) CancelWorkflow(workflowID string) {
	if mock.CancelWorkflowFunc == nil {
		panic("WorkflowHandlerMock.CancelWorkflowFunc: method is nil but WorkflowHandler.CancelWorkflow was just called")
	}
	callInfo := struct {
		WorkflowID string
	}{
		WorkflowID: workflowID,
	}
	mock.lockCancelWorkflow.Lock()
	mock.calls.CancelWorkflow = append(mock.calls.CancelWorkflow, callInfo)
	mock.lockCancelWorkflow.Unlock()
	mock.CancelWorkflowFunc(workflowID)
}

// CancelWorkflowCalls gets all the calls that were made to CancelWorkflow.
// Check the length with:
//
//	len(mockedWorkflowHandler.CancelWorkflowCalls())
func (mock *WorkflowHandlerMock) CancelWorkflowCalls() []struct {
	WorkflowID string
} {
	var calls []struct {
		WorkflowID string
	}
	mock.lockCancelWorkflow.RLock()
	calls = mock.calls.CancelWorkflow
	mock.lockCancelWorkflow.RUnlock()
	return calls
}

// HandleWorkflow calls HandleWorkflowFunc.
func (mock *WorkflowHandlerMock) HandleWorkflow(contextMoqParam context.Context, workflowMoqParam workflow.Workflow, recorder event.Recorder) {
	if mock.HandleWorkflowFunc == nil {
		panic("WorkflowHandlerMock.HandleWorkflowFunc: method is nil but WorkflowHandler.HandleWorkflow was just called")
	}
	callInfo := struct {
		ContextMoqParam  context.Context
		WorkflowMoqParam workflow.Workflow
		Recorder         event.Recorder
	}{
		ContextMoqParam:  contextMoqParam,
		WorkflowMoqParam: workflowMoqParam,
		Recorder:         recorder,
	}
	mock.lockHandleWorkflow.Lock()
	mock.calls.HandleWorkflow = append(mock.calls.HandleWorkflow, callInfo)
	mock.lockHandleWorkflow.Unlock()
	mock.HandleWorkflowFunc(contextMoqParam, workflowMoqParam, recorder)
}

// HandleWorkflowCalls gets all the calls that were made to HandleWorkflow.
// Check the length with:
//
//	len(mockedWorkflowHandler.HandleWorkflowCalls())
func (mock *WorkflowHandlerMock) HandleWorkflowCalls() []struct {
	ContextMoqParam  context.Context
	WorkflowMoqParam workflow.Workflow
	Recorder         event.Recorder
} {
	var calls []struct {
		ContextMoqParam  context.Context
		WorkflowMoqParam workflow.Workflow
		Recorder         event.Recorder
	}
	mock.lockHandleWorkflow.RLock()
	calls = mock.calls.HandleWorkflow
	mock.lockHandleWorkflow.RUnlock()
	return calls
}
