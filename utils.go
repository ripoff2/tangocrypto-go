package tangocrypto_go

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func handleAPIErrorResponse(res *http.Response) error {
	var err error
	switch res.StatusCode {
	case 400:
		br := BadRequest{}
		if err = json.NewDecoder(res.Body).Decode(&br); err != nil {
			return err
		}
		return &APIError{
			Response: br,
		}
	case 401:
		ue := UnauthorizedError{}
		if err = json.NewDecoder(res.Body).Decode(&ue); err != nil {
			return err
		}
		return &APIError{
			Response: ue,
		}
	case 403:
		fe := ForbiddenError{}
		if err = json.NewDecoder(res.Body).Decode(&fe); err != nil {
			return err
		}
		return &APIError{
			Response: fe,
		}
	case 404:
		nf := NotFound{}
		if err = json.NewDecoder(res.Body).Decode(&nf); err != nil {
			return err
		}
		return &APIError{
			Response: nf,
		}
	case 429:
		tmre := TooManyRequestsError{}
		if err = json.NewDecoder(res.Body).Decode(&tmre); err != nil {
			return err
		}
		return &APIError{
			Response: tmre,
		}
	case 500:
		ise := InternalServerError{}
		if err = json.NewDecoder(res.Body).Decode(&ise); err != nil {
			return err
		}
		return &APIError{
			Response: ise,
		}
	default:
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return &APIError{
			Response: string(data),
		}
	}
}

func (c *apiClient) handleRequest(req *http.Request) (res *http.Response, err error) {
	req.Header.Add("x-api-key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err = c.client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		return res, handleAPIErrorResponse(res)
	}

	return res, nil
}
