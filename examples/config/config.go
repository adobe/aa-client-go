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

package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/tkanos/gonfig"
)

//Configuration object for all commands
type Configuration struct {
	Analytics AnalyticsAPIConfiguration
}

// AnalyticsAPIConfiguration Adobe I/O Analytics config
type AnalyticsAPIConfiguration struct {
	Endpoint              string
	ReportSuiteID         string
	CompanyID             string
	IMSEndpoint           string
	IMSPrivateKey         string
	IMSOrgID              string
	IMSTechnicalAccountID string
	IMSClientID           string
	IMSClientSecret       string
}

type environment interface {
	Lookup(name string) (string, bool)
}

type osEnviroment struct{}

func (osEnviroment) Lookup(name string) (string, bool) {
	return os.LookupEnv(name)
}

// Read reads the configuration from environment variables and optionally from a configuration file in YAML format.
func Read() (*Configuration, error) {
	return read(osEnviroment{})
}

func read(env environment) (*Configuration, error) {
	var configuration Configuration

	// get config file
	if configFile, ok := env.Lookup("CONFIG"); ok {
		if err := gonfig.GetConf(strings.TrimSpace(configFile), &configuration); err != nil {
			return nil, fmt.Errorf("read configuration file: %v", err)
		}
	}

	// Analytics configs
	if aaPrivateKey, ok := env.Lookup("ANALYTICS_PRIVATEKEY"); ok {
		configuration.Analytics.IMSPrivateKey = aaPrivateKey
	}

	if aaClientSecret, ok := env.Lookup("ANALYTICS_CLIENTSECRET"); ok {
		configuration.Analytics.IMSClientSecret = aaClientSecret
	}

	return &configuration, nil
}
