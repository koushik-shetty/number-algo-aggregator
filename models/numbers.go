package models

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"sort"

	"fmt"

	"github.com/koushik-shetty/number-algo-aggregator/constants"
	"github.com/koushik-shetty/number-algo-aggregator/logger"
	"github.com/koushik-shetty/number-algo-aggregator/urlparser"
)

//NumbersResponse represents the response to the /numbers endpoint
type NumbersResponse struct {
	Numbers []int `json:"numbers"`
}

//NumberChan is used to receive the number data from the services
type NumberChan chan NumbersResponse

//Numbers fetches data from each of the urls in the QueryURL
func Numbers(ctx context.Context, log logger.Logger, up *urlparser.QueryURL) *NumbersResponse {
	if up == nil {
		return &NumbersResponse{
			Numbers: []int{},
		}
	}

	algoChannels := make(chan NumbersResponse, up.Length())
	errorChan := make(chan error, up.Length())
	for i := 0; i < up.Length(); i++ {
		go FetchAlgo(ctx, up.AllURL()[i], algoChannels, errorChan)
	}

	serviceNumbers := []int{}
	for i := 0; i < up.Length(); i++ {
		select {
		case data := <-algoChannels:
			serviceNumbers = append(serviceNumbers, data.Numbers...)
		case err := <-errorChan:
			log.Errorf("service error: %v", err)
		}
	}

	uniqueNum := unique(serviceNumbers)
	sort.Ints(uniqueNum)

	return &NumbersResponse{
		Numbers: uniqueNum,
	}
}

//FetchAlgo talks to a single endpoint and gets data from it if within the timeout else will error out.
func FetchAlgo(ctx context.Context, rawurl string, algoChannel NumberChan, errorChan chan error) error {
	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return err
	}

	// timedCtx will elapse after the timeout
	timedCtx, cancelFn := context.WithTimeout(ctx, constants.TIMEOUT)
	defer cancelFn()

	data := make(chan NumbersResponse, 1)
	errChan := make(chan error, 1)

	//this routine does the service call
	go func() {
		client := http.Client{}
		resp := &http.Response{}
		resp, err = client.Do(req.WithContext(timedCtx))
		if err != nil {
			errChan <- err
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}

		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			numbers := &NumbersResponse{}
			err = json.Unmarshal(body, numbers)
			if err != nil {
				errChan <- err
				return
			}
			data <- *numbers
		}
		errChan <- errors.New(string(body))
	}()

	//this selector selects between the timeout, data or error channels and provides ap
	select {
	case numbers := <-data:
		algoChannel <- numbers
	case <-timedCtx.Done():
		errorChan <- fmt.Errorf("endpoint: %v\nreason: %v", rawurl, timedCtx.Err())
	case err := <-errChan:
		errorChan <- fmt.Errorf("endpoint: %v\nreason: %v", rawurl, err)
	}
	return nil
}

func unique(algoNumbers []int) []int {
	uniqueMap := map[int]int{}
	for _, v := range algoNumbers {
		uniqueMap[v] = v
	}

	uniqueList := []int{}
	for _, v := range uniqueMap {
		uniqueList = append(uniqueList, v)
	}
	return uniqueList
}
