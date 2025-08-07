package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/constants"
	converters "github.com/itrs-group/terraform-provider-itrs-uptrends/converters/vault_item"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/helpers"
	tfsdkmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/provider/models"
)

// vaultItemResource implements the Terraform resource
type vaultItemResource struct {
	client interfaces.IVaultItem
}

// NewVaultItemResource returns a new Vault Item resource
func NewVaultItemResource(client interfaces.IVaultItem) resource.Resource {
	return &vaultItemResource{
		client: client,
	}
}

// Metadata sets the resource type name
func (r *vaultItemResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "vault_item"
}

// Schema returns the Terraform schema for this resource
func (r *vaultItemResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": rschema.StringAttribute{
				Required: true,
			},
			"vault_section_id": rschema.StringAttribute{
				Required: true,
			},
			"vault_item_type": rschema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"CredentialSet",
						"Certificate",
						"CertificateArchive",
						"File",
						"OneTimePassword",
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"notes": rschema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"value_wo": rschema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
			},
			"username": rschema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"password_wo": rschema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
			},
			"certificate_archive": rschema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]rschema.Attribute{
					"issuer": rschema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"not_before": rschema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"not_after": rschema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"password_wo": rschema.StringAttribute{
						Optional:    true,
						Sensitive:   true,
						WriteOnly:   true,
						Description: `The password for the certificate archive.`,
					},
					"archive_data_wo": rschema.StringAttribute{
						Optional:    true,
						Sensitive:   true,
						WriteOnly:   true,
						Description: `The base64 encoded certificate archive data.`,
					},
				},
			},
			"file": rschema.SingleNestedAttribute{
				Optional: true,
				Default:  nil,
				Attributes: map[string]rschema.Attribute{
					"data": rschema.StringAttribute{
						Required: true,
					},
					"name": rschema.StringAttribute{
						Required: true,
					},
				},
			},
			"one_time_password": rschema.SingleNestedAttribute{
				Optional: true,
				Default:  nil,
				Attributes: map[string]rschema.Attribute{
					"secret_wo": rschema.StringAttribute{
						Optional:  true,
						WriteOnly: true,
						Sensitive: true,
					},
					"digits": rschema.Int64Attribute{
						Required: true,
						Validators: []validator.Int64{
							int64validator.OneOf(int64(6), int64(7), int64(8)),
						},
					},
					"period": rschema.Int64Attribute{
						Required: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
					"hash_algorithm": rschema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"SHA256",
								"SHA512",
								"SHA1",
							),
						},
					},
					"secret_encoding_method_wo": rschema.StringAttribute{
						Optional:  true,
						WriteOnly: true, // this is not stored in the API
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Base32",
								"Hex",
							),
						},
					},
				},
			},
		},
	}
}

func (r *vaultItemResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config tfsdkmodels.VaultItemResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	vaultItemType := config.VaultItemType.ValueString()
	var err error

	// Check if all required attributes are provided
	required := helpers.GetRequiredAttributes(vaultItemType, constants.VaultItemResourceAttributes)
	err = helpers.ValidateRequiredAttributes("vault_item", vaultItemType, config, required)
	if err != nil {
		resp.Diagnostics.AddError("Invalid configuration", err.Error())
	}

	// Check if all attributes are allowed
	allowed := helpers.GetAllowedAttributes(vaultItemType, constants.VaultItemResourceAttributes)
	err = helpers.ValidateAllowedAttributes("vault_item", vaultItemType, config, allowed)
	if err != nil {
		resp.Diagnostics.AddError("Invalid configuration", err.Error())
	}
}

func (r *vaultItemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config tfsdkmodels.VaultItemResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate the configuration specific to the create operation
	vaultItemType := config.VaultItemType.ValueString()
	required := helpers.GetRequiredAttributes(vaultItemType, constants.VaultItemResourceAttributesCreate)
	err := helpers.ValidateRequiredAttributes("vault_item", vaultItemType, config, required)
	if err != nil {
		resp.Diagnostics.AddError("Invalid configuration", err.Error())
	}

	payload := converters.PayloadConversion(config)

	result, statusCode, err, responseBody := r.client.CreateVaultItem(payload)

	if err != nil {
		resp.Diagnostics.AddError("Error creating vault item", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to create vault item",
			fmt.Sprintf("HTTP status code: %d and response %v", statusCode, responseBody),
		)
		return
	}

	var state = converters.UpdateStateConversion(&result)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Read implements the Terraform read operation
func (r *vaultItemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state tfsdkmodels.VaultItemResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Fetch items from the API
	getVaultItem, err, _ := r.client.GetVaultItem(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading vault item", err.Error())
		return
	}

	state = converters.UpdateStateConversion(getVaultItem)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update implements the Terraform update operation
func (r *vaultItemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config tfsdkmodels.VaultItemResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state tfsdkmodels.VaultItemResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	vaultItemID := state.ID.ValueString()

	payload := converters.PayloadConversion(config)

	statusCode, msg, err := r.client.UpdateVaultItem(vaultItemID, payload)
	if err != nil {
		resp.Diagnostics.AddError("Error updating vault item", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to update vault item",
			fmt.Sprintf("HTTP %d: %s", statusCode, msg),
		)
		return
	}

	// Re-read from the server to get the latest data
	getVaultItem, _, _ := r.client.GetVaultItem(vaultItemID)

	state = converters.UpdateStateConversion(getVaultItem)

	// Update state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Delete implements the Terraform delete operation
func (r *vaultItemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state tfsdkmodels.VaultItemResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the API to delete
	vaultItemID := state.ID.ValueString()

	statusCode, msg, err := r.client.DeleteVaultItem(vaultItemID)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting vault item", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to delete vault item",
			fmt.Sprintf("HTTP %d: %s", statusCode, msg),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *vaultItemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
