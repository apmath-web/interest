package Infrastructure

import (
	"github.com/apmath-web/interests/Domain"
	"github.com/apmath-web/interests/Domain/services"
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

func (sm *serviceManager) GetCalculationService() services.CalculationService {
	clientFetch := sm.GetClientFetchService()
	calc := sm.GetCalculateService()
	cs := services.GenCalculationService(clientFetch, calc)
	return cs
}

func (sm *serviceManager) GetCalculateService() services.CalculateService {
	cs := new(services.CalculateService)
	return *cs
}
