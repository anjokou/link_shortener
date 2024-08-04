package actions_test

import (
	"link_shortener_analytics/actions"
	"link_shortener_analytics/metrics"
	"link_shortener_analytics/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

var metricsDA *mocks.MockMetricsDA
var testMetric = metrics.Metrics{}

func setup(t *testing.T) {
	ctrl := gomock.NewController(t)
	metricsDA = mocks.NewMockMetricsDA(ctrl)

	actions.InitMetricActions(metricsDA)
}

func TestCreateSuccess(t *testing.T) {
	setup(t)

	metricsDA.EXPECT().
		Create(gomock.Eq(testMetric)).
		Return(&testMetric, nil)

	returnedMetric, _ := actions.Create(&testMetric)

	assert.Equal(t, *returnedMetric, testMetric)
}

func TestGetSuccess(t *testing.T) {
	setup(t)

	metricsDA.EXPECT().
		Get(gomock.Eq(testMetric.LinkId)).
		Return(&testMetric, nil)

	returnedMetric, _ := actions.Get(testMetric.LinkId)
	assert.Equal(t, *returnedMetric, testMetric)
}

func TestGetNotFound(t *testing.T) {
	setup(t)

	metricsDA.EXPECT().
		Get(gomock.Eq(testMetric.LinkId)).
		Return(nil, metrics.ErrNotFound{})

	_, err := actions.Get(testMetric.LinkId)
	assert.ErrorType(t, err, metrics.ErrNotFound{})
}

func TestAddToAccessCountSuccess(t *testing.T) {
	setup(t)

	metricsDA.EXPECT().
		AddToAccessCount(gomock.Eq(testMetric.LinkId), gomock.Eq(uint64(1))).
		Return(nil)

	err := actions.AddToAccessCount(testMetric.LinkId, 1)
	assert.NilError(t, err)
}

func TestAddToAccessCountNotFound(t *testing.T) {
	setup(t)

	metricsDA.EXPECT().
		AddToAccessCount(gomock.Eq(testMetric.LinkId), gomock.Eq(uint64(1))).
		Return(metrics.ErrNotFound{})

	err := actions.AddToAccessCount(testMetric.LinkId, 1)
	assert.ErrorType(t, err, metrics.ErrNotFound{})
}
