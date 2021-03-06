package emr

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

// DataSource is a nested struct in emr response
type DataSource struct {
	Name           string                                  `json:"Name" xml:"Name"`
	ClusterId      string                                  `json:"ClusterId" xml:"ClusterId"`
	Capacity       int                                     `json:"Capacity" xml:"Capacity"`
	UsedSize       int                                     `json:"UsedSize" xml:"UsedSize"`
	CreateFrom     string                                  `json:"CreateFrom" xml:"CreateFrom"`
	GmtCreate      int                                     `json:"GmtCreate" xml:"GmtCreate"`
	Conf           string                                  `json:"Conf" xml:"Conf"`
	SourceType     string                                  `json:"SourceType" xml:"SourceType"`
	Modifier       string                                  `json:"Modifier" xml:"Modifier"`
	ConnectionInfo string                                  `json:"ConnectionInfo" xml:"ConnectionInfo"`
	UserId         string                                  `json:"UserId" xml:"UserId"`
	Creator        string                                  `json:"Creator" xml:"Creator"`
	GmtModified    int                                     `json:"GmtModified" xml:"GmtModified"`
	ClusterBizId   string                                  `json:"ClusterBizId" xml:"ClusterBizId"`
	ClusterName    string                                  `json:"ClusterName" xml:"ClusterName"`
	Id             string                                  `json:"Id" xml:"Id"`
	Status         string                                  `json:"Status" xml:"Status"`
	Description    string                                  `json:"Description" xml:"Description"`
	ConfigList     ConfigListInMetastoreDescribeDataSource `json:"ConfigList" xml:"ConfigList"`
}
