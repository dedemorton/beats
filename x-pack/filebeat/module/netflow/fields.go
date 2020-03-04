// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package netflow

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "netflow", asset.ModuleFieldsPri, AssetNetflow); err != nil {
		panic(err)
	}
}

// AssetNetflow returns asset data.
// This is the base64 encoded gzipped contents of module/netflow.
func AssetNetflow() string {
	return "eJw8jjFOw0AQRfs9xbtAcoAtqFCkFKAUINGazBivWGas3Ymt3B4Z4XS/eO//f+Bb7xnTGKuvhx+XW9UEUaJq5lXjVH1NINqvrcxR3DJPCeDlD2b0RtOrlqXY124wmHC+nM4fbMUb4E06vmjj/fly5G1SHnMgrh3zYBBhLFql86l3N2GdhiAm3V9SbL4Fc/OliPZj4l/ICcz3/BsAAP//pS1J+Q=="
}
