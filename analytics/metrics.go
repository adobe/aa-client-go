/*
Copyright 2020 Adobe. All rights reserved.
This file is licensed to you under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License. You may obtain a copy
of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
OF ANY KIND, either express or implied. See the License for the specific language
governing permissions and limitations under the License.
*/

package analytics

import (
	"fmt"
	"strconv"
	"strings"
)

// MetricsService handles metrics.
// Analytics docs: https://adobedocs.github.io/analytics-2.0-apis/#/metrics
type MetricsService struct {
	client *Client
}

// GetAll returns a list of metrics for a given report suite.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/metrics/getMetrics
func (s *MetricsService) GetAll(rsID, locale string, segmentable bool, expansion []string) (*[]Metric, error) {
	var params = map[string]string{}
	params["rsid"] = rsID
	if locale != "" {
		params["locale"] = locale
	}
	params["segmentable"] = strconv.FormatBool(segmentable)
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}

	var data []Metric
	err := s.client.get("/metrics", params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// GetByID returns a metric for a given report suite.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/metrics/getMetric
func (s *MetricsService) GetByID(rsID, id, locale string, expansion []string) (*Metric, error) {
	var params = map[string]string{}
	params["rsid"] = rsID
	if locale != "" {
		params["locale"] = locale
	}
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}

	var data Metric
	err := s.client.get(fmt.Sprintf("/metrics/%s", id), params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
