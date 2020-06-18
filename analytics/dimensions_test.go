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

func TestDimensionsGetAll(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/dimensions"

	raw, err := ioutil.ReadFile("./testdata/Dimensions.GetAll.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		testRequestParams(t, r, map[string]string{
			"rsid":         "rsId",
			"locale":       "en_US",
			"segmentable":  "false",
			"reportable":   "false",
			"classifiable": "false",
			"expansion":    "a,b",
		})
		fmt.Fprint(w, string(raw))
	})

	dimensions, err := testClient.Dimensions.GetAll("rsId", "en_US", false, false, false, []string{"a", "b"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(*dimensions) != 2 {
		t.Errorf("Expected %d dimensions but got %d", 2, len(*dimensions))
		return
	}
}

func TestDimensionsGetAllError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/dimensions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Dimensions.GetAll("rsId", "en_US", false, false, false, []string{"a", "b"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestDimensionsGetByID(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/dimensions/evar1"

	raw, err := ioutil.ReadFile("./testdata/Dimensions.GetByID.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		testRequestParams(t, r, map[string]string{
			"rsid":      "rsId",
			"locale":    "en_US",
			"expansion": "a,b",
		})
		fmt.Fprint(w, string(raw))
	})

	dimension, err := testClient.Dimensions.GetByID("rsId", "evar1", "en_US", []string{"a", "b"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if dimension.ID != "variables/evar1" {
		t.Errorf("Expected dimension with ID=variables/evar1 but got ID=%s", dimension.ID)
	}
}

func TestDimensionsGetByIDError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/dimensions/evar1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Dimensions.GetByID("rsId", "evar1", "en_US", []string{"a", "b"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
