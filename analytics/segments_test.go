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

func TestSegmentsGetAll(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/segments"

	raw, err := ioutil.ReadFile("./testdata/Segments.GetAll.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		testRequestParams(t, r, map[string]string{
			"rsids":                     "rsIds",
			"segmentFilter":             "segmentFilter",
			"locale":                    "en_US",
			"name":                      "name",
			"tagNames":                  "tagNames",
			"filterByPublishedSegments": "filterByPublishedSegments",
			"limit":                     "10",
			"page":                      "0",
			"sortDirection":             "sortDirection",
			"sortProperty":              "sortProperty",
			"expansion":                 "a,b",
			"includeType":               "all",
		})
		fmt.Fprint(w, string(raw))
	})

	segments, err := testClient.Segments.GetAll("rsIds", "segmentFilter", "en_US", "name", "tagNames", "filterByPublishedSegments", 10, 0, "sortDirection", "sortProperty", []string{"a", "b"}, []string{"all"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(*segments.Content) != 2 {
		t.Errorf("Expected %d dimensions but got %d", 2, len(*segments.Content))
		return
	}
}

func TestSegmentsGetAllError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/segments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Segments.GetAll("rsIds", "segmentFilter", "en_US", "name", "tagNames", "filterByPublishedSegments", 10, 0, "sortDirection", "sortProperty", []string{"a", "b"}, []string{"all"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestSegmentsGetByID(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/segments/s300003364_589ce94be4b0c29f29c4f07f"

	raw, err := ioutil.ReadFile("./testdata/Segments.GetByID.json")
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

	segment, err := testClient.Segments.GetByID("s300003364_589ce94be4b0c29f29c4f07f", "en_US", []string{"a", "b"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if segment.ID != "s300003364_589ce94be4b0c29f29c4f07f" {
		t.Errorf("Expected dimension with ID=s300003364_589ce94be4b0c29f29c4f07f but got ID=%s", segment.ID)
	}
}

func TestSegmentsGetByIDError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/segments/s300003364_589ce94be4b0c29f29c4f07f", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Segments.GetByID("s300003364_589ce94be4b0c29f29c4f07f", "en_US", []string{"a", "b"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
