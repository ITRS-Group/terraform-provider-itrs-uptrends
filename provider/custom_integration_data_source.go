package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ datasource.DataSource = &customIntegrationDataSource{}

type customIntegrationDataSource struct {
	client interfaces.ICustomIntegration
}

func NewCustomIntegrationDataSource(client interfaces.ICustomIntegration) datasource.DataSource {
	return &customIntegrationDataSource{client: client}
}

// ---------------------------------------------------------------------------
// Data source models (read-only; the write-only password is never exposed)
// ---------------------------------------------------------------------------

type customIntegrationDataModel struct {
	ID                   types.String               `tfsdk:"id"`
	Name                 types.String               `tfsdk:"name"`
	IsActive             types.Bool                 `tfsdk:"is_active"`
	Notes                types.String               `tfsdk:"notes"`
	HttpStepDefinitions  []httpStepDefinitionDataTF `tfsdk:"http_step_definitions"`
	IntegrationVariables []integrationVariableTF    `tfsdk:"integration_variables"`
}

type httpStepDefinitionDataTF struct {
	HttpStepDefinitionUsageGuid types.String            `tfsdk:"http_step_definition_usage_guid"`
	Steps                       []httpStepDataTF        `tfsdk:"steps"`
	UserDefinedFunctions        []userDefinedFunctionTF `tfsdk:"user_defined_functions"`
	Usages                      types.List              `tfsdk:"usages"`
}

type httpStepDataTF struct {
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
	Authentication             *httpStepAuthDataTF   `tfsdk:"authentication"`
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

type httpStepAuthDataTF struct {
	Id                 types.String `tfsdk:"id"`
	AuthenticationType types.String `tfsdk:"authentication_type"`
	UserName           types.String `tfsdk:"user_name"`
}

func (d *customIntegrationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_integration"
}

func (d *customIntegrationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dschema.Schema{
		Description: "Fetch a custom integration by GUID or name to reference it in other resources.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The GUID of the custom integration. Provide this or `name`.",
			},
			"name": dschema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the custom integration. Provide this or `id`. Must be unique.",
			},
			"is_active": dschema.BoolAttribute{
				Computed:    true,
				Description: "Whether the custom integration is active.",
			},
			"notes": dschema.StringAttribute{
				Computed:    true,
				Description: "Notes for the custom integration.",
			},
			"integration_variables": dschema.ListNestedAttribute{
				Computed:    true,
				Description: "Variables that can be referenced by the integration steps.",
				NestedObject: dschema.NestedAttributeObject{
					Attributes: map[string]dschema.Attribute{
						"name": dschema.StringAttribute{
							Computed:    true,
							Description: "The name of the variable.",
						},
						"value": dschema.StringAttribute{
							Computed:    true,
							Description: "The value of the variable.",
						},
						"is_value_set_in_escalation_level": dschema.BoolAttribute{
							Computed:    true,
							Description: "Whether the value is set in the escalation level.",
						},
					},
				},
			},
			"http_step_definitions": dschema.ListNestedAttribute{
				Computed:    true,
				Description: "The HTTP step definitions of the custom integration.",
				NestedObject: dschema.NestedAttributeObject{
					Attributes: map[string]dschema.Attribute{
						"http_step_definition_usage_guid": dschema.StringAttribute{
							Computed:    true,
							Description: "The usage GUID of the HTTP step definition.",
						},
						"usages": dschema.ListAttribute{
							Computed:    true,
							Description: "The usages of the HTTP step definition (e.g. `Alert`).",
							ElementType: types.StringType,
						},
						"user_defined_functions": dschema.ListNestedAttribute{
							Computed:    true,
							Description: "User defined functions usable by the steps.",
							NestedObject: dschema.NestedAttributeObject{
								Attributes: map[string]dschema.Attribute{
									"name": dschema.StringAttribute{
										Computed:    true,
										Description: "The name of the function.",
									},
									"type": dschema.StringAttribute{
										Computed:    true,
										Description: "The type of the function (e.g. `Mapping`).",
									},
									"mappings": dschema.ListNestedAttribute{
										Computed:    true,
										Description: "Key/value mappings for the function.",
										NestedObject: dschema.NestedAttributeObject{
											Attributes: map[string]dschema.Attribute{
												"key": dschema.StringAttribute{
													Computed:    true,
													Description: "The mapping key.",
												},
												"value": dschema.StringAttribute{
													Computed:    true,
													Description: "The mapping value.",
												},
											},
										},
									},
									"regex": dschema.StringAttribute{
										Computed:    true,
										Description: "The regular expression of the function.",
									},
									"jwt_signing_key": dschema.StringAttribute{
										Computed:    true,
										Description: "The JWT signing key.",
									},
									"jwt_algorithm": dschema.StringAttribute{
										Computed:    true,
										Description: "The JWT algorithm (e.g. `HS256`).",
									},
									"hash_key": dschema.StringAttribute{
										Computed:    true,
										Description: "The hash key.",
									},
									"hash_algorithm": dschema.StringAttribute{
										Computed:    true,
										Description: "The hash algorithm (e.g. `MD5`).",
									},
								},
							},
						},
						"steps": dschema.ListNestedAttribute{
							Computed:    true,
							Description: "The HTTP steps to execute.",
							NestedObject: dschema.NestedAttributeObject{
								Attributes: map[string]dschema.Attribute{
									"url": dschema.StringAttribute{
										Computed:    true,
										Description: "The URL of the request.",
									},
									"method": dschema.StringAttribute{
										Computed:    true,
										Description: "The HTTP method (e.g. `GET`, `POST`).",
									},
									"body": dschema.StringAttribute{
										Computed:    true,
										Description: "The request body.",
									},
									"body_type": dschema.StringAttribute{
										Computed:    true,
										Description: "The body type (e.g. `Raw`).",
									},
									"vault_file_id": dschema.StringAttribute{
										Computed:    true,
										Description: "The vault file id.",
									},
									"multi_part_form": dschema.ListNestedAttribute{
										Computed:    true,
										Description: "Multipart form entries.",
										NestedObject: dschema.NestedAttributeObject{
											Attributes: map[string]dschema.Attribute{
												"type": dschema.StringAttribute{
													Computed:    true,
													Description: "The type of the entry (e.g. `Text`).",
												},
												"key": dschema.StringAttribute{
													Computed:    true,
													Description: "The key of the entry.",
												},
												"value": dschema.StringAttribute{
													Computed:    true,
													Description: "The value of the entry.",
												},
											},
										},
									},
									"request_headers": dschema.ListNestedAttribute{
										Computed:    true,
										Description: "Request headers.",
										NestedObject: dschema.NestedAttributeObject{
											Attributes: map[string]dschema.Attribute{
												"key": dschema.StringAttribute{
													Computed:    true,
													Description: "The header name.",
												},
												"value": dschema.StringAttribute{
													Computed:    true,
													Description: "The header value.",
												},
											},
										},
									},
									"variables": dschema.ListNestedAttribute{
										Computed:    true,
										Description: "Variables captured from the response.",
										NestedObject: dschema.NestedAttributeObject{
											Attributes: map[string]dschema.Attribute{
												"source": dschema.StringAttribute{
													Computed:    true,
													Description: "The source of the variable (e.g. `None`).",
												},
												"property": dschema.StringAttribute{
													Computed:    true,
													Description: "The property to capture.",
												},
												"name": dschema.StringAttribute{
													Computed:    true,
													Description: "The name of the variable.",
												},
												"arguments": dschema.ListAttribute{
													Computed:    true,
													Description: "The arguments of the variable.",
													ElementType: types.StringType,
												},
											},
										},
									},
									"assertions": dschema.ListNestedAttribute{
										Computed:    true,
										Description: "Assertions to validate the response.",
										NestedObject: dschema.NestedAttributeObject{
											Attributes: map[string]dschema.Attribute{
												"source": dschema.StringAttribute{
													Computed:    true,
													Description: "The source of the assertion (e.g. `None`).",
												},
												"property": dschema.StringAttribute{
													Computed:    true,
													Description: "The property to assert.",
												},
												"comparison": dschema.StringAttribute{
													Computed:    true,
													Description: "The comparison (e.g. `Equal`).",
												},
												"target_value": dschema.StringAttribute{
													Computed:    true,
													Description: "The target value of the assertion.",
												},
											},
										},
									},
									"name": dschema.StringAttribute{
										Computed:    true,
										Description: "The name of the step.",
									},
									"use_fixed_client_certificate": dschema.BoolAttribute{
										Computed:    true,
										Description: "Whether to use a fixed client certificate.",
									},
									"authentication": dschema.SingleNestedAttribute{
										Computed:    true,
										Description: "Authentication settings for the step. The password is never exposed.",
										Attributes: map[string]dschema.Attribute{
											"id": dschema.StringAttribute{
												Computed:    true,
												Description: "The id of the authentication.",
											},
											"authentication_type": dschema.StringAttribute{
												Computed:    true,
												Description: "The authentication type (e.g. `None`, `Basic`).",
											},
											"user_name": dschema.StringAttribute{
												Computed:    true,
												Description: "The user name.",
											},
										},
									},
									"ignore_certificate_errors": dschema.BoolAttribute{
										Computed:    true,
										Description: "Whether to ignore certificate errors.",
									},
									"client_certificate_vault_item": dschema.StringAttribute{
										Computed:    true,
										Description: "The client certificate vault item.",
									},
									"delay": dschema.Int64Attribute{
										Computed:    true,
										Description: "The delay in milliseconds before the step.",
									},
									"step_type": dschema.StringAttribute{
										Computed:    true,
										Description: "The step type (e.g. `HttpRequest`).",
									},
									"retry_until_successful": dschema.BoolAttribute{
										Computed:    true,
										Description: "Whether to retry until successful.",
									},
									"max_attempts": dschema.Int64Attribute{
										Computed:    true,
										Description: "The maximum number of attempts.",
									},
									"retry_wait_milliseconds": dschema.Int64Attribute{
										Computed:    true,
										Description: "The wait time between retries in milliseconds.",
									},
									"pre_request_script": dschema.StringAttribute{
										Computed:    true,
										Description: "Script executed before the request.",
									},
									"post_response_script": dschema.StringAttribute{
										Computed:    true,
										Description: "Script executed after the response.",
									},
									"calculated_content_type": dschema.StringAttribute{
										Computed:    true,
										Description: "The calculated content type.",
									},
									"allowed_tls_versions": dschema.ListAttribute{
										Computed:    true,
										Description: "Allowed TLS versions (e.g. `Tls11`).",
										ElementType: types.StringType,
									},
									"msa_step_http_version": dschema.StringAttribute{
										Computed:    true,
										Description: "The HTTP version of the step (e.g. `Http1_1`).",
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

func (d *customIntegrationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The custom integration client was not configured. This is an internal error in the provider.")
		return
	}

	var data customIntegrationDataModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idProvided := !data.ID.IsNull() && data.ID.ValueString() != ""
	nameProvided := !data.Name.IsNull() && data.Name.ValueString() != ""

	switch {
	case idProvided && nameProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide only one of id or name to look up a custom integration.")
		return
	case !idProvided && !nameProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide either id or name to look up a custom integration.")
		return
	}

	var integration *models.CustomIntegration

	if idProvided {
		result, err := d.client.GetCustomIntegration(data.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading custom integration", err.Error())
			return
		}
		integration = result
	} else {
		integrations, err := d.client.GetCustomIntegrations()
		if err != nil {
			resp.Diagnostics.AddError("Error listing custom integrations", err.Error())
			return
		}
		name := data.Name.ValueString()
		for idx := range integrations {
			if strings.EqualFold(integrations[idx].Name, name) {
				if integration != nil {
					resp.Diagnostics.AddError("Custom integration not unique", fmt.Sprintf("More than one custom integration found with name %q", name))
					return
				}
				match := integrations[idx]
				integration = &match
			}
		}
		if integration == nil {
			resp.Diagnostics.AddError("Custom integration not found", fmt.Sprintf("No custom integration found with name %q", name))
			return
		}
	}

	state := mapCustomIntegrationToDataModel(ctx, integration, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func mapCustomIntegrationToDataModel(ctx context.Context, integration *models.CustomIntegration, diags *diag.Diagnostics) customIntegrationDataModel {
	state := customIntegrationDataModel{
		ID:       types.StringValue(integration.IntegrationGuid),
		Name:     types.StringValue(integration.Name),
		IsActive: types.BoolValue(integration.IsActive),
		Notes:    types.StringValue(integration.Notes),
	}

	for _, def := range integration.HttpStepDefinitions {
		tfDef := httpStepDefinitionDataTF{
			HttpStepDefinitionUsageGuid: types.StringValue(def.HttpStepDefinitionUsageGuid),
			Usages:                      sliceToStringList(ctx, def.Usages, diags),
		}

		for _, step := range def.Steps {
			tfStep := httpStepDataTF{
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
				tfStep.Authentication = &httpStepAuthDataTF{
					Id:                 types.StringValue(step.Authentication.Id),
					AuthenticationType: types.StringValue(step.Authentication.AuthenticationType),
					UserName:           types.StringValue(step.Authentication.UserName),
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

	return state
}
