// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "userbirthday/model"

	mock "github.com/stretchr/testify/mock"
)

// PromoRepository is an autogenerated mock type for the PromoRepository type
type PromoRepository struct {
	mock.Mock
}

// CreatePromo provides a mock function with given fields: ctx, m
func (_m *PromoRepository) CreatePromo(ctx context.Context, m model.Promo) (model.Promo, error) {
	ret := _m.Called(ctx, m)

	var r0 model.Promo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Promo) (model.Promo, error)); ok {
		return rf(ctx, m)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Promo) model.Promo); ok {
		r0 = rf(ctx, m)
	} else {
		r0 = ret.Get(0).(model.Promo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Promo) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPromoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPromoRepository creates a new instance of PromoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPromoRepository(t mockConstructorTestingTNewPromoRepository) *PromoRepository {
	mock := &PromoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
