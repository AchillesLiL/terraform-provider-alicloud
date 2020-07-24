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

// OnsMqttGroupIdCreate invokes the ons.OnsMqttGroupIdCreate API synchronously
// api document: https://help.aliyun.com/api/ons/onsmqttgroupidcreate.html
func (client *Client) OnsMqttGroupIdCreate(request *OnsMqttGroupIdCreateRequest) (response *OnsMqttGroupIdCreateResponse, err error) {
	response = CreateOnsMqttGroupIdCreateResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMqttGroupIdCreateWithChan invokes the ons.OnsMqttGroupIdCreate API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmqttgroupidcreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMqttGroupIdCreateWithChan(request *OnsMqttGroupIdCreateRequest) (<-chan *OnsMqttGroupIdCreateResponse, <-chan error) {
	responseChan := make(chan *OnsMqttGroupIdCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMqttGroupIdCreate(request)
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

// OnsMqttGroupIdCreateWithCallback invokes the ons.OnsMqttGroupIdCreate API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmqttgroupidcreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMqttGroupIdCreateWithCallback(request *OnsMqttGroupIdCreateRequest, callback func(response *OnsMqttGroupIdCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMqttGroupIdCreateResponse
		var err error
		defer close(result)
		response, err = client.OnsMqttGroupIdCreate(request)
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

// OnsMqttGroupIdCreateRequest is the request struct for api OnsMqttGroupIdCreate
type OnsMqttGroupIdCreateRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
}

// OnsMqttGroupIdCreateResponse is the response struct for api OnsMqttGroupIdCreate
type OnsMqttGroupIdCreateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

// CreateOnsMqttGroupIdCreateRequest creates a request to invoke OnsMqttGroupIdCreate API
func CreateOnsMqttGroupIdCreateRequest() (request *OnsMqttGroupIdCreateRequest) {
	request = &OnsMqttGroupIdCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMqttGroupIdCreate", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsMqttGroupIdCreateResponse creates a response to parse from OnsMqttGroupIdCreate response
func CreateOnsMqttGroupIdCreateResponse() (response *OnsMqttGroupIdCreateResponse) {
	response = &OnsMqttGroupIdCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
