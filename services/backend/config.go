package backend

import (
	"time"
)

var (
	// MySQLGetDelay is how long retrieving a record takes.
	// Using large value mostly because I cannot click the button fast enough to cause a queue.
	MySQLGetDelay = 300 * time.Millisecond

	// MySQLGetDelayStdDev is standard deviation
	MySQLGetDelayStdDev = MySQLGetDelay / 10

	// MySQLMutexDisabled controls whether there is a mutex guarding db query execution.
	// When not disabled it simulates a misconfigured connection pool of size 1.
	MySQLMutexDisabled = false
)
