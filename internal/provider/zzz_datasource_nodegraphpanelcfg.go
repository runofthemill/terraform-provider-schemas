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
	_ datasource.DataSource              = &NodeGraphPanelCfgDataSource{}
	_ datasource.DataSourceWithConfigure = &NodeGraphPanelCfgDataSource{}
)

func NewNodeGraphPanelCfgDataSource() datasource.DataSource {
	return &NodeGraphPanelCfgDataSource{}
}

// NodeGraphPanelCfgDataSource defines the data source implementation.
type NodeGraphPanelCfgDataSource struct{}

// NodeGraphPanelCfgDataSourceModel describes the data source data model.
type NodeGraphPanelCfgDataSourceModel struct {
	ArcOption *struct {
		Field types.String `tfsdk:"field" json:"field"`
		Color types.String `tfsdk:"color" json:"color"`
	} `tfsdk:"arc_option" json:"ArcOption"`
	NodeOptions *struct {
		MainStatUnit      types.String `tfsdk:"main_stat_unit" json:"mainStatUnit"`
		SecondaryStatUnit types.String `tfsdk:"secondary_stat_unit" json:"secondaryStatUnit"`
		Arcs              []struct {
			Field types.String `tfsdk:"field" json:"field"`
			Color types.String `tfsdk:"color" json:"color"`
		} `tfsdk:"arcs" json:"arcs"`
	} `tfsdk:"node_options" json:"NodeOptions"`
	EdgeOptions *struct {
		MainStatUnit      types.String `tfsdk:"main_stat_unit" json:"mainStatUnit"`
		SecondaryStatUnit types.String `tfsdk:"secondary_stat_unit" json:"secondaryStatUnit"`
	} `tfsdk:"edge_options" json:"EdgeOptions"`
	PanelOptions *struct {
		Nodes *struct {
			MainStatUnit      types.String `tfsdk:"main_stat_unit" json:"mainStatUnit"`
			SecondaryStatUnit types.String `tfsdk:"secondary_stat_unit" json:"secondaryStatUnit"`
			Arcs              []struct {
				Field types.String `tfsdk:"field" json:"field"`
				Color types.String `tfsdk:"color" json:"color"`
			} `tfsdk:"arcs" json:"arcs"`
		} `tfsdk:"nodes" json:"nodes"`
		Edges *struct {
			MainStatUnit      types.String `tfsdk:"main_stat_unit" json:"mainStatUnit"`
			SecondaryStatUnit types.String `tfsdk:"secondary_stat_unit" json:"secondaryStatUnit"`
		} `tfsdk:"edges" json:"edges"`
	} `tfsdk:"panel_options" json:"PanelOptions"`
	ToJSON types.String `tfsdk:"to_json"`
}

func (d *NodeGraphPanelCfgDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nodegraphpanelcfg"
}

func (d *NodeGraphPanelCfgDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",

		Attributes: map[string]schema.Attribute{
			"arc_option": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"field": schema.StringAttribute{
						MarkdownDescription: `Field from which to get the value. Values should be less than 1, representing fraction of a circle.`,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},

					"color": schema.StringAttribute{
						MarkdownDescription: `The color of the arc.`,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
				},
			},

			"node_options": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"main_stat_unit": schema.StringAttribute{
						MarkdownDescription: `Unit for the main stat to override what ever is set in the data frame.`,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},

					"secondary_stat_unit": schema.StringAttribute{
						MarkdownDescription: `Unit for the secondary stat to override what ever is set in the data frame.`,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},

					"arcs": schema.ListNestedAttribute{
						MarkdownDescription: `Define which fields are shown as part of the node arc (colored circle around the node).`,
						Computed:            false,
						Optional:            true,
						Required:            false,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"field": schema.StringAttribute{
									MarkdownDescription: `Field from which to get the value. Values should be less than 1, representing fraction of a circle.`,
									Computed:            false,
									Optional:            true,
									Required:            false,
								},

								"color": schema.StringAttribute{
									MarkdownDescription: `The color of the arc.`,
									Computed:            false,
									Optional:            true,
									Required:            false,
								},
							},
						},
					},
				},
			},

			"edge_options": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"main_stat_unit": schema.StringAttribute{
						MarkdownDescription: `Unit for the main stat to override what ever is set in the data frame.`,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},

					"secondary_stat_unit": schema.StringAttribute{
						MarkdownDescription: `Unit for the secondary stat to override what ever is set in the data frame.`,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
				},
			},

			"panel_options": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"nodes": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
						Attributes: map[string]schema.Attribute{
							"main_stat_unit": schema.StringAttribute{
								MarkdownDescription: `Unit for the main stat to override what ever is set in the data frame.`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"secondary_stat_unit": schema.StringAttribute{
								MarkdownDescription: `Unit for the secondary stat to override what ever is set in the data frame.`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"arcs": schema.ListNestedAttribute{
								MarkdownDescription: `Define which fields are shown as part of the node arc (colored circle around the node).`,
								Computed:            false,
								Optional:            true,
								Required:            false,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"field": schema.StringAttribute{
											MarkdownDescription: `Field from which to get the value. Values should be less than 1, representing fraction of a circle.`,
											Computed:            false,
											Optional:            true,
											Required:            false,
										},

										"color": schema.StringAttribute{
											MarkdownDescription: `The color of the arc.`,
											Computed:            false,
											Optional:            true,
											Required:            false,
										},
									},
								},
							},
						},
					},

					"edges": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
						Attributes: map[string]schema.Attribute{
							"main_stat_unit": schema.StringAttribute{
								MarkdownDescription: `Unit for the main stat to override what ever is set in the data frame.`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},

							"secondary_stat_unit": schema.StringAttribute{
								MarkdownDescription: `Unit for the secondary stat to override what ever is set in the data frame.`,
								Computed:            false,
								Optional:            true,
								Required:            false,
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

func (d *NodeGraphPanelCfgDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *NodeGraphPanelCfgDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NodeGraphPanelCfgDataSourceModel

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
