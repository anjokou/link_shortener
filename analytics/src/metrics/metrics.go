package metrics

type Metrics struct {
	LinkId        string `json:"linkId"`
	Created       int64  `json:"created"`
	TimesAccessed uint64 `json:"timesAccessed"`
}

func NewMetrics(linkId string) *Metrics {
	value := new(Metrics)
	value.LinkId = linkId
	value.TimesAccessed = 0

	return value
}

func (metric *Metrics) Copy(target *Metrics) {
	*metric = *target
}
