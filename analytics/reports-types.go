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

// Request types

// RankedRequestLocale represents a ranked request locale
type RankedRequestLocale struct {
	Language                string   `json:"language,omitempty"`
	Script                  string   `json:"script,omitempty"`
	Country                 string   `json:"country,omitempty"`
	Variant                 string   `json:"variant,omitempty"`
	ExtensionKey            []string `json:"extensionKeys,omitempty"`
	UnicodeLocaleAttributes []string `json:"unicodeLocaleAttributes,omitempty"`
	UnicodeLocaleKeys       []string `json:"unicodeLocaleKeys,omitempty"`
	Iso3Language            string   `json:"iso3Language,omitempty"`
	Iso3Country             string   `json:"iso3Country,omitempty"`
	DisplayLanguage         string   `json:"displayLanguage,omitempty"`
	DisplayScript           string   `json:"displayScript,omitempty"`
	DisplayCountry          string   `json:"displayCountry,omitempty"`
	DisplayVariant          string   `json:"displayVariant,omitempty"`
	DisplayName             string   `json:"displayName,omitempty"`
}

// RankedRequestReportFilter represents a ranked request report filter
type RankedRequestReportFilter struct {
	ID                string                          `json:"id,omitempty"`
	Type              string                          `json:"type,omitempty"`
	Dimension         string                          `json:"dimension,omitempty"`
	ItemID            string                          `json:"itemId,omitempty"`
	ItemIDs           []string                        `json:"itemIds,omitempty"`
	SegmentID         string                          `json:"segmentId,omitempty"`
	SegmentDefinition *RankedRequestSegmentDefinition `json:"segmentDefinition,omitempty"`
	DateRange         string                          `json:"dateRange,omitempty"`
	ExcludeItemIDs    []string                        `json:"excludeItemIds,omitempty"`
}

// RankedRequestSegmentDefinition represents a segment definition
type RankedRequestSegmentDefinition struct {
	Container *RankedRequestSegmentDefinitionContainer `json:"container,omitempty"`
	Func      string                                   `json:"func,omitempty"`
	Version   []int                                    `json:"version,omitempty"`
}

// RankedRequestSegmentDefinitionContainer represents a segment definition container
type RankedRequestSegmentDefinitionContainer struct {
	Func    string                                   `json:"func,omitempty"`
	Context string                                   `json:"context,omitempty"`
	Pred    *RankedRequestSegmentDefinitionPredicate `json:"pred,omitempty"`
}

// RankedRequestSegmentDefinitionPredicate represents a segment definition predicate
type RankedRequestSegmentDefinitionPredicate struct {
	Func        string                                        `json:"func,omitempty"`
	Val         *RankedRequestSegmentDefinitionPredicateValue `json:"val,omitempty"`
	Description string                                        `json:"description,omitempty"`
}

// RankedRequestSegmentDefinitionPredicateValue represents a segment definition predicate value
type RankedRequestSegmentDefinitionPredicateValue struct {
	Func string `json:"func,omitempty"`
	Name string `json:"name,omitempty"`
}

// RankedRequestSearch represents a request search
type RankedRequestSearch struct {
	Clause             string   `json:"clause,omitempty"`
	ExcludeItemIDs     []string `json:"excludeItemIds,omitempty"`
	ItemIDs            []string `json:"itemIds,omitempty"`
	IncludeSearchTotal bool     `json:"includeSearchTotal"`
	Empty              bool     `json:"empty"`
}

// RankedRequestStatistics represents the request statistics
type RankedRequestStatistics struct {
	Functions    []string `json:"functions,omitempty"`
	IngoreZeroes bool     `json:"ignoreZeroes"`
}

// RankedRequestSettings represents the request settings
type RankedRequestSettings struct {
	Limit                   int    `json:"limit"`
	Page                    int    `json:"page"`
	DimensionSort           string `json:"dimensionSort,omitempty"`
	CountRepeatInstances    bool   `json:"countRepeatInstances,omitempty"`
	ReflectRequest          bool   `json:"reflectRequest,omitempty"`
	IncludeAnomalyDetection bool   `json:"includeAnomalyDetection,omitempty"`
	IncludePercentChange    bool   `json:"includePercentChange,omitempty"`
	IncludeLatLong          bool   `json:"includeLatLong,omitempty"`
	NonesBehavior           string `json:"nonesBehavior,omitempty"`
}

// RankedRequestReportMetrics represents the report metrics
type RankedRequestReportMetrics struct {
	MetricFilters *[]RankedRequestReportFilter `json:"metricFilters,omitempty"`
	Metrics       *[]RankedRequestReportMetric `json:"metrics,omitempty"`
}

// RankedRequestReportMetric represents a report metric
type RankedRequestReportMetric struct {
	ID               string                                       `json:"id,omitempty"`
	ColumnID         string                                       `json:"columnId,omitempty"`
	Filters          []string                                     `json:"filters,omitempty"`
	Sort             string                                       `json:"sort,omitempty"`
	MetricDefinition string                                       `json:"metricDefinition,omitempty"`
	Predictive       *RankedRequestReportMetricPredictiveSettings `json:"predictive,omitempty"`
}

// RankedRequestReportMetricPredictiveSettings represents metric predictive settings
type RankedRequestReportMetricPredictiveSettings struct {
	AnomalyConfidence float64 `json:"anomalyConfidence,omitempty"`
}

// RankedRequestReportRows represents report rows
type RankedRequestReportRows struct {
	RowFilters []RankedRequestReportFilter `json:"rowFilters,omitempty"`
	Rows       []RankedRequestReportRow    `json:"rows,omitempty"`
}

// RankedRequestReportRow represents a report row
type RankedRequestReportRow struct {
	RowID   string   `json:"rowId,omitempty"`
	Filters []string `json:"filters,omitempty"`
}

// RankedRequest represents a ranked report request
type RankedRequest struct {
	ReportSuiteID   string                       `json:"rsid"`
	Dimension       string                       `json:"dimension"`
	Locale          *RankedRequestLocale         `json:"locale,omitempty"`
	GlobalFilters   *[]RankedRequestReportFilter `json:"globalFilters,omitempty"`
	Search          *RankedRequestSearch         `json:"search,omitempty"`
	Settings        *RankedRequestSettings       `json:"settings,omitempty"`
	Statistics      *RankedRequestStatistics     `json:"statistics,omitempty"`
	MetricContainer *RankedRequestReportMetrics  `json:"metricContainer,omitempty"`
	RowContainer    *RankedRequestReportRows     `json:"rowContainer,omitempty"`
	AnchorDate      string                       `json:"anchorDate,omitempty"`
}

// Response types

// RankedReportDimension represents the report dimension
type RankedReportDimension struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// RankedReportColumnError represents the report column error
type RankedReportColumnError struct {
	ColumnID         string `json:"columnId,omitempty"`
	ErrorCode        string `json:"errorCode,omitempty"`
	ErrorID          string `json:"errorId,omitempty"`
	ErrorDescription string `json:"errorDescription,omitempty"`
}

// RankedReportColumnMetaData represents the report column meta data
type RankedReportColumnMetaData struct {
	Dimension    *RankedReportDimension     `json:"dimension,omitempty"`
	ColumnIDs    []string                   `json:"columnIds,omitempty"`
	ColumnErrors *[]RankedReportColumnError `json:"columnErrors,omitempty"`
}

// RankedReportRowData represents the report row data
type RankedReportRowData struct {
	ItemID              string    `json:"itemId,omitempty"`
	Value               string    `json:"value,omitempty"`
	RowID               string    `json:"rowId,omitempty"`
	Data                []float32 `json:"data,omitempty"`
	DataExpected        []float32 `json:"dataExpected,omitempty"`
	DataUpperBound      []float32 `json:"dataUpperBound,omitempty"`
	DataLowerBound      []float32 `json:"dataLowerBound,omitempty"`
	DataAnomalyDetected bool      `json:"dataAnomalyDetected"`
	PercentChange       []float32 `json:"percentChange,omitempty"`
	Latitude            float32   `json:"latitude,omitempty"`
	Longitude           float32   `json:"longitude,omitempty"`
}

// RankedReportSummaryData represents the report summary data
type RankedReportSummaryData struct {
	*RankedReportSummaryDataTotals `json:"summaryData,omitempty"`
}

// RankedReportSummaryDataTotals represents the report summary data totals
type RankedReportSummaryDataTotals struct {
	FilteredTotals []int `json:"filteredTotals,omitempty"`
	Totals         []int `json:"totals,omitempty"`
}

// RankedReportData represents the report data
type RankedReportData struct {
	TotalPages       int                         `json:"totalPages,omitempty"`
	FirstPage        bool                        `json:"firstPage"`
	LastPage         bool                        `json:"lastPage"`
	NumberOfElements int                         `json:"numberOfElements,omitempty"`
	Number           int                         `json:"number,omitempty"`
	TotalElements    int                         `json:"totalElements"`
	Message          string                      `json:"message,omitempty"`
	Request          *RankedRequest              `json:"request,omitempty"`
	Columns          *RankedReportColumnMetaData `json:"columns,omitempty"`
	Rows             *[]RankedReportRowData      `json:"rows,omitempty"`
	SummaryData      *RankedReportSummaryData    `json:"summaryData,omitempty"`
}
