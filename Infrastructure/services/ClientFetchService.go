package services

import (
	"encoding/json"
	"errors"
	"github.com/apmath-web/interests/Domain"
	"github.com/apmath-web/interests/Infrastructure/applicationModels"
	"github.com/apmath-web/interests/Infrastructure/mappers"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var cientFecthServicePtr *clientFetchService
var oneTime sync.Once

type clientFetchService struct {
	url string
}

func (service *clientFetchService) GenURL() {
	host := os.Getenv("CLIENTS_HOST")
	port := os.Getenv("CLIENTS_PORT")
	version := os.Getenv("VERSION")
	service.url = "http://" + host + ":" + port + "/" + version + "/"
}

func Init() {
	cientFecthServicePtr = &clientFetchService{}
	cientFecthServicePtr.GenURL()
}

func GenClientFetchService() Domain.ClientFetchInterface {
	oneTime.Do(func() { Init() })
	return cientFecthServicePtr
}

func (service *clientFetchService) Fetch(id int) (Domain.PersonDomainModelInterface, error) {
	resp, err := http.Get(service.url + strconv.Itoa(id))

	if resp == nil {
		return nil, errors.New("clients service not available")
	}

	if err != nil {
		return nil, err
	}

	personApplicationModel := new(applicationModels.PersonApplicationModel)

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&personApplicationModel)
		if err != nil {
			return nil, err
		}
	}

	if resp.StatusCode == http.StatusBadRequest {
		return nil, errors.New("bad request")
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("client not found")
	}

	personDomainModel := mappers.PersonDomainModelMapper(*personApplicationModel)
	return personDomainModel, nil
}
