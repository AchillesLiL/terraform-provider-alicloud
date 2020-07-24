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

// DescribeGrantRules invokes the smartag.DescribeGrantRules API synchronously
// api document: https://help.aliyun.com/api/smartag/describegrantrules.html
func (client *Client) DescribeGrantRules(request *DescribeGrantRulesRequest) (response *DescribeGrantRulesResponse, err error) {
	response = CreateDescribeGrantRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGrantRulesWithChan invokes the smartag.DescribeGrantRules API asynchronously
// api document: https://help.aliyun.com/api/smartag/describegrantrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGrantRulesWithChan(request *DescribeGrantRulesRequest) (<-chan *DescribeGrantRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeGrantRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGrantRules(request)
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

// DescribeGrantRulesWithCallback invokes the smartag.DescribeGrantRules API asynchronously
// api document: https://help.aliyun.com/api/smartag/describegrantrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGrantRulesWithCallback(request *DescribeGrantRulesRequest, callback func(response *DescribeGrantRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGrantRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGrantRules(request)
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

// DescribeGrantRulesRequest is the request struct for api DescribeGrantRules
type DescribeGrantRulesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AssociatedCcnId      string           `position:"Query" name:"AssociatedCcnId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeGrantRulesResponse is the response struct for api DescribeGrantRules
type DescribeGrantRulesResponse struct {
	*responses.BaseResponse
	RequestId  string                         `json:"RequestId" xml:"RequestId"`
	TotalCount int                            `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                            `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                            `json:"PageSize" xml:"PageSize"`
	GrantRules GrantRulesInDescribeGrantRules `json:"GrantRules" xml:"GrantRules"`
}

// CreateDescribeGrantRulesRequest creates a request to invoke DescribeGrantRules API
func CreateDescribeGrantRulesRequest() (request *DescribeGrantRulesRequest) {
	request = &DescribeGrantRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeGrantRules", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGrantRulesResponse creates a response to parse from DescribeGrantRules response
func CreateDescribeGrantRulesResponse() (response *DescribeGrantRulesResponse) {
	response = &DescribeGrantRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
