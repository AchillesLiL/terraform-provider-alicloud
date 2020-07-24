package ons

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

// OnsTrendTopicInputTps invokes the ons.OnsTrendTopicInputTps API synchronously
// api document: https://help.aliyun.com/api/ons/onstrendtopicinputtps.html
func (client *Client) OnsTrendTopicInputTps(request *OnsTrendTopicInputTpsRequest) (response *OnsTrendTopicInputTpsResponse, err error) {
	response = CreateOnsTrendTopicInputTpsResponse()
	err = client.DoAction(request, response)
	return
}

// OnsTrendTopicInputTpsWithChan invokes the ons.OnsTrendTopicInputTps API asynchronously
// api document: https://help.aliyun.com/api/ons/onstrendtopicinputtps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsTrendTopicInputTpsWithChan(request *OnsTrendTopicInputTpsRequest) (<-chan *OnsTrendTopicInputTpsResponse, <-chan error) {
	responseChan := make(chan *OnsTrendTopicInputTpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsTrendTopicInputTps(request)
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

// OnsTrendTopicInputTpsWithCallback invokes the ons.OnsTrendTopicInputTps API asynchronously
// api document: https://help.aliyun.com/api/ons/onstrendtopicinputtps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsTrendTopicInputTpsWithCallback(request *OnsTrendTopicInputTpsRequest, callback func(response *OnsTrendTopicInputTpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsTrendTopicInputTpsResponse
		var err error
		defer close(result)
		response, err = client.OnsTrendTopicInputTps(request)
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

// OnsTrendTopicInputTpsRequest is the request struct for api OnsTrendTopicInputTps
type OnsTrendTopicInputTpsRequest struct {
	*requests.RpcRequest
	Period     requests.Integer `position:"Query" name:"Period"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	BeginTime  requests.Integer `position:"Query" name:"BeginTime"`
	Type       requests.Integer `position:"Query" name:"Type"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	Topic      string           `position:"Query" name:"Topic"`
}

// OnsTrendTopicInputTpsResponse is the response struct for api OnsTrendTopicInputTps
type OnsTrendTopicInputTpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateOnsTrendTopicInputTpsRequest creates a request to invoke OnsTrendTopicInputTps API
func CreateOnsTrendTopicInputTpsRequest() (request *OnsTrendTopicInputTpsRequest) {
	request = &OnsTrendTopicInputTpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsTrendTopicInputTps", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsTrendTopicInputTpsResponse creates a response to parse from OnsTrendTopicInputTps response
func CreateOnsTrendTopicInputTpsResponse() (response *OnsTrendTopicInputTpsResponse) {
	response = &OnsTrendTopicInputTpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
