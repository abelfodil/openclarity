// Copyright © 2024 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"

	cdx "github.com/CycloneDX/cyclonedx-go"

	"github.com/openclarity/openclarity/scanner/families"
	"github.com/openclarity/openclarity/scanner/utils/converter"
)

type Result struct {
	Metadata families.FamilyMetadata `json:"Metadata"`
	SBOM     *cdx.BOM                `json:"sbom"`
}

func NewResult(metadata families.FamilyMetadata, sbom *cdx.BOM) *Result {
	return &Result{
		Metadata: metadata,
		SBOM:     sbom,
	}
}

func (r *Result) EncodeToBytes(outputFormat string) ([]byte, error) {
	f, err := converter.StringToSbomFormat(outputFormat)
	if err != nil {
		return nil, fmt.Errorf("unable to parse output format: %w", err)
	}

	bomBytes := []byte{}
	if r != nil && r.SBOM != nil {
		bomBytes, err = converter.CycloneDxToBytes(r.SBOM, f)
		if err != nil {
			return nil, fmt.Errorf("unable to encode results to bytes: %w", err)
		}
	}

	return bomBytes, nil
}
