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

var _ datasource.DataSource = &checkpointRegionDataSource{}

// NewCheckpointRegionDataSource constructs the checkpoint region data source.
func NewCheckpointRegionDataSource(client interfaces.ICheckpoint) datasource.DataSource {
	return &checkpointRegionDataSource{client: client}
}

type checkpointRegionDataSource struct {
	client interfaces.ICheckpoint
}

func (d *checkpointRegionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_checkpoint_region"
}

func (d *checkpointRegionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Internal identifier for this data source instance.",
			},
			"regions": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Names of regions to select. If omitted, all regions are selected.",
			},
			"selected_regions_ids": schema.ListAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "IDs of regions matching the requested names, or all when none are specified.",
			},
			"regions_data": schema.ListNestedAttribute{
				Description: "Region objects returned by the API for the selected regions.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

type checkpointRegionDataSourceModel struct {
	ID                types.String `tfsdk:"id"`
	Regions           types.List   `tfsdk:"regions"`
	SelectedRegionIDs types.List   `tfsdk:"selected_regions_ids"`
	RegionsData       types.List   `tfsdk:"regions_data"`
}

type checkpointRegionModel struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (d *checkpointRegionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The checkpoint client was not configured.")
		return
	}

	var data checkpointRegionDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	regions, statusCode, responseBody, err := d.client.GetCheckpointRegions()
	if err != nil {
		resp.Diagnostics.AddError("Error listing checkpoint regions", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to list checkpoint regions",
			fmt.Sprintf("HTTP status code: %d with response body %v", statusCode, responseBody),
		)
		return
	}

	available := make([]string, 0, len(regions))
	allIDs := make([]int64, 0, len(regions))
	regionModels := make([]checkpointRegionModel, 0, len(regions))

	var requested []string
	if !data.Regions.IsNull() && !data.Regions.IsUnknown() {
		diag := data.Regions.ElementsAs(ctx, &requested, false)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	requestedSet := map[string]struct{}{}
	for _, n := range requested {
		requestedSet[strings.ToLower(n)] = struct{}{}
	}

	for _, r := range regions {
		available = append(available, r.Name)
		allIDs = append(allIDs, int64(r.Id))

		include := len(requestedSet) == 0
		if !include {
			_, include = requestedSet[strings.ToLower(r.Name)]
		}
		if !include {
			continue
		}

		regionModels = append(regionModels, checkpointRegionModel{
			ID:   types.Int64Value(int64(r.Id)),
			Name: types.StringValue(r.Name),
		})
	}

	selected := make([]int64, 0)
	if len(requested) == 0 {
		selected = allIDs
	} else {
		for _, name := range requested {
			found := false
			for _, r := range regions {
				if strings.EqualFold(r.Name, name) {
					selected = append(selected, int64(r.Id))
					found = true
				}
			}
			if !found {
				sort.Strings(available)
				resp.Diagnostics.AddError(
					fmt.Sprintf("Checkpoint region not found for name: %s", name),
					fmt.Sprintf("Checkpoint region %s does not exist. Available checkpoint regions are: %s", name, strings.Join(available, ", ")),
				)
				return
			}
		}
	}

	selectedVal, diag := types.ListValueFrom(ctx, types.Int64Type, selected)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	regionListVal, diag := types.ListValueFrom(ctx, checkpointRegionModelType(), regionModels)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	state := checkpointRegionDataSourceModel{
		ID:                types.StringValue("checkpoint_region_data_source"),
		Regions:           data.Regions,
		SelectedRegionIDs: selectedVal,
		RegionsData:       regionListVal,
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func checkpointRegionModelType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":   types.Int64Type,
			"name": types.StringType,
		},
	}
}
