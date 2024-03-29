/*
Copyright 2020 The OpenEBS Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package services

import (
	"context"
	"os"

	protos "github.com/openebs/node-disk-manager/spec/ndm"

	"k8s.io/klog/v2"
)

// Name is used to find the name of the worker node NDM is deployed on
func (n *Node) Name(ctx context.Context, null *protos.Null) (*protos.NodeName, error) {

	// Fetch the environment variable
	nodeName := os.Getenv("NODE_NAME")

	klog.Infof("Node name is : %v", nodeName)

	return &protos.NodeName{NodeName: nodeName}, nil

}
