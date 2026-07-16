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

var _ datasource.DataSource = &vaultSectionDataSource{}

func NewVaultSectionDataSource(client interfaces.IVaultSection) datasource.DataSource {
	return &vaultSectionDataSource{client: client}
}

type vaultSectionDataSource struct {
	client interfaces.IVaultSection
}

func (d *vaultSectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vault_section"
}

func (d *vaultSectionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Vault section GUID. Provide this or name.",
				Optional:    true,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Vault section name. Provide this or id.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

type vaultSectionDataSourceModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (d *vaultSectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The vault section client was not configured.")
		return
	}

	var data vaultSectionDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idProvided := !data.ID.IsNull() && data.ID.ValueString() != ""
	nameProvided := !data.Name.IsNull() && data.Name.ValueString() != ""

	switch {
	case idProvided && nameProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide only one of id or name.")
		return
	case !idProvided && !nameProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide either id or name.")
		return
	}

	var state vaultSectionResourceModel

	if idProvided {
		vs, err, msg := d.client.GetVaultSection(data.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading vault section", fmt.Sprintf("%v - %s", err, msg))
			return
		}
		state.Id = types.StringValue(vs.VaultSectionGuid)
		state.Name = types.StringValue(vs.Name)
	} else {
		sections, statusCode, responseBody, err := d.client.GetVaultSections()
		if err != nil {
			resp.Diagnostics.AddError("Error listing vault sections", err.Error())
			return
		}
		if statusCode >= 300 {
			resp.Diagnostics.AddError("Failed to list vault sections", fmt.Sprintf("HTTP %d: %s", statusCode, responseBody))
			return
		}
		name := data.Name.ValueString()
		found := false
		for _, s := range sections {
			if strings.EqualFold(s.Name, name) {
				if found {
					resp.Diagnostics.AddError("Vault section not unique", fmt.Sprintf("More than one vault section found with name %q", name))
					return
				}
				state.Id = types.StringValue(s.VaultSectionGuid)
				state.Name = types.StringValue(s.Name)
				found = true
			}
		}
		if !found {
			resp.Diagnostics.AddError("Vault section not found", fmt.Sprintf("No vault section found with name %q", name))
			return
		}
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
