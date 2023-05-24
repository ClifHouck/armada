// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package events

import (
	"context"
	"github.com/armadaproject/armada/pkg/api"
	"github.com/gogo/protobuf/types"
	"sync"
)

// Ensure, that JobEventReaderMock does implement JobEventReader.
// If this is not the case, regenerate this file with moq.
var _ JobEventReader = &JobEventReaderMock{}

// JobEventReaderMock is a mock implementation of JobEventReader.
//
//	func TestSomethingThatUsesJobEventReader(t *testing.T) {
//
//		// make and configure a mocked JobEventReader
//		mockedJobEventReader := &JobEventReaderMock{
//			CloseFunc: func()  {
//				panic("mock out the Close method")
//			},
//			GetJobEventMessageFunc: func(ctx context.Context, jobReq *api.JobSetRequest) (api.Event_GetJobSetEventsClient, error) {
//				panic("mock out the GetJobEventMessage method")
//			},
//			HealthFunc: func(ctx context.Context, empty *types.Empty) (*api.HealthCheckResponse, error) {
//				panic("mock out the Health method")
//			},
//		}
//
//		// use mockedJobEventReader in code that requires JobEventReader
//		// and then make assertions.
//
//	}
type JobEventReaderMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func()

	// GetJobEventMessageFunc mocks the GetJobEventMessage method.
	GetJobEventMessageFunc func(ctx context.Context, jobReq *api.JobSetRequest) (api.Event_GetJobSetEventsClient, error)

	// HealthFunc mocks the Health method.
	HealthFunc func(ctx context.Context, empty *types.Empty) (*api.HealthCheckResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// GetJobEventMessage holds details about calls to the GetJobEventMessage method.
		GetJobEventMessage []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// JobReq is the jobReq argument value.
			JobReq *api.JobSetRequest
		}
		// Health holds details about calls to the Health method.
		Health []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Empty is the empty argument value.
			Empty *types.Empty
		}
	}
	lockClose              sync.RWMutex
	lockGetJobEventMessage sync.RWMutex
	lockHealth             sync.RWMutex
}

// Close calls CloseFunc.
func (mock *JobEventReaderMock) Close() {
	if mock.CloseFunc == nil {
		panic("JobEventReaderMock.CloseFunc: method is nil but JobEventReader.Close was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//
//	len(mockedJobEventReader.CloseCalls())
func (mock *JobEventReaderMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// GetJobEventMessage calls GetJobEventMessageFunc.
func (mock *JobEventReaderMock) GetJobEventMessage(ctx context.Context, jobReq *api.JobSetRequest) (api.Event_GetJobSetEventsClient, error) {
	if mock.GetJobEventMessageFunc == nil {
		panic("JobEventReaderMock.GetJobEventMessageFunc: method is nil but JobEventReader.GetJobEventMessage was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		JobReq *api.JobSetRequest
	}{
		Ctx:    ctx,
		JobReq: jobReq,
	}
	mock.lockGetJobEventMessage.Lock()
	mock.calls.GetJobEventMessage = append(mock.calls.GetJobEventMessage, callInfo)
	mock.lockGetJobEventMessage.Unlock()
	return mock.GetJobEventMessageFunc(ctx, jobReq)
}

// GetJobEventMessageCalls gets all the calls that were made to GetJobEventMessage.
// Check the length with:
//
//	len(mockedJobEventReader.GetJobEventMessageCalls())
func (mock *JobEventReaderMock) GetJobEventMessageCalls() []struct {
	Ctx    context.Context
	JobReq *api.JobSetRequest
} {
	var calls []struct {
		Ctx    context.Context
		JobReq *api.JobSetRequest
	}
	mock.lockGetJobEventMessage.RLock()
	calls = mock.calls.GetJobEventMessage
	mock.lockGetJobEventMessage.RUnlock()
	return calls
}

// Health calls HealthFunc.
func (mock *JobEventReaderMock) Health(ctx context.Context, empty *types.Empty) (*api.HealthCheckResponse, error) {
	if mock.HealthFunc == nil {
		panic("JobEventReaderMock.HealthFunc: method is nil but JobEventReader.Health was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Empty *types.Empty
	}{
		Ctx:   ctx,
		Empty: empty,
	}
	mock.lockHealth.Lock()
	mock.calls.Health = append(mock.calls.Health, callInfo)
	mock.lockHealth.Unlock()
	return mock.HealthFunc(ctx, empty)
}

// HealthCalls gets all the calls that were made to Health.
// Check the length with:
//
//	len(mockedJobEventReader.HealthCalls())
func (mock *JobEventReaderMock) HealthCalls() []struct {
	Ctx   context.Context
	Empty *types.Empty
} {
	var calls []struct {
		Ctx   context.Context
		Empty *types.Empty
	}
	mock.lockHealth.RLock()
	calls = mock.calls.Health
	mock.lockHealth.RUnlock()
	return calls
}
