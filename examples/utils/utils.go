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

package utils

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/adobe/ims-go/ims"
	"github.com/hashicorp/go-retryablehttp"

	"github.com/adobe/aa-client-go/analytics"
	"github.com/adobe/aa-client-go/examples/config"
)

// PrintJSON pretty prints the JSON data
func PrintJSON(data interface{}) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

// NewAnalyticsClient returns a new Analytics client
func NewAnalyticsClient(aaConfig *config.AnalyticsAPIConfiguration) (*analytics.Client, error) {
	// Exchange JWT for Auth Token
	responseToken, err := getAuthToken(aaConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("IMS AuthToken received, expires in %v", responseToken.ExpiresIn)

	// Create resilient client
	httpClient := newHTTPClient()

	// Create Analytics client
	client, err := analytics.NewClient(&analytics.Config{
		HTTPClient:  httpClient,
		BaseURL:     aaConfig.Endpoint,
		ClientID:    aaConfig.IMSClientID,
		OrgID:       aaConfig.IMSOrgID,
		AccessToken: responseToken.AccessToken,
		CompanyID:   aaConfig.CompanyID,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Analytics client created for [ClientID:%s, CompanyID:%s]", aaConfig.IMSClientID, aaConfig.CompanyID)

	return client, nil
}

// rateLimitPolicy implements a retryablehttp.CheckRetry for handling 429, 5xx status codes
func rateLimitPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	var (
		// A regular expression to match the error returned by net/http when the
		// configured number of redirects is exhausted. This error isn't typed
		// specifically so we resort to matching on the error string.
		redirectsErrorRe = regexp.MustCompile(`stopped after \d+ redirects\z`)

		// A regular expression to match the error returned by net/http when the
		// scheme specified in the URL is invalid. This error isn't typed
		// specifically so we resort to matching on the error string.
		schemeErrorRe = regexp.MustCompile(`unsupported protocol scheme`)
	)

	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	if err != nil {
		if v, ok := err.(*url.Error); ok {
			// Don't retry if the error was due to too many redirects.
			if redirectsErrorRe.MatchString(v.Error()) {
				return false, nil
			}

			// Don't retry if the error was due to an invalid protocol scheme.
			if schemeErrorRe.MatchString(v.Error()) {
				return false, nil
			}

			// Don't retry if the error was due to TLS cert verification failure.
			if _, ok := v.Err.(x509.UnknownAuthorityError); ok {
				return false, nil
			}
		}

		// The error is likely recoverable so retry.
		return true, nil
	}

	// Check the response code. We retry on 500-range responses to allow
	// the server time to recover, as 500's are typically not permanent
	// errors and may relate to outages on the server side. This will catch
	// invalid response codes as well, like 0 and 999.
	if resp.StatusCode == 0 || resp.StatusCode == 429 || (resp.StatusCode >= 500 && resp.StatusCode != 501) {
		return true, nil
	}

	return false, nil
}

// newHTTPClient create a new resilient HTTP client
func newHTTPClient() *http.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = 10
	retryClient.CheckRetry = rateLimitPolicy

	return retryClient.StandardClient()
}

// getAuthToken returns an IMS AuthToken in exchange of a JWT
func getAuthToken(aaConfig *config.AnalyticsAPIConfiguration) (*ims.ExchangeJWTResponse, error) {
	c, err := ims.NewClient(&ims.ClientConfig{
		URL: aaConfig.IMSEndpoint,
	})

	if err != nil {
		return nil, err
	}

	// get token response
	r, err := c.ExchangeJWT(&ims.ExchangeJWTRequest{
		PrivateKey:   []byte(aaConfig.IMSPrivateKey),
		Expiration:   time.Now().Add(12 * time.Hour),
		Issuer:       aaConfig.IMSOrgID,
		Subject:      aaConfig.IMSTechnicalAccountID,
		ClientID:     aaConfig.IMSClientID,
		ClientSecret: aaConfig.IMSClientSecret,
		MetaScope:    []ims.MetaScope{ims.MetaScopeAnalyticsBulkIngest},
	})

	return r, err
}
