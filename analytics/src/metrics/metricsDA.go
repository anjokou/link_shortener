package metrics

type MetricsDA interface {
	Create(value Metrics) (*Metrics, error)
	Get(linkId string) (*Metrics, error)
	AddToAccessCount(linkId string, amount uint64) error
}

type ErrNotFound struct {
}

func (err ErrNotFound) Error() string {
	return "no metrics for requested link id"
}
