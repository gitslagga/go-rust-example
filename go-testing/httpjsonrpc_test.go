package go_testing

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

// adapt HTTP connection to ReadWriteCloser
type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }

// our service
type CakeBaker struct{}

func (cb *CakeBaker) BakeIt(n int, msg *string) error {
	*msg = fmt.Sprintf("your cake has been bacon (%d)", n)
	return nil
}

func TestHTTPServer(t *testing.T) {

	fmt.Printf("TestHTTPServer\n")

	cb := &CakeBaker{}

	server := rpc.NewServer()
	server.Register(cb)

	listener, e := net.Listen("tcp", ":4321")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer listener.Close()

	go http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/bake-me-a-cake" {
			serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(200)
			err := server.ServeRequest(serverCodec)
			if err != nil {
				log.Printf("Error while serving JSON request: %v", err)
				http.Error(w, "Error while serving JSON request, details have been logged.", 500)
				return
			}
		}

	}))

	resp, err := http.Post("http://localhost:4321/bake-me-a-cake", "application/json", bytes.NewBufferString(
		`{"jsonrpc":"2.0","id":1,"method":"CakeBaker.BakeIt","params":[10]}`,
	))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("returned JSON: %s\n", string(b))

}
