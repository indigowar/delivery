// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package menu

import (
	"context"
	"github.com/google/uuid"
	"github.com/indigowar/delivery/internal/entities"
	"sync"
)

// Ensure, that DishStorageMock does implement DishStorage.
// If this is not the case, regenerate this file with moq.
var _ DishStorage = &DishStorageMock{}

// DishStorageMock is a mock implementation of DishStorage.
//
//	func TestSomethingThatUsesDishStorage(t *testing.T) {
//
//		// make and configure a mocked DishStorage
//		mockedDishStorage := &DishStorageMock{
//			AddFunc: func(ctx context.Context, dish *entities.Dish) (*entities.Dish, error) {
//				panic("mock out the Add method")
//			},
//			GetFunc: func(ctx context.Context, id uuid.UUID) (*entities.Dish, error) {
//				panic("mock out the Get method")
//			},
//			GetSetFunc: func(ctx context.Context, ids uuid.UUIDs) ([]*entities.Dish, error) {
//				panic("mock out the GetSet method")
//			},
//			RemoveFunc: func(ctx context.Context, id uuid.UUID) (*entities.Dish, error) {
//				panic("mock out the Remove method")
//			},
//		}
//
//		// use mockedDishStorage in code that requires DishStorage
//		// and then make assertions.
//
//	}
type DishStorageMock struct {
	// AddFunc mocks the Add method.
	AddFunc func(ctx context.Context, dish *entities.Dish) (*entities.Dish, error)

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, id uuid.UUID) (*entities.Dish, error)

	// GetSetFunc mocks the GetSet method.
	GetSetFunc func(ctx context.Context, ids uuid.UUIDs) ([]*entities.Dish, error)

	// RemoveFunc mocks the Remove method.
	RemoveFunc func(ctx context.Context, id uuid.UUID) (*entities.Dish, error)

	// calls tracks calls to the methods.
	calls struct {
		// Add holds details about calls to the Add method.
		Add []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Dish is the dish argument value.
			Dish *entities.Dish
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// GetSet holds details about calls to the GetSet method.
		GetSet []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Ids is the ids argument value.
			Ids uuid.UUIDs
		}
		// Remove holds details about calls to the Remove method.
		Remove []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
	}
	lockAdd    sync.RWMutex
	lockGet    sync.RWMutex
	lockGetSet sync.RWMutex
	lockRemove sync.RWMutex
}

// Add calls AddFunc.
func (mock *DishStorageMock) Add(ctx context.Context, dish *entities.Dish) (*entities.Dish, error) {
	if mock.AddFunc == nil {
		panic("DishStorageMock.AddFunc: method is nil but DishStorage.Add was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Dish *entities.Dish
	}{
		Ctx:  ctx,
		Dish: dish,
	}
	mock.lockAdd.Lock()
	mock.calls.Add = append(mock.calls.Add, callInfo)
	mock.lockAdd.Unlock()
	return mock.AddFunc(ctx, dish)
}

// AddCalls gets all the calls that were made to Add.
// Check the length with:
//
//	len(mockedDishStorage.AddCalls())
func (mock *DishStorageMock) AddCalls() []struct {
	Ctx  context.Context
	Dish *entities.Dish
} {
	var calls []struct {
		Ctx  context.Context
		Dish *entities.Dish
	}
	mock.lockAdd.RLock()
	calls = mock.calls.Add
	mock.lockAdd.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *DishStorageMock) Get(ctx context.Context, id uuid.UUID) (*entities.Dish, error) {
	if mock.GetFunc == nil {
		panic("DishStorageMock.GetFunc: method is nil but DishStorage.Get was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedDishStorage.GetCalls())
func (mock *DishStorageMock) GetCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// GetSet calls GetSetFunc.
func (mock *DishStorageMock) GetSet(ctx context.Context, ids uuid.UUIDs) ([]*entities.Dish, error) {
	if mock.GetSetFunc == nil {
		panic("DishStorageMock.GetSetFunc: method is nil but DishStorage.GetSet was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Ids uuid.UUIDs
	}{
		Ctx: ctx,
		Ids: ids,
	}
	mock.lockGetSet.Lock()
	mock.calls.GetSet = append(mock.calls.GetSet, callInfo)
	mock.lockGetSet.Unlock()
	return mock.GetSetFunc(ctx, ids)
}

// GetSetCalls gets all the calls that were made to GetSet.
// Check the length with:
//
//	len(mockedDishStorage.GetSetCalls())
func (mock *DishStorageMock) GetSetCalls() []struct {
	Ctx context.Context
	Ids uuid.UUIDs
} {
	var calls []struct {
		Ctx context.Context
		Ids uuid.UUIDs
	}
	mock.lockGetSet.RLock()
	calls = mock.calls.GetSet
	mock.lockGetSet.RUnlock()
	return calls
}

// Remove calls RemoveFunc.
func (mock *DishStorageMock) Remove(ctx context.Context, id uuid.UUID) (*entities.Dish, error) {
	if mock.RemoveFunc == nil {
		panic("DishStorageMock.RemoveFunc: method is nil but DishStorage.Remove was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockRemove.Lock()
	mock.calls.Remove = append(mock.calls.Remove, callInfo)
	mock.lockRemove.Unlock()
	return mock.RemoveFunc(ctx, id)
}

// RemoveCalls gets all the calls that were made to Remove.
// Check the length with:
//
//	len(mockedDishStorage.RemoveCalls())
func (mock *DishStorageMock) RemoveCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockRemove.RLock()
	calls = mock.calls.Remove
	mock.lockRemove.RUnlock()
	return calls
}