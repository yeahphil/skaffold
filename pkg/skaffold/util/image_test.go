/*
Copyright 2018 The Skaffold Authors

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

package util

import (
	"testing"

	"github.com/GoogleContainerTools/skaffold/testutil"
)

func TestImageReplaceDefaultRepo(t *testing.T) {
	tests := []struct {
		name          string
		image         string
		defaultRepo   string
		expectedImage string
	}{
		{
			name:          "basic GCR concatenation",
			image:         "gcr.io/some/registry",
			defaultRepo:   "gcr.io/default",
			expectedImage: "gcr.io/default/gcr.io/some/registry",
		},
		{
			name:          "no default repo set",
			image:         "gcr.io/some/registry",
			expectedImage: "gcr.io/some/registry",
		},
		{
			name:          "provided image has defaultRepo prefix",
			image:         "gcr.io/default/registry",
			defaultRepo:   "gcr.io/default",
			expectedImage: "gcr.io/default/registry",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testutil.CheckDeepEqual(t, test.expectedImage, SubstituteDefaultRepoIntoImage(test.defaultRepo, test.image))
		})
	}
}
