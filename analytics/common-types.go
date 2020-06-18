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

// Common/Shared types

// Owner represents an owner
type Owner struct {
	ID    int    `json:"id"`
	Name  string `json:"name,omitempty"`
	Login string `json:"login,omitempty"`
}

// Sort represents sorting properties
type Sort struct {
	Direction  string `json:"direction,omitempty"`
	Property   string `json:"property,omitempty"`
	IgnoreCase bool   `json:"ignoreCase"`
	Ascending  bool   `json:"ascending"`
}

// TaggedComponent represents a tagged component
type TaggedComponent struct {
	ComponentType string `json:"componentType,omitempty"`
	ComponentID   string `json:"componentId,omitempty"`
	Tags          *[]Tag `json:"tags,omitempty"`
}

// Tag represents a tag
type Tag struct {
	ID          string             `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Components  *[]TaggedComponent `json:"components,omitempty"`
}
