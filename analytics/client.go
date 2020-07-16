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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Config holds configuration values
type Config struct {
	HTTPClient  *http.Client
	BaseURL     string
	ClientID    string
	OrgID       string
	AccessToken string
	CompanyID   string
}

// Auth holds authentication information
type auth struct {
	imsClientID    string
	imsOrgID       string
	imsAccessToken string
	companyID      string
}

// Client is used to make HTTP requests against the Analytics API 2.0.
// It wraps a HTTP client and handles authentication.
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	auth       *auth

	// Services used for communicating to different parts of the API.
	CalculatedMetrics *CalculatedMetricsService
	Collections       *CollectionsService
	DateRanges        *DateRangesService
	Dimensions        *DimensionsService
	Metrics           *MetricsService
	Reports           *ReportsService
	Segments          *SegmentsService
	Users             *UsersService
}

// NewClient returns a new Analytics API 2.0 client.
func NewClient(config *Config) (*Client, error) {
	var httpClient *http.Client
	if config.HTTPClient != nil {
		httpClient = config.HTTPClient
	} else {
		httpClient = http.DefaultClient
	}

	parsedBaseURL, err := verifyBaseURL(config.BaseURL)
	if err != nil {
		return nil, err
	}

	auth := &auth{
		imsClientID:    config.ClientID,
		imsOrgID:       config.OrgID,
		imsAccessToken: config.AccessToken,
		companyID:      config.CompanyID,
	}
	err = verifyAuth(auth)
	if err != nil {
		return nil, err
	}

	c := &Client{
		httpClient: httpClient,
		baseURL:    parsedBaseURL,
		auth:       auth,
	}

	c.CalculatedMetrics = &CalculatedMetricsService{client: c}
	c.Collections = &CollectionsService{client: c}
	c.DateRanges = &DateRangesService{client: c}
	c.Dimensions = &DimensionsService{client: c}
	c.Metrics = &MetricsService{client: c}
	c.Reports = &ReportsService{client: c}
	c.Segments = &SegmentsService{client: c}
	c.Users = &UsersService{client: c}

	return c, nil
}

// verifyBaseURL verifies the specified URL
func verifyBaseURL(baseURL string) (*url.URL, error) {
	if strings.HasSuffix(baseURL, "/") {
		baseURL = strings.TrimSuffix(baseURL, "/")
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("malformed URL")
	}

	if parsedBaseURL.Scheme == "" {
		return nil, fmt.Errorf("missing URL scheme")
	}

	if parsedBaseURL.Host == "" {
		return nil, fmt.Errorf("missing URL host")
	}
	return parsedBaseURL, nil
}

// verifyAuth verifies the specified auth
func verifyAuth(auth *auth) error {
	if auth.imsClientID == "" {
		return fmt.Errorf("missing ClientID")
	}
	if auth.imsAccessToken == "" {
		return fmt.Errorf("missing AccessToken")
	}
	if auth.imsOrgID == "" {
		return fmt.Errorf("missing OrgID")
	}
	if auth.companyID == "" {
		return fmt.Errorf("missing CompanyID")
	}
	return nil
}

// GetBaseURL returns the API base URL.
func (client *Client) GetBaseURL() string {
	return client.baseURL.String()
}

// get is a convenience method to send an HTTP GET request
func (client *Client) get(path string, params map[string]string, body io.Reader, model interface{}) error {
	return apiRequest(client, http.MethodGet, path, params, body, model)
}

// post is a convenience method to send an HTTP POST request
func (client *Client) post(path string, params map[string]string, body io.Reader, model interface{}) error {
	return apiRequest(client, http.MethodPost, path, params, body, model)
}

// apiRequest does a HTTP request and unmarshals the response into the specified model.
func apiRequest(client *Client, method, path string, params map[string]string, body io.Reader, model interface{}) error {
	resp, respErr := request(client, method, path, params, body)
	if respErr != nil {
		return respErr
	}
	defer resp.Body.Close()

	bodyStr, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return bodyErr
	}

	jsonErr := json.Unmarshal(bodyStr, &model)
	if jsonErr != nil {
		return jsonErr
	}

	return jsonErr
}

// request does a HTTP request with the specified client, method, path, params and body.
func request(client *Client, method, path string, params map[string]string, body io.Reader) (*http.Response, error) {
	// ensure path has prefix
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}

	// URL format is <API_URL>/<COMPANY_ID>/<PATH>
	// join <baseURL.Path>/<CompanyID>/<Path>
	rel := &url.URL{Path: fmt.Sprintf("%s/%s/%s", client.baseURL.Path, client.auth.companyID, path)}
	u := client.baseURL.ResolveReference(rel)
	addParams(u, params)

	httpClient := client.httpClient

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Set required headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.auth.imsAccessToken))
	req.Header.Set("x-api-key", client.auth.imsClientID)
	req.Header.Set("x-gw-ims-org-id", client.auth.imsOrgID)
	req.Header.Set("x-proxy-global-company-id", client.auth.companyID)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = checkResponse(res)
	if err != nil {
		// we return the response in case the caller wants to inspect it
		return res, err
	}

	return res, err
}

// addParams adds the specified params to the URL.
func addParams(u *url.URL, params map[string]string) string {
	q, _ := url.ParseQuery(u.RawQuery)

	for k, v := range params {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()
	return u.String()
}

// checkResponse checks the API response for errors.
// A response is considered an error if it has a status code outside the 200 range.
// The caller is responsible to analyze the response body.
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	err := fmt.Errorf("received unexpected status code %d", r.StatusCode)
	return err
}
