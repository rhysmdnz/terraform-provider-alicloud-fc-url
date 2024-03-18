// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

// func TestAccFcTriggerUrlDataSource(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { testAccPreCheck(t) },
// 		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Read testing
// 			{
// 				Config: testAccExampleDataSourceConfig,
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr("data.alicloud-fc-url_trigger_url.test", "id", "example-id"),
// 				),
// 			},
// 		},
// 	})
// }

// const testAccExampleDataSourceConfig = `
// provider "alicloud-fc-url" {
// 	region = "cn-shanghai"
//   }

// data "alicloud-fc-url_trigger_url" "test" {
// 	service_name  = "some-value"
// 	function_name = "some-value"
// 	trigger_name  = "some-value"
// }
// `
