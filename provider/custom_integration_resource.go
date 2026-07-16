package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ resource.Resource = &customIntegrationResource{}

type customIntegrationResource struct {
	client interfaces.ICustomIntegration
}

func NewCustomIntegrationResource(client interfaces.ICustomIntegration) resource.Resource {
	return &customIntegrationResource{client: client}
}

// ---------------------------------------------------------------------------
// Terraform models
// ---------------------------------------------------------------------------

type customIntegrationModel struct {
	ID                   types.String            `tfsdk:"id"`
	Name                 types.String            `tfsdk:"name"`
	IsActive             types.Bool              `tfsdk:"is_active"`
	Notes                types.String            `tfsdk:"notes"`
	HttpStepDefinitions  []httpStepDefinitionTF  `tfsdk:"http_step_definitions"`
	IntegrationVariables []integrationVariableTF `tfsdk:"integration_variables"`
}

type httpStepDefinitionTF struct {
	HttpStepDefinitionUsageGuid types.String            `tfsdk:"http_step_definition_usage_guid"`
	Steps                       []httpStepTF            `tfsdk:"steps"`
	UserDefinedFunctions        []userDefinedFunctionTF `tfsdk:"user_defined_functions"`
	Usages                      types.List              `tfsdk:"usages"`
}

type httpStepTF struct {
	Url                        types.String          `tfsdk:"url"`
	Method                     types.String          `tfsdk:"method"`
	Body                       types.String          `tfsdk:"body"`
	BodyType                   types.String          `tfsdk:"body_type"`
	VaultFileId                types.String          `tfsdk:"vault_file_id"`
	MultiPartForm              []multiPartFormItemTF `tfsdk:"multi_part_form"`
	RequestHeaders             []keyValueTF          `tfsdk:"request_headers"`
	Variables                  []httpStepVariableTF  `tfsdk:"variables"`
	Assertions                 []httpStepAssertionTF `tfsdk:"assertions"`
	Name                       types.String          `tfsdk:"name"`
	UseFixedClientCertificate  types.Bool            `tfsdk:"use_fixed_client_certificate"`
	Authentication             *httpStepAuthTF       `tfsdk:"authentication"`
	IgnoreCertificateErrors    types.Bool            `tfsdk:"ignore_certificate_errors"`
	ClientCertificateVaultItem types.String          `tfsdk:"client_certificate_vault_item"`
	Delay                      types.Int64           `tfsdk:"delay"`
	StepType                   types.String          `tfsdk:"step_type"`
	RetryUntilSuccessful       types.Bool            `tfsdk:"retry_until_successful"`
	MaxAttempts                types.Int64           `tfsdk:"max_attempts"`
	RetryWaitMilliseconds      types.Int64           `tfsdk:"retry_wait_milliseconds"`
	PreRequestScript           types.String          `tfsdk:"pre_request_script"`
	PostResponseScript         types.String          `tfsdk:"post_response_script"`
	CalculatedContentType      types.String          `tfsdk:"calculated_content_type"`
	AllowedTlsVersions         types.List            `tfsdk:"allowed_tls_versions"`
	MsaStepHttpVersion         types.String          `tfsdk:"msa_step_http_version"`
}

type multiPartFormItemTF struct {
	Type  types.String `tfsdk:"type"`
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type keyValueTF struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type httpStepVariableTF struct {
	Source    types.String `tfsdk:"source"`
	Property  types.String `tfsdk:"property"`
	Name      types.String `tfsdk:"name"`
	Arguments types.List   `tfsdk:"arguments"`
}

type httpStepAssertionTF struct {
	Source      types.String `tfsdk:"source"`
	Property    types.String `tfsdk:"property"`
	Comparison  types.String `tfsdk:"comparison"`
	TargetValue types.String `tfsdk:"target_value"`
}

type httpStepAuthTF struct {
	Id                 types.String `tfsdk:"id"`
	AuthenticationType types.String `tfsdk:"authentication_type"`
	UserName           types.String `tfsdk:"user_name"`
	PasswordWO         types.String `tfsdk:"password_wo"`
	PasswordWOVersion  types.Int64  `tfsdk:"password_wo_version"`
}

type userDefinedFunctionTF struct {
	Name          types.String `tfsdk:"name"`
	Type          types.String `tfsdk:"type"`
	Mappings      []keyValueTF `tfsdk:"mappings"`
	Regex         types.String `tfsdk:"regex"`
	JwtSigningKey types.String `tfsdk:"jwt_signing_key"`
	JwtAlgorithm  types.String `tfsdk:"jwt_algorithm"`
	HashKey       types.String `tfsdk:"hash_key"`
	HashAlgorithm types.String `tfsdk:"hash_algorithm"`
}

type integrationVariableTF struct {
	Name                        types.String `tfsdk:"name"`
	Value                       types.String `tfsdk:"value"`
	IsValueSetInEscalationLevel types.Bool   `tfsdk:"is_value_set_in_escalation_level"`
}

func (r *customIntegrationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "itrs-uptrends_custom_integration"
}

func (r *customIntegrationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Description: "Manages a custom integration in the Uptrends monitoring platform.",
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier (GUID) of the custom integration.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": rschema.StringAttribute{
				Required:    true,
				Description: "The name of the custom integration.",
			},
			"is_active": rschema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Whether the custom integration is active. Defaults to `true`. The API requires a non-null value, so a value is always sent.",
			},
			"notes": rschema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "Notes for the custom integration. The API requires a non-null value, so an empty string is sent when unset.",
			},
			"integration_variables": rschema.ListNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Variables that can be referenced by the integration steps. Always sent to the API (as an empty array when unset).",
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: rschema.NestedAttributeObject{
					Attributes: map[string]rschema.Attribute{
						"name": rschema.StringAttribute{
							Required:    true,
							Description: "The name of the variable.",
						},
						"value": rschema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The value of the variable.",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"is_value_set_in_escalation_level": rschema.BoolAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Whether the value is set in the escalation level.",
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"http_step_definitions": rschema.ListNestedAttribute{
				Required:    true,
				Description: "The HTTP step definitions of the custom integration. At least one definition is required by the API.",
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
				NestedObject: rschema.NestedAttributeObject{
					Attributes: map[string]rschema.Attribute{
						"http_step_definition_usage_guid": rschema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The usage GUID of the HTTP step definition.",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"usages": rschema.ListAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The usages of the HTTP step definition (e.g. `Alert`).",
							ElementType: types.StringType,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
						},
						"user_defined_functions": rschema.ListNestedAttribute{
							Optional:    true,
							Computed:    true,
							Description: "User defined functions usable by the steps.",
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							NestedObject: rschema.NestedAttributeObject{
								Attributes: map[string]rschema.Attribute{
									"name": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The name of the function.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"type": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The type of the function (e.g. `Mapping`).",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"mappings": rschema.ListNestedAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Key/value mappings for the function.",
										PlanModifiers: []planmodifier.List{
											listplanmodifier.UseStateForUnknown(),
										},
										NestedObject: rschema.NestedAttributeObject{
											Attributes: map[string]rschema.Attribute{
												"key": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The mapping key.",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
												"value": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The mapping value.",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
											},
										},
									},
									"regex": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The regular expression of the function.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"jwt_signing_key": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The JWT signing key.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"jwt_algorithm": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The JWT algorithm (e.g. `HS256`).",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"hash_key": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The hash key.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"hash_algorithm": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The hash algorithm (e.g. `MD5`).",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
						},
						"steps": rschema.ListNestedAttribute{
							Required:    true,
							Description: "The HTTP steps to execute. At least one step is required for a usable definition.",
							Validators: []validator.List{
								listvalidator.SizeAtLeast(1),
							},
							NestedObject: rschema.NestedAttributeObject{
								Attributes: map[string]rschema.Attribute{
									"url": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The URL of the request.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"method": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The HTTP method (e.g. `GET`, `POST`).",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"body": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The request body.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"body_type": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The body type (e.g. `Raw`).",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"vault_file_id": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The vault file id.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"multi_part_form": rschema.ListNestedAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Multipart form entries.",
										PlanModifiers: []planmodifier.List{
											listplanmodifier.UseStateForUnknown(),
										},
										NestedObject: rschema.NestedAttributeObject{
											Attributes: map[string]rschema.Attribute{
												"type": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The type of the entry (e.g. `Text`).",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
												"key": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The key of the entry.",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
												"value": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The value of the entry.",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
											},
										},
									},
									"request_headers": rschema.ListNestedAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Request headers.",
										PlanModifiers: []planmodifier.List{
											listplanmodifier.UseStateForUnknown(),
										},
										NestedObject: rschema.NestedAttributeObject{
											Attributes: map[string]rschema.Attribute{
												"key": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The header name.",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
												"value": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The header value.",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
											},
										},
									},
									"variables": rschema.ListNestedAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Variables captured from the response.",
										PlanModifiers: []planmodifier.List{
											listplanmodifier.UseStateForUnknown(),
										},
										NestedObject: rschema.NestedAttributeObject{
											Attributes: map[string]rschema.Attribute{
												"source": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The source of the variable (e.g. `None`).",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
												"property": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The property to capture.",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
												"name": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The name of the variable.",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
												"arguments": rschema.ListAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The arguments of the variable.",
													ElementType: types.StringType,
													PlanModifiers: []planmodifier.List{
														listplanmodifier.UseStateForUnknown(),
													},
												},
											},
										},
									},
									"assertions": rschema.ListNestedAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Assertions to validate the response.",
										PlanModifiers: []planmodifier.List{
											listplanmodifier.UseStateForUnknown(),
										},
										NestedObject: rschema.NestedAttributeObject{
											Attributes: map[string]rschema.Attribute{
												"source": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The source of the assertion (e.g. `None`).",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
												"property": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The property to assert.",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
												"comparison": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The comparison (e.g. `Equal`).",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
												"target_value": rschema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "The target value of the assertion.",
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
												},
											},
										},
									},
									"name": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The name of the step.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"use_fixed_client_certificate": rschema.BoolAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Whether to use a fixed client certificate.",
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.UseStateForUnknown(),
										},
									},
									"authentication": rschema.SingleNestedAttribute{
										Optional:    true,
										Description: "Authentication settings for the step.",
										Attributes: map[string]rschema.Attribute{
											"id": rschema.StringAttribute{
												Optional:    true,
												Computed:    true,
												Description: "The id of the authentication.",
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"authentication_type": rschema.StringAttribute{
												Optional:    true,
												Computed:    true,
												Description: "The authentication type (e.g. `None`, `Basic`).",
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"user_name": rschema.StringAttribute{
												Optional:    true,
												Computed:    true,
												Description: "The user name.",
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"password_wo": rschema.StringAttribute{
												Optional:    true,
												WriteOnly:   true,
												Sensitive:   true,
												Description: "The password (write-only). This value is sent to the API but never stored in Terraform state. Increment `password_wo_version` to re-send a changed value.",
											},
											"password_wo_version": rschema.Int64Attribute{
												Optional:    true,
												Computed:    true,
												Description: "Version of the `password_wo` field. Increment to re-send the password without changing other attributes.",
												PlanModifiers: []planmodifier.Int64{
													int64planmodifier.UseStateForUnknown(),
												},
											},
										},
									},
									"ignore_certificate_errors": rschema.BoolAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Whether to ignore certificate errors.",
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.UseStateForUnknown(),
										},
									},
									"client_certificate_vault_item": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The client certificate vault item.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"delay": rschema.Int64Attribute{
										Optional:    true,
										Computed:    true,
										Description: "The delay in milliseconds before the step.",
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
									},
									"step_type": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The step type (e.g. `HttpRequest`).",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"retry_until_successful": rschema.BoolAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Whether to retry until successful.",
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.UseStateForUnknown(),
										},
									},
									"max_attempts": rschema.Int64Attribute{
										Optional:    true,
										Computed:    true,
										Description: "The maximum number of attempts.",
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
									},
									"retry_wait_milliseconds": rschema.Int64Attribute{
										Optional:    true,
										Computed:    true,
										Description: "The wait time between retries in milliseconds.",
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
									},
									"pre_request_script": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Script executed before the request.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"post_response_script": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Script executed after the response.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"calculated_content_type": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The calculated content type.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"allowed_tls_versions": rschema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Allowed TLS versions (e.g. `Tls11`).",
										ElementType: types.StringType,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.UseStateForUnknown(),
										},
									},
									"msa_step_http_version": rschema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The HTTP version of the step (e.g. `Http1_1`).",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *customIntegrationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config customIntegrationModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := buildCustomIntegrationPayload(ctx, &config, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	created, err := r.client.CreateCustomIntegration(payload)
	if err != nil {
		resp.Diagnostics.AddError("Error creating custom integration", err.Error())
		return
	}

	state := mapCustomIntegrationToState(ctx, created, &config, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *customIntegrationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state customIntegrationModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	integration, err := r.client.GetCustomIntegration(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading custom integration", err.Error())
		return
	}

	newState := mapCustomIntegrationToState(ctx, integration, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}

func (r *customIntegrationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// The payload is built from the configuration (not the plan) because the
	// nested Terraform model uses plain Go slices/pointers that cannot decode the
	// unknown values an Optional+Computed plan may contain. Configuration values
	// are always known or null, so they decode cleanly.
	var config customIntegrationModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state customIntegrationModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	integrationGuid := state.ID.ValueString()
	payload := buildCustomIntegrationPayload(ctx, &config, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	payload.IntegrationGuid = integrationGuid

	// http_step_definition_usage_guid is computed and never present in config, so
	// carry the existing GUIDs over from prior state by index. Sending an empty
	// GUID would make the API allocate a new one, which conflicts with the value
	// the plan preserved via UseStateForUnknown.
	applyUsageGuids(payload, &state)

	updated, err := r.client.UpdateCustomIntegration(integrationGuid, payload)
	if err != nil {
		resp.Diagnostics.AddError("Error updating custom integration", err.Error())
		return
	}

	newState := mapCustomIntegrationToState(ctx, updated, &config, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}

func (r *customIntegrationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state customIntegrationModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteCustomIntegration(state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Error deleting custom integration", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *customIntegrationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// ---------------------------------------------------------------------------
// Conversions: Terraform model -> API payload
// ---------------------------------------------------------------------------

func buildCustomIntegrationPayload(ctx context.Context, model *customIntegrationModel, diags *diag.Diagnostics) *models.CustomIntegration {
	payload := &models.CustomIntegration{
		Name:     model.Name.ValueString(),
		IsActive: model.IsActive.ValueBool(),
		Notes:    model.Notes.ValueString(),
		// HttpStepDefinitions and IntegrationVariables are required by the API and
		// must be serialized as non-null arrays even when empty. PUT applies no
		// server-side defaulting, so the full object is always sent.
		HttpStepDefinitions:  []models.HttpStepDefinition{},
		IntegrationVariables: []models.IntegrationVariable{},
	}

	for _, def := range model.HttpStepDefinitions {
		apiDef := models.HttpStepDefinition{
			HttpStepDefinitionUsageGuid: def.HttpStepDefinitionUsageGuid.ValueString(),
			Usages:                      stringListToSlice(ctx, def.Usages, diags),
		}

		for _, step := range def.Steps {
			apiStep := models.HttpStep{
				Url:                        step.Url.ValueString(),
				Method:                     step.Method.ValueString(),
				Body:                       step.Body.ValueString(),
				BodyType:                   step.BodyType.ValueString(),
				VaultFileId:                step.VaultFileId.ValueString(),
				Name:                       step.Name.ValueString(),
				UseFixedClientCertificate:  step.UseFixedClientCertificate.ValueBool(),
				IgnoreCertificateErrors:    step.IgnoreCertificateErrors.ValueBool(),
				ClientCertificateVaultItem: step.ClientCertificateVaultItem.ValueString(),
				Delay:                      int(step.Delay.ValueInt64()),
				StepType:                   step.StepType.ValueString(),
				RetryUntilSuccessful:       step.RetryUntilSuccessful.ValueBool(),
				MaxAttempts:                int(step.MaxAttempts.ValueInt64()),
				RetryWaitMilliseconds:      int(step.RetryWaitMilliseconds.ValueInt64()),
				PreRequestScript:           step.PreRequestScript.ValueString(),
				PostResponseScript:         step.PostResponseScript.ValueString(),
				CalculatedContentType:      step.CalculatedContentType.ValueString(),
				AllowedTlsVersions:         stringListToSlice(ctx, step.AllowedTlsVersions, diags),
				MsaStepHttpVersion:         step.MsaStepHttpVersion.ValueString(),
			}

			for _, mpf := range step.MultiPartForm {
				apiStep.MultiPartForm = append(apiStep.MultiPartForm, models.MultiPartFormItem{
					Type:  mpf.Type.ValueString(),
					Key:   mpf.Key.ValueString(),
					Value: mpf.Value.ValueString(),
				})
			}
			for _, h := range step.RequestHeaders {
				apiStep.RequestHeaders = append(apiStep.RequestHeaders, models.IntegrationKeyValue{
					Key:   h.Key.ValueString(),
					Value: h.Value.ValueString(),
				})
			}
			for _, v := range step.Variables {
				apiStep.Variables = append(apiStep.Variables, models.HttpStepVariable{
					Source:    v.Source.ValueString(),
					Property:  v.Property.ValueString(),
					Name:      v.Name.ValueString(),
					Arguments: stringListToSlice(ctx, v.Arguments, diags),
				})
			}
			for _, a := range step.Assertions {
				apiStep.Assertions = append(apiStep.Assertions, models.HttpStepAssertion{
					Source:      a.Source.ValueString(),
					Property:    a.Property.ValueString(),
					Comparison:  a.Comparison.ValueString(),
					TargetValue: a.TargetValue.ValueString(),
				})
			}
			if step.Authentication != nil {
				apiStep.Authentication = &models.HttpStepAuthentication{
					Id:                 step.Authentication.Id.ValueString(),
					AuthenticationType: step.Authentication.AuthenticationType.ValueString(),
					UserName:           step.Authentication.UserName.ValueString(),
					Password:           step.Authentication.PasswordWO.ValueString(),
				}
				if !step.Authentication.PasswordWO.IsNull() {
					apiStep.Authentication.PasswordSpecified = true
				} else {
					apiStep.Authentication.PasswordSpecified = false
				}
			}

			apiDef.Steps = append(apiDef.Steps, apiStep)
		}

		for _, fn := range def.UserDefinedFunctions {
			apiFn := models.UserDefinedFunction{
				Name:          fn.Name.ValueString(),
				Type:          fn.Type.ValueString(),
				Regex:         fn.Regex.ValueString(),
				JwtSigningKey: fn.JwtSigningKey.ValueString(),
				JwtAlgorithm:  fn.JwtAlgorithm.ValueString(),
				HashKey:       fn.HashKey.ValueString(),
				HashAlgorithm: fn.HashAlgorithm.ValueString(),
			}
			for _, m := range fn.Mappings {
				apiFn.Mappings = append(apiFn.Mappings, models.IntegrationKeyValue{
					Key:   m.Key.ValueString(),
					Value: m.Value.ValueString(),
				})
			}
			apiDef.UserDefinedFunctions = append(apiDef.UserDefinedFunctions, apiFn)
		}

		payload.HttpStepDefinitions = append(payload.HttpStepDefinitions, apiDef)
	}

	for _, v := range model.IntegrationVariables {
		payload.IntegrationVariables = append(payload.IntegrationVariables, models.IntegrationVariable{
			Name:                        v.Name.ValueString(),
			Value:                       v.Value.ValueString(),
			IsValueSetInEscalationLevel: v.IsValueSetInEscalationLevel.ValueBool(),
		})
	}

	return payload
}

func stringListToSlice(ctx context.Context, list types.List, diags *diag.Diagnostics) []string {
	if list.IsNull() || list.IsUnknown() {
		return nil
	}
	var result []string
	diags.Append(list.ElementsAs(ctx, &result, false)...)
	return result
}

// ---------------------------------------------------------------------------
// Conversions: API response -> Terraform model
// ---------------------------------------------------------------------------

func mapCustomIntegrationToState(ctx context.Context, integration *models.CustomIntegration, source *customIntegrationModel, diags *diag.Diagnostics) customIntegrationModel {
	state := customIntegrationModel{
		ID:       types.StringValue(integration.IntegrationGuid),
		Name:     types.StringValue(integration.Name),
		IsActive: types.BoolValue(integration.IsActive),
		Notes:    types.StringValue(integration.Notes),
	}

	for _, def := range integration.HttpStepDefinitions {
		tfDef := httpStepDefinitionTF{
			HttpStepDefinitionUsageGuid: types.StringValue(def.HttpStepDefinitionUsageGuid),
			Usages:                      sliceToStringList(ctx, def.Usages, diags),
		}

		for _, step := range def.Steps {
			tfStep := httpStepTF{
				Url:                        types.StringValue(step.Url),
				Method:                     types.StringValue(step.Method),
				Body:                       types.StringValue(step.Body),
				BodyType:                   types.StringValue(step.BodyType),
				VaultFileId:                types.StringValue(step.VaultFileId),
				Name:                       types.StringValue(step.Name),
				UseFixedClientCertificate:  types.BoolValue(step.UseFixedClientCertificate),
				IgnoreCertificateErrors:    types.BoolValue(step.IgnoreCertificateErrors),
				ClientCertificateVaultItem: types.StringValue(step.ClientCertificateVaultItem),
				Delay:                      types.Int64Value(int64(step.Delay)),
				StepType:                   types.StringValue(step.StepType),
				RetryUntilSuccessful:       types.BoolValue(step.RetryUntilSuccessful),
				MaxAttempts:                types.Int64Value(int64(step.MaxAttempts)),
				RetryWaitMilliseconds:      types.Int64Value(int64(step.RetryWaitMilliseconds)),
				PreRequestScript:           types.StringValue(step.PreRequestScript),
				PostResponseScript:         types.StringValue(step.PostResponseScript),
				CalculatedContentType:      types.StringValue(step.CalculatedContentType),
				AllowedTlsVersions:         sliceToStringList(ctx, step.AllowedTlsVersions, diags),
				MsaStepHttpVersion:         types.StringValue(step.MsaStepHttpVersion),
			}

			for _, mpf := range step.MultiPartForm {
				tfStep.MultiPartForm = append(tfStep.MultiPartForm, multiPartFormItemTF{
					Type:  types.StringValue(mpf.Type),
					Key:   types.StringValue(mpf.Key),
					Value: types.StringValue(mpf.Value),
				})
			}
			for _, h := range step.RequestHeaders {
				tfStep.RequestHeaders = append(tfStep.RequestHeaders, keyValueTF{
					Key:   types.StringValue(h.Key),
					Value: types.StringValue(h.Value),
				})
			}
			for _, v := range step.Variables {
				tfStep.Variables = append(tfStep.Variables, httpStepVariableTF{
					Source:    types.StringValue(v.Source),
					Property:  types.StringValue(v.Property),
					Name:      types.StringValue(v.Name),
					Arguments: sliceToStringList(ctx, v.Arguments, diags),
				})
			}
			for _, a := range step.Assertions {
				tfStep.Assertions = append(tfStep.Assertions, httpStepAssertionTF{
					Source:      types.StringValue(a.Source),
					Property:    types.StringValue(a.Property),
					Comparison:  types.StringValue(a.Comparison),
					TargetValue: types.StringValue(a.TargetValue),
				})
			}
			if step.Authentication != nil {
				tfStep.Authentication = &httpStepAuthTF{
					Id:                 types.StringValue(step.Authentication.Id),
					AuthenticationType: types.StringValue(step.Authentication.AuthenticationType),
					UserName:           types.StringValue(step.Authentication.UserName),
					PasswordWO:         types.StringNull(),
					PasswordWOVersion:  types.Int64Null(),
				}
			}

			tfDef.Steps = append(tfDef.Steps, tfStep)
		}

		for _, fn := range def.UserDefinedFunctions {
			tfFn := userDefinedFunctionTF{
				Name:          types.StringValue(fn.Name),
				Type:          types.StringValue(fn.Type),
				Regex:         types.StringValue(fn.Regex),
				JwtSigningKey: types.StringValue(fn.JwtSigningKey),
				JwtAlgorithm:  types.StringValue(fn.JwtAlgorithm),
				HashKey:       types.StringValue(fn.HashKey),
				HashAlgorithm: types.StringValue(fn.HashAlgorithm),
			}
			for _, m := range fn.Mappings {
				tfFn.Mappings = append(tfFn.Mappings, keyValueTF{
					Key:   types.StringValue(m.Key),
					Value: types.StringValue(m.Value),
				})
			}
			tfDef.UserDefinedFunctions = append(tfDef.UserDefinedFunctions, tfFn)
		}

		state.HttpStepDefinitions = append(state.HttpStepDefinitions, tfDef)
	}

	for _, v := range integration.IntegrationVariables {
		state.IntegrationVariables = append(state.IntegrationVariables, integrationVariableTF{
			Name:                        types.StringValue(v.Name),
			Value:                       types.StringValue(v.Value),
			IsValueSetInEscalationLevel: types.BoolValue(v.IsValueSetInEscalationLevel),
		})
	}

	// Overlay the values the user actually configured back onto the state. For
	// Optional+Computed attributes that were set in configuration, the plan locks
	// to the configured value, so the post-apply state must match it exactly even
	// if the API echoes a normalized value. This also keeps the write-only
	// password null and preserves password_wo_version. Genuinely computed values
	// (usage GUID, authentication id) are left as returned by the API.
	reconcileWithConfig(&state, source)

	return state
}

func sliceToStringList(ctx context.Context, values []string, diags *diag.Diagnostics) types.List {
	if len(values) == 0 {
		return types.ListNull(types.StringType)
	}
	list, d := types.ListValueFrom(ctx, types.StringType, values)
	diags.Append(d...)
	return list
}

// applyUsageGuids carries the computed http_step_definition_usage_guid values
// from prior state into the API payload, matching definitions by index. These
// GUIDs are never present in configuration, and sending an empty value would
// make the API allocate new ones on every update.
func applyUsageGuids(payload *models.CustomIntegration, state *customIntegrationModel) {
	if payload == nil || state == nil {
		return
	}
	for di := range payload.HttpStepDefinitions {
		if di >= len(state.HttpStepDefinitions) {
			break
		}
		stateDef := state.HttpStepDefinitions[di]

		guid := stateDef.HttpStepDefinitionUsageGuid
		if !guid.IsNull() && !guid.IsUnknown() {
			payload.HttpStepDefinitions[di].HttpStepDefinitionUsageGuid = guid.ValueString()
		}

		// authentication.id is computed (never present in config). Carry it over
		// from prior state so the API does not reset the authentication when the
		// user leaves the password (and the rest of the block) unchanged.
		for si := range payload.HttpStepDefinitions[di].Steps {
			if si >= len(stateDef.Steps) {
				break
			}
			stateAuth := stateDef.Steps[si].Authentication
			payloadAuth := payload.HttpStepDefinitions[di].Steps[si].Authentication
			if stateAuth == nil || payloadAuth == nil {
				continue
			}
			if !stateAuth.Id.IsNull() && !stateAuth.Id.IsUnknown() {
				payloadAuth.Id = stateAuth.Id.ValueString()
			}
		}
	}
}

// reconcileWithConfig overlays the values the user configured (source) onto the
// state built from the API response, matching definitions and steps by index.
// For Optional+Computed attributes the plan locks to the configured value, so
// the state must reflect that value even if the API normalizes it; otherwise
// Terraform reports an "inconsistent result after apply" error (masked as a
// "sensitive attribute" error because the list contains the write-only
// password). Only values that are known in configuration are overlaid, so
// fields the user omitted keep the value returned by the API. Genuinely
// computed values (http_step_definition_usage_guid, authentication.id) are not
// overlaid. The write-only password is always left null in state.
func reconcileWithConfig(state, source *customIntegrationModel) {
	if source == nil {
		return
	}
	for di := range state.HttpStepDefinitions {
		if di >= len(source.HttpStepDefinitions) {
			break
		}
		srcDef := source.HttpStepDefinitions[di]
		stDef := &state.HttpStepDefinitions[di]

		overlayList(&stDef.Usages, srcDef.Usages)
		if srcDef.UserDefinedFunctions != nil {
			stDef.UserDefinedFunctions = srcDef.UserDefinedFunctions
		}

		for si := range stDef.Steps {
			if si >= len(srcDef.Steps) {
				break
			}
			srcStep := srcDef.Steps[si]
			stStep := &stDef.Steps[si]

			overlayString(&stStep.Url, srcStep.Url)
			overlayString(&stStep.Method, srcStep.Method)
			overlayString(&stStep.Body, srcStep.Body)
			overlayString(&stStep.BodyType, srcStep.BodyType)
			overlayString(&stStep.VaultFileId, srcStep.VaultFileId)
			overlayString(&stStep.Name, srcStep.Name)
			overlayString(&stStep.ClientCertificateVaultItem, srcStep.ClientCertificateVaultItem)
			overlayString(&stStep.StepType, srcStep.StepType)
			overlayString(&stStep.PreRequestScript, srcStep.PreRequestScript)
			overlayString(&stStep.PostResponseScript, srcStep.PostResponseScript)
			overlayString(&stStep.CalculatedContentType, srcStep.CalculatedContentType)
			overlayString(&stStep.MsaStepHttpVersion, srcStep.MsaStepHttpVersion)
			overlayBool(&stStep.UseFixedClientCertificate, srcStep.UseFixedClientCertificate)
			overlayBool(&stStep.IgnoreCertificateErrors, srcStep.IgnoreCertificateErrors)
			overlayBool(&stStep.RetryUntilSuccessful, srcStep.RetryUntilSuccessful)
			overlayInt(&stStep.Delay, srcStep.Delay)
			overlayInt(&stStep.MaxAttempts, srcStep.MaxAttempts)
			overlayInt(&stStep.RetryWaitMilliseconds, srcStep.RetryWaitMilliseconds)
			overlayList(&stStep.AllowedTlsVersions, srcStep.AllowedTlsVersions)

			if srcStep.MultiPartForm != nil {
				stStep.MultiPartForm = srcStep.MultiPartForm
			}
			if srcStep.RequestHeaders != nil {
				stStep.RequestHeaders = srcStep.RequestHeaders
			}
			if srcStep.Variables != nil {
				stStep.Variables = srcStep.Variables
			}
			if srcStep.Assertions != nil {
				stStep.Assertions = srcStep.Assertions
			}

			srcAuth := srcStep.Authentication
			// `authentication` is a non-computed (Optional) block, so its presence
			// in state must match the configuration.
			if srcAuth == nil {
				stStep.Authentication = nil
				continue
			}
			if stStep.Authentication == nil {
				authCopy := *srcAuth
				authCopy.PasswordWO = types.StringNull()
				stStep.Authentication = &authCopy
				continue
			}
			stAuth := stStep.Authentication
			overlayString(&stAuth.AuthenticationType, srcAuth.AuthenticationType)
			overlayString(&stAuth.UserName, srcAuth.UserName)
			overlayString(&stAuth.Id, srcAuth.Id)
			stAuth.PasswordWOVersion = srcAuth.PasswordWOVersion
			stAuth.PasswordWO = types.StringNull()
		}
	}
}

func overlayString(dst *types.String, src types.String) {
	if !src.IsNull() && !src.IsUnknown() {
		*dst = src
	}
}

func overlayBool(dst *types.Bool, src types.Bool) {
	if !src.IsNull() && !src.IsUnknown() {
		*dst = src
	}
}

func overlayInt(dst *types.Int64, src types.Int64) {
	if !src.IsNull() && !src.IsUnknown() {
		*dst = src
	}
}

func overlayList(dst *types.List, src types.List) {
	if !src.IsNull() && !src.IsUnknown() {
		*dst = src
	}
}
