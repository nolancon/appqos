package appqos

// AppQoS API Calls + Marshalling

import (
	//	"crypto/ecdsa"
	//	"crypto/rsa"
	"crypto/tls"
	//	"crypto/x509"
	//	"encoding/json"
	//	"fmt"
	//	"io/ioutil"
	//	"k8s.io/apimachinery/pkg/api/errors"
	"net/http"
	//	"reflect"
	//	"strconv"
	//	"strings"
)

// AppQoSClient is used by the operator to become a client to AppQoS
type AppQoSClient struct {
	client *http.Client
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
