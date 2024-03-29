// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package auth

import (
	"sync"
)

// Ensure, that ShortLiveTokenGeneratorMock does implement ShortLiveTokenGenerator.
// If this is not the case, regenerate this file with moq.
var _ ShortLiveTokenGenerator = &ShortLiveTokenGeneratorMock{}

// ShortLiveTokenGeneratorMock is a mock implementation of ShortLiveTokenGenerator.
//
//	func TestSomethingThatUsesShortLiveTokenGenerator(t *testing.T) {
//
//		// make and configure a mocked ShortLiveTokenGenerator
//		mockedShortLiveTokenGenerator := &ShortLiveTokenGeneratorMock{
//			GenerateFunc: func(payload TokenPayload) (string, error) {
//				panic("mock out the Generate method")
//			},
//		}
//
//		// use mockedShortLiveTokenGenerator in code that requires ShortLiveTokenGenerator
//		// and then make assertions.
//
//	}
type ShortLiveTokenGeneratorMock struct {
	// GenerateFunc mocks the Generate method.
	GenerateFunc func(payload TokenPayload) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Generate holds details about calls to the Generate method.
		Generate []struct {
			// Payload is the payload argument value.
			Payload TokenPayload
		}
	}
	lockGenerate sync.RWMutex
}

// Generate calls GenerateFunc.
func (mock *ShortLiveTokenGeneratorMock) Generate(payload TokenPayload) (string, error) {
	if mock.GenerateFunc == nil {
		panic("ShortLiveTokenGeneratorMock.GenerateFunc: method is nil but ShortLiveTokenGenerator.Generate was just called")
	}
	callInfo := struct {
		Payload TokenPayload
	}{
		Payload: payload,
	}
	mock.lockGenerate.Lock()
	mock.calls.Generate = append(mock.calls.Generate, callInfo)
	mock.lockGenerate.Unlock()
	return mock.GenerateFunc(payload)
}

// GenerateCalls gets all the calls that were made to Generate.
// Check the length with:
//
//	len(mockedShortLiveTokenGenerator.GenerateCalls())
func (mock *ShortLiveTokenGeneratorMock) GenerateCalls() []struct {
	Payload TokenPayload
} {
	var calls []struct {
		Payload TokenPayload
	}
	mock.lockGenerate.RLock()
	calls = mock.calls.Generate
	mock.lockGenerate.RUnlock()
	return calls
}
