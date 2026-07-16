package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// operatorResource implements the resource.Resource interface.
// It holds a reference to an IOperator client, which is configured by the provider.
type operatorResource struct {
	client interfaces.IOperator
}

// NewOperatorResource is a constructor function that Terraform will call
// when initializing this resource within your provider.
func NewOperatorResource(operator interfaces.IOperator) resource.Resource {
	return &operatorResource{
		client: operator,
	}
}

// operatorResourceModel maps Terraform schema attributes to typed fields
// which we’ll use to build requests and set state.

type operatorResourceModel struct {
	ID                             types.String `tfsdk:"id"`
	FullName                       types.String `tfsdk:"full_name"`
	Email                          types.String `tfsdk:"email"`
	Password                       types.String `tfsdk:"password_wo"`
	PasswordVersion                types.Int64  `tfsdk:"password_wo_version"`
	MobilePhone                    types.String `tfsdk:"mobile_phone"`
	OutgoingPhoneNumberId          types.Int64  `tfsdk:"outgoing_phone_number_id"`
	OutgoingPhoneNumberIdSpecified types.Bool   `tfsdk:"outgoing_phone_number_id_specified"`
	IsAccountAdministrator         types.Bool   `tfsdk:"is_account_administrator"`
	BackupEmail                    types.String `tfsdk:"backup_email"`
	IsOnDuty                       types.Bool   `tfsdk:"is_on_duty"`
	CultureName                    types.String `tfsdk:"culture_name"`
	CultureNameSpecified           types.Bool   `tfsdk:"culture_name_specified"`
	TimeZoneId                     types.Int64  `tfsdk:"time_zone_id"`
	TimeZoneIdSpecified            types.Bool   `tfsdk:"time_zone_id_specified"`
	SmsProvider                    types.String `tfsdk:"sms_provider"`
	UseNumericSender               types.Bool   `tfsdk:"use_numeric_sender"`
	UseNumericSenderSpecified      types.Bool   `tfsdk:"use_numeric_sender_specified"`
	AllowNativeLogin               types.Bool   `tfsdk:"allow_native_login"`
	AllowNativeLoginSpecified      types.Bool   `tfsdk:"allow_native_login_specified"`
	AllowSingleSignon              types.Bool   `tfsdk:"allow_single_signon"`
	AllowSingleSignonSpecified     types.Bool   `tfsdk:"allow_single_signon_specified"`
	DefaultDashboard               types.String `tfsdk:"default_dashboard"`
	SetupMode                      types.String `tfsdk:"setup_mode"`
	OperatorRole                   types.String `tfsdk:"operator_role"`
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
				Computed:    true,
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
				Description: "Password for the operator. Write-only field, not stored in state. Required when setup_mode is 'Manual' (default) and allow_native_login is true (default). Must not be provided when setup_mode is 'Invitation' or allow_native_login is false.",
				Sensitive:   true,
				Optional:    true,
				WriteOnly:   true,
			},
			"password_wo_version": rschema.Int64Attribute{
				Description: "Version of the password for the operator.",
				Optional:    true,
			},
			"mobile_phone": rschema.StringAttribute{
				Description: "Mobile phone number for the operator. Empty string is accepted.",
				Required:    true,
			},
			"outgoing_phone_number_id": rschema.Int64Attribute{
				Description: "The outgoing phone number ID for the operator.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"outgoing_phone_number_id_specified": rschema.BoolAttribute{
				Description: "Whether the outgoing_phone_number_id value should be sent to the API.",
				Optional:    true,
				WriteOnly:   true,
			},
			"is_account_administrator": rschema.BoolAttribute{
				Description: "Whether the operator is an account administrator.",
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"backup_email": rschema.StringAttribute{
				Description: "Backup email address for the operator.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_on_duty": rschema.BoolAttribute{
				Description: "Whether the operator is currently on duty.",
				Required:    true,
			},
			"culture_name": rschema.StringAttribute{
				Description: "The culture name (locale) for the operator.",
				Optional:    true,
				WriteOnly:   true,
			},
			"culture_name_specified": rschema.BoolAttribute{
				Description: "Whether the culture_name value should be sent to the API.",
				Optional:    true,
				WriteOnly:   true,
			},
			"time_zone_id": rschema.Int64Attribute{
				Description: "The time zone ID for the operator.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"time_zone_id_specified": rschema.BoolAttribute{
				Description: "Whether the time_zone_id value should be sent to the API.",
				Optional:    true,
				WriteOnly:   true,
			},
			"sms_provider": rschema.StringAttribute{
				Description: "SMS provider for the operator.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"use_numeric_sender": rschema.BoolAttribute{
				Description: "Whether to use a numeric sender for SMS.",
				Optional:    true,
				WriteOnly:   true,
			},
			"use_numeric_sender_specified": rschema.BoolAttribute{
				Description: "Whether the use_numeric_sender value should be sent to the API.",
				Optional:    true,
				WriteOnly:   true,
			},
			"allow_native_login": rschema.BoolAttribute{
				Description: "Whether native login is allowed for the operator. Defaults to true. When false, password must not be provided and the account must have SSO enabled.",
				Optional:    true,
				WriteOnly:   true,
			},
			"allow_native_login_specified": rschema.BoolAttribute{
				Description: "Whether the allow_native_login value should be sent to the API.",
				Optional:    true,
				WriteOnly:   true,
			},
			"allow_single_signon": rschema.BoolAttribute{
				Description: "Whether single sign-on is allowed for the operator.",
				Optional:    true,
				WriteOnly:   true,
			},
			"allow_single_signon_specified": rschema.BoolAttribute{
				Description: "Whether the allow_single_signon value should be sent to the API.",
				Optional:    true,
				WriteOnly:   true,
			},
			"default_dashboard": rschema.StringAttribute{
				Description: "Default dashboard for the operator.",
				Required:    true,
			},
			"setup_mode": rschema.StringAttribute{
				Description: "The setup mode of the operator.",
				Optional:    true,
				WriteOnly:   true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Manual",
						"Invitation",
					),
				},
			},
			"operator_role": rschema.StringAttribute{
				Description: "Role assigned to the operator.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
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

	setupMode := "Manual"
	if !config.SetupMode.IsNull() && !config.SetupMode.IsUnknown() {
		setupMode = config.SetupMode.ValueString()
	}

	allowNativeLogin := true
	if !config.AllowNativeLogin.IsNull() && !config.AllowNativeLogin.IsUnknown() {
		allowNativeLogin = config.AllowNativeLogin.ValueBool()
	}

	passwordAvailable := setupMode == "Manual" && allowNativeLogin

	if passwordAvailable && config.Password.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Password",
			"Password is required when setup_mode is 'Manual' and allow_native_login is true (both are defaults).",
		)
		return
	}

	if !passwordAvailable && !config.Password.IsNull() {
		if setupMode == "Invitation" {
			resp.Diagnostics.AddError(
				"Password not allowed",
				"Password has to be empty when setup_mode is 'Invitation'.",
			)
		} else {
			resp.Diagnostics.AddError(
				"Password not allowed",
				"Password has to be empty when allow_native_login is false.",
			)
		}
		return
	}

	createReq := models.OperatorRequest{
		FullName:                       config.FullName.ValueString(),
		Email:                          config.Email.ValueString(),
		MobilePhone:                    config.MobilePhone.ValueString(),
		OutgoingPhoneNumberIdSpecified: config.OutgoingPhoneNumberIdSpecified.ValueBool(),
		BackupEmail:                    config.BackupEmail.ValueString(),
		IsOnDuty:                       config.IsOnDuty.ValueBool(),
		CultureNameSpecified:           config.CultureNameSpecified.ValueBool(),
		TimeZoneIdSpecified:            config.TimeZoneIdSpecified.ValueBool(),
		SmsProvider:                    config.SmsProvider.ValueString(),
		UseNumericSenderSpecified:      config.UseNumericSenderSpecified.ValueBool(),
		AllowNativeLoginSpecified:      config.AllowNativeLoginSpecified.ValueBool(),
		AllowSingleSignonSpecified:     config.AllowSingleSignonSpecified.ValueBool(),
		DefaultDashboard:               config.DefaultDashboard.ValueString(),
		SetupMode:                      config.SetupMode.ValueString(),
		OperatorRole:                   config.OperatorRole.ValueString(),
	}

	if !config.Password.IsNull() {
		password := config.Password.ValueString()
		createReq.Password = &password
	}
	if !config.OutgoingPhoneNumberId.IsNull() {
		v := int(config.OutgoingPhoneNumberId.ValueInt64())
		createReq.OutgoingPhoneNumberId = &v
	}
	if !config.CultureName.IsNull() {
		v := config.CultureName.ValueString()
		createReq.CultureName = &v
	}
	if !config.TimeZoneId.IsNull() {
		v := int(config.TimeZoneId.ValueInt64())
		createReq.TimeZoneId = &v
	}
	if !config.UseNumericSender.IsNull() {
		v := config.UseNumericSender.ValueBool()
		createReq.UseNumericSender = &v
	}
	if !config.AllowNativeLogin.IsNull() {
		v := config.AllowNativeLogin.ValueBool()
		createReq.AllowNativeLogin = &v
	}
	if !config.AllowSingleSignon.IsNull() {
		v := config.AllowSingleSignon.ValueBool()
		createReq.AllowSingleSignon = &v
	}

	result, statusCode, responseBody, err := r.client.CreateOperator(createReq)
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

	var state operatorResourceModel
	state.ID = types.StringValue(result.OperatorGuid)
	state.FullName = types.StringValue(result.FullName)
	state.Email = types.StringValue(result.Email)
	state.MobilePhone = types.StringValue(result.MobilePhone)
	state.OutgoingPhoneNumberId = types.Int64Value(int64(result.OutgoingPhoneNumberId))
	state.IsAccountAdministrator = types.BoolValue(result.IsAccountAdministrator)
	state.BackupEmail = types.StringValue(result.BackupEmail)
	state.IsOnDuty = types.BoolValue(result.IsOnDuty)
	state.TimeZoneId = types.Int64Value(int64(result.TimeZoneId))
	state.SmsProvider = types.StringValue(result.SmsProvider)
	state.DefaultDashboard = types.StringValue(result.DefaultDashboard)
	role := result.OperatorRole
	if role == "" {
		role = "Unspecified"
	}
	state.OperatorRole = types.StringValue(role)

	if !config.PasswordVersion.IsNull() {
		state.PasswordVersion = types.Int64Value(config.PasswordVersion.ValueInt64())
	}

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

	passwordVersion := types.Int64Null()
	if !state.PasswordVersion.IsNull() {
		passwordVersion = types.Int64Value(state.PasswordVersion.ValueInt64())
	}

	operator, statusCode, responseBody, err := r.client.GetOperator(operatorID)
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
	state.OutgoingPhoneNumberId = types.Int64Value(int64(operator.OutgoingPhoneNumberId))
	state.IsAccountAdministrator = types.BoolValue(operator.IsAccountAdministrator)
	state.BackupEmail = types.StringValue(operator.BackupEmail)
	state.IsOnDuty = types.BoolValue(operator.IsOnDuty)
	state.TimeZoneId = types.Int64Value(int64(operator.TimeZoneId))
	state.SmsProvider = types.StringValue(operator.SmsProvider)
	state.DefaultDashboard = types.StringValue(operator.DefaultDashboard)
	state.PasswordVersion = passwordVersion
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

	setupMode := "Manual"
	if !config.SetupMode.IsNull() && !config.SetupMode.IsUnknown() {
		setupMode = config.SetupMode.ValueString()
	}

	allowNativeLogin := true
	if !config.AllowNativeLogin.IsNull() && !config.AllowNativeLogin.IsUnknown() {
		allowNativeLogin = config.AllowNativeLogin.ValueBool()
	}

	passwordAvailable := setupMode == "Manual" && allowNativeLogin

	if !passwordAvailable && !config.Password.IsNull() {
		if setupMode == "Invitation" {
			resp.Diagnostics.AddError(
				"Password not allowed",
				"Password has to be empty when setup_mode is 'Invitation'.",
			)
		} else {
			resp.Diagnostics.AddError(
				"Password not allowed",
				"Password has to be empty when allow_native_login is false.",
			)
		}
		return
	}

	updateReq := models.OperatorRequest{
		FullName:                       config.FullName.ValueString(),
		Email:                          config.Email.ValueString(),
		MobilePhone:                    config.MobilePhone.ValueString(),
		OutgoingPhoneNumberIdSpecified: config.OutgoingPhoneNumberIdSpecified.ValueBool(),
		BackupEmail:                    config.BackupEmail.ValueString(),
		IsOnDuty:                       config.IsOnDuty.ValueBool(),
		CultureNameSpecified:           config.CultureNameSpecified.ValueBool(),
		TimeZoneIdSpecified:            config.TimeZoneIdSpecified.ValueBool(),
		SmsProvider:                    config.SmsProvider.ValueString(),
		UseNumericSenderSpecified:      config.UseNumericSenderSpecified.ValueBool(),
		AllowNativeLoginSpecified:      config.AllowNativeLoginSpecified.ValueBool(),
		AllowSingleSignonSpecified:     config.AllowSingleSignonSpecified.ValueBool(),
		DefaultDashboard:               config.DefaultDashboard.ValueString(),
		SetupMode:                      config.SetupMode.ValueString(),
		OperatorRole:                   config.OperatorRole.ValueString(),
	}

	if config.BackupEmail.IsNull() || config.BackupEmail.ValueString() == "" {
		updateReq.BackupEmail = ""
	}

	if !config.Password.IsNull() {
		password := config.Password.ValueString()
		updateReq.Password = &password
	}
	if !config.OutgoingPhoneNumberId.IsNull() {
		v := int(config.OutgoingPhoneNumberId.ValueInt64())
		updateReq.OutgoingPhoneNumberId = &v
	}
	if !config.CultureName.IsNull() {
		v := config.CultureName.ValueString()
		updateReq.CultureName = &v
	}
	if !config.TimeZoneId.IsNull() {
		v := int(config.TimeZoneId.ValueInt64())
		updateReq.TimeZoneId = &v
	}
	if !config.UseNumericSender.IsNull() {
		v := config.UseNumericSender.ValueBool()
		updateReq.UseNumericSender = &v
	}
	if !config.AllowNativeLogin.IsNull() {
		v := config.AllowNativeLogin.ValueBool()
		updateReq.AllowNativeLogin = &v
	}
	if !config.AllowSingleSignon.IsNull() {
		v := config.AllowSingleSignon.ValueBool()
		updateReq.AllowSingleSignon = &v
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

	operator, _, _, _ := r.client.GetOperator(operatorID)
	if operator != nil {
		state.FullName = types.StringValue(operator.FullName)
		state.Email = types.StringValue(operator.Email)
		state.MobilePhone = types.StringValue(operator.MobilePhone)
		state.OutgoingPhoneNumberId = types.Int64Value(int64(operator.OutgoingPhoneNumberId))
		state.IsAccountAdministrator = types.BoolValue(operator.IsAccountAdministrator)
		state.BackupEmail = types.StringValue(operator.BackupEmail)
		state.IsOnDuty = types.BoolValue(operator.IsOnDuty)
		state.TimeZoneId = types.Int64Value(int64(operator.TimeZoneId))
		state.SmsProvider = types.StringValue(operator.SmsProvider)
		state.DefaultDashboard = types.StringValue(operator.DefaultDashboard)
		state.OperatorRole = types.StringValue(operator.OperatorRole)
	}
	if !config.PasswordVersion.IsNull() {
		state.PasswordVersion = types.Int64Value(config.PasswordVersion.ValueInt64())
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
	Client interfaces.IOperator
}
