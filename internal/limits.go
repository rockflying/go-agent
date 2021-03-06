package internal

import "time"

const (
	// app behavior
	connectBackoff            = 20 * time.Second
	harvestPeriod             = 60 * time.Second
	collectorTimeout          = 20 * time.Second
	appDataChanSize           = 200
	failedMetricAttemptsLimit = 5
	failedEventsAttemptsLimit = 10

	// transaction behavior
	maxStackTraceFrames = 100
	maxTxnErrors        = 5

	// harvest data
	maxMetrics       = 2 * 1000
	maxCustomEvents  = 10 * 1000
	maxTxnEvents     = 10 * 1000
	maxErrorEvents   = 100
	maxHarvestErrors = 20

	// attributes
	attributeKeyLengthLimit   = 255
	attributeValueLengthLimit = 255
	attributeUserLimit        = 64
	attributeAgentLimit       = 255 - attributeUserLimit
	customEventAttributeLimit = 64

	// Limits affecting Config validation are found in the config package.

	// Runtime metrics should not depend on the sampler period, but the
	// period must be the same across instances.  For that reason, this
	// value should not be changed without notifying customers that they
	// must update all instance simultaneously for valid runtime metrics.
	runtimeSamplerPeriod = 60 * time.Second
)
