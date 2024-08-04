package actions

import "link_shortener_analytics/metrics"

var metricsStorage metrics.MetricsDA

func InitMetricActions(injectDA metrics.MetricsDA) {
	metricsStorage = injectDA
}

func Create(value *metrics.Metrics) (*metrics.Metrics, error) {
	return metricsStorage.Create(*value)
}

func Get(linkId string) (*metrics.Metrics, error) {
	return metricsStorage.Get(linkId)
}

func AddToAccessCount(linkId string, amount uint64) error {
	return metricsStorage.AddToAccessCount(linkId, amount)
}
