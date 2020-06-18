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

// SegmentDefinitionContainerPredicateValue represents a segment definition predicate value
type SegmentDefinitionContainerPredicateValue struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Func        string `json:"func,omitempty"`
}

// SegmentDefinitionContainerPredicate represents a segment definition predicate
type SegmentDefinitionContainerPredicate struct {
	Val  *SegmentDefinitionContainerPredicateValue `json:"val,omitempty"`
	Str  string                                    `json:"str,omitempty"`
	Func string                                    `json:"func,omitempty"`
}

// SegmentDefinitionContainer represents a segment definition container
type SegmentDefinitionContainer struct {
	Context string `json:"context,omitempty"`
	Func    string `json:"func,omitempty"`
	Pred    string `json:"pred,omitempty"`
}

// SegmentDefinition represents a segment definition
type SegmentDefinition struct {
	Container *SegmentDefinitionContainer `json:"container,omitempty"`
	Func      string                      `json:"func,omitempty"`
	Version   []int                       `json:"version"`
}

// SegmentCompatibility represents segment compatibility settings
type SegmentCompatibility struct {
	Valid              bool     `json:"valid"`
	Message            string   `json:"message,omitempty"`
	ValidatorVersion   string   `json:"validator_version,omitempty"`
	SupportedProducts  []string `json:"supported_products,omitempty"`
	SupportedSchema    []string `json:"supported_schema,omitempty"`
	SuppportedFeatures []string `json:"supported_features,omitempty"`
}

// Segment represents a segment
type Segment struct {
	ID                     string                `json:"id,omitempty"`
	Name                   string                `json:"name,omitempty"`
	Description            string                `json:"description,omitempty"`
	ReportSuiteID          string                `json:"rsid,omitempty"`
	ReportSuiteName        string                `json:"reportSuiteName,omitempty"`
	Owner                  *Owner                `json:"owner,omitempty"`
	Definition             *SegmentDefinition    `json:"definition,omitempty"`
	Compatibility          *SegmentCompatibility `json:"compatibility,omitempty"`
	DefinitionLastModified string                `json:"definitionLastModified,omitempty"`
	Categories             []string              `json:"categories,omitempty"`
	SiteTitle              string                `json:"siteTitle,omitempty"`
	Tags                   *[]Tag                `json:"tags,omitempty"`
	Modified               string                `json:"modified,omitempty"`
	Created                string                `json:"created,omitempty"`
}

// Segments represents a page of segments
type Segments struct {
	Content          *[]Segment `json:"content,omitempty"`
	Number           int        `json:"number"`
	Size             int        `json:"size"`
	NumberOfElements int        `json:"numberOfElements"`
	TotalElements    int        `json:"totalElements"`
	PreviousPage     bool       `json:"previousPage"`
	FirstPage        bool       `json:"firstPage"`
	NextPage         bool       `json:"nextPage"`
	LastPage         bool       `json:"lastPage"`
	Sort             *[]Sort    `json:"sort,omitempty"`
	TotalPages       int        `json:"totalPages"`
}
