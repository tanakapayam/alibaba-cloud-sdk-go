package emr

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

// GetJobInputStatisticInfo invokes the emr.GetJobInputStatisticInfo API synchronously
// api document: https://help.aliyun.com/api/emr/getjobinputstatisticinfo.html
func (client *Client) GetJobInputStatisticInfo(request *GetJobInputStatisticInfoRequest) (response *GetJobInputStatisticInfoResponse, err error) {
	response = CreateGetJobInputStatisticInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetJobInputStatisticInfoWithChan invokes the emr.GetJobInputStatisticInfo API asynchronously
// api document: https://help.aliyun.com/api/emr/getjobinputstatisticinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetJobInputStatisticInfoWithChan(request *GetJobInputStatisticInfoRequest) (<-chan *GetJobInputStatisticInfoResponse, <-chan error) {
	responseChan := make(chan *GetJobInputStatisticInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJobInputStatisticInfo(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetJobInputStatisticInfoWithCallback invokes the emr.GetJobInputStatisticInfo API asynchronously
// api document: https://help.aliyun.com/api/emr/getjobinputstatisticinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetJobInputStatisticInfoWithCallback(request *GetJobInputStatisticInfoRequest, callback func(response *GetJobInputStatisticInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJobInputStatisticInfoResponse
		var err error
		defer close(result)
		response, err = client.GetJobInputStatisticInfo(request)
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

// GetJobInputStatisticInfoRequest is the request struct for api GetJobInputStatisticInfo
type GetJobInputStatisticInfoRequest struct {
	*requests.RpcRequest
	FromDatetime    string           `position:"Query" name:"FromDatetime"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	ToDatetime      string           `position:"Query" name:"ToDatetime"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
}

// GetJobInputStatisticInfoResponse is the response struct for api GetJobInputStatisticInfo
type GetJobInputStatisticInfoResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Total        int          `json:"Total" xml:"Total"`
	PageNumber   int          `json:"PageNumber" xml:"PageNumber"`
	PageSize     int          `json:"PageSize" xml:"PageSize"`
	JobInputList JobInputList `json:"JobInputList" xml:"JobInputList"`
}

// CreateGetJobInputStatisticInfoRequest creates a request to invoke GetJobInputStatisticInfo API
func CreateGetJobInputStatisticInfoRequest() (request *GetJobInputStatisticInfoRequest) {
	request = &GetJobInputStatisticInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "GetJobInputStatisticInfo", "emr", "openAPI")
	return
}

// CreateGetJobInputStatisticInfoResponse creates a response to parse from GetJobInputStatisticInfo response
func CreateGetJobInputStatisticInfoResponse() (response *GetJobInputStatisticInfoResponse) {
	response = &GetJobInputStatisticInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
