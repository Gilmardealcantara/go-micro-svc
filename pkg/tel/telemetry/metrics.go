package telemetry

func RecordMetric(name string, value float64) {
	App.RecordCustomMetric(name, value)
}

func IncrementMetric(name string) {
	App.RecordCustomMetric(name, 1)
}

func DecrementMetric(name string) {
	App.RecordCustomMetric(name, -1)
}
