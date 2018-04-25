package csb

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

// UpdateOrder invokes the csb.UpdateOrder API synchronously
// api document: https://help.aliyun.com/api/csb/updateorder.html
func (client *Client) UpdateOrder(request *UpdateOrderRequest) (response *UpdateOrderResponse, err error) {
	response = CreateUpdateOrderResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateOrderWithChan invokes the csb.UpdateOrder API asynchronously
// api document: https://help.aliyun.com/api/csb/updateorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateOrderWithChan(request *UpdateOrderRequest) (<-chan *UpdateOrderResponse, <-chan error) {
	responseChan := make(chan *UpdateOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateOrder(request)
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

// UpdateOrderWithCallback invokes the csb.UpdateOrder API asynchronously
// api document: https://help.aliyun.com/api/csb/updateorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateOrderWithCallback(request *UpdateOrderRequest, callback func(response *UpdateOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateOrderResponse
		var err error
		defer close(result)
		response, err = client.UpdateOrder(request)
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

// UpdateOrderRequest is the request struct for api UpdateOrder
type UpdateOrderRequest struct {
	*requests.RpcRequest
	CsbId requests.Integer `position:"Query" name:"CsbId"`
	Data  string           `position:"Body" name:"Data"`
}

// UpdateOrderResponse is the response struct for api UpdateOrder
type UpdateOrderResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateOrderRequest creates a request to invoke UpdateOrder API
func CreateUpdateOrderRequest() (request *UpdateOrderRequest) {
	request = &UpdateOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "UpdateOrder", "CSB", "openAPI")
	return
}

// CreateUpdateOrderResponse creates a response to parse from UpdateOrder response
func CreateUpdateOrderResponse() (response *UpdateOrderResponse) {
	response = &UpdateOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}