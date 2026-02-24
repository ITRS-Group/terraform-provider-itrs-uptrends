package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &operatorDataSource{}
)

// NewOperatorDataSource is a helper function to simplify the provider implementation.
func NewOperatorDataSource(operator interfaces.IOperator) datasource.DataSource {
	return &operatorDataSource{
		client: operator,
	}
}

// operatorDataSource is the data source implementation.
type operatorDataSource struct {
	client interfaces.IOperator
}

// Metadata returns the data source type name.
func (d *operatorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_operator"
}

// Schema defines the schema for the data source.
func (d *operatorDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique identifier (GUID) of the operator. Provide this or full_name.",
				Optional:    true,
				Computed:    true,
			},
			"full_name": schema.StringAttribute{
				Description: "Full name of the operator. Provide this or id.",
				Optional:    true,
				Computed:    true,
			},
			"email": schema.StringAttribute{
				Description: "Primary email address for the operator.",
				Computed:    true,
			},
			"mobile_phone": schema.StringAttribute{
				Description: "Mobile phone number for the operator.",
				Computed:    true,
			},
			"backup_email": schema.StringAttribute{
				Description: "Backup email address for the operator.",
				Computed:    true,
			},
			"is_on_duty": schema.BoolAttribute{
				Description: "Whether the operator is currently on duty.",
				Computed:    true,
			},
			"sms_provider": schema.StringAttribute{
				Description: "SMS provider for the operator.",
				Computed:    true,
			},
			"default_dashboard": schema.StringAttribute{
				Description: "Default dashboard for the operator.",
				Computed:    true,
			},
			"operator_role": schema.StringAttribute{
				Description: "Role assigned to the operator.",
				Computed:    true,
			},
			"is_account_administrator": schema.BoolAttribute{
				Description: "Whether the operator is an account administrator.",
				Computed:    true,
			},
		},
	}
}

// operatorDataSourceModel maps the Terraform schema data.
type operatorDataSourceModel struct {
	ID                     types.String `tfsdk:"id"`
	FullName               types.String `tfsdk:"full_name"`
	Email                  types.String `tfsdk:"email"`
	MobilePhone            types.String `tfsdk:"mobile_phone"`
	BackupEmail            types.String `tfsdk:"backup_email"`
	IsOnDuty               types.Bool   `tfsdk:"is_on_duty"`
	SmsProvider            types.String `tfsdk:"sms_provider"`
	DefaultDashboard       types.String `tfsdk:"default_dashboard"`
	OperatorRole           types.String `tfsdk:"operator_role"`
	IsAccountAdministrator types.Bool   `tfsdk:"is_account_administrator"`
}

// Read refreshes the Terraform state with the latest data.
func (d *operatorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The operator client was not configured. This is an internal error in the provider.")
		return
	}

	var data operatorDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idProvided := !data.ID.IsNull() && data.ID.ValueString() != ""
	nameProvided := !data.FullName.IsNull() && data.FullName.ValueString() != ""

	switch {
	case idProvided && nameProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide only one of id or full_name to look up an operator.")
		return
	case !idProvided && !nameProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide either id or full_name to look up an operator.")
		return
	}

	var operatorID string
	var operatorName string

	if idProvided {
		operatorID = data.ID.ValueString()
	}
	if nameProvided {
		operatorName = data.FullName.ValueString()
	}

	var operator *models.OperatorResponse
	if operatorID != "" {
		found, statusCode, responseBody, err := d.client.GetOperator(operatorID)
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
		operator = found
	} else {
		operators, statusCode, responseBody, err := d.client.GetOperators()
		if err != nil {
			resp.Diagnostics.AddError("Error listing operators", err.Error())
			return
		}
		if statusCode >= 300 {
			resp.Diagnostics.AddError(
				"Failed to list operators",
				fmt.Sprintf("HTTP status code: %d with response body %v", statusCode, responseBody),
			)
			return
		}
		found := false
		for idx := range operators {
			if strings.EqualFold(operators[idx].FullName, operatorName) {
				if found {
					resp.Diagnostics.AddError("Operator not unique", fmt.Sprintf("More than one operator found with full_name %q", operatorName))
					return
				}
				operator = &operators[idx]
				found = true
			}
		}
		if !found {
			resp.Diagnostics.AddError("Operator not found", fmt.Sprintf("No operator found with full_name %q", operatorName))
			return
		}
	}

	// Map the response to state
	var state operatorDataSourceModel
	state.ID = types.StringValue(operator.OperatorGuid)
	state.FullName = types.StringValue(operator.FullName)
	state.Email = types.StringValue(operator.Email)
	state.MobilePhone = types.StringValue(operator.MobilePhone)
	state.BackupEmail = types.StringValue(operator.BackupEmail)
	state.IsOnDuty = types.BoolValue(operator.IsOnDuty)
	state.SmsProvider = types.StringValue(operator.SmsProvider)
	state.DefaultDashboard = types.StringValue(operator.DefaultDashboard)
	state.IsAccountAdministrator = types.BoolValue(operator.IsAccountAdministrator)

	role := operator.OperatorRole
	if strings.TrimSpace(role) == "" {
		role = "Unspecified"
	}
	state.OperatorRole = types.StringValue(role)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
