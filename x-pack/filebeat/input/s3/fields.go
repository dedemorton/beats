// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package s3

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "s3", asset.ModuleFieldsPri, AssetS3); err != nil {
		panic(err)
	}
}

// AssetS3 returns asset data.
// This is the base64 encoded gzipped contents of input/s3.
func AssetS3() string {
	return "eJyskz9v3DAMxXd9iofs8eLNQ4Gia9Ah186BzqJs1bJpUHQMf/tCipP6kH8IrhoMSBT5I9+TbzHQ1iByZwANGqnBTeTuxgCOUith1sBTg28GAO64gw8UXYIXHpFqhGletDKAUCSbqMGZ1Brs95qSd4vJjtSAvU+k5QjQbaaMnrr94A1gXr96gg+R9mxoTxCaWZQcYpgISa1ogi19/KPlpGq22l8AB9pWFvc5k89/qNXf93dgX6CliTL32oe2L2eFv9oEIesqY3ZB7ZoOgto1vSPo9zVdJWiqL2brhJf548lOdS5ss3Zh8iyjzTeqPX5kHDnnpR1IH/LmJfa2oB+g8/ppR3rW81TvdaG9zZ+Q8kuEkEqgR3JFk+pVM0/GPAy0/ddensp+0suzwW3kxR0sLvt3TP6RY1fZPAs/Bkfy9Yd8nLH0+FILngWn+vKXEeoCT1/H3Je8UuSVu5X5GwAA//9i6E0j"
}
