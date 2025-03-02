package mediamtx

import (
	"encoding/json"
	"net/http"
)

type HLSMuxer struct {
	Path        string
	Created     string
	LastRequest string
	BytesSent   int
}

type HLSListResponse struct {
	PageCount int
	ItemCount int
	Items     []HLSMuxer
}

func (m MediamtxAPI) GetHLSMuxers() ([]HLSMuxer, error) {
	resp, err := http.Get(m.apiUrl + "/v3/hlsmuxers/list")
	if err != nil {
		return nil, err
	}

	var respBody HLSListResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody.Items, nil
}

func (m MediamtxAPI) GetHLSMuxer(name string) (HLSMuxer, error) {
	resp, err := http.Get(m.apiUrl + "/v3/hlsmuxers/get/" + name)
	if err != nil {
		return HLSMuxer{}, err
	}

	var respBody HLSMuxer
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return HLSMuxer{}, err
	}

	return respBody, nil
}
