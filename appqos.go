package appqos

// AppQoS API Calls + Marshalling

import (
	"bytes"
	//	"crypto/ecdsa"
	//	"crypto/rsa"
	"crypto/tls"
	//	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"net/http"
	//	"reflect"
	"strconv"
	//	"strings"
)

const (
	poolsEndpoint         = "/pools"
	appsEndpoint          = "/apps"
	powerProfilesEndpoint = "/power_profiles"
	username              = "admin"
	passwd                = "password"
)

// OperatorRmdClient is used by the operator to become a client to RMD
type AppQoSClient struct {
	client *http.Client
}

// Pools
type Pool struct {
	Name         *string       `json:"name,omitempty"`
	ID           *int          `json:"id,omitempty"`
	Apps         *[]App        `json:"apps,omitempty"`
	Cbm          *int          `json:"cbm,omitempty"`
	Mba          *int          `json:"mba,omitempty"`
	MbaBw        *int          `json:"mba_bw,omitempty"`
	Cores        *[]int        `json:"cores,omitempty"`
	PowerProfile *PowerProfile `json:"power_profiles,omitempty"`
}

// NewDefaultAppQosClient returns a default client for testing and debugging
func NewDefaultAppQoSClient() *AppQoSClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defaultClient := &http.Client{Transport: tr}
	rmdClient := &AppQoSClient{
		client: defaultClient,
	}
	return rmdClient
}

// GetPools /pools
func (ac *AppQoSClient) GetPools(address string) ([]Pool, error) {
	httpString := fmt.Sprintf("%s%s", address, poolsEndpoint)

	req, err := http.NewRequest("GET", httpString, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, passwd)
	resp, err := ac.client.Do(req)
	if err != nil {
		return nil, err
	}
	receivedJSON, err := ioutil.ReadAll(resp.Body) //This reads raw request body
	if err != nil {
		return nil, err
	}

	allPools := make([]Pool, 0)
	err = json.Unmarshal([]byte(receivedJSON), &allPools)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	return allPools, nil
}

// GetPool /pools/{id}
func (ac *AppQoSClient) GetPool(address string, id int) (*Pool, error) {
	httpString := fmt.Sprintf("%s%s%s%s", address, poolsEndpoint, "/", strconv.Itoa(id))
	pool := &Pool{}
	req, err := http.NewRequest("GET", httpString, nil)
	if err != nil {
		return pool, err
	}

	req.SetBasicAuth(username, passwd)
	resp, err := ac.client.Do(req)
	if err != nil {
		return pool, err
	}
	receivedJSON, err := ioutil.ReadAll(resp.Body) //This reads raw request body
	if err != nil {
		return pool, err
	}

	err = json.Unmarshal([]byte(receivedJSON), pool)
	if err != nil {
		return pool, err
	}

	resp.Body.Close()

	return pool, nil
}

// PostPools /pools
// PostWorkload posts workload data from RmdWorkload to RMD
func (ac *AppQoSClient) PostPool(pool *Pool, address string) (string, error) {
	postFailedErr := errors.NewServiceUnavailable("Response status code error")

	payloadBytes, err := json.Marshal(pool)
	if err != nil {
		return "Failed to marshal payload data", err
	}
	body := bytes.NewReader(payloadBytes)

	httpString := fmt.Sprintf("%s%s%s", address, "/", poolsEndpoint)
	req, err := http.NewRequest("POST", httpString, body)
	if err != nil {
		return "Failed to create new http post request", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)
	resp, err := ac.client.Do(req)
	if err != nil {
		return "Failed to set header for http post request", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respStr := buf.String()

	if resp.StatusCode != 201 {
		errStr := fmt.Sprintf("%s%v", "Fail: ", respStr)
		return errStr, postFailedErr
	}

	defer resp.Body.Close()

	successStr := fmt.Sprintf("%s%v", "Success: ", resp.StatusCode)

	return successStr, nil
}

// PutPool /pools/{id}

// DeletePools /pools/{id}

// Power Profiles
type PowerProfile struct {
	ID      string
	Name    string
	MinFreq int
	MaxFreq int
	Epp     string
}

// GetPowerProfiles /power_profiles

// GetPowerProfile /power_profiles/{id}

// PostPowerProfiles /power_profiles

// PutPowerProfiles /power_profiles/{id}

// DeletePowerProfiles /power_profiles/{id}

// Apps
type App struct {
	ID    string
	Name  string
	Cores []string
	Pid   []string
}

// GetApps /apps

// GetApp /apps/{id}

// PostApps /apps

// PutApps /apps/{id}

// DeleteApps /apps/{id}
