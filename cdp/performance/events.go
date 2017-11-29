package performance

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// EventMetrics current values of the metrics.
type EventMetrics struct {
	Metrics []*Metric `json:"metrics"` // Current values of the metrics.
	Title   string    `json:"title"`   // Timestamp title.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventPerformanceMetrics,
}
