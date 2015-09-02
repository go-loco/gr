package gr

import (
	"net"
	"strconv"
	"sync"
)

type pool struct {
	minConns int
	tcpAddr  *net.TCPAddr
	queue    queue
	channel  chan bool
	mutex    sync.Mutex
}

func newPool(config *Config) (*pool, error) {

	servAddr := config.Address + ":" + strconv.Itoa(config.Port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		return nil, err
	}

	p := &pool{
		config.MinConnections,
		tcpAddr,
		queue{},
		make(chan bool),
		sync.Mutex{},
	}

	p.initMinConns()

	return p, nil

}

func (p *pool) initMinConns() {
	for i := 0; i < p.minConns; i++ {
		p.putNewConn()
	}
}

func (p *pool) putNewConn() {
	if conn, err := p.connect(); err == nil {
		p.put(conn)
	}
}

func (p *pool) connect() (*net.TCPConn, error) {
	return net.DialTCP("tcp", nil, p.tcpAddr)
}

func (p *pool) get() *net.TCPConn {

	var conn interface{}

	for conn == nil {

		p.mutex.Lock()
		conn = p.queue.dequeue()
		p.mutex.Unlock()

		if conn == nil {
			go p.putNewConn()
			<-p.channel
		}
	}

	return conn.(*net.TCPConn)
}

func (p *pool) put(conn *net.TCPConn) {

	if conn != nil {

		p.mutex.Lock()

		if p.queue.size < p.minConns {
			p.queue.enqueue(conn)

			select {
			case p.channel <- true:
			default:
			}

		} else {
			defer conn.Close()
		}

		p.mutex.Unlock()

	}

}
