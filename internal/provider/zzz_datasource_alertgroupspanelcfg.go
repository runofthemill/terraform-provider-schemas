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
	_ datasource.DataSource              = &AlertGroupsPanelCfgDataSource{}
	_ datasource.DataSourceWithConfigure = &AlertGroupsPanelCfgDataSource{}
)

func NewAlertGroupsPanelCfgDataSource() datasource.DataSource {
	return &AlertGroupsPanelCfgDataSource{}
}

// AlertGroupsPanelCfgDataSource defines the data source implementation.
type AlertGroupsPanelCfgDataSource struct{}

// AlertGroupsPanelCfgDataSourceModel describes the data source data model.
type AlertGroupsPanelCfgDataSourceModel struct {
	PanelOptions *struct {
		Labels       types.String `tfsdk:"labels" json:"labels"`
		Alertmanager types.String `tfsdk:"alertmanager" json:"alertmanager"`
		ExpandAll    types.Bool   `tfsdk:"expand_all" json:"expandAll"`
	} `tfsdk:"panel_options" json:"PanelOptions"`
	ToJSON types.String `tfsdk:"to_json"`
}

func (d *AlertGroupsPanelCfgDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_alertgroupspanelcfg"
}

func (d *AlertGroupsPanelCfgDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",

		Attributes: map[string]schema.Attribute{
			"panel_options": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"labels": schema.StringAttribute{
						MarkdownDescription: `Comma-separated list of values used to filter alert results`,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"alertmanager": schema.StringAttribute{
						MarkdownDescription: `Name of the alertmanager used as a source for alerts`,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"expand_all": schema.BoolAttribute{
						MarkdownDescription: `Expand all alert groups by default`,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
				},
			},

			"to_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "This datasource rendered as JSON",
			},
		},
	}
}

func (d *AlertGroupsPanelCfgDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *AlertGroupsPanelCfgDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AlertGroupsPanelCfgDataSourceModel

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
