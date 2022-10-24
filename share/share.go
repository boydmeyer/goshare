package share

import (
	"fmt"
	"net"
	"net/http"

	qr "github.com/Baozisoftware/qrcode-terminal-go"
)

type Share struct {
	localIP			string
	Port			string
	Directory		string
	HideQR			bool
	URL				string
}

// Creates a new Share Module to start a mux server
func New(port string, directory string, hideQR bool) (*Share, error) {
	ip, err := getLocalIP()
	if err != nil {
		return nil, err
	}
	s := Share{
		localIP: ip,
		Port: port,
		Directory: directory,
		HideQR: hideQR,
	}
	s.URL = fmt.Sprintf("http://%s:%v", s.localIP, s.Port)
	return &s, nil
}

// Start Sharing Server
func (s *Share) StartServer() {
	//Print Sharing URL
	fmt.Println(s.URL)

	// Show QR-Code in terminal
	if !s.HideQR {
		obj := qr.New()
		obj.Get(s.URL).Print()
	}

	// Start Mux Server
	m := http.NewServeMux()
	m.Handle("/", http.FileServer(http.Dir(s.Directory)))
	panic(http.ListenAndServe(":"+s.Port, m))
}

// Local IP Address
func getLocalIP() (string, error) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addresses {
		ip := address.(*net.IPNet)
		if ip.IP.To4() != nil && !ip.IP.IsLoopback() {
			return ip.IP.String(), nil
		}
	}
	return "127.0.0.1", nil
}
