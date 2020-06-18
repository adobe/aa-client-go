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

// SegmentsService handles segments.
// Analytics docs: https://adobedocs.github.io/analytics-2.0-apis/#/segments
type SegmentsService struct {
	client *Client
}

// GetAll returns all segments.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/segments/segments_getSegments
func (s *SegmentsService) GetAll(rsids, segmentFilter, locale, name, tagNames, filterByPublishedSegments string,
	limit, page int64, sortDirection, sortProperty string,
	expansion []string, includeType []string) (*Segments, error) {

	var params = map[string]string{}
	params["rsids"] = rsids
	if segmentFilter != "" {
		params["segmentFilter"] = segmentFilter
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
	if filterByPublishedSegments != "" {
		params["filterByPublishedSegments"] = filterByPublishedSegments
	}
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

	var data Segments
	err := s.client.get("/segments", params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// GetByID returns a single segment.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/segments/segments_getSegment
func (s *SegmentsService) GetByID(id, locale string, expansion []string) (*Segment, error) {
	var params = map[string]string{}
	if locale != "" {
		params["locale"] = locale
	}
	if len(expansion) > 0 {
		params["expansion"] = strings.Join(expansion[:], ",")
	}

	var data Segment
	err := s.client.get(fmt.Sprintf("/segments/%s", id), params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err

}
