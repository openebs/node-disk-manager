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

package features

import (
	"reflect"
	"testing"
)

func TestFeatureGateIsEnabled(t *testing.T) {
	testFG := make(featureFlag)
	testFG["feature1"] = false
	testFG["feature2"] = true
	tests := map[string]struct {
		fg      featureFlag
		feature Feature
		want    bool
	}{
		"when feature gate is empty": {
			fg:      nil,
			feature: "test",
			want:    false,
		},
		"when feature gate does not have the feature": {
			fg:      testFG,
			feature: "feature3",
			want:    false,
		},
		"when feature gate has the feature and feature is disabled": {
			fg:      testFG,
			feature: "feature1",
			want:    false,
		},
		"when feature gate has the feature and feature is enabled": {
			fg:      testFG,
			feature: "feature2",
			want:    true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.fg.IsEnabled(tt.feature); got != tt.want {
				t.Errorf("IsEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetFeatureFlag(t *testing.T) {
	F1 := Feature("FeatureGate1")
	F2 := Feature("FeatureGate2")
	F3 := Feature("FeatureGate3")
	supportedFeatures = []Feature{
		F1, F2, F3,
	}
	type args struct {
		features []string
	}
	tests := map[string]struct {
		args    args
		want    featureFlag
		wantErr bool
	}{
		"empty feature gate slice": {
			args: args{
				features: nil,
			},
			want:    make(featureFlag),
			wantErr: false,
		},
		"a single feature is added": {
			args: args{
				features: []string{"FeatureGate1"},
			},
			want: map[Feature]bool{
				F1: true,
			},
			wantErr: false,
		},
		"a single feature is set in disabled mode": {
			args: args{
				features: []string{"FeatureGate1=false"},
			},
			want: map[Feature]bool{
				F1: false,
			},
			wantErr: false,
		},
		"feature that is not present in the supported feature": {
			args: args{
				features: []string{"WrongFeatureGate"},
			},
			want:    make(featureFlag),
			wantErr: true,
		},
		"multiple non-default features in enabled and disabled state": {
			args: args{
				features: []string{"FeatureGate1", "FeatureGate2=false", "FeatureGate3=true"},
			},
			want: featureFlag{
				F1: true,
				F2: false,
				F3: true,
			},
			wantErr: false,
		},
		"wrong format in one feature gate": {
			args: args{
				features: []string{"FeatureGate1", "FeatureGate2=true=true"},
			},
			want: featureFlag{
				F1: true,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			fg := make(featureFlag)
			err := fg.SetFeatureFlag(tt.args.features)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetFeatureFlag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fg, tt.want) && err == nil {
				t.Errorf("SetFeatureFlag() got = %v, want %v", fg, tt.want)
			}
		})
	}
}

func TestBasicFeatureDependencies(t *testing.T) {
	firstFeature := Feature("FeatureGate1")
	secondFeature := Feature("FeatureGate2")
	featureDependsOnFirst := Feature("FeatureGateNeeds1")
	featureDependsOnSecond := Feature("FeatureGateNeeds2")
	featureDependsOnBoth := Feature("FeatureGateNeeds1And2")

	supportedFeatures = []Feature{
		firstFeature,
		secondFeature,
		featureDependsOnFirst,
		featureDependsOnSecond,
		featureDependsOnBoth,
	}

	featureDependencies = map[Feature][]Feature{
		featureDependsOnFirst:  {firstFeature},
		featureDependsOnSecond: {secondFeature},
		featureDependsOnBoth:   {firstFeature, secondFeature},
	}

	type args struct {
		features []string
	}

	tests := map[string]struct {
		args args
		want featureFlag
	}{
		"unmet dependency should fail": {
			args: args{
				features: []string{string(featureDependsOnFirst)},
			},
			want: map[Feature]bool{
				featureDependsOnFirst: false,
			},
		},
		"met dependency should succeed": {
			args: args{
				features: []string{string(firstFeature), string(featureDependsOnFirst)},
			},
			want: map[Feature]bool{
				firstFeature:          true,
				featureDependsOnFirst: true,
			},
		},
		"one of two unmet dependency should fail": {
			args: args{
				features: []string{string(firstFeature), string(featureDependsOnBoth)},
			},
			want: map[Feature]bool{
				firstFeature:         true,
				featureDependsOnBoth: false,
			},
		},
		"all dependencies met should succeed": {
			args: args{
				features: []string{
					string(firstFeature),
					string(secondFeature),
					string(featureDependsOnBoth)},
			},
			want: map[Feature]bool{
				firstFeature:         true,
				secondFeature:        true,
				featureDependsOnBoth: true,
			},
		},
		"multiple features can have same dependencies": {
			args: args{
				features: []string{
					string(firstFeature),
					string(secondFeature),
					string(featureDependsOnFirst),
					string(featureDependsOnSecond),
					string(featureDependsOnBoth)},
			},
			want: map[Feature]bool{
				firstFeature:           true,
				secondFeature:          true,
				featureDependsOnFirst:  true,
				featureDependsOnSecond: true,
				featureDependsOnBoth:   true,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			flags := make(featureFlag)
			err := flags.SetFeatureFlag(test.args.features)
			if err != nil {
				t.Errorf("SetFeatureFlag() error = %v", err)
				return
			}
			if !reflect.DeepEqual(flags, test.want) {
				t.Errorf("SetFeatureFlag() got = %v, want %v", flags, test.want)
			}
		})
	}

}

func TestNestedFeatureDependencies(t *testing.T) {
	firstFeature := Feature("FeatureGate1")
	secondFeature := Feature("FeatureGate2")
	thirdFeatureNeedsBoth := Feature("FeatureGate3Needs1And2")
	fourthFeatureNeedsThird := Feature("FeatureGate4Needs3")

	supportedFeatures = []Feature{
		firstFeature,
		secondFeature,
		thirdFeatureNeedsBoth,
		fourthFeatureNeedsThird,
	}

	featureDependencies = map[Feature][]Feature{
		thirdFeatureNeedsBoth:   {firstFeature, secondFeature},
		fourthFeatureNeedsThird: {thirdFeatureNeedsBoth},
	}

	type args struct {
		init   []string
		action []string
	}

	type want struct {
		initResults   featureFlag
		actionResults featureFlag
	}

	tests := map[string]struct {
		args args
		want want
	}{
		"satisfied nested features should succeed": {
			args: args{
				init:   []string{string(firstFeature), string(secondFeature), string(thirdFeatureNeedsBoth)},
				action: []string{string(fourthFeatureNeedsThird)},
			},
			want: want{
				initResults: map[Feature]bool{
					firstFeature:          true,
					secondFeature:         true,
					thirdFeatureNeedsBoth: true,
				},
				actionResults: map[Feature]bool{
					firstFeature:            true,
					secondFeature:           true,
					thirdFeatureNeedsBoth:   true,
					fourthFeatureNeedsThird: true,
				},
			},
		},
		"unsatisfied nested dependency should fail": {
			args: args{
				init:   []string{string(firstFeature)},
				action: []string{string(thirdFeatureNeedsBoth), string(fourthFeatureNeedsThird)},
			},
			want: want{
				initResults: map[Feature]bool{
					firstFeature: true,
				},
				actionResults: map[Feature]bool{
					firstFeature:            true,
					thirdFeatureNeedsBoth:   false,
					fourthFeatureNeedsThird: false,
				},
			},
		},
		"turning off dependency should turn off dependant": {
			args: args{
				init:   []string{string(firstFeature), string(secondFeature), string(thirdFeatureNeedsBoth)},
				action: []string{string(firstFeature + "=false")},
			},
			want: want{
				initResults: map[Feature]bool{
					firstFeature:          true,
					secondFeature:         true,
					thirdFeatureNeedsBoth: true,
				},
				actionResults: map[Feature]bool{
					firstFeature:          false,
					secondFeature:         true,
					thirdFeatureNeedsBoth: false,
				},
			},
		},
		"turning off dependency should turn off nested dependants": {
			args: args{
				init:   []string{string(firstFeature), string(secondFeature), string(thirdFeatureNeedsBoth), string(fourthFeatureNeedsThird)},
				action: []string{string(firstFeature + "=false")},
			},
			want: want{
				initResults: map[Feature]bool{
					firstFeature:            true,
					secondFeature:           true,
					thirdFeatureNeedsBoth:   true,
					fourthFeatureNeedsThird: true,
				},
				actionResults: map[Feature]bool{
					firstFeature:            false,
					secondFeature:           true,
					thirdFeatureNeedsBoth:   false,
					fourthFeatureNeedsThird: false,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			flags := make(featureFlag)
			initErr := flags.SetFeatureFlag(test.args.init)
			if initErr != nil {
				t.Errorf("SetFeatureFlag() error initializing test = %v", initErr)
				return
			}
			if !reflect.DeepEqual(flags, test.want.initResults) {
				t.Errorf("SetFeatureFlag() initializing test got = %v, want %v", flags, test.want)
				return
			}
			err := flags.SetFeatureFlag(test.args.action)
			if err != nil {
				t.Errorf("SetFeatureFlag() error = %v", err)
				return
			}
			if !reflect.DeepEqual(flags, test.want.actionResults) {
				t.Errorf("SetFeatureFlag() got = %v, want %v", flags, test.want)
			}
		})
	}

}
