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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &ElasticsearchDataQueryDataSource{}
	_ datasource.DataSourceWithConfigure = &ElasticsearchDataQueryDataSource{}
)

func NewElasticsearchDataQueryDataSource() datasource.DataSource {
	return &ElasticsearchDataQueryDataSource{}
}

// ElasticsearchDataQueryDataSource defines the data source implementation.
type ElasticsearchDataQueryDataSource struct{}

// ElasticsearchDataQueryDataSourceModel describes the data source data model.
type ElasticsearchDataQueryDataSourceModel struct {
	Alias     types.String `tfsdk:"alias" json:"alias"`
	Query     types.String `tfsdk:"query" json:"query"`
	TimeField types.String `tfsdk:"time_field" json:"timeField"`
	RefId     types.String `tfsdk:"ref_id" json:"refId"`
	Hide      types.Bool   `tfsdk:"hide" json:"hide"`
	Key       types.String `tfsdk:"key" json:"key"`
	QueryType types.String `tfsdk:"query_type" json:"queryType"`
	ToJSON    types.String `tfsdk:"to_json"`
}

func (d *ElasticsearchDataQueryDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_elasticsearchdataquery"
}

func (d *ElasticsearchDataQueryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",

		Attributes: map[string]schema.Attribute{
			"alias": schema.StringAttribute{
				MarkdownDescription: `Alias pattern`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},

			"query": schema.StringAttribute{
				MarkdownDescription: `Lucene query`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},

			"time_field": schema.StringAttribute{
				MarkdownDescription: `Name of time field`,
				Computed:            false,
				Optional:            true,
				Required:            false,
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

func (d *ElasticsearchDataQueryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *ElasticsearchDataQueryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ElasticsearchDataQueryDataSourceModel

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
