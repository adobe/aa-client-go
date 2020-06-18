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

func TestMetricsGetAll(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/metrics"

	raw, err := ioutil.ReadFile("./testdata/Metrics.GetAll.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, apiEndpoint)
		testRequestParams(t, r, map[string]string{
			"rsid":        "rsId",
			"locale":      "en_US",
			"segmentable": "false",
			"expansion":   "a,b",
		})
		fmt.Fprint(w, string(raw))
	})

	metrics, err := testClient.Metrics.GetAll("rsId", "en_US", false, []string{"a", "b"})
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	if len(*metrics) != 2 {
		t.Errorf("Expected %d metrics but got %d", 2, len(*metrics))
		return
	}
}

func TestMetricsGetAllErro(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Metrics.GetAll("rsId", "en_US", false, []string{"a", "b"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestMetricsGetByID(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/metrics/pageviews"

	raw, err := ioutil.ReadFile("./testdata/Metrics.GetByID.json")
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

	metric, err := testClient.Metrics.GetByID("rsId", "pageviews", "en_US", []string{"a", "b"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if metric.ID != "metrics/pageviews" {
		t.Errorf("Expected metric with ID=metrics/pageviews but got ID=%s", metric.ID)
	}
}

func TestMetricsGetByIDError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/metrics/pageviews", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Metrics.GetByID("rsId", "pageviews", "en_US", []string{"a", "b"})
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
