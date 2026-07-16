package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// rumWebsiteResource implements the Terraform resource interface.
type rumWebsiteResource struct {
	client interfaces.IRumWebsite
}

var _ resource.ResourceWithValidateConfig = &rumWebsiteResource{}

// NewRumWebsiteResource creates a new instance of the operator group resource.
func NewRumWebsiteResource(client interfaces.IRumWebsite) resource.Resource {
	return &rumWebsiteResource{
		client: client,
	}
}

// Metadata returns the resource type name.
func (r *rumWebsiteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "itrs-uptrends_rum_website"
}

// Schema defines the schema for the resource.
func (r *rumWebsiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier of the rum website",
			},
			"description": rschema.StringAttribute{
				Required:    true,
				Description: "The description of the rum website",
			},
			"url": rschema.StringAttribute{
				Required:    true,
				Description: "The URL of the monitored website",
			},
			"is_spa": rschema.BoolAttribute{
				Optional:    true,
				Description: "Indicates whether the website is a Single Page Application (SPA)",
			},
			"include_url_fragment": rschema.BoolAttribute{
				Optional:    true,
				Description: "Specifies whether to include the URL fragment (hash) in monitoring",
			},
			"rum_script": rschema.StringAttribute{
				Computed:    true,
				Description: "The RUM script to be added to the website",
			},
		},
	}
}

// ValidateConfig enforces cross-field constraints.
func (r *rumWebsiteResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config rumWebsiteModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !config.IncludeUrlFragment.IsNull() && !config.IncludeUrlFragment.IsUnknown() && config.IncludeUrlFragment.ValueBool() {
		if config.IsSpa.IsNull() || config.IsSpa.IsUnknown() || !config.IsSpa.ValueBool() {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"When 'include_url_fragment' is true, 'is_spa' must be true.",
			)
		}
	}
}

// rumWebsiteModel maps resource schema data.
type rumWebsiteModel struct {
	Id                 types.String `tfsdk:"id"`
	Description        types.String `tfsdk:"description"`
	Url                types.String `tfsdk:"url"`
	IsSpa              types.Bool   `tfsdk:"is_spa"`
	IncludeUrlFragment types.Bool   `tfsdk:"include_url_fragment"`
	RumScript          types.String `tfsdk:"rum_script"`
}

// Create is called when the resource is created.
func (r *rumWebsiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan rumWebsiteModel

	// Retrieve the plan.
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the underlying CreateRumWebsite method.
	result, msg, err := r.client.CreateRumWebsite(&models.RumWebsite{
		Description:        plan.Description.ValueString(),
		Url:                plan.Url.ValueString(),
		IsSpa:              plan.IsSpa.ValueBool(),
		IncludeUrlFragment: plan.IncludeUrlFragment.ValueBool(),
		RumScript:          plan.RumScript.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating Rum Website", fmt.Sprintf("Error: %v. Message: %s", err, msg))
		return
	}

	// Update the plan with data from the response.
	plan.Id = types.StringValue(result.RumWebsiteGuid)
	plan.Description = types.StringValue(result.Description)
	plan.Url = types.StringValue(result.Url)
	plan.IsSpa = types.BoolValue(result.IsSpa)
	plan.IncludeUrlFragment = types.BoolValue(result.IncludeUrlFragment)
	plan.RumScript = types.StringValue(result.RumScript)
	// Set the state.
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Read is called to refresh the Terraform state.
func (r *rumWebsiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state rumWebsiteModel

	// Retrieve the state.
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the underlying GetRumWebsite method.
	result, msg, err := r.client.GetRumWebsite(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading Rum Website", fmt.Sprintf("Error: %v. Message: %s", err, msg))
		return
	}

	// Update the state with the refreshed data.
	state.Description = types.StringValue(result.Description)
	state.Url = types.StringValue(result.Url)
	state.IsSpa = types.BoolValue(result.IsSpa)
	state.IncludeUrlFragment = types.BoolValue(result.IncludeUrlFragment)
	state.RumScript = types.StringValue(result.RumScript)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the state.
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update is called when the resource is modified.
func (r *rumWebsiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state rumWebsiteModel
	// Retrieve the existing state to obtain the computed ID.
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config rumWebsiteModel
	// Retrieve the plan.
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	rumWebsiteID := state.Id.ValueString()

	// Use the merged plan.Id for the update API call.
	result, err := r.client.UpdateRumWebsite(&models.RumWebsite{
		RumWebsiteGuid:     rumWebsiteID,
		Description:        config.Description.ValueString(),
		Url:                config.Url.ValueString(),
		IsSpa:              config.IsSpa.ValueBool(),
		IncludeUrlFragment: config.IncludeUrlFragment.ValueBool(),
		RumScript:          config.RumScript.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error updating Rum Website", fmt.Sprintf("Error: %v", err))
		return
	}
	if result != "" {
		resp.Diagnostics.AddWarning("Rum Website Update Warning", fmt.Sprintf("Update response: %s", result))

	}
	state.Id = types.StringValue(rumWebsiteID)
	state.Description = types.StringValue(config.Description.ValueString())
	state.Url = types.StringValue(config.Url.ValueString())
	state.IsSpa = types.BoolValue(config.IsSpa.ValueBool())
	state.IncludeUrlFragment = types.BoolValue(config.IncludeUrlFragment.ValueBool())
	state.RumScript = types.StringValue(config.RumScript.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the state in a single, consistent operation.
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Delete is called when the resource is destroyed.
func (r *rumWebsiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state rumWebsiteModel

	// Retrieve the state.
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the underlying DeleteRumWebsite method.
	msg, err := r.client.DeleteRumWebsite(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting Rum Website", fmt.Sprintf("Error: %v. Message: %s", err, msg))
		return
	}

	// Remove the resource from state.
	resp.State.RemoveResource(ctx)
}

func (r *rumWebsiteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
