package main

import (
	"fmt"
	"github.com/rtt/Go-Solr"
)

/*
 * README
 * ------
 * This example shows an update Query being performed. An update document
 * is built and sent off to Solr.
 */

func main() {

	// init a connection
	s, err := solr.Init("192.168.70.220", 18983, "song_daemon")

	if err != nil {
		fmt.Println(err)
		return
	}

	// build an update document, in this case adding two documents
	f := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{"id": "artist-2000001", "artist_id": "2000001", "artist_name": "TEST 2000001"},
			map[string]interface{}{"id": "artist-2000002", "artist_id": "2000002", "artist_name": "TEST 2000002"},
			map[string]interface{}{"id": "artist-2000003", "artist_id": "2000003", "artist_name": "TEST 2000003"},
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
