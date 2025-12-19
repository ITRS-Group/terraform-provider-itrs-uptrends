package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	converters "github.com/itrs-group/terraform-provider-itrs-uptrends/converters/vault_item"
	tfsdkmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/provider/models"
)

var _ datasource.DataSource = &vaultItemDataSource{}

func NewVaultItemDataSource(client interfaces.IVaultItem) datasource.DataSource {
	return &vaultItemDataSource{client: client}
}

type vaultItemDataSource struct {
	client interfaces.IVaultItem
}

func (d *vaultItemDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vault_item"
}

func (d *vaultItemDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of the vault item.",
			},
			"name": dschema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The name of the vault item.",
			},
			"vault_section_id": dschema.StringAttribute{
				Computed: true,
			},
			"vault_item_type": dschema.StringAttribute{
				Computed: true,
			},
			"notes": dschema.StringAttribute{
				Computed: true,
			},
			"value_wo": dschema.StringAttribute{
				Computed: true,
			},
			"username": dschema.StringAttribute{
				Computed: true,
			},
			"password_wo": dschema.StringAttribute{
				Computed: true,
			},
			"certificate_archive": dschema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]dschema.Attribute{
					"issuer": dschema.StringAttribute{
						Computed: true,
					},
					"not_before": dschema.StringAttribute{
						Computed: true,
					},
					"not_after": dschema.StringAttribute{
						Computed: true,
					},
					"password_wo": dschema.StringAttribute{
						Computed: true,
					},
					"archive_data_wo": dschema.StringAttribute{
						Computed: true,
					},
				},
			},
			"file": dschema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]dschema.Attribute{
					"data": dschema.StringAttribute{
						Computed: true,
					},
					"name": dschema.StringAttribute{
						Computed: true,
					},
				},
			},
			"one_time_password": dschema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]dschema.Attribute{
					"secret_wo": dschema.StringAttribute{
						Computed: true,
					},
					"digits": dschema.Int64Attribute{
						Computed: true,
					},
					"period": dschema.Int64Attribute{
						Computed: true,
					},
					"hash_algorithm": dschema.StringAttribute{
						Computed: true,
					},
					"secret_encoding_method_wo": dschema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *vaultItemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The vault item client was not configured.")
		return
	}

	var data tfsdkmodels.VaultItemResourceModel
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

	var state tfsdkmodels.VaultItemResourceModel

	if idProvided {
		item, err, msg := d.client.GetVaultItem(data.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading vault item", fmt.Sprintf("%v - %s", err, msg))
			return
		}
		state = converters.UpdateStateConversion(item)
	} else {
		items, statusCode, responseBody, err := d.client.GetVaultItems()
		if err != nil {
			resp.Diagnostics.AddError("Error listing vault items", err.Error())
			return
		}
		if statusCode >= 300 {
			resp.Diagnostics.AddError("Failed to list vault items", fmt.Sprintf("HTTP %d: %s", statusCode, responseBody))
			return
		}
		name := data.Name.ValueString()
		found := false
		for idx := range items {
			if strings.EqualFold(items[idx].Name, name) {
				if found {
					resp.Diagnostics.AddError("Vault item not unique", fmt.Sprintf("More than one vault item found with name %q", name))
					return
				}
				state = converters.UpdateStateConversion(&items[idx])
				found = true
			}
		}
		if !found {
			resp.Diagnostics.AddError("Vault item not found", fmt.Sprintf("No vault item found with name %q", name))
			return
		}
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
