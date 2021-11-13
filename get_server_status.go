/*
Copyright © 2021 Thomas Meitz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ksqldb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ServerStatusResponse
type ServerStatusResponse struct {
	IsHealthy *bool `json:"isHealthy"`
	Details   struct {
		Metastore struct {
			IsHealthy *bool `json:"isHealthy"`
		} `json:"metastore"`
		Kafka struct {
			IsHealthy *bool `json:"isHealthy"`
		} `json:"kafka"`
	} `json:"details"`
	KsqlServiceID string `json:"ksqlServiceId"`
}

// ServerInfo provides information about your server
func (c *KsqldbClient) GetServerStatus() (*ServerStatusResponse, error) {
	info := ServerStatusResponse{}
	url := (*c.http).GetUrl(HEALTHCHECK_ENDPOINT)

	res, err := (*c.http).Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't get healthcheck informations: %v", err)
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, fmt.Errorf("could not read response body: %v", readErr)
	}

	if err := json.Unmarshal(body, &info); err != nil {
		return nil, fmt.Errorf("could not parse the response: %w", err)
	}

	return &info, nil
}
