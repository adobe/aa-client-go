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

// User represents an Analytics user
type User struct {
	CompanyID      int    `json:"companyId"`
	LoginID        int    `json:"loginId"`
	Login          string `json:"login,omitempty"`
	ChangePassword bool   `json:"changePassword"`
	CreateDate     string `json:"createDate,omitempty"`
	Disabled       bool   `json:"disabled"`
	Email          string `json:"email,omitempty"`
	FirstName      string `json:"firstName,omitempty"`
	FullName       string `json:"fullName,omitempty"`
	IMSUserID      string `json:"imsUserId,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	LastAccess     string `json:"lastAccess,omitempty"`
	PhoneNumber    string `json:"phoneNumber,omitempty"`
	TempLoginEnd   string `json:"tempLoginEnd,omitempty"`
	Title          string `json:"title,omitempty"`
}

// Users represents a page of Analytics users
type Users struct {
	Content          *[]User `json:"content,omitempty"`
	Number           int     `json:"number"`
	Size             int     `json:"size"`
	NumberOfElements int     `json:"numberOfElements"`
	TotalElements    int     `json:"totalElements"`
	PreviousPage     bool    `json:"previousPage"`
	FirstPage        bool    `json:"firstPage"`
	NextPage         bool    `json:"nextPage"`
	LastPage         bool    `json:"lastPage"`
	Sort             *[]Sort `json:"sort,omitempty"`
	TotalPages       int     `json:"totalPages"`
}
