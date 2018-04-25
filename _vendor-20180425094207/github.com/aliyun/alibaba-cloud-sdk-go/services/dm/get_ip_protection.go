package dm

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

// GetIpProtection invokes the dm.GetIpProtection API synchronously
// api document: https://help.aliyun.com/api/dm/getipprotection.html
func (client *Client) GetIpProtection(request *GetIpProtectionRequest) (response *GetIpProtectionResponse, err error) {
	response = CreateGetIpProtectionResponse()
	err = client.DoAction(request, response)
	return
}

// GetIpProtectionWithChan invokes the dm.GetIpProtection API asynchronously
// api document: https://help.aliyun.com/api/dm/getipprotection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetIpProtectionWithChan(request *GetIpProtectionRequest) (<-chan *GetIpProtectionResponse, <-chan error) {
	responseChan := make(chan *GetIpProtectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetIpProtection(request)
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

// GetIpProtectionWithCallback invokes the dm.GetIpProtection API asynchronously
// api document: https://help.aliyun.com/api/dm/getipprotection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetIpProtectionWithCallback(request *GetIpProtectionRequest, callback func(response *GetIpProtectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetIpProtectionResponse
		var err error
		defer close(result)
		response, err = client.GetIpProtection(request)
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

// GetIpProtectionRequest is the request struct for api GetIpProtection
type GetIpProtectionRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
}

// GetIpProtectionResponse is the response struct for api GetIpProtection
type GetIpProtectionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	IpProtection string `json:"IpProtection" xml:"IpProtection"`
}

// CreateGetIpProtectionRequest creates a request to invoke GetIpProtection API
func CreateGetIpProtectionRequest() (request *GetIpProtectionRequest) {
	request = &GetIpProtectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "GetIpProtection", "", "")
	return
}

// CreateGetIpProtectionResponse creates a response to parse from GetIpProtection response
func CreateGetIpProtectionResponse() (response *GetIpProtectionResponse) {
	response = &GetIpProtectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}