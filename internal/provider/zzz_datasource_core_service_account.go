// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     terraform
// Using jennies:
//     TerraformDataSourceJenny
//     LatestJenny
//
// Run 'go generate ./' from repository root to regenerate.

package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &CoreServiceAccountDataSource{}
	_ datasource.DataSourceWithConfigure = &CoreServiceAccountDataSource{}
)

func NewCoreServiceAccountDataSource() datasource.DataSource {
	return &CoreServiceAccountDataSource{}
}

// CoreServiceAccountDataSource defines the data source implementation.
type CoreServiceAccountDataSource struct{}

type CoreServiceAccountDataSourceModel_AccessControl struct {
}

func (m CoreServiceAccountDataSourceModel_AccessControl) MarshalJSON() ([]byte, error) {
	type jsonCoreServiceAccountDataSourceModel_AccessControl struct {
	}

	model := &jsonCoreServiceAccountDataSourceModel_AccessControl{}
	return json.Marshal(model)
}

type CoreServiceAccountDataSourceModel struct {
	ToJSON        types.String                                     `tfsdk:"to_json"`
	Id            types.Int64                                      `tfsdk:"id"`
	OrgId         types.Int64                                      `tfsdk:"org_id"`
	Name          types.String                                     `tfsdk:"name"`
	Login         types.String                                     `tfsdk:"login"`
	IsDisabled    types.Bool                                       `tfsdk:"is_disabled"`
	Role          types.String                                     `tfsdk:"role"`
	Tokens        types.Int64                                      `tfsdk:"tokens"`
	AvatarUrl     types.String                                     `tfsdk:"avatar_url"`
	AccessControl *CoreServiceAccountDataSourceModel_AccessControl `tfsdk:"access_control"`
	Teams         types.List                                       `tfsdk:"teams"`
	Created       types.Int64                                      `tfsdk:"created"`
	Updated       types.Int64                                      `tfsdk:"updated"`
}

func (m CoreServiceAccountDataSourceModel) MarshalJSON() ([]byte, error) {
	type jsonCoreServiceAccountDataSourceModel struct {
		Id            int64       `json:"id"`
		OrgId         int64       `json:"orgId"`
		Name          string      `json:"name"`
		Login         string      `json:"login"`
		IsDisabled    bool        `json:"isDisabled"`
		Role          string      `json:"role"`
		Tokens        int64       `json:"tokens"`
		AvatarUrl     string      `json:"avatarUrl"`
		AccessControl interface{} `json:"accessControl,omitempty"`
		Teams         []string    `json:"teams,omitempty"`
		Created       *int64      `json:"created,omitempty"`
		Updated       *int64      `json:"updated,omitempty"`
	}
	attr_id := m.Id.ValueInt64()
	attr_orgid := m.OrgId.ValueInt64()
	attr_name := m.Name.ValueString()
	attr_login := m.Login.ValueString()
	attr_isdisabled := m.IsDisabled.ValueBool()
	attr_role := m.Role.ValueString()
	attr_tokens := m.Tokens.ValueInt64()
	attr_avatarurl := m.AvatarUrl.ValueString()
	var attr_accesscontrol interface{}
	if m.AccessControl != nil {
		attr_accesscontrol = m.AccessControl
	}
	attr_teams := []string{}
	for _, v := range m.Teams.Elements() {
		attr_teams = append(attr_teams, v.(types.String).ValueString())
	}
	attr_created := m.Created.ValueInt64()
	attr_updated := m.Updated.ValueInt64()

	model := &jsonCoreServiceAccountDataSourceModel{
		Id:            attr_id,
		OrgId:         attr_orgid,
		Name:          attr_name,
		Login:         attr_login,
		IsDisabled:    attr_isdisabled,
		Role:          attr_role,
		Tokens:        attr_tokens,
		AvatarUrl:     attr_avatarurl,
		AccessControl: attr_accesscontrol,
		Teams:         attr_teams,
		Created:       &attr_created,
		Updated:       &attr_updated,
	}
	return json.Marshal(model)
}

func (d *CoreServiceAccountDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_core_service_account"
}

func (d *CoreServiceAccountDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: `ID is the unique identifier of the service account in the database.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"org_id": schema.Int64Attribute{
				MarkdownDescription: `OrgId is the ID of an organisation the service account belongs to.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the service account.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"login": schema.StringAttribute{
				MarkdownDescription: `Login of the service account.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"is_disabled": schema.BoolAttribute{
				MarkdownDescription: `IsDisabled indicates if the service account is disabled.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"role": schema.StringAttribute{
				MarkdownDescription: `Role is the Grafana organization role of the service account which can be 'Viewer', 'Editor', 'Admin'.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"tokens": schema.Int64Attribute{
				MarkdownDescription: `Tokens is the number of active tokens for the service account.
Tokens are used to authenticate the service account against Grafana.`,
				Computed: false,
				Optional: false,
				Required: true,
			},
			"avatar_url": schema.StringAttribute{
				MarkdownDescription: `AvatarUrl is the service account's avatar URL. It allows the frontend to display a picture in front
of the service account.`,
				Computed: false,
				Optional: false,
				Required: true,
			},
			"access_control": schema.SingleNestedAttribute{
				MarkdownDescription: `AccessControl metadata associated with a given resource.`,
				Computed:            true,
				Optional:            true,
				Required:            false,
			},
			"teams": schema.ListAttribute{
				MarkdownDescription: `Teams is a list of teams the service account belongs to.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
				ElementType:         types.StringType,
			},
			"created": schema.Int64Attribute{
				MarkdownDescription: `Created indicates when the service account was created.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"updated": schema.Int64Attribute{
				MarkdownDescription: `Updated indicates when the service account was updated.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},

			"to_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "This datasource rendered as JSON",
			},
		},
	}
}

func (d *CoreServiceAccountDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *CoreServiceAccountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CoreServiceAccountDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	d.applyDefaults(&data)
	JSONConfig, err := json.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("JSON marshalling error", err.Error())
		return
	}

	// Not sure about that
	data.ToJSON = types.StringValue(string(JSONConfig))

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *CoreServiceAccountDataSource) applyDefaults(data *CoreServiceAccountDataSourceModel) {
	if data.AccessControl == nil {
		data.AccessControl = &CoreServiceAccountDataSourceModel_AccessControl{}
	}
}
