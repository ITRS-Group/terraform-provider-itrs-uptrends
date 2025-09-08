package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
	"github.com/samber/lo"
)

// operatorResource implements the resource.Resource interface.
// It holds a reference to an IOperator client, which is configured by the provider.
type operatorResource struct {
	client interfaces.IOperatorFacade
}

// NewOperatorResource is a constructor function that Terraform will call
// when initializing this resource within your provider.
func NewOperatorResource(operator interfaces.IOperatorFacade) resource.Resource {
	return &operatorResource{
		client: operator,
	}
}

// operatorResourceModel maps Terraform schema attributes to typed fields
// which we’ll use to build requests and set state.

type operatorResourceModel struct {
	ID                     types.String `tfsdk:"id"`
	FullName               types.String `tfsdk:"full_name"`
	Email                  types.String `tfsdk:"email"`
	Password               types.String `tfsdk:"password_wo"`
	MobilePhone            types.String `tfsdk:"mobile_phone"`
	BackupEmail            types.String `tfsdk:"backup_email"`
	IsOnDuty               types.Bool   `tfsdk:"is_on_duty"`
	SmsProvider            types.String `tfsdk:"sms_provider"`
	DefaultDashboard       types.String `tfsdk:"default_dashboard"`
	OperatorRole           types.String `tfsdk:"operator_role"`
	IsAccountAdministrator types.Bool   `tfsdk:"is_account_administrator"`
}

// Metadata sets the resource type name used in Terraform configurations.
func (r *operatorResource) Metadata(
	_ context.Context,
	_ resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = "itrs-uptrends_operator"
}

// Schema defines the fields that can be set or read by this resource.
//
// Note the use of rschema.Schema from "resource/schema" (aliased as rschema).
func (r *operatorResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = rschema.Schema{
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Description: "The unique identifier of the operator.",
				// Marked as Computed because the server sets it upon creation.
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"full_name": rschema.StringAttribute{
				Description: "Full name of the operator.",
				Required:    true,
			},
			"email": rschema.StringAttribute{
				Description: "Primary email address for the operator.",
				Required:    true,
			},
			"password_wo": rschema.StringAttribute{
				Description: "Password for the operator. Write-only field, not stored in state. Required during creation, optional during updates.",
				Sensitive:   true,
				Optional:    true,
				WriteOnly:   true,
			},
			"mobile_phone": rschema.StringAttribute{
				Description: "Mobile phone number for the operator.",
				Required:    true,
			},
			"backup_email": rschema.StringAttribute{
				Description: "Backup email address for the operator.",
				Required:    true,
			},
			"is_on_duty": rschema.BoolAttribute{
				Description: "Whether the operator is currently on duty.",
				Required:    true,
			},
			"sms_provider": rschema.StringAttribute{
				Description: "SMS provider for the operator.",
				Optional:    true,
			},
			"default_dashboard": rschema.StringAttribute{
				Description: "Default dashboard for the operator.",
				Required:    true,
			},
			"operator_role": rschema.StringAttribute{
				Description: "Role assigned to the operator.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_account_administrator": rschema.BoolAttribute{
				Description: "Whether the operator is an account administrator.",
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Configure receives the parent provider’s configured data (which includes the client).
func (r *operatorResource) Configure(
	_ context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}
	// In your provider implementation, you’ll store a reference to client.IOperator in ProviderData.
	if p, ok := req.ProviderData.(providerData); ok {
		r.client = p.Client
	}
}

func (r *operatorResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var config operatorResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Ensure password is provided during creation
	if config.Password.IsNull() {
		resp.Diagnostics.AddError("Missing Password", "Password is required during operator creation.")
		return
	}

	createReq := models.OperatorExtendedRequest{
		FullName:         config.FullName.ValueString(),
		Email:            config.Email.ValueString(),
		Password:         config.Password.ValueString(), // Use the password
		MobilePhone:      config.MobilePhone.ValueString(),
		BackupEmail:      config.BackupEmail.ValueStringPointer(),
		IsOnDuty:         config.IsOnDuty.ValueBool(),
		SmsProvider:      config.SmsProvider.ValueString(),
		DefaultDashboard: config.DefaultDashboard.ValueString(),
		OperatorRole:     lo.Ternary(config.OperatorRole.IsNull(), "Unspecified", config.OperatorRole.ValueString()),
	}

	result, statusCode, err, responseBody := r.client.CreateOperator(createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating operator", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to create operator",
			fmt.Sprintf("HTTP status code: %d and response %v", statusCode, responseBody),
		)
		return
	}

	// Set final state after successful creation
	var state operatorResourceModel
	state.ID = types.StringValue(result.OperatorGuid)
	state.FullName = types.StringValue(result.FullName)
	state.Email = types.StringValue(result.Email)
	state.MobilePhone = types.StringValue(result.MobilePhone)
	state.BackupEmail = types.StringValue(result.BackupEmail)
	state.IsOnDuty = types.BoolValue(result.IsOnDuty)
	state.SmsProvider = types.StringValue(result.SmsProvider)
	state.DefaultDashboard = types.StringValue(result.DefaultDashboard)
	state.IsAccountAdministrator = types.BoolValue(false)
	state.OperatorRole = types.StringValue(lo.Ternary(result.OperatorRole == "", "Unspecified", result.OperatorRole))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	resp.Diagnostics.AddWarning(
		"When you create an operator, it has a default authorization of 'AccountAccess'. Import the following itrs-uptrends_operator_authorization resource to manage this.",
		fmt.Sprintf(`import {
  to = itrs-uptrends_operator_authorization.default_authorization_for%s
  id = "%s:AccountAccess"
  provider = itrs-uptrends.uptrendsauthenticated
}`, result.OperatorGuid, result.OperatorGuid),
	)
}

// Read refreshes the Terraform state by fetching the operator’s latest data.
func (r *operatorResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state operatorResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	operatorID := state.ID.ValueString()
	operator, statusCode, err, responseBody := r.client.GetOperator(operatorID)
	if err != nil {
		resp.Diagnostics.AddError("Error reading operator", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to read operator",
			fmt.Sprintf("HTTP status code: %d with response body %v", statusCode, responseBody),
		)
		return
	}

	state.FullName = types.StringValue(operator.FullName)
	state.Email = types.StringValue(operator.Email)
	state.MobilePhone = types.StringValue(operator.MobilePhone)
	state.BackupEmail = types.StringValue(operator.BackupEmail)
	state.IsOnDuty = types.BoolValue(operator.IsOnDuty)
	state.SmsProvider = types.StringValue(operator.SmsProvider)
	state.DefaultDashboard = types.StringValue(operator.DefaultDashboard)
	state.IsAccountAdministrator = types.BoolValue(operator.IsAccountAdministrator)

	// Ensure we default operator_role to "Unspecified" if empty
	role := operator.OperatorRole
	if role == "" {
		role = "Unspecified"
	}
	state.OperatorRole = types.StringValue(role)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update modifies the operator by calling UpdateOperator.
func (r *operatorResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var config operatorResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state operatorResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	operatorID := state.ID.ValueString()

	updateReq := models.OperatorExtendedUpdateRequest{
		FullName:         config.FullName.ValueString(),
		Email:            config.Email.ValueString(),
		MobilePhone:      config.MobilePhone.ValueString(),
		BackupEmail:      config.BackupEmail.ValueString(),
		IsOnDuty:         config.IsOnDuty.ValueBool(),
		SmsProvider:      config.SmsProvider.ValueString(),
		DefaultDashboard: config.DefaultDashboard.ValueString(),
		OperatorRole:     lo.Ternary(config.OperatorRole.IsNull(), "Unspecified", config.OperatorRole.ValueString()),

		// ... your other Specified flags ...
		// The next fields are set to true because they are required true by the Put call of API
		OutgoingPhoneNumberIdSpecified: true,
		CultureNameSpecified:           true,
		TimeZoneIdSpecified:            true,
		UseNumericSenderSpecified:      true,
		AllowNativeLoginSpecified:      true,
		AllowSingleSignonSpecified:     true,
	}

	// Handle the optional password field
	if !config.Password.IsNull() {
		password := config.Password.ValueString()
		updateReq.Password = &password // Set the password if provided
	} else {
		updateReq.Password = nil // Leave the password as nil if not provided
	}

	statusCode, msg, err := r.client.UpdateOperator(operatorID, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating operator", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to update operator",
			fmt.Sprintf("HTTP %d: %s", statusCode, msg),
		)
		return
	}

	// Re-read from the server to get the latest data
	operator, _, _, _ := r.client.GetOperator(operatorID)
	if operator != (models.OperatorExtendedGetResponse{}) {
		state.FullName = types.StringValue(operator.FullName)
		state.Email = types.StringValue(operator.Email)
		state.MobilePhone = types.StringValue(operator.MobilePhone)
		state.BackupEmail = types.StringValue(operator.BackupEmail)
		state.IsOnDuty = types.BoolValue(operator.IsOnDuty)
		state.SmsProvider = types.StringValue(operator.SmsProvider)
		state.DefaultDashboard = types.StringValue(operator.DefaultDashboard)
		state.OperatorRole = types.StringValue(lo.Ternary(operator.OperatorRole == "", "Unspecified", operator.OperatorRole))
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Delete removes the operator by calling DeleteOperator and then clears state.
func (r *operatorResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state operatorResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	operatorID := state.ID.ValueString()
	statusCode, msg, err := r.client.DeleteOperator(operatorID)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting operator", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to delete operator",
			fmt.Sprintf("HTTP %d: %s", statusCode, msg),
		)
		return
	}

	// Remove resource from Terraform state to finalize deletion
	resp.State.RemoveResource(ctx)
}

func (r *operatorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// providerData is an example struct representing data stored in the provider.
// Typically, your provider might store an IOperator client to share with resources.
type providerData struct {
	Client interfaces.IOperatorFacade
}
