package metrics

import "sync"

type InMemoryMetricsDA struct {
	mapMutex   sync.Mutex
	metricsMap map[string]*Metrics
}

func NewInMemoryMetricsDA() *InMemoryMetricsDA {
	service := new(InMemoryMetricsDA)
	service.metricsMap = make(map[string]*Metrics)
	return service
}

func (da *InMemoryMetricsDA) Create(value Metrics) (*Metrics, error) {
	newValue := NewMetrics(value.LinkId)
	newValue.Copy(&value)

	returnValue := NewMetrics(value.LinkId)
	returnValue.Copy(newValue)

	da.mapMutex.Lock()
	da.metricsMap[newValue.LinkId] = newValue
	da.mapMutex.Unlock()

	return returnValue, nil
}

func (da *InMemoryMetricsDA) Get(linkId string) (*Metrics, error) {
	var storedValue *Metrics
	var found bool
	var returnError error = nil
	returnValue := NewMetrics(linkId)

	da.mapMutex.Lock()

	storedValue, found = da.metricsMap[linkId]
	if found {
		returnValue.Copy(storedValue)
	} else {
		returnValue = nil
		returnError = ErrNotFound{}
	}

	da.mapMutex.Unlock()

	return returnValue, returnError
}

func (da *InMemoryMetricsDA) AddToAccessCount(linkId string, amount uint64) error {
	var returnError error = nil

	da.mapMutex.Lock()
	storedValue, found := da.metricsMap[linkId]
	if found {
		storedValue.TimesAccessed += amount
	} else {
		returnError = &ErrNotFound{}
	}
	da.mapMutex.Unlock()

	return returnError
}
