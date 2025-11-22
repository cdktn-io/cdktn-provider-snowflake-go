// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebook


type NotebookFrom struct {
	// Identifier of the stage where the .ipynb file is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#stage Notebook#stage}
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// Location of the .ipynb file in the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#path Notebook#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

