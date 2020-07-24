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

// DescribeSagWanSnat invokes the smartag.DescribeSagWanSnat API synchronously
// api document: https://help.aliyun.com/api/smartag/describesagwansnat.html
func (client *Client) DescribeSagWanSnat(request *DescribeSagWanSnatRequest) (response *DescribeSagWanSnatResponse, err error) {
	response = CreateDescribeSagWanSnatResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSagWanSnatWithChan invokes the smartag.DescribeSagWanSnat API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesagwansnat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSagWanSnatWithChan(request *DescribeSagWanSnatRequest) (<-chan *DescribeSagWanSnatResponse, <-chan error) {
	responseChan := make(chan *DescribeSagWanSnatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSagWanSnat(request)
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

// DescribeSagWanSnatWithCallback invokes the smartag.DescribeSagWanSnat API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesagwansnat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSagWanSnatWithCallback(request *DescribeSagWanSnatRequest, callback func(response *DescribeSagWanSnatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSagWanSnatResponse
		var err error
		defer close(result)
		response, err = client.DescribeSagWanSnat(request)
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

// DescribeSagWanSnatRequest is the request struct for api DescribeSagWanSnat
type DescribeSagWanSnatRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
}

// DescribeSagWanSnatResponse is the response struct for api DescribeSagWanSnat
type DescribeSagWanSnatResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	Snat       string      `json:"Snat" xml:"Snat"`
	TaskStates []TaskState `json:"TaskStates" xml:"TaskStates"`
}

// CreateDescribeSagWanSnatRequest creates a request to invoke DescribeSagWanSnat API
func CreateDescribeSagWanSnatRequest() (request *DescribeSagWanSnatRequest) {
	request = &DescribeSagWanSnatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSagWanSnat", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSagWanSnatResponse creates a response to parse from DescribeSagWanSnat response
func CreateDescribeSagWanSnatResponse() (response *DescribeSagWanSnatResponse) {
	response = &DescribeSagWanSnatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
