package provider

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
)

var _ datasource.DataSource = &checkpointDataSource{}

// NewCheckpointDataSource constructs the checkpoint data source.
func NewCheckpointDataSource(client interfaces.ICheckpoint) datasource.DataSource {
	return &checkpointDataSource{client: client}
}

type checkpointDataSource struct {
	client interfaces.ICheckpoint
}

func (d *checkpointDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_checkpoint"
}

func (d *checkpointDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Internal identifier for this data source instance.",
			},
			"checkpoints": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Names of checkpoints to select. If omitted, all checkpoints are selected.",
			},
			"selected_checkpoints_ids": schema.ListAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "IDs of checkpoints matching the requested names, or all when none are specified.",
			},
			"checkpoints_data": schema.ListNestedAttribute{
				Description: "Checkpoint objects returned by the API for the selected checkpoints.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed: true,
						},
						"attributes": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"checkpoint_name": schema.StringAttribute{Computed: true},
								"code":            schema.StringAttribute{Computed: true},
								"ipv4_addresses": schema.ListAttribute{
									ElementType: types.StringType,
									Computed:    true,
								},
								"ipv6_addresses": schema.ListAttribute{
									ElementType: types.StringType,
									Computed:    true,
								},
								"is_primary_checkpoint": schema.BoolAttribute{Computed: true},
								"supports_ipv6":         schema.BoolAttribute{Computed: true},
								"has_high_availability": schema.BoolAttribute{Computed: true},
							},
						},
						"links": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"self": schema.StringAttribute{Computed: true},
							},
						},
					},
				},
			},
		},
	}
}

type checkpointAttributesModel struct {
	CheckpointName      types.String `tfsdk:"checkpoint_name"`
	Code                types.String `tfsdk:"code"`
	Ipv4Addresses       types.List   `tfsdk:"ipv4_addresses"`
	Ipv6Addresses       types.List   `tfsdk:"ipv6_addresses"`
	IsPrimaryCheckpoint types.Bool   `tfsdk:"is_primary_checkpoint"`
	SupportsIpv6        types.Bool   `tfsdk:"supports_ipv6"`
	HasHighAvailability types.Bool   `tfsdk:"has_high_availability"`
}

type checkpointLinksModel struct {
	Self types.String `tfsdk:"self"`
}

type checkpointDataModel struct {
	ID         types.Int64               `tfsdk:"id"`
	Type       types.String              `tfsdk:"type"`
	Attributes checkpointAttributesModel `tfsdk:"attributes"`
	Links      checkpointLinksModel      `tfsdk:"links"`
}

type checkpointDataSourceModel struct {
	ID                    types.String `tfsdk:"id"`
	Checkpoints           types.List   `tfsdk:"checkpoints"`
	SelectedCheckpointIDs types.List   `tfsdk:"selected_checkpoints_ids"`
	CheckpointObjects     types.List   `tfsdk:"checkpoints_data"`
}

func (d *checkpointDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The checkpoint client was not configured.")
		return
	}

	var data checkpointDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	checkpointResp, statusCode, responseBody, err := d.client.GetCheckpoints()
	if err != nil {
		resp.Diagnostics.AddError("Error listing checkpoints", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to list checkpoints",
			fmt.Sprintf("HTTP status code: %d with response body %v", statusCode, responseBody),
		)
		return
	}

	availableNames := make([]string, 0, len(checkpointResp.Data))
	allIDs := make([]int64, 0, len(checkpointResp.Data))
	checkpointModels := make([]checkpointDataModel, 0, len(checkpointResp.Data))

	// Map of requested checkpoint names for quick lookup (lowercased).
	var requestedNames []string
	if !data.Checkpoints.IsNull() && !data.Checkpoints.IsUnknown() {
		diag := data.Checkpoints.ElementsAs(ctx, &requestedNames, false)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	requestedSet := map[string]struct{}{}
	for _, n := range requestedNames {
		requestedSet[strings.ToLower(n)] = struct{}{}
	}

	for _, cp := range checkpointResp.Data {
		availableNames = append(availableNames, cp.Attributes.CheckpointName)
		allIDs = append(allIDs, int64(cp.Id))

		include := len(requestedSet) == 0
		if !include {
			_, include = requestedSet[strings.ToLower(cp.Attributes.CheckpointName)]
		}
		if !include {
			continue
		}

		ipv4, diag := types.ListValueFrom(ctx, types.StringType, cp.Attributes.Ipv4Addresses)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
		ipv6, diag := types.ListValueFrom(ctx, types.StringType, cp.Attributes.IpV6Addresses)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}

		linkValue := checkpointLinksModel{Self: types.StringNull()}
		if cp.Links != nil {
			linkValue.Self = types.StringValue(cp.Links.Self)
		}

		checkpointModels = append(checkpointModels, checkpointDataModel{
			ID:   types.Int64Value(int64(cp.Id)),
			Type: types.StringValue(cp.Type),
			Attributes: checkpointAttributesModel{
				CheckpointName:      types.StringValue(cp.Attributes.CheckpointName),
				Code:                types.StringValue(cp.Attributes.Code),
				Ipv4Addresses:       ipv4,
				Ipv6Addresses:       ipv6,
				IsPrimaryCheckpoint: types.BoolValue(cp.Attributes.IsPrimaryCheckpoint),
				SupportsIpv6:        types.BoolValue(cp.Attributes.SupportsIpv6),
				HasHighAvailability: types.BoolValue(cp.Attributes.HasHighAvailability),
			},
			Links: linkValue,
		})
	}

	selectedIDs := make([]int64, 0)
	if len(requestedNames) == 0 {
		selectedIDs = allIDs
	} else {
		for _, name := range requestedNames {
			found := false
			for _, cp := range checkpointResp.Data {
				if strings.EqualFold(cp.Attributes.CheckpointName, name) {
					selectedIDs = append(selectedIDs, int64(cp.Id))
					found = true
				}
			}
			if !found {
				sort.Strings(availableNames)
				resp.Diagnostics.AddError(
					fmt.Sprintf("Checkpoint not found for name: %s", name),
					fmt.Sprintf("Checkpoint %s does not exist. Available checkpoints are: %s", name, strings.Join(availableNames, ", ")),
				)
				return
			}
		}
	}

	selectedIDsVal, diag := types.ListValueFrom(ctx, types.Int64Type, selectedIDs)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	checkpointListVal, diag := types.ListValueFrom(ctx, checkpointDataModelType(), checkpointModels)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	state := checkpointDataSourceModel{
		ID:                    types.StringValue("checkpoint_data_source"),
		Checkpoints:           data.Checkpoints,
		SelectedCheckpointIDs: selectedIDsVal,
		CheckpointObjects:     checkpointListVal,
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func checkpointDataModelType() attr.Type {
	attrType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"checkpoint_name":       types.StringType,
			"code":                  types.StringType,
			"ipv4_addresses":        types.ListType{ElemType: types.StringType},
			"ipv6_addresses":        types.ListType{ElemType: types.StringType},
			"is_primary_checkpoint": types.BoolType,
			"supports_ipv6":         types.BoolType,
			"has_high_availability": types.BoolType,
		},
	}
	linkType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"self": types.StringType,
		},
	}
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.Int64Type,
			"type":       types.StringType,
			"attributes": attrType,
			"links":      linkType,
		},
	}
}
