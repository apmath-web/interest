package Infrastructure

import (
	"encoding/json"
	"errors"
	"github.com/apmath-web/interests/Domain"
	"github.com/apmath-web/interests/Infrastructure/Mappers"
	"github.com/apmath-web/interests/Infrastructure/applicationModels"
	"net/http"
	"strconv"
)

type clientFetchService struct {
	url string
}

func GenClientFetchService(host string, port string, version string) Domain.ClientFetchInterface {
	var instantiated *clientFetchService
	instantiated = &clientFetchService{"http://" + host + ":" + port + "/" + version + "/"}
	return instantiated
}

func (clfs *clientFetchService) Fetch(id int) (Domain.PersonDomainModelInterface, error) {
	resp, err := http.Get(clfs.url + strconv.Itoa(id))
	if resp == nil {
		return nil, errors.New("clients service not available")
	}
	if err != nil {
		return nil, err
	}
	person := new(applicationModels.PersonApplicationModel)

	if resp.StatusCode == http.StatusOK {
		dec := json.NewDecoder(resp.Body)
		err := dec.Decode(&person)
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
	pdm := Mappers.PersonApplicationMapper(*person)
	return pdm, nil
}
