// Copyright 2023 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package verificationmanagers is used to verify the agent packages
package verificationmanagers

import "github.com/aws/amazon-ssm-agent/agent/log"

type IVerificationManager interface {
	// VerifySignature verifies the agent binary signature
	VerifySignature(log log.T, signaturePath string, artifactsPath string, fileExtension string) error
}