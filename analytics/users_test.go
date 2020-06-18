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

package analytics_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestUsersGetAll(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/users"

	raw, err := ioutil.ReadFile("./testdata/Users.GetAll.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		fmt.Fprint(w, string(raw))
	})

	users, err := testClient.Users.GetAll(10, 0)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if users.Content == nil {
		t.Error("Expected users list but was is nil")
	}

	if len(*users.Content) != 2 {
		t.Errorf("Expected %d users but got %d", 2, len(*users.Content))
	}

}

func TestUsersGetAllError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Users.GetAll(10, 0)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestUsersGetCurrent(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/users/me"

	raw, err := ioutil.ReadFile("./testdata/Users.GetCurrent.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		fmt.Fprint(w, string(raw))
	})

	user, err := testClient.Users.GetCurrent()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if user.LoginID != 1 {
		t.Errorf("Expected loginId=1 but was loginId=%d", user.LoginID)
	}
}

func TestUsersGetCurrentError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/users/me", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Users.GetCurrent()
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
