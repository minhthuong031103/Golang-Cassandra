package main

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true
	// cluster.Port = 9043
	cluster.Keyspace = "user_data"
	cluster.Consistency = gocql.LocalOne // Example consistency level: QUORUM

	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra well initialized")
}
