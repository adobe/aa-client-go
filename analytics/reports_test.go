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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/adobe/aa-client-go/analytics"
)

func TestReportsRun(t *testing.T) {
	setup()
	defer teardown()

	apiEndpoint := baseURL + "/reports"

	req, err := ioutil.ReadFile("./testdata/Reports.Run.Request.json")
	if err != nil {
		t.Error(err.Error())
	}

	var rankedRequest analytics.RankedRequest
	json.Unmarshal(req, &rankedRequest)

	raw, err := ioutil.ReadFile("./testdata/Reports.Run.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(apiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, apiEndpoint)
		testRequestBody(t, r, req)
		fmt.Fprint(w, string(raw))
	})

	report, err := testClient.Reports.Run(&rankedRequest)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(*report.Rows) != 68 {
		t.Errorf("Expected %d report rows but got %d", 68, len(*report.Rows))
	}
}

func TestReportsRunError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(baseURL+"/reports", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := testClient.Reports.Run(nil)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
