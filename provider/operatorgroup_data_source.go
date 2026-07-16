package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
)

var _ datasource.DataSource = &operatorGroupDataSource{}

func NewOperatorGroupDataSource(client interfaces.IOperatorGroup) datasource.DataSource {
	return &operatorGroupDataSource{client: client}
}

type operatorGroupDataSource struct {
	client interfaces.IOperatorGroup
}

func (d *operatorGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_operatorgroup"
}

func (d *operatorGroupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Operator group GUID. Provide this or description.",
				Optional:    true,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the operator group. Provide this or id.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

type operatorGroupDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
}

func (d *operatorGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The operator group client was not configured.")
		return
	}

	var data operatorGroupDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idProvided := !data.ID.IsNull() && data.ID.ValueString() != ""
	descProvided := !data.Description.IsNull() && data.Description.ValueString() != ""

	switch {
	case idProvided && descProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide only one of id or description.")
		return
	case !idProvided && !descProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide either id or description.")
		return
	}

	var state operatorGroupModel

	if idProvided {
		result, err, msg := d.client.GetOperatorGroup(data.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading operator group", fmt.Sprintf("%v - %s", err, msg))
			return
		}
		state.Id = types.StringValue(result.OperatorGroupGuid)
		state.Description = types.StringValue(result.Description)
	} else {
		// find by description
		groups, statusCode, responseBody, err := d.client.GetOperatorGroups()
		if err != nil {
			resp.Diagnostics.AddError("Error listing operator groups", err.Error())
			return
		}
		if statusCode >= 300 {
			resp.Diagnostics.AddError("Failed to list operator groups", fmt.Sprintf("HTTP %d: %s", statusCode, responseBody))
			return
		}
		desc := data.Description.ValueString()
		found := false
		for _, g := range groups {
			if strings.EqualFold(g.Description, desc) {
				if found {
					resp.Diagnostics.AddError("Operator group not unique", fmt.Sprintf("More than one operator group found with description %q", desc))
					return
				}
				state.Id = types.StringValue(g.OperatorGroupGuid)
				state.Description = types.StringValue(g.Description)
				found = true
			}
		}
		if !found {
			resp.Diagnostics.AddError("Operator group not found", fmt.Sprintf("No operator group found with description %q", desc))
			return
		}
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
