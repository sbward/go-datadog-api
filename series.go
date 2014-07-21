/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import "encoding/json"

// DataPoint is a tuple of [UNIX timestamp, value]. This has to use interface{}
// because the value could be non-integer.
type DataPoint struct {
	Time  time.Time
	Value float64
}

func (p *Datapoint) MarshalJSON() ([]byte, error) {
	return json.Marshal([2]interface{}{
		p.Time,
		p.Value,
	})
}

// Metric represents a collection of data points that we might send or receive
// on one single metric line.
type Metric struct {
	Metric string      `json:"metric"`
	Points []DataPoint `json:"points"`
	Type   string      `json:"type"`
	Host   string      `json:"host"`
	Tags   []string    `json:"tags"`
}

// reqPostSeries from /api/v1/series
type reqPostSeries struct {
	Series []Metric `json:"series"`
}

// PostSeries takes as input a slice of metrics and then posts them up to the
// server for posting data.
func (self *Client) PostMetrics(series []Metric) error {
	return self.doJsonRequest("POST", "/v1/series",
		reqPostSeries{Series: series}, nil)
}

func (self *Client) PostMetricsResp(series []Metric) (map[string]interface{}, error) {
	var resp map[string]interface{}
	err := self.doJsonRequest("POST", "/v1/series",
		reqPostSeries{Series: series}, &resp)
	return resp, err
}
