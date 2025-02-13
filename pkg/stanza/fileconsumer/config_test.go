// Copyright The OpenTelemetry Authors
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

package fileconsumer

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/operatortest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/testutil"
)

func TestUnmarshal(t *testing.T) {
	operatortest.ConfigUnmarshalTests{
		DefaultConfig: newMockOperatorConfig(NewConfig()),
		TestsFile:     filepath.Join(".", "testdata", "config.yaml"),
		Tests: []operatortest.ConfigUnmarshalTest{
			{
				Name: "include_one",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "one.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_multi",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "one.log", "two.log", "three.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_glob",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_glob_double_asterisk",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "**.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_glob_double_asterisk_nested",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "directory/**/*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_glob_double_asterisk_prefix",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "**/directory/**/*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_inline",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "a.log", "b.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_string",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "aString")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_one",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "one.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_multi",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "one.log", "two.log", "three.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_glob",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "not*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_glob_double_asterisk",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "not**.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_glob_double_asterisk_nested",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "directory/**/not*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_glob_double_asterisk_prefix",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "**/directory/**/not*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_inline",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "a.log", "b.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_string",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "aString")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "poll_interval_no_units",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.PollInterval = time.Second
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "poll_interval_1s",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.PollInterval = time.Second
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "poll_interval_1ms",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.PollInterval = time.Millisecond
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "poll_interval_1000ms",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.PollInterval = time.Second
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_no_units",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1000)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_1kb_lower",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1000)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_1KB",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1000)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_1kib_lower",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1024)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_1KiB",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1024)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_float",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1100)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "multiline_line_start_string",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					newSplit := helper.NewSplitterConfig()
					newSplit.Multiline.LineStartPattern = "Start"
					cfg.Splitter = newSplit
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "multiline_line_start_special",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					newSplit := helper.NewSplitterConfig()
					newSplit.Multiline.LineStartPattern = "%"
					cfg.Splitter = newSplit
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "multiline_line_end_string",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					newSplit := helper.NewSplitterConfig()
					newSplit.Multiline.LineEndPattern = "Start"
					cfg.Splitter = newSplit
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "multiline_line_end_special",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					newSplit := helper.NewSplitterConfig()
					newSplit.Multiline.LineEndPattern = "%"
					cfg.Splitter = newSplit
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "start_at_string",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.StartAt = "beginning"
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "max_concurrent_large",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.MaxConcurrentFiles = 9223372036854775807
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "max_log_size_mib_lower",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.MaxLogSize = helper.ByteSize(1048576)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "max_log_size_mib_upper",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.MaxLogSize = helper.ByteSize(1048576)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "max_log_size_mb_upper",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.MaxLogSize = helper.ByteSize(1048576)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "max_log_size_mb_lower",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.MaxLogSize = helper.ByteSize(1048576)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "encoding_lower",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Splitter.EncodingConfig = helper.EncodingConfig{Encoding: "utf-16le"}
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "encoding_upper",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Splitter.EncodingConfig = helper.EncodingConfig{Encoding: "UTF-16lE"}
					return newMockOperatorConfig(cfg)
				}(),
			},
		},
	}.Run(t)
}

func TestBuild(t *testing.T) {
	t.Parallel()

	basicConfig := func() *Config {
		cfg := NewConfig()
		cfg.Include = []string{"/var/log/testpath.*"}
		cfg.Exclude = []string{"/var/log/testpath.ex*"}
		cfg.PollInterval = 10 * time.Millisecond
		return cfg
	}

	cases := []struct {
		name             string
		modifyBaseConfig func(*Config)
		errorRequirement require.ErrorAssertionFunc
		validate         func(*testing.T, *Manager)
	}{
		{
			"Basic",
			func(f *Config) {},
			require.NoError,
			func(t *testing.T, f *Manager) {
				require.Equal(t, f.finder.Include, []string{"/var/log/testpath.*"})
				require.Equal(t, f.pollInterval, 10*time.Millisecond)
			},
		},
		{
			"BadIncludeGlob",
			func(f *Config) {
				f.Include = []string{"["}
			},
			require.Error,
			nil,
		},
		{
			"BadExcludeGlob",
			func(f *Config) {
				f.Include = []string{"["}
			},
			require.Error,
			nil,
		},
		{
			"MultilineConfiguredStartAndEndPatterns",
			func(f *Config) {
				f.Splitter = helper.NewSplitterConfig()
				f.Splitter.Multiline = helper.MultilineConfig{
					LineEndPattern:   "Exists",
					LineStartPattern: "Exists",
				}
			},
			require.Error,
			nil,
		},
		{
			"MultilineConfiguredStartPattern",
			func(f *Config) {
				f.Splitter = helper.NewSplitterConfig()
				f.Splitter.Multiline = helper.MultilineConfig{
					LineStartPattern: "START.*",
				}
			},
			require.NoError,
			func(t *testing.T, f *Manager) {},
		},
		{
			"MultilineConfiguredEndPattern",
			func(f *Config) {
				f.Splitter = helper.NewSplitterConfig()
				f.Splitter.Multiline = helper.MultilineConfig{
					LineEndPattern: "END.*",
				}
			},
			require.NoError,
			func(t *testing.T, f *Manager) {},
		},
		{
			"InvalidEncoding",
			func(f *Config) {
				f.Splitter.EncodingConfig = helper.EncodingConfig{Encoding: "UTF-3233"}
			},
			require.Error,
			nil,
		},
		{
			"LineStartAndEnd",
			func(f *Config) {
				f.Splitter = helper.NewSplitterConfig()
				f.Splitter.Multiline = helper.MultilineConfig{
					LineStartPattern: ".*",
					LineEndPattern:   ".*",
				}
			},
			require.Error,
			nil,
		},
		{
			"NoLineStartOrEnd",
			func(f *Config) {
				f.Splitter = helper.NewSplitterConfig()
				f.Splitter.Multiline = helper.MultilineConfig{}
			},
			require.NoError,
			func(t *testing.T, f *Manager) {},
		},
		{
			"InvalidLineStartRegex",
			func(f *Config) {
				f.Splitter = helper.NewSplitterConfig()
				f.Splitter.Multiline = helper.MultilineConfig{
					LineStartPattern: "(",
				}
			},
			require.Error,
			nil,
		},
		{
			"InvalidLineEndRegex",
			func(f *Config) {
				f.Splitter = helper.NewSplitterConfig()
				f.Splitter.Multiline = helper.MultilineConfig{
					LineEndPattern: "(",
				}
			},
			require.Error,
			nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			cfg := basicConfig()
			tc.modifyBaseConfig(cfg)

			nopEmit := func(_ context.Context, _ *FileAttributes, _ []byte) {}

			input, err := cfg.Build(testutil.Logger(t), nopEmit)
			tc.errorRequirement(t, err)
			if err != nil {
				return
			}

			tc.validate(t, input)
		})
	}
}

func NewTestConfig() *Config {
	cfg := NewConfig()
	cfg.Include = []string{"i1", "i2"}
	cfg.Exclude = []string{"e1", "e2"}
	cfg.Splitter = helper.NewSplitterConfig()
	cfg.Splitter.Multiline = helper.MultilineConfig{
		LineStartPattern: "start",
		LineEndPattern:   "end",
	}
	cfg.FingerprintSize = 1024
	cfg.Splitter.EncodingConfig = helper.EncodingConfig{Encoding: "utf16"}
	return cfg
}

func TestMapStructureDecodeConfigWithHook(t *testing.T) {
	expect := NewTestConfig()
	cfgMap := map[string]interface{}{
		"attributes":    map[string]interface{}{},
		"resource":      map[string]interface{}{},
		"include":       expect.Include,
		"exclude":       expect.Exclude,
		"poll_interval": 200 * time.Millisecond,
		"multiline": map[string]interface{}{
			"line_start_pattern": expect.Splitter.Multiline.LineStartPattern,
			"line_end_pattern":   expect.Splitter.Multiline.LineEndPattern,
		},
		"force_flush_period":   500 * time.Millisecond,
		"include_file_name":    true,
		"include_file_path":    false,
		"start_at":             "end",
		"fingerprint_size":     "1024",
		"max_log_size":         "1mib",
		"max_concurrent_files": 1024,
		"encoding":             "utf16",
	}

	var actual Config
	dc := &mapstructure.DecoderConfig{Result: &actual, DecodeHook: operatortest.JSONUnmarshalerHook()}
	ms, err := mapstructure.NewDecoder(dc)
	require.NoError(t, err)
	err = ms.Decode(cfgMap)
	require.NoError(t, err)
	require.Equal(t, expect, &actual)
}

func TestMapStructureDecodeConfig(t *testing.T) {
	expect := NewTestConfig()
	cfgMap := map[string]interface{}{
		"attributes":    map[string]interface{}{},
		"resource":      map[string]interface{}{},
		"include":       expect.Include,
		"exclude":       expect.Exclude,
		"poll_interval": 200 * time.Millisecond,
		"multiline": map[string]interface{}{
			"line_start_pattern": expect.Splitter.Multiline.LineStartPattern,
			"line_end_pattern":   expect.Splitter.Multiline.LineEndPattern,
		},
		"include_file_name":    true,
		"include_file_path":    false,
		"start_at":             "end",
		"fingerprint_size":     1024,
		"max_log_size":         1024 * 1024,
		"max_concurrent_files": 1024,
		"encoding":             "utf16",
		"force_flush_period":   500 * time.Millisecond,
	}

	var actual Config
	err := mapstructure.Decode(cfgMap, &actual)
	require.NoError(t, err)
	require.Equal(t, expect, &actual)
}
