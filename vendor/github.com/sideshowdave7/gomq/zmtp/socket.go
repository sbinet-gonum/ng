package zmtp

import "errors"

// Socket is a ZMTP socket
type Socket interface {
	Type() SocketType
	IsSocketTypeCompatible(socketType SocketType) bool
	IsCommandTypeValid(name string) bool
}

// NewSocket returns a new ZMTP socket
func NewSocket(socketType SocketType) (Socket, error) {
	switch socketType {
	case ClientSocketType:
		return clientSocket{}, nil
	case ServerSocketType:
		return serverSocket{}, nil
	case PullSocketType:
		return pullSocket{}, nil
	case PushSocketType:
		return pushSocket{}, nil
	case DealerSocketType:
		return dealerSocket{}, nil
	case RouterSocketType:
		return routerSocket{}, nil
	case PubSocketType:
		return pubSocket{}, nil
	case SubSocketType:
		return subSocket{}, nil
	case ReqSocketType:
		return reqSocket{}, nil
	case RepSocketType:
		return repSocket{}, nil
	default:
		return nil, errors.New("Invalid socket type")
	}
}

type clientSocket struct{}

// Type returns the Socket's type
func (clientSocket) Type() SocketType {
	return ClientSocketType
}

// IsSocketTypeCompatible checks if the socket is compatible with
// another socket type.
func (clientSocket) IsSocketTypeCompatible(socketType SocketType) bool {
	return socketType == ServerSocketType
}

// IsCommandTypeValid returns if a command is valid for this socket.
func (clientSocket) IsCommandTypeValid(name string) bool {
	return false
}

type serverSocket struct{}

// Type returns the Socket's type
func (serverSocket) Type() SocketType {
	return ServerSocketType
}

// IsSocketTypeCompatible checks if the socket is compatible with
// another socket type.
func (serverSocket) IsSocketTypeCompatible(socketType SocketType) bool {
	return socketType == ClientSocketType
}

// IsCommandTypeValid returns if a command is valid for this socket.
func (serverSocket) IsCommandTypeValid(name string) bool {
	return false
}

type pullSocket struct{}

// Type returns the Socket's type
func (pullSocket) Type() SocketType {
	return PullSocketType
}

// IsSocketTypeCompatible checks if the socket is compatible with
// another socket type.
func (pullSocket) IsSocketTypeCompatible(socketType SocketType) bool {
	return socketType == PushSocketType
}

// IsCommandTypeValid returns if a command is valid for this socket.
func (pullSocket) IsCommandTypeValid(name string) bool {
	// FIXME(sbinet)
	return false
}

type pushSocket struct{}

// Type returns the Socket's type
func (pushSocket) Type() SocketType {
	return PushSocketType
}

// IsSocketTypeCompatible checks if the socket is compatible with
// another socket type.
func (pushSocket) IsSocketTypeCompatible(socketType SocketType) bool {
	return socketType == PullSocketType
}

// IsCommandTypeValid returns if a command is valid for this socket.
func (pushSocket) IsCommandTypeValid(name string) bool {
	// FIXME(sbinet)
	return false
}

type dealerSocket struct{}

func (dealerSocket) Type() SocketType {
	return DealerSocketType
}

func (dealerSocket) IsSocketTypeCompatible(socketType SocketType) bool {
	return socketType == RouterSocketType
}

// IsCommandTypeValid returns if a command is valid for this socket.
func (dealerSocket) IsCommandTypeValid(name string) bool {
	// FIXME
	return false
}

type routerSocket struct{}

func (routerSocket) Type() SocketType {
	return RouterSocketType
}

func (routerSocket) IsSocketTypeCompatible(socketType SocketType) bool {
	return socketType == RepSocketType ||
		socketType == DealerSocketType ||
		socketType == RouterSocketType
}

// IsCommandTypeValid returns if a command is valid for this socket.
func (routerSocket) IsCommandTypeValid(name string) bool {
	// FIXME
	return false
}

type pubSocket struct{}

func (pubSocket) Type() SocketType {
	return PubSocketType
}

func (pubSocket) IsSocketTypeCompatible(socketType SocketType) bool {
	return socketType == SubSocketType
}

func (pubSocket) IsCommandTypeValid(name string) bool {
	// FIXME
	return false
}

type subSocket struct{}

func (subSocket) Type() SocketType {
	return SubSocketType
}

func (subSocket) IsSocketTypeCompatible(socketType SocketType) bool {
	return socketType == PubSocketType
}

func (subSocket) IsCommandTypeValid(name string) bool {
	// FIXME
	return false
}

type reqSocket struct{}

func (reqSocket) Type() SocketType {
	return ReqSocketType
}

func (reqSocket) IsSocketTypeCompatible(socketType SocketType) bool {
	return socketType == RepSocketType
}

func (reqSocket) IsCommandTypeValid(name string) bool {
	// FIXME
	return false
}

type repSocket struct{}

func (repSocket) Type() SocketType {
	return RepSocketType
}

func (repSocket) IsSocketTypeCompatible(socketType SocketType) bool {
	return socketType == ReqSocketType
}

func (repSocket) IsCommandTypeValid(name string) bool {
	// FIXME
	return false
}
