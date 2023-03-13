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
	_ datasource.DataSource              = &PieChartPanelCfgDataSource{}
	_ datasource.DataSourceWithConfigure = &PieChartPanelCfgDataSource{}
)

func NewPieChartPanelCfgDataSource() datasource.DataSource {
	return &PieChartPanelCfgDataSource{}
}

// PieChartPanelCfgDataSource defines the data source implementation.
type PieChartPanelCfgDataSource struct{}

// PieChartPanelCfgDataSourceModel describes the data source data model.
type PieChartPanelCfgDataSourceModel struct {
	PieChartType          types.String `tfsdk:"pie_chart_type" json:"PieChartType"`
	PieChartLabels        types.String `tfsdk:"pie_chart_labels" json:"PieChartLabels"`
	PieChartLegendValues  types.String `tfsdk:"pie_chart_legend_values" json:"PieChartLegendValues"`
	PieChartLegendOptions *struct {
		Values      types.List   `tfsdk:"values" json:"values"`
		DisplayMode types.String `tfsdk:"display_mode" json:"displayMode"`
		Placement   types.String `tfsdk:"placement" json:"placement"`
		ShowLegend  types.Bool   `tfsdk:"show_legend" json:"showLegend"`
		AsTable     types.Bool   `tfsdk:"as_table" json:"asTable"`
		IsVisible   types.Bool   `tfsdk:"is_visible" json:"isVisible"`
		SortBy      types.String `tfsdk:"sort_by" json:"sortBy"`
		SortDesc    types.Bool   `tfsdk:"sort_desc" json:"sortDesc"`
		Width       types.Number `tfsdk:"width" json:"width"`
		Calcs       types.List   `tfsdk:"calcs" json:"calcs"`
	} `tfsdk:"pie_chart_legend_options" json:"PieChartLegendOptions"`
	PanelOptions *struct {
		PieType       types.String `tfsdk:"pie_type" json:"pieType"`
		DisplayLabels types.List   `tfsdk:"display_labels" json:"displayLabels"`
		Tooltip       *struct {
			Mode types.String `tfsdk:"mode" json:"mode"`
			Sort types.String `tfsdk:"sort" json:"sort"`
		} `tfsdk:"tooltip" json:"tooltip"`
		ReduceOptions *struct {
			Values types.Bool   `tfsdk:"values" json:"values"`
			Limit  types.Number `tfsdk:"limit" json:"limit"`
			Calcs  types.List   `tfsdk:"calcs" json:"calcs"`
			Fields types.String `tfsdk:"fields" json:"fields"`
		} `tfsdk:"reduce_options" json:"reduceOptions"`
		Text *struct {
			TitleSize types.Number `tfsdk:"title_size" json:"titleSize"`
			ValueSize types.Number `tfsdk:"value_size" json:"valueSize"`
		} `tfsdk:"text" json:"text"`
		Legend *struct {
			Values      types.List   `tfsdk:"values" json:"values"`
			DisplayMode types.String `tfsdk:"display_mode" json:"displayMode"`
			Placement   types.String `tfsdk:"placement" json:"placement"`
			ShowLegend  types.Bool   `tfsdk:"show_legend" json:"showLegend"`
			AsTable     types.Bool   `tfsdk:"as_table" json:"asTable"`
			IsVisible   types.Bool   `tfsdk:"is_visible" json:"isVisible"`
			SortBy      types.String `tfsdk:"sort_by" json:"sortBy"`
			SortDesc    types.Bool   `tfsdk:"sort_desc" json:"sortDesc"`
			Width       types.Number `tfsdk:"width" json:"width"`
			Calcs       types.List   `tfsdk:"calcs" json:"calcs"`
		} `tfsdk:"legend" json:"legend"`
		Orientation types.String `tfsdk:"orientation" json:"orientation"`
	} `tfsdk:"panel_options" json:"PanelOptions"`
	PanelFieldConfig *struct {
		HideFrom *struct {
			Tooltip types.Bool `tfsdk:"tooltip" json:"tooltip"`
			Legend  types.Bool `tfsdk:"legend" json:"legend"`
			Viz     types.Bool `tfsdk:"viz" json:"viz"`
		} `tfsdk:"hide_from" json:"hideFrom"`
	} `tfsdk:"panel_field_config" json:"PanelFieldConfig"`
	ToJSON types.String `tfsdk:"to_json"`
}

func (d *PieChartPanelCfgDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_piechartpanelcfg"
}

func (d *PieChartPanelCfgDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",

		Attributes: map[string]schema.Attribute{
			"pie_chart_type": schema.StringAttribute{
				MarkdownDescription: `Select the pie chart display style.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},

			"pie_chart_labels": schema.StringAttribute{
				MarkdownDescription: `Select labels to display on the pie chart.
 - Name - The series or field name.
 - Percent - The percentage of the whole.
 - Value - The raw numerical value.`,
				Computed: false,
				Optional: false,
				Required: true,
			},

			"pie_chart_legend_values": schema.StringAttribute{
				MarkdownDescription: `Select values to display in the legend.
 - Percent: The percentage of the whole.
 - Value: The raw numerical value.`,
				Computed: false,
				Optional: false,
				Required: true,
			},

			"pie_chart_legend_options": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"values": schema.ListAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						ElementType:         types.StringType,
					},

					"display_mode": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"placement": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"show_legend": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"as_table": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},

					"is_visible": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},

					"sort_by": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},

					"sort_desc": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},

					"width": schema.NumberAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},

					"calcs": schema.ListAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						ElementType:         types.StringType,
					},
				},
			},

			"panel_options": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"pie_type": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"display_labels": schema.ListAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						ElementType:         types.StringType,
					},

					"tooltip": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},

							"sort": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
						},
					},

					"reduce_options": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"values": schema.BoolAttribute{
								MarkdownDescription: `If true show each row value`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"limit": schema.NumberAttribute{
								MarkdownDescription: `if showing all values limit`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"calcs": schema.ListAttribute{
								MarkdownDescription: `When !values, pick one value for the whole field`,
								Computed:            false,
								Optional:            false,
								Required:            true,
								ElementType:         types.StringType,
							},

							"fields": schema.StringAttribute{
								MarkdownDescription: `Which fields to show.  By default this is only numeric fields`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
						},
					},

					"text": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
						Attributes: map[string]schema.Attribute{
							"title_size": schema.NumberAttribute{
								MarkdownDescription: `Explicit title text size`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"value_size": schema.NumberAttribute{
								MarkdownDescription: `Explicit value text size`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
						},
					},

					"legend": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"values": schema.ListAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
								ElementType:         types.StringType,
							},

							"display_mode": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},

							"placement": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},

							"show_legend": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},

							"as_table": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"is_visible": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"sort_by": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"sort_desc": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"width": schema.NumberAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"calcs": schema.ListAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
								ElementType:         types.StringType,
							},
						},
					},

					"orientation": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
				},
			},

			"panel_field_config": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"hide_from": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
						Attributes: map[string]schema.Attribute{
							"tooltip": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},

							"legend": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},

							"viz": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
						},
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

func (d *PieChartPanelCfgDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *PieChartPanelCfgDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PieChartPanelCfgDataSourceModel

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
