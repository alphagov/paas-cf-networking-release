package handlers

import (
	"net/http"
	"policy-server/store"
	"io/ioutil"
)

//go:generate counterfeiter -o fakes/egress_policy_mapper.go --fake-name EgressPolicyMapper . egressPolicyMapper
type egressPolicyMapper interface {
	AsStoreEgressPolicy(bytes []byte) ([]store.EgressPolicy, error)
}

type EgressPolicyCreate struct {
	Store egressPolicyStore
	Mapper egressPolicyMapper
}


func (e *EgressPolicyCreate) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	requestBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	storeEgressPolicies, err := e.Mapper.AsStoreEgressPolicy(requestBytes)
	if err != nil {
		panic(err)
	}

	err = e.Store.Create(storeEgressPolicies)
	if err != nil {
		panic(err)
	}
}
