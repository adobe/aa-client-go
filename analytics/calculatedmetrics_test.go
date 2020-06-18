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

func TestCalculatedMetricsGetAll(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/calculatedmetrics"

	raw, err := ioutil.ReadFile("./testdata/CalculatedMetrics.GetAll.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		testRequestParams(t, r, map[string]string{
			"rsids":          "1,2",
			"ownerId":        "ownerId",
			"filterByIds":    "filterByIds",
			"toBeUsedInRsid": "toBeUsedInRsid",
			"locale":         "en_US",
			"name":           "name",
			"tagNames":       "tagNames",
			"favorite":       "false",
			"approved":       "false",
			"limit":          "10",
			"page":           "0",
			"sortDirection":  "asc",
			"sortProperty":   "date",
			"expansion":      "a,b",
			"includeType":    "all",
		})
		fmt.Fprint(w, string(raw))
	})

	collections, err := testClient.CalculatedMetrics.GetAll(
		"1,2", "ownerId", "filterByIds", "toBeUsedInRsid",
		"en_US", "name", "tagNames",
		false, false, 10, 0,
		"asc", "date",
		[]string{"a", "b"}, []string{"all"},
	)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(*collections.Content) != 2 {
		t.Errorf("Expected %d collections but got %d", 2, len(*collections.Content))
		return
	}
}

func TestCalculatedMetricsGetAllError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/calculatedmetrics", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.CalculatedMetrics.GetAll(
		"1,2", "ownerId", "filterByIds", "toBeUsedInRsid",
		"en_US", "name", "tagNames",
		false, false, 10, 0,
		"asc", "date",
		[]string{"a", "b"}, []string{"all"},
	)

	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestCalculatedMetricsGetByID(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/calculatedmetrics/cm300003364_5ae7447df118f061698ddc31"

	raw, err := ioutil.ReadFile("./testdata/CalculatedMetrics.GetByID.json")
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

	metric, err := testClient.CalculatedMetrics.GetByID("cm300003364_5ae7447df118f061698ddc31", "en_US", []string{"a", "b"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if metric.ID != "cm300003364_5ae7447df118f061698ddc31" {
		t.Errorf("Expected calculated metric with ID=cm300003364_5ae7447df118f061698ddc31 but got ID=%s", metric.ID)
	}
}

func TestCalculatedMetricsGetByIDError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/calculatedmetrics/cm300003364_5ae7447df118f061698ddc31", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.CalculatedMetrics.GetByID("cm300003364_5ae7447df118f061698ddc31", "en_US", []string{"a", "b"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
