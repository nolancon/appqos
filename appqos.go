package appqos

// AppQoS API Calls + Marshalling

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"net/http"
	"strconv"
)

const (
	poolsEndpoint         = "/pools"
	appsEndpoint          = "/apps"
	powerProfilesEndpoint = "/power_profiles"
	username              = "admin"
	passwd                = "password"
)

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
func (ac *AppQoSClient) PostPool(pool *Pool, address string) (string, error) {
	postFailedErr := errors.NewServiceUnavailable("Response status code error")

	payloadBytes, err := json.Marshal(pool)
	if err != nil {
		return "Failed to marshal payload data", err
	}
	body := bytes.NewReader(payloadBytes)

	httpString := fmt.Sprintf("%s%s", address, poolsEndpoint)
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
func (ac *AppQoSClient) PutPool(pool *Pool, address string, poolID int) (string, error) {
	patchFailedErr := errors.NewServiceUnavailable("Response status code error")

	payloadBytes, err := json.Marshal(pool)
	if err != nil {
		return "Failed to marshal payload data", err
	}
	body := bytes.NewReader(payloadBytes)

	httpString := fmt.Sprintf("%s%s%s%s", address, poolsEndpoint, "/", strconv.Itoa(poolID))
	req, err := http.NewRequest("PUT", httpString, body)
	if err != nil {
		return "Failed to create new http patch request", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)
	resp, err := ac.client.Do(req)
	if err != nil {
		return "Failed to set header for http patch request", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respStr := buf.String()
	if resp.StatusCode != 200 {
		errStr := fmt.Sprintf("%s%v", "Fail: ", respStr)
		return errStr, patchFailedErr

	}
	defer resp.Body.Close()

	successStr := fmt.Sprintf("%s%v", "Success: ", resp.StatusCode)
	return successStr, nil
}

// DeletePool /pools/{id}
func (ac *AppQoSClient) DeletePool(address string, poolID int) error {
	httpString := fmt.Sprintf("%s%s%s%s", address, poolsEndpoint, "/", strconv.Itoa(poolID))
	req, err := http.NewRequest("DELETE", httpString, nil)
	if err != nil {
		return err

	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)
	resp, err := ac.client.Do(req)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if resp.StatusCode != 200 {
		deleteFailedErr := errors.NewServiceUnavailable(buf.String())
		return deleteFailedErr
	}

	defer resp.Body.Close()
	return nil
}

// GetPowerProfiles /power_profiles
func (ac *AppQoSClient) GetPowerProfiles(address string) ([]PowerProfile, error) {
	httpString := fmt.Sprintf("%s%s", address, powerProfilesEndpoint)

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

	allPowerProfiles := make([]PowerProfile, 0)
	err = json.Unmarshal([]byte(receivedJSON), &allPowerProfiles)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	return allPowerProfiles, nil
}

// GetPowerProfile /power_profiles
func (ac *AppQoSClient) GetPowerProfile(address string, id int) (*PowerProfile, error) {
	httpString := fmt.Sprintf("%s%s%s%s", address, powerProfilesEndpoint, "/", strconv.Itoa(id))
	powerProfile := &PowerProfile{}
	req, err := http.NewRequest("GET", httpString, nil)
	if err != nil {
		return powerProfile, err
	}

	req.SetBasicAuth(username, passwd)
	resp, err := ac.client.Do(req)
	if err != nil {
		return powerProfile, err
	}
	receivedJSON, err := ioutil.ReadAll(resp.Body) //This reads raw request body
	if err != nil {
		return powerProfile, err
	}

	err = json.Unmarshal([]byte(receivedJSON), powerProfile)
	if err != nil {
		return powerProfile, err
	}

	resp.Body.Close()

	return powerProfile, nil
}

// PostPowerProfile /power_profile
func (ac *AppQoSClient) PostPowerProfile(powerProfile *PowerProfile, address string) (string, error) {
	postFailedErr := errors.NewServiceUnavailable("Response status code error")

	payloadBytes, err := json.Marshal(powerProfile)
	if err != nil {
		return "Failed to marshal payload data", err
	}
	body := bytes.NewReader(payloadBytes)

	httpString := fmt.Sprintf("%s%s", address, powerProfilesEndpoint)
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

// PutPowerProfile /power_profiles/{id}
func (ac *AppQoSClient) PutPowerProfile(powerProfile *PowerProfile, address string, powerProfileID int) (string, error) {
	patchFailedErr := errors.NewServiceUnavailable("Response status code error")

	payloadBytes, err := json.Marshal(powerProfile)
	if err != nil {
		return "Failed to marshal payload data", err
	}
	body := bytes.NewReader(payloadBytes)

	httpString := fmt.Sprintf("%s%s%s%s", address, powerProfilesEndpoint, "/", strconv.Itoa(powerProfileID))
	req, err := http.NewRequest("PUT", httpString, body)
	if err != nil {
		return "Failed to create new http patch request", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)
	resp, err := ac.client.Do(req)
	if err != nil {
		return "Failed to set header for http patch request", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respStr := buf.String()
	if resp.StatusCode != 200 {
		errStr := fmt.Sprintf("%s%v", "Fail: ", respStr)
		return errStr, patchFailedErr

	}
	defer resp.Body.Close()

	successStr := fmt.Sprintf("%s%v", "Success: ", resp.StatusCode)
	return successStr, nil
}

// DeletePowerProfile /power_profiles/{id}
func (ac *AppQoSClient) DeletePowerProfile(address string, powerProfileID int) error {
	httpString := fmt.Sprintf("%s%s%s%s", address, powerProfilesEndpoint, "/", strconv.Itoa(powerProfileID))
	req, err := http.NewRequest("DELETE", httpString, nil)
	if err != nil {
		return err

	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)
	resp, err := ac.client.Do(req)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if resp.StatusCode != 200 {
		deleteFailedErr := errors.NewServiceUnavailable(buf.String())
		return deleteFailedErr
	}

	defer resp.Body.Close()
	return nil
}
