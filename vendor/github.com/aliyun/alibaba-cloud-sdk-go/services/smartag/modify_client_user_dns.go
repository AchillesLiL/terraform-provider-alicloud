package smartag

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

// ModifyClientUserDNS invokes the smartag.ModifyClientUserDNS API synchronously
// api document: https://help.aliyun.com/api/smartag/modifyclientuserdns.html
func (client *Client) ModifyClientUserDNS(request *ModifyClientUserDNSRequest) (response *ModifyClientUserDNSResponse, err error) {
	response = CreateModifyClientUserDNSResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyClientUserDNSWithChan invokes the smartag.ModifyClientUserDNS API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifyclientuserdns.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyClientUserDNSWithChan(request *ModifyClientUserDNSRequest) (<-chan *ModifyClientUserDNSResponse, <-chan error) {
	responseChan := make(chan *ModifyClientUserDNSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyClientUserDNS(request)
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

// ModifyClientUserDNSWithCallback invokes the smartag.ModifyClientUserDNS API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifyclientuserdns.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyClientUserDNSWithCallback(request *ModifyClientUserDNSRequest, callback func(response *ModifyClientUserDNSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyClientUserDNSResponse
		var err error
		defer close(result)
		response, err = client.ModifyClientUserDNS(request)
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

// ModifyClientUserDNSRequest is the request struct for api ModifyClientUserDNS
type ModifyClientUserDNSRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RecoveredDNS         *[]string        `position:"Query" name:"RecoveredDNS"  type:"Repeated"`
	AppDNS               *[]string        `position:"Query" name:"AppDNS"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// ModifyClientUserDNSResponse is the response struct for api ModifyClientUserDNS
type ModifyClientUserDNSResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyClientUserDNSRequest creates a request to invoke ModifyClientUserDNS API
func CreateModifyClientUserDNSRequest() (request *ModifyClientUserDNSRequest) {
	request = &ModifyClientUserDNSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifyClientUserDNS", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyClientUserDNSResponse creates a response to parse from ModifyClientUserDNS response
func CreateModifyClientUserDNSResponse() (response *ModifyClientUserDNSResponse) {
	response = &ModifyClientUserDNSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
