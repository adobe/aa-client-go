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

// CollectionCalendarType represents a collection calendar type
type CollectionCalendarType struct {
	RSID       string `json:"rsid,omitempty"`
	Type       string `json:"type,omitempty"`
	AnchorDate string `json:"anchorDate,omitempty"`
}

// Collection represents a collection
type Collection struct {
	Name               string                  `json:"name,omitempty"`
	TimezoneZoneInfo   string                  `json:"timezoneZoneInfo,omitempty"`
	ParentRSID         string                  `json:"parentRsid,omitempty"`
	CollectionItemType string                  `json:"collectionItemType,omitempty"`
	Currency           string                  `json:"currency,omitempty"`
	CalendarType       *CollectionCalendarType `json:"calendarType,omitempty"`
	RSID               string                  `json:"rsid,omitempty"`
}

// Collections represents a page of collections
type Collections struct {
	Content          *[]Collection `json:"content,omitempty"`
	Number           int           `json:"number"`
	Size             int           `json:"size"`
	NumberOfElements int           `json:"numberOfElements"`
	TotalElements    int           `json:"totalElements"`
	PreviousPage     bool          `json:"previousPage"`
	FirstPage        bool          `json:"firstPage"`
	NextPage         bool          `json:"nextPage"`
	LastPage         bool          `json:"lastPage"`
	Sort             *[]Sort       `json:"sort,omitempty"`
	TotalPages       int           `json:"totalPages"`
}
