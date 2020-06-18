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

// DimensionsService handles dimensions.
// Analytics docs: https://adobedocs.github.io/analytics-2.0-apis/#/dimensions
type DimensionsService struct {
	client *Client
}

// GetAll returns a list of dimensions for a given report suite.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/dimensions/dimensions_getDimensions
func (s *DimensionsService) GetAll(rsID, locale string, segmentable, reportable, classifiable bool, expansion []string) (*[]Dimension, error) {
	var params = map[string]string{}
	params["rsid"] = rsID
	if locale != "" {
		params["locale"] = locale
	}
	params["segmentable"] = strconv.FormatBool(segmentable)
	params["reportable"] = strconv.FormatBool(reportable)
	params["classifiable"] = strconv.FormatBool(classifiable)
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}

	var data []Dimension
	err := s.client.get("/dimensions", params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// GetByID returns a dimension for a given report suite.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/dimensions/dimensions_getDimension
func (s *DimensionsService) GetByID(rsID, id, locale string, expansion []string) (*Dimension, error) {
	var params = map[string]string{}
	params["rsid"] = rsID
	if locale != "" {
		params["locale"] = locale
	}
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}

	var data Dimension
	err := s.client.get(fmt.Sprintf("/dimensions/%s", id), params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
