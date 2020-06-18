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

// CalculatedMetricsService handles calculated metrics.
// Analytics docs: https://adobedocs.github.io/analytics-2.0-apis/#/calculatedmetrics
type CalculatedMetricsService struct {
	client *Client
}

// GetAll returns a list of calculated metrics that match the given filters.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/calculatedmetrics/findCalculatedMetrics
func (s *CalculatedMetricsService) GetAll(rsids, ownerID, filterByIds, toBeUsedInRsid, locale, name, tagNames string,
	favorite, approved bool,
	limit, page int64,
	sortDirection, sortProperty string,
	expansion, includeType []string) (*CalculatedMetrics, error) {

	var params = map[string]string{}
	if rsids != "" {
		params["rsids"] = rsids
	}
	if ownerID != "" {
		params["ownerId"] = ownerID
	}
	if filterByIds != "" {
		params["filterByIds"] = filterByIds
	}
	if toBeUsedInRsid != "" {
		params["toBeUsedInRsid"] = toBeUsedInRsid
	}
	if locale != "" {
		params["locale"] = locale
	}
	if name != "" {
		params["name"] = name
	}
	if tagNames != "" {
		params["tagNames"] = tagNames
	}
	params["favorite"] = strconv.FormatBool(favorite)
	params["approved"] = strconv.FormatBool(approved)
	params["limit"] = strconv.FormatInt(limit, 10)
	params["page"] = strconv.FormatInt(page, 10)
	if sortDirection != "" {
		params["sortDirection"] = sortDirection
	}
	if sortProperty != "" {
		params["sortProperty"] = sortProperty
	}
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}
	if len(includeType) > 0 {
		params["includeType"] = strings.Join(includeType[:], ",")
	}

	var data CalculatedMetrics
	err := s.client.get("/calculatedmetrics", params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// GetByID returns a single calculated metric by ID.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/calculatedmetrics/findOneCalculatedMetric
func (s *CalculatedMetricsService) GetByID(id, locale string, expansion []string) (*CalculatedMetric, error) {
	var params = map[string]string{}
	if locale != "" {
		params["locale"] = locale
	}
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}

	var data CalculatedMetric
	err := s.client.get(fmt.Sprintf("/calculatedmetrics/%s", id), params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
