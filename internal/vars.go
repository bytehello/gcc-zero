package internal

import "time"

const (
	endpointsSeparator = ","
	autoSyncInterval   = time.Minute
	dialTimeout        = 5 * time.Second
	dialKeepAliveTime  = 5 * time.Second
)
