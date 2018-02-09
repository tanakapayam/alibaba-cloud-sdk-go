package mts

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

func (client *Client) UpdateCensorPipeline(request *UpdateCensorPipelineRequest) (response *UpdateCensorPipelineResponse, err error) {
	response = CreateUpdateCensorPipelineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdateCensorPipelineWithChan(request *UpdateCensorPipelineRequest) (<-chan *UpdateCensorPipelineResponse, <-chan error) {
	responseChan := make(chan *UpdateCensorPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCensorPipeline(request)
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

func (client *Client) UpdateCensorPipelineWithCallback(request *UpdateCensorPipelineRequest, callback func(response *UpdateCensorPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCensorPipelineResponse
		var err error
		defer close(result)
		response, err = client.UpdateCensorPipeline(request)
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

type UpdateCensorPipelineRequest struct {
	*requests.RpcRequest
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	Name                 string           `position:"Query" name:"Name"`
	State                string           `position:"Query" name:"State"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type UpdateCensorPipelineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Pipeline  struct {
		Id           string `json:"Id" xml:"Id"`
		Name         string `json:"Name" xml:"Name"`
		State        string `json:"State" xml:"State"`
		Priority     int    `json:"Priority" xml:"Priority"`
		NotifyConfig struct {
			Topic string `json:"Topic" xml:"Topic"`
			Queue string `json:"Queue" xml:"Queue"`
		} `json:"NotifyConfig" xml:"NotifyConfig"`
	} `json:"Pipeline" xml:"Pipeline"`
}

func CreateUpdateCensorPipelineRequest() (request *UpdateCensorPipelineRequest) {
	request = &UpdateCensorPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateCensorPipeline", "mts", "openAPI")
	return
}

func CreateUpdateCensorPipelineResponse() (response *UpdateCensorPipelineResponse) {
	response = &UpdateCensorPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}