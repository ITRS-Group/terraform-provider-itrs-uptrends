package converters

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	jsonmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/helpers"
	tfsdkmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/provider/models"
)

// UpdateStateConversion maps a json MonitorResponse to a tfsdk MonitorModel.
func UpdateStateConversion(monitor *jsonmodels.MonitorResponse) tfsdkmodels.MonitorModel {
	var state tfsdkmodels.MonitorModel

	state.MonitorGuid = types.StringValue(monitor.MonitorGuid)
	state.Name = types.StringValue(monitor.Name)
	state.MonitorType = types.StringValue(monitor.MonitorType)
	state.GenerateAlert = types.BoolValue(monitor.GenerateAlert)
	state.IsActive = types.BoolValue(monitor.IsActive)

	state.MonitorMode = types.StringValue(monitor.MonitorMode)
	state.UsePrimaryCheckpointsOnly = types.BoolValue(monitor.UsePrimaryCheckpointsOnly)

	if monitor.CheckInterval != nil {
		value := types.Int64Value(int64(*monitor.CheckInterval))
		state.CheckInterval = value
	} else {
		state.CheckInterval = types.Int64Null()
	}

	if monitor.CheckIntervalSeconds != nil {
		value := types.Int64Value(int64(*monitor.CheckIntervalSeconds))
		state.CheckIntervalSeconds = value
	} else {
		state.CheckIntervalSeconds = types.Int64Null()
	}

	if monitor.UseConcurrentMonitoring != nil {
		useConcurrentValue := types.BoolValue(*monitor.UseConcurrentMonitoring)
		state.UseConcurrentMonitoring = useConcurrentValue
	} else {
		state.UseConcurrentMonitoring = types.BoolNull()
	}

	state.Notes = types.StringValue(monitor.Notes)

	if monitor.CustomMetrics != nil {
		var convCM []tfsdkmodels.CustomMetricModel
		for _, cm := range *monitor.CustomMetrics {
			convCM = append(convCM, tfsdkmodels.CustomMetricModel{
				Name:         types.StringValue(cm.Name),
				VariableName: types.StringValue(cm.VariableName),
			})
		}
		state.CustomMetrics = &convCM
	} else {
		state.CustomMetrics = nil
	}

	if len(monitor.CustomFields) > 0 {
		var convCF []tfsdkmodels.CustomFieldModel
		for _, cf := range monitor.CustomFields {
			convCF = append(convCF, tfsdkmodels.CustomFieldModel{
				Name:  types.StringValue(cf.Name),
				Value: types.StringValue(cf.Value),
			})
		}
		state.CustomFields = &convCF
	} else {
		state.CustomFields = &[]tfsdkmodels.CustomFieldModel{}
	}

	state.SelectedCheckpoints = convertSelectedCheckpointsFromJSON(monitor.SelectedCheckpoints)

	if monitor.SelfServiceTransactionScript != nil {
		normalized, err := helpers.NormalizeJSON(*monitor.SelfServiceTransactionScript)
		if err != nil {
			state.SelfServiceTransactionScript = types.StringValue(*monitor.SelfServiceTransactionScript)
		} else {
			state.SelfServiceTransactionScript = types.StringValue(normalized)
		}
	} else {
		state.SelfServiceTransactionScript = types.StringNull()
	}

	if monitor.PostmanCollectionJson != nil {
		normalized, err := helpers.NormalizeJSON(*monitor.PostmanCollectionJson)
		if err != nil {
			state.PostmanCollectionJson = types.StringValue(*monitor.PostmanCollectionJson)
		} else {
			state.PostmanCollectionJson = types.StringValue(normalized)
		}
	} else {
		state.PostmanCollectionJson = types.StringNull()
	}

	if monitor.MultiStepApiTransactionScript != nil {
		normalized, err := helpers.NormalizeJSON(*monitor.MultiStepApiTransactionScript)
		if err != nil {
			state.MultiStepApiTransactionScript = types.StringValue(*monitor.MultiStepApiTransactionScript)
		} else {
			state.MultiStepApiTransactionScript = types.StringValue(normalized)
		}
	} else {
		state.MultiStepApiTransactionScript = types.StringNull()
	}

	if monitor.BlockGoogleAnalytics != nil {
		value := types.BoolValue(*monitor.BlockGoogleAnalytics)
		state.BlockGoogleAnalytics = value
	} else {
		state.BlockGoogleAnalytics = types.BoolNull()
	}

	if monitor.BlockUptrendsRum != nil {
		value := types.BoolValue(*monitor.BlockUptrendsRum)
		state.BlockUptrendsRum = value
	} else {
		state.BlockUptrendsRum = types.BoolNull()
	}

	if monitor.BlockUrls != nil && len(*monitor.BlockUrls) > 0 {
		state.BlockUrls = convertStringSliceToList(*monitor.BlockUrls)
	} else {
		state.BlockUrls = types.ListNull(types.StringType)
	}

	if monitor.RequestHeaders != nil {
		var convRH []tfsdkmodels.RequestHeaderModel
		for _, rh := range *monitor.RequestHeaders {
			convRH = append(convRH, tfsdkmodels.RequestHeaderModel{
				Name:  types.StringValue(rh.Name),
				Value: types.StringValue(rh.Value),
			})
		}
		state.RequestHeaders = &convRH
	} else {
		state.RequestHeaders = nil
	}

	if monitor.PredefinedVariables != nil {
		var convPV []tfsdkmodels.PredefinedVariablesModel
		for _, pv := range *monitor.PredefinedVariables {
			convPV = append(convPV, tfsdkmodels.PredefinedVariablesModel{
				Key:   types.StringValue(pv.Key),
				Value: types.StringValue(pv.Value),
			})
		}
		state.PredefinedVariables = &convPV
	} else {
		state.PredefinedVariables = nil
	}

	if monitor.UserAgent != nil {
		value := types.StringValue(*monitor.UserAgent)
		state.UserAgent = value
	} else {
		state.UserAgent = types.StringNull()
	}

	if monitor.Username != nil {
		value := types.StringValue(*monitor.Username)
		state.Username = value
	} else {
		state.Username = types.StringNull()
	}

	if monitor.Password != nil {
		value := types.StringValue(*monitor.Password)
		state.Password = value
	} else {
		state.Password = types.StringNull()
	}

	if monitor.NameForPhoneAlerts != nil {
		value := types.StringValue(*monitor.NameForPhoneAlerts)
		state.NameForPhoneAlerts = value
	} else {
		state.NameForPhoneAlerts = types.StringNull()
	}

	if monitor.AuthenticationType != nil {
		value := types.StringValue(*monitor.AuthenticationType)
		state.AuthenticationType = value
	} else {
		state.AuthenticationType = types.StringNull()
	}

	if monitor.ThrottlingOptions != nil {
		state.ThrottlingOptions = convertThrottlingOptionsFromJSON(*monitor.ThrottlingOptions)
	} else {
		state.ThrottlingOptions = nil
	}

	if monitor.DnsBypasses != nil {
		var convDB []tfsdkmodels.DnsBypassModel
		for _, db := range *monitor.DnsBypasses {
			convDB = append(convDB, tfsdkmodels.DnsBypassModel{
				Source: types.StringValue(db.Source),
				Target: types.StringValue(db.Target),
			})
		}
		state.DnsBypasses = &convDB
	} else {
		state.DnsBypasses = nil
	}

	if monitor.CertificateName != nil {
		value := types.StringValue(*monitor.CertificateName)
		state.CertificateName = value
	} else {
		state.CertificateName = types.StringNull()
	}

	if monitor.CertificateOrganization != nil {
		value := types.StringValue(*monitor.CertificateOrganization)
		state.CertificateOrganization = value
	} else {
		state.CertificateOrganization = types.StringNull()
	}

	if monitor.CertificateOrganizationalUnit != nil {
		value := types.StringValue(*monitor.CertificateOrganizationalUnit)
		state.CertificateOrganizationalUnit = value
	} else {
		state.CertificateOrganizationalUnit = types.StringNull()
	}

	if monitor.CertificateSerialNumber != nil {
		value := types.StringValue(*monitor.CertificateSerialNumber)
		state.CertificateSerialNumber = value
	} else {
		state.CertificateSerialNumber = types.StringNull()
	}

	if monitor.CertificateFingerprint != nil {
		value := types.StringValue(*monitor.CertificateFingerprint)
		state.CertificateFingerprint = value
	} else {
		state.CertificateFingerprint = types.StringNull()
	}

	if monitor.CertificateIssuerName != nil {
		value := types.StringValue(*monitor.CertificateIssuerName)
		state.CertificateIssuerName = value
	} else {
		state.CertificateIssuerName = types.StringNull()
	}

	if monitor.CertificateIssuerCompanyName != nil {
		value := types.StringValue(*monitor.CertificateIssuerCompanyName)
		state.CertificateIssuerCompanyName = value
	} else {
		state.CertificateIssuerCompanyName = types.StringNull()
	}

	if monitor.CertificateIssuerOrganizationalUnit != nil {
		value := types.StringValue(*monitor.CertificateIssuerOrganizationalUnit)
		state.CertificateIssuerOrganizationalUnit = value
	} else {
		state.CertificateIssuerOrganizationalUnit = types.StringNull()
	}

	if monitor.CertificateExpirationWarningDays != nil {
		value := types.Int64Value(int64(*monitor.CertificateExpirationWarningDays))
		state.CertificateExpirationWarningDays = value
	} else {
		state.CertificateExpirationWarningDays = types.Int64Null()
	}

	if monitor.CheckCertificateErrors != nil {
		value := types.BoolValue(*monitor.CheckCertificateErrors)
		state.CheckCertificateErrors = value
	} else {
		state.CheckCertificateErrors = types.BoolNull()
	}

	if monitor.IgnoreExternalElements != nil {
		value := types.BoolValue(*monitor.IgnoreExternalElements)
		state.IgnoreExternalElements = value
	} else {
		state.IgnoreExternalElements = types.BoolNull()
	}

	if monitor.DomainGroupGuid != nil {
		value := types.StringValue(*monitor.DomainGroupGuid)
		state.DomainGroupGuid = value
	} else {
		state.DomainGroupGuid = types.StringNull()
	}

	if monitor.DomainGroupGuidSpecified != nil {
		value := types.BoolValue(*monitor.DomainGroupGuidSpecified)
		state.DomainGroupGuidSpecified = value
	} else {
		state.DomainGroupGuidSpecified = types.BoolNull()
	}

	if monitor.DnsServer != nil {
		value := types.StringValue(*monitor.DnsServer)
		state.DnsServer = value
	} else {
		state.DnsServer = types.StringNull()
	}

	if monitor.DnsQuery != nil {
		value := types.StringValue(*monitor.DnsQuery)
		state.DnsQuery = value
	} else {
		state.DnsQuery = types.StringNull()
	}

	if monitor.DnsExpectedResult != nil {
		value := types.StringValue(*monitor.DnsExpectedResult)
		state.DnsExpectedResult = value
	} else {
		state.DnsExpectedResult = types.StringNull()
	}

	if monitor.DnsTestValue != nil {
		value := types.StringValue(*monitor.DnsTestValue)
		state.DnsTestValue = value
	} else {
		state.DnsTestValue = types.StringNull()
	}

	if monitor.Port != nil {
		value := types.Int64Value(int64(*monitor.Port))
		state.Port = value
	} else {
		state.Port = types.Int64Null()
	}

	if monitor.IpVersion != nil {
		value := types.StringValue(*monitor.IpVersion)
		state.IpVersion = value
	} else {
		state.IpVersion = types.StringNull()
	}

	if monitor.DatabaseName != nil {
		value := types.StringValue(*monitor.DatabaseName)
		state.DatabaseName = value
	} else {
		state.DatabaseName = types.StringNull()
	}

	if monitor.NetworkAddress != nil {
		value := types.StringValue(*monitor.NetworkAddress)
		state.NetworkAddress = value
	} else {
		state.NetworkAddress = types.StringNull()
	}

	if monitor.ImapSecureConnection != nil {
		value := types.BoolValue(*monitor.ImapSecureConnection)
		state.ImapSecureConnection = value
	} else {
		state.ImapSecureConnection = types.BoolNull()
	}

	if monitor.UseW3CTotalTime != nil {
		value := types.BoolValue(*monitor.UseW3CTotalTime)
		state.UseW3CTotalTime = value
	} else {
		state.UseW3CTotalTime = types.BoolNull()
	}

	if monitor.SftpAction != nil {
		value := types.StringValue(*monitor.SftpAction)
		state.SftpAction = value
	} else {
		state.SftpAction = types.StringNull()
	}

	if monitor.SftpActionPath != nil {
		value := types.StringValue(*monitor.SftpActionPath)
		state.SftpActionPath = value
	} else {
		state.SftpActionPath = types.StringNull()
	}

	if monitor.HttpMethod != nil {
		value := types.StringValue(*monitor.HttpMethod)
		state.HttpMethod = value
	} else {
		state.HttpMethod = types.StringNull()
	}

	if monitor.HttpVersion != nil {
		value := types.StringValue(*monitor.HttpVersion)
		state.HttpVersion = value
	} else {
		state.HttpVersion = types.StringNull()
	}

	if monitor.TlsVersion != nil {
		value := types.StringValue(*monitor.TlsVersion)
		state.TlsVersion = value
	} else {
		state.TlsVersion = types.StringNull()
	}

	if monitor.RequestBody != nil {
		value := types.StringValue(*monitor.RequestBody)
		state.RequestBody = value
	} else {

		state.RequestBody = types.StringNull()
	}
	if monitor.Url != nil {
		value := types.StringValue(*monitor.Url)
		state.Url = value
	} else {
		state.Url = types.StringNull()
	}

	if monitor.BrowserType != nil {
		value := types.StringValue(*monitor.BrowserType)
		state.BrowserType = value
	} else {
		state.BrowserType = types.StringNull()
	}

	if monitor.BrowserWindowDimensions != nil {
		state.BrowserWindowDimensions = convertBrowserWindowDimensionsFromJSON(*monitor.BrowserWindowDimensions)
	} else {
		state.BrowserWindowDimensions = nil
	}

	if monitor.ConcurrentUnconfirmedErrorThreshold != nil {
		value := types.Int64Value(int64(*monitor.ConcurrentUnconfirmedErrorThreshold))
		state.ConcurrentUnconfirmedErrorThreshold = value
	} else {
		state.ConcurrentUnconfirmedErrorThreshold = types.Int64Null()
	}

	if monitor.ConcurrentConfirmedErrorThreshold != nil {
		value := types.Int64Value(int64(*monitor.ConcurrentConfirmedErrorThreshold))
		state.ConcurrentConfirmedErrorThreshold = value
	} else {
		state.ConcurrentConfirmedErrorThreshold = types.Int64Null()
	}

	if monitor.ErrorConditions != nil {
		var convEC []tfsdkmodels.ErrorConditionModel
		for _, ec := range *monitor.ErrorConditions {
			convEC = append(convEC, tfsdkmodels.ErrorConditionModel{
				ErrorConditionType: types.StringValue(ec.ErrorConditionType),
				Value:              types.StringValue(ec.Value),
				Percentage: func() types.String {
					if ec.Percentage != nil {
						return types.StringValue(*ec.Percentage)
					}
					return types.StringNull()
				}(),
				Level: func() types.String {
					if ec.Level != nil {
						return types.StringValue(*ec.Level)
					}
					return types.StringNull()
				}(),
				MatchType: func() types.String {
					if ec.MatchType != nil {
						return types.StringValue(*ec.MatchType)
					}
					return types.StringNull()
				}(),
				Effect: func() types.String {
					if ec.Effect != nil {
						return types.StringValue(*ec.Effect)
					}
					return types.StringNull()
				}(),
			})
		}
		state.ErrorConditions = &convEC
	} else {
		state.ErrorConditions = nil
	}

	if monitor.CreatedDate != "" {
		value := types.StringValue(monitor.CreatedDate)
		state.CreatedDate = value
	} else {
		state.CreatedDate = types.StringNull()
	}

	return state
}

// Helper conversion functions
func convertStringSliceToList(strs []string) types.List {
	if strs == nil || len(strs) == 0 {
		return types.ListNull(types.StringType)
	}
	values := make([]attr.Value, len(strs))
	for i, s := range strs {
		values[i] = types.StringValue(s)
	}
	list, _ := types.ListValue(types.StringType, values)
	return list
}

func convertSelectedCheckpointsFromJSON(sc jsonmodels.SelectedCheckpoints) *tfsdkmodels.SelectedCheckpointsModel {
	model := &tfsdkmodels.SelectedCheckpointsModel{}

	if sc.Checkpoints != nil && len(*sc.Checkpoints) > 0 {
		values := make([]attr.Value, len(*sc.Checkpoints))
		for i, cp := range *sc.Checkpoints {
			values[i] = types.Int64Value(int64(cp))
		}
		list, _ := types.ListValue(types.Int64Type, values)
		model.Checkpoints = list
	} else {
		model.Checkpoints = types.ListNull(types.Int64Type)
	}

	if sc.Regions != nil && len(*sc.Regions) > 0 {
		values := make([]attr.Value, len(*sc.Regions))
		for i, r := range *sc.Regions {
			values[i] = types.Int64Value(int64(r))
		}
		list, _ := types.ListValue(types.Int64Type, values)
		model.Regions = list
	} else {
		model.Regions = types.ListNull(types.Int64Type)
	}

	if sc.ExcludeLocations != nil && len(*sc.ExcludeLocations) > 0 {
		values := make([]attr.Value, len(*sc.ExcludeLocations))
		for i, el := range *sc.ExcludeLocations {
			values[i] = types.Int64Value(int64(el))
		}
		list, _ := types.ListValue(types.Int64Type, values)
		model.ExcludeLocations = list
	} else {
		model.ExcludeLocations = types.ListNull(types.Int64Type)
	}

	return model
}

func convertThrottlingOptionsFromJSON(to jsonmodels.ThrottlingOptions) *tfsdkmodels.ThrottlingOptionsModel {
	return &tfsdkmodels.ThrottlingOptionsModel{
		ThrottlingType: types.StringValue(to.ThrottlingType),
		ThrottlingValue: func() types.String {
			if to.ThrottlingValue != nil {
				return types.StringValue(*to.ThrottlingValue)
			}
			return types.StringNull()
		}(),
		ThrottlingSpeedUp: func() types.Int64 {
			if to.ThrottlingSpeedUp != nil {
				return types.Int64Value(int64(*to.ThrottlingSpeedUp))
			}
			return types.Int64Null()
		}(),
		ThrottlingSpeedDown: func() types.Int64 {
			if to.ThrottlingSpeedDown != nil {
				return types.Int64Value(int64(*to.ThrottlingSpeedDown))
			}
			return types.Int64Null()
		}(),
		ThrottlingLatency: func() types.Int64 {
			if to.ThrottlingLatency != nil {
				return types.Int64Value(int64(*to.ThrottlingLatency))
			}
			return types.Int64Null()
		}(),
	}
}

func convertBrowserWindowDimensionsFromJSON(bwd jsonmodels.BrowserWindowDimensions) *tfsdkmodels.BrowserWindowDimensionsModel {
	return &tfsdkmodels.BrowserWindowDimensionsModel{
		IsMobile:     types.BoolValue(bwd.IsMobile),
		Width:        types.Int64Value(int64(bwd.Width)),
		Height:       types.Int64Value(int64(bwd.Height)),
		PixelRatio:   types.Int64Value(int64(bwd.PixelRatio)),
		MobileDevice: types.StringValue(bwd.MobileDevice),
	}
}
