// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	fc20230330 "github.com/alibabacloud-go/fc-20230330/v3/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	credentials "github.com/aliyun/credentials-go/credentials"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &AlicloudFcUrlProvider{}
var _ provider.ProviderWithFunctions = &AlicloudFcUrlProvider{}

// AlicloudFcUrlProvider defines the provider implementation.
type AlicloudFcUrlProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// AlicloudFcUrlProviderModel describes the provider data model.
type AlicloudFcUrlProviderModel struct {
	// AccessKey             types.String `tfsdk:"access_key"`
	// SecretKey             types.String `tfsdk:"secret_key"`
	// SecurityToken         types.String `tfsdk:"security_token"`
	// EcsRoleName           types.String `tfsdk:"ecs_role_name"`
	Region types.String `tfsdk:"region"`
	// AccountId             types.String `tfsdk:"account_id"`
	// SharedCredentialsFile types.String `tfsdk:"shared_credentials_file"`
	// Profile               types.String `tfsdk:"profile"`
	// ClientReadTimeout     types.String `tfsdk:"client_read_timeout"`
	// ClientConnectTimeout  types.String `tfsdk:"client_connect_timeout"`
}

func (p *AlicloudFcUrlProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "alicloud-fc-url"
	resp.Version = p.version
}

func (p *AlicloudFcUrlProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// "access_key": schema.StringAttribute{
			// 	MarkdownDescription: "Example provider attribute",
			// 	Optional:            true,
			// },
			// "secret_key": schema.StringAttribute{
			// 	MarkdownDescription: "Example provider attribute",
			// 	Optional:            true,
			// },
			// "security_token": schema.StringAttribute{
			// 	MarkdownDescription: "Example provider attribute",
			// 	Optional:            true,
			// },
			// "ecs_role_name": schema.StringAttribute{
			// 	MarkdownDescription: "Example provider attribute",
			// 	Optional:            true,
			// },
			"region": schema.StringAttribute{
				MarkdownDescription: "Example provider attribute",
				Required:            true,
			},
			// "account_id": schema.StringAttribute{
			// 	MarkdownDescription: "Example provider attribute",
			// 	Optional:            true,
			// },
			// "shared_credentials_file": schema.StringAttribute{
			// 	MarkdownDescription: "Example provider attribute",
			// 	Optional:            true,
			// },
			// "profile": schema.StringAttribute{
			// 	MarkdownDescription: "Example provider attribute",
			// 	Optional:            true,
			// },
			// "client_read_timeout": schema.StringAttribute{
			// 	MarkdownDescription: "Example provider attribute",
			// 	Optional:            true,
			// },
			// "client_connect_timeout": schema.StringAttribute{
			// 	MarkdownDescription: "Example provider attribute",
			// 	Optional:            true,
			// },
		},
	}
}

func (p *AlicloudFcUrlProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data AlicloudFcUrlProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	credentialClient, err := credentials.NewCredential(nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("unable to initialize the FC client: %#v", err.Error()), "hi")
		return
	}

	config := &openapi.Config{
		Credential: credentialClient,
		RegionId:   data.Region.ValueStringPointer(),
	}

	sts, err := sts20150401.NewClient(config)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("unable to initialize the FC client: %#v", err.Error()), "hi")
		return
	}

	response, err := sts.GetCallerIdentity()
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("unable to initialize the FC client: %#v", err.Error()), "hi")
		return
	}

	accountId := response.Body.AccountId

	endpoint := fmt.Sprintf("%s.%s.fc.aliyuncs.com", *accountId, data.Region.ValueString())

	config.Endpoint = tea.String(endpoint)

	client, err := fc20230330.NewClient(config)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("unable to initialize the FC client: %#v", err.Error()), "hi")
		return
	}

	resp.DataSourceData = client
}

func (p *AlicloudFcUrlProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *AlicloudFcUrlProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewFcTriggerUrlDataSource,
	}
}

func (p *AlicloudFcUrlProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AlicloudFcUrlProvider{
			version: version,
		}
	}
}
