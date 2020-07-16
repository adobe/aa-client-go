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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/adobe/aa-client-go/analytics"
)

var (
	baseURL    string
	testMux    *http.ServeMux
	testConfig *analytics.Config
	testClient *analytics.Client
	testServer *httptest.Server
)

func setup() {
	// mock server
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	// test config
	baseURL = "/api/aaCompanyId"
	testConfig = &analytics.Config{
		BaseURL:     string(testServer.URL + "/api"),
		ClientID:    "imsClientId",
		OrgID:       "imsOrgId",
		AccessToken: "imsAuthToken",
		CompanyID:   "aaCompanyId",
	}

	// client configured to use test server
	testClient, _ = analytics.NewClient(testConfig)
}

func teardown() {
	testServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testRequestURL(t *testing.T, r *http.Request, want string) {
	if got := r.URL.String(); !strings.HasPrefix(got, want) {
		t.Errorf("Request URL: %v, want %v", got, want)
	}
}

func testRequestParams(t *testing.T, r *http.Request, want map[string]string) {
	params := r.URL.Query()

	if len(params) != len(want) {
		t.Errorf("Request params: %d, want %d", len(params), len(want))
	}

	for key, val := range want {
		if got := params.Get(key); val != got {
			t.Errorf("Request params: %s, want %s", got, val)
		}

	}
}

func testRequestBody(t *testing.T, r *http.Request, want []byte) {
	data, _ := ioutil.ReadAll(r.Body)
	var buf bytes.Buffer
	json.Indent(&buf, data, "", "  ")
	got := buf.Bytes()

	if eq, _ := jsonBytesEqual(got, want); !eq {
		t.Errorf("Request Body:\n%v\nwant:\n%v", string(got), string(want))
	}
}

func jsonBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j, j2), nil
}

func TestNewClientNoScheme(t *testing.T) {
	_, err := analytics.NewClient(
		&analytics.Config{
			BaseURL:     "domain.com/api",
			ClientID:    "imsClientId",
			OrgID:       "imsOrgId",
			AccessToken: "imsAccessToken",
			CompanyID:   "aaCompanyId",
		},
	)
	if err == nil {
		t.Fatalf("expected error")
	}
	if err.Error() != "missing URL scheme" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewClientNoHost(t *testing.T) {
	_, err := analytics.NewClient(
		&analytics.Config{
			BaseURL:     "https:///api",
			ClientID:    "imsClientId",
			OrgID:       "imsOrgId",
			AccessToken: "imsAccessToken",
			CompanyID:   "aaCompanyId",
		},
	)
	if err == nil {
		t.Fatalf("expected error")
	}
	if err.Error() != "missing URL host" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewClientMalformedURL(t *testing.T) {
	_, err := analytics.NewClient(
		&analytics.Config{
			BaseURL:     ":",
			ClientID:    "imsClientId",
			OrgID:       "imsOrgId",
			AccessToken: "imsAccessToken",
			CompanyID:   "aaCompanyId",
		},
	)
	if err == nil {
		t.Fatalf("expected error")
	}
	if err.Error() != "malformed URL" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewClientNoClientID(t *testing.T) {
	err := testNewClientInvalidAuth(t, "", "OrgID", "AccessToken", "CompanyID")
	if err.Error() != "missing ClientID" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewClientNoIMSOrgID(t *testing.T) {
	err := testNewClientInvalidAuth(t, "ClientID", "", "AccessToken", "CompanyID")
	if err.Error() != "missing OrgID" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewClientNoIMSAuthToken(t *testing.T) {
	err := testNewClientInvalidAuth(t, "ClientID", "OrgID", "", "CompanyID")
	if err.Error() != "missing AccessToken" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewClientNoCompanyID(t *testing.T) {
	err := testNewClientInvalidAuth(t, "ClientID", "OrgID", "AccessToken", "")
	if err.Error() != "missing CompanyID" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func testNewClientInvalidAuth(t *testing.T, clientID, orgID, accessToken, companyID string) error {
	_, err := analytics.NewClient(
		&analytics.Config{
			BaseURL:     "https://domain.com/api",
			ClientID:    clientID,
			OrgID:       orgID,
			AccessToken: accessToken,
			CompanyID:   companyID,
		},
	)
	if err == nil {
		t.Fatalf("expected error")
	}
	return err
}

func TestNewClientTrimSuffix(t *testing.T) {
	c, err := analytics.NewClient(
		&analytics.Config{
			BaseURL:     "https://domain.com/api/",
			ClientID:    "imsClientId",
			OrgID:       "imsOrgId",
			AccessToken: "imsAccessToken",
			CompanyID:   "companyId",
		},
	)
	if err != nil {
		t.Fatalf("unexpected error")
	}
	if strings.HasSuffix(c.GetBaseURL(), "/") {
		t.Fatalf("unexpected trailing slash")
	}
}
