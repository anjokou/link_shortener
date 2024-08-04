package analytics

import (
	syserrors "link_shortener/errors"
	"link_shortener/links"
	"link_shortener/settings"
	"net/rpc"
	"strconv"
	"time"
)

type RemoteAnalytics struct {
	errorHandler syserrors.ErrorHandler
	rpcClient    *rpc.Client
}

type CreateArgs struct {
	LinkId           string
	CreatedTimestamp int64
}

type AddAccessArgs struct {
	LinkId    string
	AddAmount uint64
}

func CreateRemoteAnalytics(errorHandler syserrors.ErrorHandler) *RemoteAnalytics {
	service := new(RemoteAnalytics)
	service.errorHandler = errorHandler

	analyticsServerAddress := settings.AnalyticsHost() + ":" + strconv.Itoa(settings.AnalyticsPort())
	client, err := rpc.Dial("tcp", analyticsServerAddress)
	if err != nil {
		panic(err)
	}

	service.rpcClient = client

	return service
}

func (engine *RemoteAnalytics) Close() {
	engine.rpcClient.Close()
}

func (engine *RemoteAnalytics) Created(link links.Link) {
	var result int
	args := &CreateArgs{
		LinkId:           link.Id,
		CreatedTimestamp: time.Now().UnixNano() / int64(time.Millisecond),
	}

	err := engine.rpcClient.Call("MetricsRpcActions.AddLink", args, &result)

	if err != nil {
		engine.errorHandler.AnalyticsServiceError(err)
	}
}

func (engine *RemoteAnalytics) Accessed(link links.Link, timesAccessed int) {

	var result int
	args := &AddAccessArgs{
		LinkId:    link.Id,
		AddAmount: 1,
	}

	err := engine.rpcClient.Call("MetricsRpcActions.AddToAccessedCount", args, &result)

	if err != nil {
		engine.errorHandler.AnalyticsServiceError(err)
	}
}
