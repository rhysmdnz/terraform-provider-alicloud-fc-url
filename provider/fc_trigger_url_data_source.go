// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	fc20230330 "github.com/alibabacloud-go/fc-20230330/v3/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &FcTriggerUrlDataSource{}

func NewFcTriggerUrlDataSource() datasource.DataSource {
	return &FcTriggerUrlDataSource{}
}

// FcTriggerUrlDataSource defines the data source implementation.
type FcTriggerUrlDataSource struct {
	client *fc20230330.Client
}

// FcTriggerUrlDataSourceModel describes the data source data model.
type FcTriggerUrlDataSourceModel struct {
	ServiceName          types.String `tfsdk:"service_name"`
	FunctionName         types.String `tfsdk:"function_name"`
	TriggerName          types.String `tfsdk:"trigger_name"`
	Id                   types.String `tfsdk:"id"`
	SourceArn            types.String `tfsdk:"source_arn"`
	Type                 types.String `tfsdk:"type"`
	InvocationRole       types.String `tfsdk:"invocation_role"`
	Config               types.String `tfsdk:"config"`
	URLInternet          types.String `tfsdk:"url_internet"`
	URLIntranet          types.String `tfsdk:"url_intranet"`
	CreationTime         types.String `tfsdk:"creation_time"`
	LastModificationTime types.String `tfsdk:"last_modification_time"`
}

func (d *FcTriggerUrlDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_trigger_url"
}

func (d *FcTriggerUrlDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example data source",

		Attributes: map[string]schema.Attribute{
			"service_name": schema.StringAttribute{
				MarkdownDescription: "Example configurable attribute",
				Required:            true,
			},
			"function_name": schema.StringAttribute{
				MarkdownDescription: "Example identifier",
				Required:            true,
			},
			"trigger_name": schema.StringAttribute{
				MarkdownDescription: "Example identifier",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"source_arn": schema.StringAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
			"invocation_role": schema.StringAttribute{
				Computed: true,
			},
			"config": schema.StringAttribute{
				Computed: true,
			},
			"url_internet": schema.StringAttribute{
				Computed: true,
			},
			"url_intranet": schema.StringAttribute{
				Computed: true,
			},
			"creation_time": schema.StringAttribute{
				Computed: true,
			},
			"last_modification_time": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *FcTriggerUrlDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*fc20230330.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *fc20230330.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *FcTriggerUrlDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data FcTriggerUrlDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := d.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }
	serviceName := data.ServiceName.ValueString()
	functionName := data.FunctionName.ValueString()
	triggerName := data.TriggerName.ValueString()

	newFunctionName := fmt.Sprintf("%s$%s", serviceName, functionName)

	tflog.Warn(ctx, "read a data source")

	// request := fc.NewGetTriggerInput(serviceName, functionName, triggerName)
	response, err := d.client.GetTrigger(&newFunctionName, &triggerName)
	if err != nil {
		resp.Diagnostics.AddError("It went wrong", err.Error())
		return
		// return WrapErrorf(err, DataDefaultErrorMsg, "alicloud_fc_triggers", "ListTriggers", FcGoSdk)
	}

	data.Id = types.StringPointerValue(response.Body.TriggerId)
	tflog.Warn(ctx, "read a data source")
	data.SourceArn = types.StringPointerValue(response.Body.SourceArn)
	tflog.Warn(ctx, "read a data source")
	data.Type = types.StringPointerValue(response.Body.TriggerType)
	tflog.Warn(ctx, "read a data source")
	data.URLInternet = types.StringPointerValue(response.Body.HttpTrigger.UrlInternet)
	tflog.Warn(ctx, "read a data source")
	data.URLIntranet = types.StringPointerValue(response.Body.HttpTrigger.UrlIntranet)
	tflog.Warn(ctx, "read a data source")
	data.InvocationRole = types.StringPointerValue(response.Body.InvocationRole)
	tflog.Warn(ctx, "read a data source")
	data.Config = types.StringPointerValue(response.Body.TriggerConfig)
	tflog.Warn(ctx, "read a data source")
	data.CreationTime = types.StringPointerValue(response.Body.CreatedTime)
	tflog.Warn(ctx, "read a data source")
	data.LastModificationTime = types.StringPointerValue(response.Body.LastModifiedTime)
	tflog.Warn(ctx, "read a data source")

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	// data.Id = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
