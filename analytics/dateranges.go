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

// DateRangesService handles date ranges.
// Analytics docs: https://adobedocs.github.io/analytics-2.0-apis/#/dateranges
type DateRangesService struct {
	client *Client
}

// GetAll returns a list of date ranges for the user.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/dateranges/getDateRanges
func (s *DateRangesService) GetAll(locale, filterByIDs string, limit, page int64, expansion, includeType []string) (*DateRanges, error) {
	var params = map[string]string{}
	if locale != "" {
		params["locale"] = locale
	}
	if filterByIDs != "" {
		params["filterByIds"] = filterByIDs
	}
	params["limit"] = strconv.FormatInt(limit, 10)
	params["page"] = strconv.FormatInt(page, 10)
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}
	if len(includeType) > 0 {
		params["includeType"] = strings.Join(includeType[:], ",")
	}

	var data DateRanges
	err := s.client.get("/dateranges", params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// GetByID returns configuration for a date range.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/dateranges/getDateRange
func (s *DateRangesService) GetByID(id, locale string, expansion []string) (*DateRange, error) {
	var params = map[string]string{}
	if locale != "" {
		params["locale"] = locale
	}
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}

	var data DateRange
	err := s.client.get(fmt.Sprintf("/dateranges/%s", id), params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
