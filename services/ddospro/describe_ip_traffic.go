package ddospro

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeIpTraffic(request *DescribeIpTrafficRequest) (response *DescribeIpTrafficResponse, err error) {
	response = CreateDescribeIpTrafficResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeIpTrafficWithChan(request *DescribeIpTrafficRequest) (<-chan *DescribeIpTrafficResponse, <-chan error) {
	responseChan := make(chan *DescribeIpTrafficResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIpTraffic(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeIpTrafficWithCallback(request *DescribeIpTrafficRequest, callback func(response *DescribeIpTrafficResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIpTrafficResponse
		var err error
		defer close(result)
		response, err = client.DescribeIpTraffic(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeIpTrafficRequest struct {
	*requests.RpcRequest
	Ip              string           `position:"Query" name:"Ip"`
	EndDateMillis   requests.Integer `position:"Query" name:"EndDateMillis"`
	StartDateMillis requests.Integer `position:"Query" name:"StartDateMillis"`
}

type DescribeIpTrafficResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      struct {
		PpsDrop   []string `json:"PpsDrop" xml:"PpsDrop"`
		PpsTotal  []string `json:"PpsTotal" xml:"PpsTotal"`
		BpsDrop   []string `json:"BpsDrop" xml:"BpsDrop"`
		BpsTotal  []string `json:"BpsTotal" xml:"BpsTotal"`
		TimeScope struct {
			Interval int `json:"Interval" xml:"Interval"`
			Start    int `json:"Start" xml:"Start"`
		} `json:"TimeScope" xml:"TimeScope"`
	} `json:"Data" xml:"Data"`
}

func CreateDescribeIpTrafficRequest() (request *DescribeIpTrafficRequest) {
	request = &DescribeIpTrafficRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "DescribeIpTraffic", "", "")
	return
}

func CreateDescribeIpTrafficResponse() (response *DescribeIpTrafficResponse) {
	response = &DescribeIpTrafficResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
