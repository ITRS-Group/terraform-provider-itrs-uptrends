package provider

import (
	"context"
	"log"
	"runtime"
	"strings"

	"github.com/itrs-group/terraform-provider-itrs-uptrends/general"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource" // Added import for resources
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/client"
	api "github.com/itrs-group/terraform-provider-itrs-uptrends/client/api"
)

type UptrendsProvider struct {
	operator                               *api.Operator
	operatorGroup                          *api.OperatorGroup
	membership                             *api.Membership
	operatorFacade                         *api.OperatorFacade
	monitorGroupMembership                 *api.MonitorGroupMember
	alertDefinition                        *api.AlertDefinition
	alertDefinitionMonitorMember           *api.AlertDefinitionMonitorMember
	monitor                                *api.Monitor
	monitorGroup                           *api.MonitorGroupClient
	alertDefinitionOperatorMembership      *api.AlertDefinitionOperatorMembership
	alertDefinitionOperatorGroupMembership *api.AlertDefinitionOperatorGroupMembership
	operatorGroupPermission                *api.OperatorGroupPermission
	operatorPermission                     *api.OperatorPermission
	vaultItem                              *api.VaultItem
	vaultSection                           *api.VaultSection
}

const defaultBaseUrl = "https://api.uptrends.com/v4"

// Ensure UptrendsProvider implements the provider.Provider interface.
var _ provider.Provider = &UptrendsProvider{}

func New() provider.Provider {
	return &UptrendsProvider{}
}

func (p *UptrendsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "itrs-uptrends"
	resp.Version = strings.ReplaceAll(general.NewBuildVersion, "__NEW_BUILD_VERSION__", "1.0.0")
}

func (p *UptrendsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Username for the Uptrends API.",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "Password for the Uptrends API.",
			},
			"debug": schema.BoolAttribute{
				Optional:    true,
				Description: "Enable debug mode.",
			},
			"baseurl": schema.StringAttribute{
				Optional:    true,
				Description: "Custom API URL. Defaults to " + defaultBaseUrl + " if not provided.",
			},
		},
	}
}

func (p *UptrendsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config struct {
		Username types.String `tfsdk:"username"`
		Password types.String `tfsdk:"password"`
		Debug    types.Bool   `tfsdk:"debug"`
		BaseUrl  types.String `tfsdk:"baseurl"`
	}
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {

		return
	}

	log.Printf("UptrendsProvider data received from %s", config.Username.ValueString())

	var baseAPIUrl string
	if config.BaseUrl.IsNull() || config.BaseUrl.ValueString() == "" {
		baseAPIUrl = defaultBaseUrl
	} else {
		baseAPIUrl = config.BaseUrl.ValueString()
	}

	var urlSource = client.NewUrlSource(baseAPIUrl)
	platform := runtime.GOOS
	var header = client.GenerateBasicAuthHeader(config.Username.ValueString(), config.Password.ValueString())

	p.operator = api.NewOperator(urlSource.OperatorURL(), header, general.NewBuildVersion, platform)
	p.operatorGroup = api.NewOperatorGroup(urlSource.OperatorGroupURL(), header, general.NewBuildVersion, platform)
	p.membership = api.NewMembership(urlSource.OperatorGroupURL(), header, general.NewBuildVersion, platform)
	p.operatorFacade = api.NewOperatorFacade(p.operator, p.operatorGroup, p.membership)
	p.monitor = api.NewMonitorClient(header, urlSource.MonitorURL(), general.NewBuildVersion, platform)
	p.monitorGroup = api.NewMonitorGroupClient(urlSource.MonitorGroupURL(), header, general.NewBuildVersion, platform)
	p.monitorGroupMembership = api.NewMonitorGroupMember(urlSource.MonitorGroupURL(), header, general.NewBuildVersion, platform)
	p.alertDefinition = api.NewAlertDefinition(urlSource.AlertDefinitionURL(), header, general.NewBuildVersion, platform)
	p.alertDefinitionMonitorMember = api.NewAlertDefinitionMonitorMember(urlSource.AlertDefinitionURL(), header, general.NewBuildVersion, platform)
	p.alertDefinitionOperatorMembership = api.NewAlertDefinitionOperatorMembership(urlSource.AlertDefinitionURL(), header, general.NewBuildVersion, platform)
	p.alertDefinitionOperatorGroupMembership = api.NewAlertDefinitionOperatorGroupMembership(urlSource.AlertDefinitionURL(), header, general.NewBuildVersion, platform)
	p.operatorGroupPermission = api.NewOperatorGroupPermission(urlSource.OperatorGroupURL(), header, general.NewBuildVersion, platform)
	p.operatorPermission = api.NewOperatorPermission(urlSource.OperatorURL(), header, general.NewBuildVersion, platform)
	p.vaultItem = api.NewVaultItem(urlSource.VaultItemURL(), header, general.NewBuildVersion, platform)
	p.vaultSection = api.NewVaultSection(urlSource.VaultSectionURL(), header, general.NewBuildVersion, platform)

}

func (p *UptrendsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		p.createAlertDefinition,
		p.createAlertDefinitionOperatorMembershipResource,
		p.createAlertDefinitionOperatorGroupMembershipResource,
		p.createAlertDefinitionMonitorMember,
		p.createMembershipResource,
		p.createMonitorgroupMembershipResource,
		p.createMonitorGroupResource,
		p.createMonitorResource,
		p.createOperatorGroupResource,
		p.createOperatorResource,
		p.createOperatorGroupPermissionResource,
		p.createOperatorPermissionResource,
		p.createVaultItemResource,
		p.createVaultSectionResource,
	}
}

func (p *UptrendsProvider) createAlertDefinition() resource.Resource {
	return NewAlertdefinitionResource(p.alertDefinition)
}

func (p *UptrendsProvider) createAlertDefinitionOperatorMembershipResource() resource.Resource {
	return NewAlertDefinitionOperatorMembershipResource(p.alertDefinitionOperatorMembership)
}

func (p *UptrendsProvider) createAlertDefinitionOperatorGroupMembershipResource() resource.Resource {
	return NewAlertDefinitionOperatorGroupMembershipResource(p.alertDefinitionOperatorGroupMembership)
}

func (p *UptrendsProvider) createAlertDefinitionMonitorMember() resource.Resource {
	return NewAlertDefinitionMonitorMembershipResource(p.alertDefinitionMonitorMember)
}

func (p *UptrendsProvider) createMembershipResource() resource.Resource {
	return NewMembershipResource(p.membership)
}

func (p *UptrendsProvider) createMonitorgroupMembershipResource() resource.Resource {
	return NewMonitorgroupMembershipResource(p.monitorGroupMembership)
}

func (p *UptrendsProvider) createMonitorGroupResource() resource.Resource {
	return NewMonitorGroupResource(p.monitorGroup)
}

func (p *UptrendsProvider) createMonitorResource() resource.Resource {
	return NewMonitorResource(p.monitor)
}

func (p *UptrendsProvider) createOperatorResource() resource.Resource {
	return NewOperatorResource(p.operatorFacade)
}

func (p *UptrendsProvider) createOperatorGroupResource() resource.Resource {
	return NewOperatorGroupResource(p.operatorGroup)
}

func (p *UptrendsProvider) createOperatorGroupPermissionResource() resource.Resource {
	return NewOperatorGroupPermissionResource(p.operatorGroupPermission)
}

func (p *UptrendsProvider) createOperatorPermissionResource() resource.Resource {
	return NewOperatorPermissionResource(p.operatorPermission)
}

func (p *UptrendsProvider) createVaultItemResource() resource.Resource {
	return NewVaultItemResource(p.vaultItem)
}

func (p *UptrendsProvider) createVaultSectionResource() resource.Resource {
	return NewVaultSectionResource(p.vaultSection)
}

func (p *UptrendsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
