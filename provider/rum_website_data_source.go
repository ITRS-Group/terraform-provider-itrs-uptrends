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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &rumWebsiteDataSource{}
)

// NewRumWebsiteDataSource constructs the rum website data source.
func NewRumWebsiteDataSource(client interfaces.IRumWebsite) datasource.DataSource {
	return &rumWebsiteDataSource{client: client}
}

type rumWebsiteDataSource struct {
	client interfaces.IRumWebsite
}

// Metadata returns the data source type name.
func (d *rumWebsiteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rum_website"
}

// Schema defines the schema for the data source.
func (d *rumWebsiteDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "RUM website GUID. Provide this or description.",
				Optional:    true,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "RUM website description (name). Provide this or id.",
				Optional:    true,
				Computed:    true,
			},
			"url": schema.StringAttribute{
				Description: "The URL of the monitored website.",
				Computed:    true,
			},
			"is_spa": schema.BoolAttribute{
				Description: "Indicates whether the website is a Single Page Application (SPA).",
				Computed:    true,
			},
			"include_url_fragment": schema.BoolAttribute{
				Description: "Specifies whether to include the URL fragment (hash) in monitoring.",
				Computed:    true,
			},
			"rum_script": schema.StringAttribute{
				Description: "The RUM script to be added to the website.",
				Computed:    true,
			},
		},
	}
}

type rumWebsiteDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Description        types.String `tfsdk:"description"`
	Url                types.String `tfsdk:"url"`
	IsSpa              types.Bool   `tfsdk:"is_spa"`
	IncludeUrlFragment types.Bool   `tfsdk:"include_url_fragment"`
	RumScript          types.String `tfsdk:"rum_script"`
}

// Read refreshes the Terraform state with the latest data.
func (d *rumWebsiteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The rum website client was not configured. This is an internal error in the provider.")
		return
	}

	var data rumWebsiteDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idProvided := !data.Id.IsNull() && data.Id.ValueString() != ""
	descriptionProvided := !data.Description.IsNull() && data.Description.ValueString() != ""

	switch {
	case idProvided && descriptionProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide only one of id or description to look up a rum website.")
		return
	case !idProvided && !descriptionProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide either id or description to look up a rum website.")
		return
	}

	if idProvided {
		result, msg, err := d.client.GetRumWebsite(data.Id.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading rum website", fmt.Sprintf("Error: %v. Message: %s", err, msg))
			return
		}

		data.Id = types.StringValue(result.RumWebsiteGuid)
		data.Description = types.StringValue(result.Description)
		data.Url = types.StringValue(result.Url)
		data.IsSpa = types.BoolValue(result.IsSpa)
		data.IncludeUrlFragment = types.BoolValue(result.IncludeUrlFragment)
		data.RumScript = types.StringValue(result.RumScript)
		diags = resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	}

	rumWebsites, statusCode, responseBody, err := d.client.GetRumWebsites()
	if err != nil {
		resp.Diagnostics.AddError("Error listing rum websites", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to list rum websites",
			fmt.Sprintf("HTTP status code: %d with response body %v", statusCode, responseBody),
		)
		return
	}

	name := data.Description.ValueString()
	found := false
	for idx := range rumWebsites {
		if strings.EqualFold(rumWebsites[idx].Description, name) {
			if found {
				resp.Diagnostics.AddError("Rum website not unique", fmt.Sprintf("More than one rum website found with description %q", name))
				return
			}
			data.Id = types.StringValue(rumWebsites[idx].RumWebsiteGuid)
			data.Description = types.StringValue(rumWebsites[idx].Description)
			data.Url = types.StringValue(rumWebsites[idx].Url)
			data.IsSpa = types.BoolValue(rumWebsites[idx].IsSpa)
			data.IncludeUrlFragment = types.BoolValue(rumWebsites[idx].IncludeUrlFragment)
			data.RumScript = types.StringValue(rumWebsites[idx].RumScript)
			found = true
		}
	}

	if !found {
		resp.Diagnostics.AddError("Rum website not found", fmt.Sprintf("No rum website found with description %q", name))
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
