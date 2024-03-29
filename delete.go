package main

import (
	"fmt"
	"github.com/rtt/Go-Solr"
)

/*
 * README
 * ------
 * This example shows an update (delete) Query being performed. An update (delete) document
 * is built and sent off to Solr.
 */

func main() {

	// init a connection
	s, err := solr.Init("192.168.70.220", 18983, "song_daemon")

	if err != nil {
		fmt.Println(err)
		return
	}

	// build an update document, in this case removing two documents
	f := map[string]interface{}{
		"delete": map[string]interface{}{
			"id": "artist-2000002",
		},
	}

	// send off the update (2nd parameter indicates we also want to commit the operation)
	resp, err := s.Update(f, true)

	if err != nil {
		fmt.Println("error =>", err)
	} else {
		fmt.Println("resp =>", resp)
	}

}
