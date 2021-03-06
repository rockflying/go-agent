package internal

const (
	apdexRollup = "Apdex"
	apdexPrefix = "Apdex/"

	webRollup        = "WebTransaction"
	backgroundRollup = "OtherTransaction/all"

	errorsAll        = "Errors/all"
	errorsWeb        = "Errors/allWeb"
	errorsBackground = "Errors/allOther"
	errorsPrefix     = "Errors/"

	// "HttpDispatcher" metric is used for the overview graph, and
	// therefore should only be made for web transactions.
	dispatcherMetric = "HttpDispatcher"

	queueMetric = "WebFrontend/QueueTime"

	webMetricPrefix        = "WebTransaction/Go"
	backgroundMetricPrefix = "OtherTransaction/Go"

	instanceReporting = "Instance/Reporting"

	// https://newrelic.atlassian.net/wiki/display/eng/Custom+Events+in+New+Relic+Agents
	customEventsSeen = "Supportability/Events/Customer/Seen"
	customEventsSent = "Supportability/Events/Customer/Sent"

	// https://source.datanerd.us/agents/agent-specs/blob/master/Transaction-Events-PORTED.md
	txnEventsSeen = "Supportability/AnalyticsEvents/TotalEventsSeen"
	txnEventsSent = "Supportability/AnalyticsEvents/TotalEventsSent"

	// https://source.datanerd.us/agents/agent-specs/blob/master/Error-Events.md
	errorEventsSeen = "Supportability/Events/TransactionError/Seen"
	errorEventsSent = "Supportability/Events/TransactionError/Sent"

	supportabilityDropped = "Supportability/MetricsDropped"

	customSegmentPrefix = "Custom/"

	// source.datanerd.us/agents/agent-specs/blob/master/Datastore-Metrics-PORTED.md
	datastoreAll   = "Datastore/all"
	datastoreWeb   = "Datastore/allWeb"
	datastoreOther = "Datastore/allOther"

	// source.datanerd.us/agents/agent-specs/blob/master/APIs/external_segment.md
	// source.datanerd.us/agents/agent-specs/blob/master/APIs/external_cat.md
	// source.datanerd.us/agents/agent-specs/blob/master/Cross-Application-Tracing-PORTED.md
	externalAll   = "External/all"
	externalWeb   = "External/allWeb"
	externalOther = "External/allOther"

	// Runtime/System Metrics
	memoryPhysical       = "Memory/Physical"
	cpuUserUtilization   = "CPU/User/Utilization"
	cpuSystemUtilization = "CPU/System/Utilization"
	cpuUserTime          = "CPU/User Time"
	cpuSystemTime        = "CPU/System Time"
	runGoroutine         = "Go/Runtime/Goroutines"
	gcPauseFraction      = "GC/System/Pause Fraction"
	gcPauses             = "GC/System/Pauses"
)

type datastoreProductMetrics struct {
	All   string // Datastore/{datastore}/all
	Web   string // Datastore/{datastore}/allWeb
	Other string // Datastore/{datastore}/allOther
}

func datastoreProductMetric(key datastoreMetricKey) datastoreProductMetrics {
	d, ok := datastoreProductMetricsCache[key.Product]
	if ok {
		return d
	}
	return datastoreProductMetrics{
		All:   "Datastore/" + string(key.Product) + "/all",
		Web:   "Datastore/" + string(key.Product) + "/allWeb",
		Other: "Datastore/" + string(key.Product) + "/allOther",
	}
}

// Datastore/operation/{datastore}/{operation}
func datastoreOperationMetric(key datastoreMetricKey) string {
	return "Datastore/operation/" + string(key.Product) +
		"/" + key.Operation
}

// Datastore/statement/{datastore}/{table}/{operation}
func datastoreStatementMetric(key datastoreMetricKey) string {
	return "Datastore/statement/" + string(key.Product) +
		"/" + key.Collection +
		"/" + key.Operation
}

// External/{host}/all
func externalHostMetric(key externalMetricKey) string {
	return "External/" + key.Host + "/all"
}

// ExternalApp/{host}/{external_id}/all
func externalAppMetric(key externalMetricKey) string {
	return "ExternalApp/" + key.Host +
		"/" + key.ExternalCrossProcessID + "/all"
}

// ExternalTransaction/{host}/{external_id}/{external_txnname}
func externalTransactionMetric(key externalMetricKey) string {
	return "ExternalTransaction/" + key.Host +
		"/" + key.ExternalCrossProcessID +
		"/" + key.ExternalTransactionName
}
