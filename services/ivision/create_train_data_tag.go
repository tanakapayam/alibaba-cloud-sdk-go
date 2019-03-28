package ivision

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

// CreateTrainDataTag invokes the ivision.CreateTrainDataTag API synchronously
// api document: https://help.aliyun.com/api/ivision/createtraindatatag.html
func (client *Client) CreateTrainDataTag(request *CreateTrainDataTagRequest) (response *CreateTrainDataTagResponse, err error) {
	response = CreateCreateTrainDataTagResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTrainDataTagWithChan invokes the ivision.CreateTrainDataTag API asynchronously
// api document: https://help.aliyun.com/api/ivision/createtraindatatag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTrainDataTagWithChan(request *CreateTrainDataTagRequest) (<-chan *CreateTrainDataTagResponse, <-chan error) {
	responseChan := make(chan *CreateTrainDataTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTrainDataTag(request)
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

// CreateTrainDataTagWithCallback invokes the ivision.CreateTrainDataTag API asynchronously
// api document: https://help.aliyun.com/api/ivision/createtraindatatag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTrainDataTagWithCallback(request *CreateTrainDataTagRequest, callback func(response *CreateTrainDataTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTrainDataTagResponse
		var err error
		defer close(result)
		response, err = client.CreateTrainDataTag(request)
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

// CreateTrainDataTagRequest is the request struct for api CreateTrainDataTag
type CreateTrainDataTagRequest struct {
	*requests.RpcRequest
	ProjectId string           `position:"Query" name:"ProjectId"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	TagItems  string           `position:"Query" name:"TagItems"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	DataId    string           `position:"Query" name:"DataId"`
}

// CreateTrainDataTagResponse is the response struct for api CreateTrainDataTag
type CreateTrainDataTagResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	TrainData TrainData `json:"TrainData" xml:"TrainData"`
}

// CreateCreateTrainDataTagRequest creates a request to invoke CreateTrainDataTag API
func CreateCreateTrainDataTagRequest() (request *CreateTrainDataTagRequest) {
	request = &CreateTrainDataTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "CreateTrainDataTag", "ivision", "openAPI")
	return
}

// CreateCreateTrainDataTagResponse creates a response to parse from CreateTrainDataTag response
func CreateCreateTrainDataTagResponse() (response *CreateTrainDataTagResponse) {
	response = &CreateTrainDataTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}