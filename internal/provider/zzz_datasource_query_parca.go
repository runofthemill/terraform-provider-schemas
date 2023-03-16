// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     terraform
// Using jennies:
//     TerraformDataSourceJenny
//     ComposableLatestMajorsOrXJenny
//
// Run 'go generate ./' from repository root to regenerate.

package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure that the imports are used to avoid compiler errors.
var _ attr.Value
var _ diag.Diagnostic

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &QueryParcaDataSource{}
	_ datasource.DataSourceWithConfigure = &QueryParcaDataSource{}
)

func NewQueryParcaDataSource() datasource.DataSource {
	return &QueryParcaDataSource{}
}

// QueryParcaDataSource defines the data source implementation.
type QueryParcaDataSource struct{}

type QueryParcaDataSourceModel struct {
	ToJSON        types.String `tfsdk:"to_json"`
	LabelSelector types.String `tfsdk:"label_selector"`
	ProfileTypeId types.String `tfsdk:"profile_type_id"`
	RefId         types.String `tfsdk:"ref_id"`
	Hide          types.Bool   `tfsdk:"hide"`
	Key           types.String `tfsdk:"key"`
	QueryType     types.String `tfsdk:"query_type"`
}

func (m QueryParcaDataSourceModel) MarshalJSON() ([]byte, error) {
	type jsonQueryParcaDataSourceModel struct {
		LabelSelector string  `json:"labelSelector"`
		ProfileTypeId string  `json:"profileTypeId"`
		RefId         string  `json:"refId"`
		Hide          *bool   `json:"hide,omitempty"`
		Key           *string `json:"key,omitempty"`
		QueryType     *string `json:"queryType,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_labelselector := m.LabelSelector.ValueString()
	attr_profiletypeid := m.ProfileTypeId.ValueString()
	attr_refid := m.RefId.ValueString()
	attr_hide := m.Hide.ValueBool()
	attr_key := m.Key.ValueString()
	attr_querytype := m.QueryType.ValueString()

	model := &jsonQueryParcaDataSourceModel{
		LabelSelector: attr_labelselector,
		ProfileTypeId: attr_profiletypeid,
		RefId:         attr_refid,
		Hide:          &attr_hide,
		Key:           &attr_key,
		QueryType:     &attr_querytype,
	}
	return json.Marshal(model)
}

func (m QueryParcaDataSourceModel) ApplyDefaults() QueryParcaDataSourceModel {
	if m.LabelSelector.IsNull() {
		m.LabelSelector = types.StringValue(`{}`)
	}
	return m
}

func (d *QueryParcaDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_query_parca"
}

func (d *QueryParcaDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",
		Attributes: map[string]schema.Attribute{
			"label_selector": schema.StringAttribute{
				MarkdownDescription: `Specifies the query label selectors. Defaults to "{}".`,
				Computed:            true,
				Optional:            true,
				Required:            false,
			},
			"profile_type_id": schema.StringAttribute{
				MarkdownDescription: `Specifies the type of profile to query.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"ref_id": schema.StringAttribute{
				MarkdownDescription: `A - Z`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"hide": schema.BoolAttribute{
				MarkdownDescription: `true if query is disabled (ie should not be returned to the dashboard)`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"key": schema.StringAttribute{
				MarkdownDescription: `Unique, guid like, string used in explore mode`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"query_type": schema.StringAttribute{
				MarkdownDescription: `Specify the query flavor
TODO make this required and give it a default`,
				Computed: false,
				Optional: true,
				Required: false,
			},

			"to_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "This datasource rendered as JSON",
			},
		},
	}
}

func (d *QueryParcaDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *QueryParcaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data QueryParcaDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

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
