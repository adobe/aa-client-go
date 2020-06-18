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

// CollectionsService handles collections.
// Analytics docs: https://adobedocs.github.io/analytics-2.0-apis/#/collections
type CollectionsService struct {
	client *Client
}

// GetAll returns a list of report suites that match the given filters.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/collections/findAll
func (s *CollectionsService) GetAll(rsids, rsidContains string, limit, page int64, expansion []string) (*Collections, error) {
	var params = map[string]string{}
	if rsids != "" {
		params["rsids"] = rsids
	}
	if rsidContains != "" {
		params["rsidContains"] = rsidContains
	}
	params["limit"] = strconv.FormatInt(limit, 10)
	params["page"] = strconv.FormatInt(page, 10)
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}

	var data Collections
	err := s.client.get("/collections/suites", params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// GetByID returns a report suite by ID.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/collections/findOne
func (s *CollectionsService) GetByID(id string, expansion []string) (*Collection, error) {
	var params = map[string]string{}
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}

	var data Collection
	err := s.client.get(fmt.Sprintf("/collections/suites/%s", id), params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
