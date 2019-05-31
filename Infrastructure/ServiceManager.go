package Infrastructure

import (
	"github.com/apmath-web/expenses/Domain"
	"os"
	"sync"
)

type serviceManager struct {
}

func (sm *serviceManager) GetClientFetchService() Domain.ClientFetchInterface {
	host := os.Getenv("CLIENTS_HOST")
	port := os.Getenv("CLIENTS_PORT")
	version := os.Getenv("CLIENTS_VERSION")
	return GenClientFetchService(host, port, version)
}

var sm *serviceManager
var once sync.Once

func GetServiceManager() *serviceManager {
	once.Do(func() {
		sm = &serviceManager{}
	})
	return sm
}
