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

func TestDateRangesGetAll(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/dateranges"

	raw, err := ioutil.ReadFile("./testdata/DateRanges.GetAll.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		testRequestParams(t, r, map[string]string{
			"locale":      "en_US",
			"filterByIds": "1,2",
			"limit":       "10",
			"page":        "0",
			"expansion":   "a,b",
			"includeType": "all",
		})
		fmt.Fprint(w, string(raw))
	})

	dateranges, err := testClient.DateRanges.GetAll("en_US", "1,2", 10, 0, []string{"a", "b"}, []string{"all"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(*dateranges.Content) != 2 {
		t.Errorf("Expected %d dateranges but got %d", 2, len(*dateranges.Content))
		return
	}
}

func TestDateRangesGetAllError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/dateranges", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.DateRanges.GetAll("en_US", "1,2", 10, 0, []string{"a", "b"}, []string{"all"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestDateRangesGetByID(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/dateranges/57a9ad685fe707f55ffb68f5"

	raw, err := ioutil.ReadFile("./testdata/DateRanges.GetByID.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		testRequestParams(t, r, map[string]string{
			"locale":    "en_US",
			"expansion": "a,b",
		})
		fmt.Fprint(w, string(raw))
	})

	daterange, err := testClient.DateRanges.GetByID("57a9ad685fe707f55ffb68f5", "en_US", []string{"a", "b"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if daterange.ID != "57a9ad685fe707f55ffb68f5" {
		t.Errorf("Expected dimension with ID=57a9ad685fe707f55ffb68f5 but got ID=%s", daterange.ID)
	}
}

func TestDateRangesGetByIDError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/dateranges/?", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.DateRanges.GetByID("?", "en_US", []string{"a", "b"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
