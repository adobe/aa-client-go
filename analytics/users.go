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
	"strconv"
)

// UsersService handles users.
// Analytics docs: https://adobedocs.github.io/analytics-2.0-apis/#/users
type UsersService struct {
	client *Client
}

// GetAll returns a list of users for the current users login company.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/users/findAll
func (s *UsersService) GetAll(limit, page int64) (*Users, error) {
	var params = map[string]string{}
	params["limit"] = strconv.FormatInt(limit, 10)
	params["page"] = strconv.FormatInt(page, 10)

	var data Users
	err := s.client.get("/users", params, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// GetCurrent returns the current user.
// API docs: https://adobedocs.github.io/analytics-2.0-apis/#/users/getCurrentUser
func (s *UsersService) GetCurrent() (*User, error) {
	var data User
	err := s.client.get("/users/me", map[string]string{}, nil, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
