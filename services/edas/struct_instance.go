package edas

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

// Instance is a nested struct in edas response
type Instance struct {
	InstanceId   string `json:"InstanceId" xml:"InstanceId"`
	InstanceName string `json:"InstanceName" xml:"InstanceName"`
	EcuId        string `json:"EcuId" xml:"EcuId"`
	VpcId        string `json:"VpcId" xml:"VpcId"`
	VpcName      string `json:"VpcName" xml:"VpcName"`
	Expired      bool   `json:"Expired" xml:"Expired"`
	Status       string `json:"Status" xml:"Status"`
	RegionId     string `json:"RegionId" xml:"RegionId"`
	Cpu          int    `json:"Cpu" xml:"Cpu"`
	Mem          int    `json:"Mem" xml:"Mem"`
	PublicIp     string `json:"PublicIp" xml:"PublicIp"`
	InnerIp      string `json:"InnerIp" xml:"InnerIp"`
	PrivateIp    string `json:"PrivateIp" xml:"PrivateIp"`
	Eip          string `json:"Eip" xml:"Eip"`
}