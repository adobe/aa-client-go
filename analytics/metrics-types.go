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

// Response types

// Metric represents a metric
type Metric struct {
	ID                  string   `json:"id,omitempty"`
	Title               string   `json:"title,omitempty"`
	Name                string   `json:"name,omitempty"`
	Type                string   `json:"type,omitempty"`
	ExtraTitleInfo      string   `json:"extraTitleInfo,omitempty"`
	Category            string   `json:"category,omitempty"`
	Categories          []string `json:"categories,omitempty"`
	Support             []string `json:"support,omitempty"`
	Allocation          bool     `json:"allocation"`
	Precision           int      `json:"precision"`
	Calculated          bool     `json:"calculated"`
	Segmentable         bool     `json:"segmentable"`
	Description         string   `json:"description,omitempty"`
	Polarity            string   `json:"polarity,omitempty"`
	HelpLink            string   `json:"helpLink,omitempty"`
	AllowedForReporting bool     `json:"allowedForReporting"`
	Tags                *[]Tag   `json:"tags,omitempty"`
}
