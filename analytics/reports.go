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
	"encoding/json"
	"strings"
)

// ReportsService handles reports.
// Analytics docs: https://www.adobe.io/apis/experiencecloud/analytics/docs.html#!AdobeDocs/analytics-2.0-apis/master/reporting-guide.md
type ReportsService struct {
	client *Client
}

// Run runs a report for the passed RankedRequest.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/reports
func (s *ReportsService) Run(rankedRequest *RankedRequest) (*RankedReportData, error) {
	reqJSON, _ := json.Marshal(rankedRequest)
	reqBody := strings.NewReader(string(reqJSON))

	var data RankedReportData
	err := s.client.post("/reports", map[string]string{}, reqBody, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
