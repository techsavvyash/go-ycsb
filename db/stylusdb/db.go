package stylusdb

import (
	"fmt"
	"net"

	"github.com/magiconair/properties"
	"github.com/pingcap/go-ycsb/pkg/ycsb"
)

type stylusDBCreator struct{}
type stylusDB {
	conn *net.Conn
}

func (c stylusDBCreator) Create(p *properties.Properties) (ycsb.DB, error) {
	// return the TCP Client
	conn, err := net.Dial("tcp", "localhost:6767")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return nil, err
	}
	return &stylusDB{
		conn: conn,
	}, err
}

func init() {
	ycsb.RegisterDBCreator("stylusdb", stylusDBCreator{})
}
