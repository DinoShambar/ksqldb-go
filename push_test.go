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

package ksqldb_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thmeitz/ksqldb-go"
	mocknet "github.com/thmeitz/ksqldb-go/mocks/net"
)

func TestPush_EmptyQuery(t *testing.T) {
	rowChannel := make(chan ksqldb.Row)
	headerChannel := make(chan ksqldb.Header, 1)
	m := mocknet.HTTPClient{}
	kcl, _ := ksqldb.NewClient(&m)
	kcl.EnableParseSQL(true)
	err := kcl.Push(context.TODO(), ksqldb.QueryOptions{Sql: ""}, rowChannel, headerChannel)
	require.NotNil(t, err)
	require.Equal(t, "empty ksql query", err.Error())
}
func TestPush_ParseSQLError(t *testing.T) {
	rowChannel := make(chan ksqldb.Row)
	headerChannel := make(chan ksqldb.Header, 1)
	m := mocknet.HTTPClient{}
	kcl, _ := ksqldb.NewClient(&m)
	kcl.EnableParseSQL(true)
	err := kcl.Push(context.TODO(), ksqldb.QueryOptions{Sql: "select * from bla"}, rowChannel, headerChannel)
	require.NotNil(t, err)
	require.Equal(t, "1 sql syntax error(s) found", err.Error())
}
