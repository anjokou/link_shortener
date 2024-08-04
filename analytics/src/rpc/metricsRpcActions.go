package appRpc

import (
	"link_shortener_analytics/actions"
	"link_shortener_analytics/metrics"
)

type MetricsRpcActions struct{}

type CreateArgs struct {
	LinkId           string
	CreatedTimestamp int64
}

type AddAccessArgs struct {
	LinkId    string
	AddAmount uint64
}

func (a *MetricsRpcActions) AddLink(args *CreateArgs, reply *int) error {
	metric := metrics.NewMetrics(args.LinkId)
	metric.Created = args.CreatedTimestamp

	_, err := actions.Create(metric)

	return err
}

func (a *MetricsRpcActions) AddToAccessedCount(args *AddAccessArgs, reply *int) error {
	return actions.AddToAccessCount(args.LinkId, args.AddAmount)
}
