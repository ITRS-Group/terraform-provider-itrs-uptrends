package converters

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	jsonmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
	tfsdkmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/provider/models"
)

// PayloadConversion converts the tfsdk MonitorModel to the JSON MonitorRequest.
// For optional fields (pointer types), the value is assigned only when not nil and not null.
func PayloadConversion(config tfsdkmodels.MonitorModel) jsonmodels.MonitorRequest {

	payload := jsonmodels.MonitorRequest{
		Name:                      config.Name.ValueString(),
		MonitorType:               config.MonitorType.ValueString(),
		GenerateAlert:             config.GenerateAlert.ValueBool(),
		IsActive:                  config.IsActive.ValueBool(),
		MonitorMode:               config.MonitorMode.ValueString(),
		UsePrimaryCheckpointsOnly: config.UsePrimaryCheckpointsOnly.ValueBool(),
	}

	if !config.CheckInterval.IsNull() {
		v := int(config.CheckInterval.ValueInt64())
		payload.CheckInterval = &v
	}

	if !config.CheckIntervalSeconds.IsNull() {
		v := int(config.CheckIntervalSeconds.ValueInt64())
		payload.CheckIntervalSeconds = &v
	}

	if !config.UseConcurrentMonitoring.IsNull() {
		val := config.UseConcurrentMonitoring.ValueBool()
		payload.UseConcurrentMonitoring = &val
	}

	if !config.Notes.IsNull() {
		notes := config.Notes.ValueString()
		payload.Notes = &notes
	}

	if !config.MonitorGuid.IsNull() {
		guid := config.MonitorGuid.ValueString()
		payload.MonitorGuid = &guid
	}

	if config.CustomMetrics != nil {
		var convCM []jsonmodels.CustomMetric
		for _, cm := range *config.CustomMetrics {
			convCM = append(convCM, jsonmodels.CustomMetric{
				Name:         cm.Name.ValueString(),
				VariableName: cm.VariableName.ValueString(),
			})
		}
		payload.CustomMetrics = &convCM
	}

	if config.CustomFields != nil {
		var convCF []jsonmodels.CustomField
		for _, cf := range *config.CustomFields {
			convCF = append(convCF, jsonmodels.CustomField{
				Name:  cf.Name.ValueString(),
				Value: cf.Value.ValueString(),
			})
		}
		payload.CustomFields = convCF
	}

	if config.SelectedCheckpoints != nil {
		payload.SelectedCheckpoints = convertSelectedCheckpoints(*config.SelectedCheckpoints)
	}

	if !config.SelfServiceTransactionScript.IsNull() {
		v := config.SelfServiceTransactionScript.ValueString()
		payload.SelfServiceTransactionScript = &v
	}

	if !config.MultiStepApiTransactionScript.IsNull() {
		v := config.MultiStepApiTransactionScript.ValueString()
		payload.MultiStepApiTransactionScript = &v
	}

	if !config.BlockGoogleAnalytics.IsNull() {
		v := config.BlockGoogleAnalytics.ValueBool()
		payload.BlockGoogleAnalytics = &v
	}

	if !config.BlockUptrendsRum.IsNull() {
		v := config.BlockUptrendsRum.ValueBool()
		payload.BlockUptrendsRum = &v
	}

	if !config.BlockUrls.IsNull() && !config.BlockUrls.IsUnknown() {
		blockUrls := convertStringList(config.BlockUrls)
		payload.BlockUrls = &blockUrls
	} else {
		payload.BlockUrls = nil
	}

	if config.RequestHeaders != nil {
		var convRH []jsonmodels.RequestHeader
		for _, rh := range *config.RequestHeaders {
			convRH = append(convRH, jsonmodels.RequestHeader{
				Name:  rh.Name.ValueString(),
				Value: rh.Value.ValueString(),
			})
		}
		payload.RequestHeaders = &convRH
	}

	if config.PredefinedVariables != nil {
		var convPV []jsonmodels.PredefinedVariables
		for _, rh := range *config.PredefinedVariables {
			convPV = append(convPV, jsonmodels.PredefinedVariables{
				Key:   rh.Key.ValueString(),
				Value: rh.Value.ValueString(),
			})
		}
		payload.PredefinedVariables = &convPV
	}

	if !config.UserAgent.IsNull() {
		v := config.UserAgent.ValueString()
		payload.UserAgent = &v
	}

	if !config.Username.IsNull() {
		v := config.Username.ValueString()
		payload.Username = &v
	}

	if !config.Password.IsNull() {
		v := config.Password.ValueString()
		payload.Password = &v
	}

	if !config.NameForPhoneAlerts.IsNull() {
		v := config.NameForPhoneAlerts.ValueString()
		payload.NameForPhoneAlerts = &v
	}

	if !config.AuthenticationType.IsNull() {
		v := config.AuthenticationType.ValueString()
		payload.AuthenticationType = &v
	}

	if config.ThrottlingOptions != nil {
		payload.ThrottlingOptions = convertThrottlingOptions(*config.ThrottlingOptions)
	}

	if config.DnsBypasses != nil {
		var convDB []jsonmodels.DnsBypass
		for _, db := range *config.DnsBypasses {
			convDB = append(convDB, jsonmodels.DnsBypass{
				Source: db.Source.ValueString(),
				Target: db.Target.ValueString(),
			})
		}
		payload.DnsBypasses = &convDB
	}

	if !config.CertificateName.IsNull() {
		v := config.CertificateName.ValueString()
		payload.CertificateName = &v
	}

	if !config.CertificateOrganization.IsNull() {
		v := config.CertificateOrganization.ValueString()
		payload.CertificateOrganization = &v
	}

	if !config.CertificateOrganizationalUnit.IsNull() {
		v := config.CertificateOrganizationalUnit.ValueString()
		payload.CertificateOrganizationalUnit = &v
	}

	if !config.CertificateSerialNumber.IsNull() {
		v := config.CertificateSerialNumber.ValueString()
		payload.CertificateSerialNumber = &v
	}

	if !config.CertificateFingerprint.IsNull() {
		v := config.CertificateFingerprint.ValueString()
		payload.CertificateFingerprint = &v
	}

	if !config.CertificateIssuerName.IsNull() {
		v := config.CertificateIssuerName.ValueString()
		payload.CertificateIssuerName = &v
	}

	if !config.CertificateIssuerCompanyName.IsNull() {
		v := config.CertificateIssuerCompanyName.ValueString()
		payload.CertificateIssuerCompanyName = &v
	}

	if !config.CertificateIssuerOrganizationalUnit.IsNull() {
		v := config.CertificateIssuerOrganizationalUnit.ValueString()
		payload.CertificateIssuerOrganizationalUnit = &v
	}

	if !config.CertificateExpirationWarningDays.IsNull() {
		v := int(config.CertificateExpirationWarningDays.ValueInt64())
		payload.CertificateExpirationWarningDays = &v
	}

	if !config.CheckCertificateErrors.IsNull() {
		v := config.CheckCertificateErrors.ValueBool()
		payload.CheckCertificateErrors = &v
	}

	if !config.IgnoreExternalElements.IsNull() {
		v := config.IgnoreExternalElements.ValueBool()
		payload.IgnoreExternalElements = &v
	}

	if !config.DomainGroupGuid.IsNull() {
		v := config.DomainGroupGuid.ValueString()
		payload.DomainGroupGuid = &v
	}

	if !config.DomainGroupGuidSpecified.IsNull() {
		v := config.DomainGroupGuidSpecified.ValueBool()
		payload.DomainGroupGuidSpecified = &v
	}

	if !config.DnsServer.IsNull() {
		v := config.DnsServer.ValueString()
		payload.DnsServer = &v
	}

	if !config.DnsQuery.IsNull() {
		v := config.DnsQuery.ValueString()
		payload.DnsQuery = &v
	}

	if !config.DnsExpectedResult.IsNull() {
		v := config.DnsExpectedResult.ValueString()
		payload.DnsExpectedResult = &v
	}

	if !config.DnsTestValue.IsNull() {
		v := config.DnsTestValue.ValueString()
		payload.DnsTestValue = &v
	}

	if !config.Port.IsNull() {
		v := int(config.Port.ValueInt64())
		payload.Port = &v
	}

	if !config.IpVersion.IsNull() {
		v := config.IpVersion.ValueString()
		payload.IpVersion = &v
	}

	if !config.DatabaseName.IsNull() {
		v := config.DatabaseName.ValueString()
		payload.DatabaseName = &v
	}

	if !config.NetworkAddress.IsNull() {
		v := config.NetworkAddress.ValueString()
		payload.NetworkAddress = &v
	}

	if !config.ImapSecureConnection.IsNull() {
		v := config.ImapSecureConnection.ValueBool()
		payload.ImapSecureConnection = &v
	}

	if !config.SftpAction.IsNull() {
		v := config.SftpAction.ValueString()
		payload.SftpAction = &v
	}

	if !config.SftpActionPath.IsNull() {
		v := config.SftpActionPath.ValueString()
		payload.SftpActionPath = &v
	}

	if !config.HttpMethod.IsNull() {
		v := config.HttpMethod.ValueString()
		payload.HttpMethod = &v
	}

	if !config.HttpVersion.IsNull() {
		v := config.HttpVersion.ValueString()
		payload.HttpVersion = &v
	}

	if !config.TlsVersion.IsNull() {
		v := config.TlsVersion.ValueString()
		payload.TlsVersion = &v
	}

	if !config.RequestBody.IsNull() {
		v := config.RequestBody.ValueString()
		payload.RequestBody = &v
	}

	if !config.Url.IsNull() {
		v := config.Url.ValueString()
		payload.Url = &v
	}

	if !config.BrowserType.IsNull() {
		v := config.BrowserType.ValueString()
		payload.BrowserType = &v
	}

	if config.BrowserWindowDimensions != nil {
		payload.BrowserWindowDimensions = convertBrowserWindowDimensions(*config.BrowserWindowDimensions)
	}

	if !config.ConcurrentUnconfirmedErrorThreshold.IsNull() {
		v := int(config.ConcurrentUnconfirmedErrorThreshold.ValueInt64())
		payload.ConcurrentUnconfirmedErrorThreshold = &v
	}

	if !config.ConcurrentConfirmedErrorThreshold.IsNull() {
		v := int(config.ConcurrentConfirmedErrorThreshold.ValueInt64())
		payload.ConcurrentConfirmedErrorThreshold = &v
	}

	if !config.UseW3CTotalTime.IsNull() {
		val := config.UseW3CTotalTime.ValueBool()
		payload.UseW3CTotalTime = &val
	}

	if !config.PostmanCollectionJson.IsNull() {
		v := config.PostmanCollectionJson.ValueString()
		payload.PostmanCollectionJson = &v
	}

	if config.ErrorConditions != nil {
		var convEC []jsonmodels.ErrorCondition
		for _, ec := range *config.ErrorConditions {
			var percentagePtr *string
			if !ec.Percentage.IsNull() {
				val := ec.Percentage.ValueString()
				percentagePtr = &val
			}

			var levelPtr *string
			if !ec.Level.IsNull() {
				val := ec.Level.ValueString()
				levelPtr = &val
			}

			var matchTypePtr *string
			if !ec.MatchType.IsNull() {
				val := ec.MatchType.ValueString()
				matchTypePtr = &val
			}

			var effectPtr *string
			if !ec.Effect.IsNull() {
				val := ec.Effect.ValueString()
				effectPtr = &val
			}

			convEC = append(convEC, jsonmodels.ErrorCondition{
				ErrorConditionType: ec.ErrorConditionType.ValueString(),
				Value:              ec.Value.ValueString(),
				Percentage:         percentagePtr,
				Level:              levelPtr,
				MatchType:          matchTypePtr,
				Effect:             effectPtr,
			})
		}
		payload.ErrorConditions = &convEC
	}

	return payload
}

func convertStringList(list types.List) []string {
	var result []string
	if list.IsNull() || list.IsUnknown() {
		return result
	}

	listElements := list.Elements()
	for _, elem := range listElements {
		if stringVal, ok := elem.(types.String); ok {
			result = append(result, stringVal.ValueString())
		}
	}
	return result
}

func convertSelectedCheckpoints(sc tfsdkmodels.SelectedCheckpointsModel) jsonmodels.SelectedCheckpoints {
	var selectedCheckpoints = jsonmodels.SelectedCheckpoints{}
	var checkpoints []int
	if !sc.Checkpoints.IsNull() {
		elements := sc.Checkpoints.Elements()
		for _, elem := range elements {
			if intVal, ok := elem.(types.Int64); ok {
				checkpoints = append(checkpoints, int(intVal.ValueInt64()))
			}
		}
		selectedCheckpoints.Checkpoints = &checkpoints
	}

	var regions []int
	if !sc.Regions.IsNull() {
		elements := sc.Regions.Elements()
		for _, elem := range elements {
			if intVal, ok := elem.(types.Int64); ok {
				regions = append(regions, int(intVal.ValueInt64()))
			}
		}
		selectedCheckpoints.Regions = &regions
	}

	var exclude []int
	if !sc.ExcludeLocations.IsNull() {
		elements := sc.ExcludeLocations.Elements()
		for _, elem := range elements {
			if intVal, ok := elem.(types.Int64); ok {
				exclude = append(exclude, int(intVal.ValueInt64()))
			}
		}
		selectedCheckpoints.ExcludeLocations = &exclude
	}

	return selectedCheckpoints
}

func convertThrottlingOptions(to tfsdkmodels.ThrottlingOptionsModel) *jsonmodels.ThrottlingOptions {
	return &jsonmodels.ThrottlingOptions{
		ThrottlingType:      to.ThrottlingType.ValueString(),
		ThrottlingValue:     func() *string { v := to.ThrottlingValue.ValueString(); return &v }(),
		ThrottlingSpeedUp:   func() *int { v := int(to.ThrottlingSpeedUp.ValueInt64()); return &v }(),
		ThrottlingSpeedDown: func() *int { v := int(to.ThrottlingSpeedDown.ValueInt64()); return &v }(),
		ThrottlingLatency:   func() *int { v := int(to.ThrottlingLatency.ValueInt64()); return &v }(),
	}
}

func convertBrowserWindowDimensions(bwd tfsdkmodels.BrowserWindowDimensionsModel) *jsonmodels.BrowserWindowDimensions {
	return &jsonmodels.BrowserWindowDimensions{
		IsMobile:     bwd.IsMobile.ValueBool(),
		Width:        int(bwd.Width.ValueInt64()),
		Height:       int(bwd.Height.ValueInt64()),
		PixelRatio:   int(bwd.PixelRatio.ValueInt64()),
		MobileDevice: bwd.MobileDevice.ValueString(),
	}
}
