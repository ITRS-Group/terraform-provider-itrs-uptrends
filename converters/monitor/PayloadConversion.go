package converters

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/types"
	jsonmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
	tfsdkmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/provider/models"
)

// PayloadConversion converts the tfsdk MonitorModel to the JSON MonitorRequest.
// For optional fields (pointer types), the value is assigned only when not nil and not null.
func PayloadConversion(config tfsdkmodels.MonitorModel) jsonmodels.MonitorRequest {

	payload := jsonmodels.MonitorRequest{
		Name:          config.Name.ValueString(),
		MonitorType:   config.MonitorType.ValueString(),
		GenerateAlert: config.GenerateAlert.ValueBool(),
		IsActive:      config.IsActive.ValueBool(),
		MonitorMode:   config.MonitorMode.ValueString(),
	}

	if !config.UsePrimaryCheckpointsOnly.IsNull() {
		val := config.UsePrimaryCheckpointsOnly.ValueBool()
		payload.UsePrimaryCheckpointsOnly = &val
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

	if config.MsaSteps != nil {
		payload.MsaSteps = convertMsaSteps(*config.MsaSteps)
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

func convertMsaSteps(steps []tfsdkmodels.MsaStepModel) *[]jsonmodels.MsaStep {
	result := make([]jsonmodels.MsaStep, 0, len(steps))
	for _, s := range steps {
		step := jsonmodels.MsaStep{}

		if !s.Name.IsNull() && !s.Name.IsUnknown() {
			v := s.Name.ValueString()
			step.Name = &v
		}
		if !s.StepType.IsNull() && !s.StepType.IsUnknown() {
			v := s.StepType.ValueString()
			step.StepType = &v
		}
		if !s.Url.IsNull() && !s.Url.IsUnknown() {
			v := s.Url.ValueString()
			step.Url = &v
		}
		if !s.Method.IsNull() && !s.Method.IsUnknown() {
			v := s.Method.ValueString()
			step.Method = &v
		}
		if !s.Body.IsNull() && !s.Body.IsUnknown() {
			v := s.Body.ValueString()
			step.Body = &v
		}
		if !s.BodyType.IsNull() && !s.BodyType.IsUnknown() {
			v := s.BodyType.ValueString()
			step.BodyType = &v
		}
		if !s.VaultFileId.IsNull() && !s.VaultFileId.IsUnknown() {
			v := s.VaultFileId.ValueString()
			step.VaultFileId = &v
		}
		if !s.ClientCertificateVaultItem.IsNull() && !s.ClientCertificateVaultItem.IsUnknown() {
			v := s.ClientCertificateVaultItem.ValueString()
			step.ClientCertificateVaultItem = &v
		}
		if !s.PreRequestScript.IsNull() && !s.PreRequestScript.IsUnknown() {
			v := s.PreRequestScript.ValueString()
			step.PreRequestScript = &v
		}
		if !s.PostResponseScript.IsNull() && !s.PostResponseScript.IsUnknown() {
			v := s.PostResponseScript.ValueString()
			step.PostResponseScript = &v
		}
		if !s.CalculatedContentType.IsNull() && !s.CalculatedContentType.IsUnknown() {
			v := s.CalculatedContentType.ValueString()
			step.CalculatedContentType = &v
		}
		if !s.MsaStepHttpVersion.IsNull() && !s.MsaStepHttpVersion.IsUnknown() {
			v := s.MsaStepHttpVersion.ValueString()
			step.MsaStepHttpVersion = &v
		}
		if !s.UseFixedClientCertificate.IsNull() && !s.UseFixedClientCertificate.IsUnknown() {
			v := s.UseFixedClientCertificate.ValueBool()
			step.UseFixedClientCertificate = &v
		}
		if !s.IgnoreCertificateErrors.IsNull() && !s.IgnoreCertificateErrors.IsUnknown() {
			v := s.IgnoreCertificateErrors.ValueBool()
			step.IgnoreCertificateErrors = &v
		}
		if !s.RetryUntilSuccessful.IsNull() && !s.RetryUntilSuccessful.IsUnknown() {
			v := s.RetryUntilSuccessful.ValueBool()
			step.RetryUntilSuccessful = &v
		}
		if !s.Delay.IsNull() && !s.Delay.IsUnknown() {
			v := int(s.Delay.ValueInt64())
			step.Delay = &v
		}
		if !s.MaxAttempts.IsNull() && !s.MaxAttempts.IsUnknown() {
			v := int(s.MaxAttempts.ValueInt64())
			step.MaxAttempts = &v
		}
		if !s.RetryWaitMilliseconds.IsNull() && !s.RetryWaitMilliseconds.IsUnknown() {
			v := int(s.RetryWaitMilliseconds.ValueInt64())
			step.RetryWaitMilliseconds = &v
		}
		if !s.AllowedTlsVersions.IsNull() && !s.AllowedTlsVersions.IsUnknown() {
			tls := convertStringList(s.AllowedTlsVersions)
			step.AllowedTlsVersions = &tls
		}

		if s.MultiPartForm != nil {
			var parts []jsonmodels.MsaFormPart
			for _, p := range *s.MultiPartForm {
				parts = append(parts, jsonmodels.MsaFormPart{
					Type:  p.Type.ValueString(),
					Key:   p.Key.ValueString(),
					Value: p.Value.ValueString(),
				})
			}
			step.MultiPartForm = &parts
		}

		if s.RequestHeaders != nil {
			var headers []jsonmodels.MsaHttpHeaderValue
			for _, h := range *s.RequestHeaders {
				headers = append(headers, jsonmodels.MsaHttpHeaderValue{
					Key:   h.Key.ValueString(),
					Value: h.Value.ValueString(),
				})
			}
			step.RequestHeaders = &headers
		}

		if s.Assertions != nil {
			var assertions []jsonmodels.MsaAssertion
			for _, a := range *s.Assertions {
				assertions = append(assertions, jsonmodels.MsaAssertion{
					Source:      a.Source.ValueString(),
					Property:    a.Property.ValueString(),
					Comparison:  a.Comparison.ValueString(),
					TargetValue: a.TargetValue.ValueString(),
				})
			}
			step.Assertions = &assertions
		}

		if s.Variables != nil {
			step.Variables = convertMsaVariables(*s.Variables)
		}

		if s.Authentication != nil {
			step.Authentication = convertMsaAuthentication(*s.Authentication)
		}

		result = append(result, step)
	}
	return &result
}

func convertMsaVariables(vars []tfsdkmodels.MsaVariableModel) *[]jsonmodels.MsaVariableDefinition {
	result := make([]jsonmodels.MsaVariableDefinition, 0, len(vars))
	for _, v := range vars {
		def := jsonmodels.MsaVariableDefinition{
			Source:   v.Source.ValueString(),
			Property: v.Property.ValueString(),
			Name:     v.Name.ValueString(),
		}
		if !v.Arguments.IsNull() && !v.Arguments.IsUnknown() && v.Arguments.ValueString() != "" {
			var args []jsonmodels.MsaVariableDefinition
			if err := json.Unmarshal([]byte(v.Arguments.ValueString()), &args); err == nil {
				def.Arguments = args
			}
		}
		result = append(result, def)
	}
	return &result
}

func convertMsaAuthentication(a tfsdkmodels.MsaAuthenticationModel) *jsonmodels.MsaAuthenticationInfo {
	info := &jsonmodels.MsaAuthenticationInfo{
		AuthenticationType: a.AuthenticationType.ValueString(),
	}
	if !a.Id.IsNull() {
		v := a.Id.ValueString()
		info.Id = &v
	}
	if !a.UserName.IsNull() {
		v := a.UserName.ValueString()
		info.UserName = &v
	}
	if !a.Password.IsNull() {
		v := a.Password.ValueString()
		info.Password = &v
		specified := true
		info.PasswordSpecified = &specified
	}
	return info
}
