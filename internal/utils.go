/*
Copyright © 2021 Thomas Meitz <thme219@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Parts of this apiclient are borrowed from Zalando Skipper
https://github.com/zalando/skipper/blob/master/net/httpclient.go

Zalando licence: MIT
https://github.com/zalando/skipper/blob/master/LICENSE
*/

package internal

import (
	"fmt"
	"net/url"
	"strings"
)

// ValidateUrl checks the url; url must not contain a trailing slash
func GetUrl(path string) (*url.URL, bool, error) {
	trimmedPath := strings.TrimSuffix(path, "/")
	u, err := url.Parse(trimmedPath)
	if err != nil {
		return nil, false, fmt.Errorf("can't parse url: %w", err)
	}
	if u.Host == "" {
		return nil, false, fmt.Errorf("invalid host name given")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, false, fmt.Errorf("invalid url scheme given")
	}
	return u, true, nil
}
