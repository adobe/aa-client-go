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

func TestCollectionsGetAll(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/collections/suites"

	raw, err := ioutil.ReadFile("./testdata/Collections.GetAll.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		testRequestParams(t, r, map[string]string{
			"rsids":        "a,b",
			"rsidContains": "prod",
			"limit":        "10",
			"page":         "0",
			"expansion":    "a,b",
		})
		fmt.Fprint(w, string(raw))
	})

	collections, err := testClient.Collections.GetAll("a,b", "prod", 10, 0, []string{"a", "b"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(*collections.Content) != 2 {
		t.Errorf("Expected %d collections but got %d", 2, len(*collections.Content))
		return
	}
}

func TestCollectionsGetAllError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/collections/suites", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Collections.GetAll("rsIds", "rsidContains", 10, 0, []string{"a", "b"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestCollectionsGetByID(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/collections/suites/amc.aem.prod"

	raw, err := ioutil.ReadFile("./testdata/Collections.GetByID.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		testRequestParams(t, r, map[string]string{
			"expansion": "a,b",
		})
		fmt.Fprint(w, string(raw))
	})

	collection, err := testClient.Collections.GetByID("amc.aem.prod", []string{"a", "b"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if collection.RSID != "amc.aem.prod" {
		t.Errorf("Expected collection with RSID=amc.aem.prod but got RSID=%s", collection.RSID)
	}
}

func TestCollectionsGetByIDError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/collections/suites/amc.aem.prod", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Collections.GetByID("amc.aem.prod", []string{"a", "b"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
