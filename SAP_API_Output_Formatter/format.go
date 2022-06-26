package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-project-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToProject(raw []byte, l *logger.Logger) ([]Project, error) {
	pm := &responses.Project{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Project. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	project := make([]Project, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		project = append(project, Project{
			ProjectInternalID:             data.ProjectInternalID,
			ProjectExternalID:             data.ProjectExternalID,
			ProjectDescription:            data.ProjectDescription,
			ProjectLangBsdDescription:     data.ProjectLangBsdDescription,
			ProjectProfileCode:            data.ProjectProfileCode,
			CompanyCode:                   data.CompanyCode,
			ControllingArea:               data.ControllingArea,
			FunctionalArea:                data.FunctionalArea,
			ProfitCenter:                  data.ProfitCenter,
			PlannedStartDate:              data.PlannedStartDate,
			PlannedEndDate:                data.PlannedEndDate,
			WorkCenterLocation:            data.WorkCenterLocation,
			ResponsiblePerson:             data.ResponsiblePerson,
			ResponsiblePersonName:         data.ResponsiblePersonName,
			ApplicantCode:                 data.ApplicantCode,
			ApplicantName:                 data.ApplicantName,
			CreationDate:                  data.CreationDate,
			LastChangeDate:                data.LastChangeDate,
			BasicDatesLastScheduledDate:   data.BasicDatesLastScheduledDate,
			FcstdDatesLastScheduledDate:   data.FcstdDatesLastScheduledDate,
			FactoryCalendar:               data.FactoryCalendar,
			SchedulingDurationUnit:        data.SchedulingDurationUnit,
			ForecastedStartDate:           data.ForecastedStartDate,
			ForecastedEndDate:             data.ForecastedEndDate,
			BusinessArea:                  data.BusinessArea,
			Plant:                         data.Plant,
			Currency:                      data.Currency,
			BudgetProfile:                 data.BudgetProfile,
			PlanningProfile:               data.PlanningProfile,
			InvestmentProfile:             data.InvestmentProfile,
			ProjInterestCalcProfile:       data.ProjInterestCalcProfile,
			ResultAnalysisInternalID:      data.ResultAnalysisInternalID,
			NetworkProfile:                data.NetworkProfile,
			WBSSchedulingProfile:          data.WBSSchedulingProfile,
			PlanningMethForProjBasicDate:  data.PlanningMethForProjBasicDate,
			PlanningMethForProjFcstdDate:  data.PlanningMethForProjFcstdDate,
			NetworkAssignmentType:         data.NetworkAssignmentType,
			WBSIsStatisticalWBSElement:    data.WBSIsStatisticalWBSElement,
			WBSIsMarkedForIntegratedPlng:  data.WBSIsMarkedForIntegratedPlng,
			ProjectHasOwnStock:            data.ProjectHasOwnStock,
			InventorySpecialStockValnType: data.InventorySpecialStockValnType,
			WBSIsMarkedForAutomReqmtGrpg:  data.WBSIsMarkedForAutomReqmtGrpg,
			SalesOrganization:             data.SalesOrganization,
			DistributionChannel:           data.DistributionChannel,
			Language:                      data.Language,
			WBSElementExternalID:          data.WBSElementExternalID,
			Division:                      data.Division,
			JointVenture:                  data.JointVenture,
			JointVentureEquityType:        data.JointVentureEquityType,
			JointVentureObjectType:        data.JointVentureObjectType,
			StatusProfile:                 data.StatusProfile,
			WBSStatusProfile:              data.WBSStatusProfile,
			SimulationProfile:             data.SimulationProfile,
			SchedulingScenario:            data.SchedulingScenario,
			DistributionProfile:           data.DistributionProfile,
			PartnerDeterminationProcedure: data.PartnerDeterminationProcedure,
			FreeDefinedTableFieldSemantic: data.FreeDefinedTableFieldSemantic,
			FreeDefinedAttribute01:        data.FreeDefinedAttribute01,
			FreeDefinedAttribute02:        data.FreeDefinedAttribute02,
			FreeDefinedAttribute03:        data.FreeDefinedAttribute03,
			FreeDefinedAttribute04:        data.FreeDefinedAttribute04,
			FreeDefinedQuantity1:          data.FreeDefinedQuantity1,
			FreeDefinedQuantity1Unit:      data.FreeDefinedQuantity1Unit,
			FreeDefinedQuantity2:          data.FreeDefinedQuantity2,
			FreeDefinedQuantity2Unit:      data.FreeDefinedQuantity2Unit,
			FreeDefinedAmount1:            data.FreeDefinedAmount1,
			FreeDefinedAmount1Currency:    data.FreeDefinedAmount1Currency,
			FreeDefinedAmount2:            data.FreeDefinedAmount2,
			FreeDefinedAmount2Currency:    data.FreeDefinedAmount2Currency,
			FreeDefinedDate1:              data.FreeDefinedDate1,
			FreeDefinedDate2:              data.FreeDefinedDate2,
			FreeDefinedIndicator1:         data.FreeDefinedIndicator1,
			FreeDefinedIndicator2:         data.FreeDefinedIndicator2,
			ToWBSElement:                  data.ToWBSElement.Deferred.URI,
		})
	}

	return project, nil
}

func ConvertToToWBSElement(raw []byte, l *logger.Logger) ([]ToWBSElement, error) {
	pm := &responses.ToWBSElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToWBSElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toWBSElement := make([]ToWBSElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toWBSElement = append(toWBSElement, ToWBSElement{
			WBSElementInternalID:           data.WBSElementInternalID,
			WBSElementExternalID:           data.WBSElementExternalID,
			WBSElementShortID:              data.WBSElementShortID,
			WBSDescription:                 data.WBSDescription,
			WBSElementLangBsdDescription:   data.WBSElementLangBsdDescription,
			ResponsiblePerson:              data.ResponsiblePerson,
			ResponsiblePersonName:          data.ResponsiblePersonName,
			ApplicantCode:                  data.ApplicantCode,
			ApplicantName:                  data.ApplicantName,
			CompanyCode:                    data.CompanyCode,
			BusinessArea:                   data.BusinessArea,
			ControllingArea:                data.ControllingArea,
			FunctionalArea:                 data.FunctionalArea,
			ProfitCenter:                   data.ProfitCenter,
			ResponsibleCostCenter:          data.ResponsibleCostCenter,
			Plant:                          data.Plant,
			FreeDefinedTableFieldSemantic:  data.FreeDefinedTableFieldSemantic,
			FactoryCalendar:                data.FactoryCalendar,
			PriorityCode:                   data.PriorityCode,
			Currency:                       data.Currency,
			CostingSheet:                   data.CostingSheet,
			CostCenter:                     data.CostCenter,
			RequestingCostCenter:           data.RequestingCostCenter,
			ProjectInternalID:              data.ProjectInternalID,
			WBSElementIsBillingElement:     data.WBSElementIsBillingElement,
			InvestmentProfile:              data.InvestmentProfile,
			WBSIsStatisticalWBSElement:     data.WBSIsStatisticalWBSElement,
			WBSIsAccountAssignmentElement:  data.WBSIsAccountAssignmentElement,
			ProjectType:                    data.ProjectType,
			WBSElementIsPlanningElement:    data.WBSElementIsPlanningElement,
			WorkCenterLocation:             data.WorkCenterLocation,
			ResultAnalysisInternalID:       data.ResultAnalysisInternalID,
			TaxJurisdiction:                data.TaxJurisdiction,
			FunctionalLocation:             data.FunctionalLocation,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			WBSIsMarkedForIntegratedPlng:   data.WBSIsMarkedForIntegratedPlng,
			Equipment:                      data.Equipment,
			ProjectObjectChangeNumber:      data.ProjectObjectChangeNumber,
			WBSElementHierarchyLevel:       data.WBSElementHierarchyLevel,
			ReferenceElement:               data.ReferenceElement,
			ProjInterestCalcProfile:        data.ProjInterestCalcProfile,
			Language:                       data.Language,
			IsMarkedForDeletion:            data.IsMarkedForDeletion,
			BasicStartDate:                 data.BasicStartDate,
			ForecastedStartDate:            data.ForecastedStartDate,
			ActualStartDate:                data.ActualStartDate,
			BasicEndDate:                   data.BasicEndDate,
			ForecastedEndDate:              data.ForecastedEndDate,
			ActualEndDate:                  data.ActualEndDate,
			BasicDuration:                  data.BasicDuration,
			BasicDurationUnit:              data.BasicDurationUnit,
			ForecastedDuration:             data.ForecastedDuration,
			ForecastedDurationUnit:         data.ForecastedDurationUnit,
			ActualDuration:                 data.ActualDuration,
			ActualDurationUnit:             data.ActualDurationUnit,
			SchedldBasicEarliestStartDate:  data.SchedldBasicEarliestStartDate,
			ScheduledBasicLatestEndDate:    data.ScheduledBasicLatestEndDate,
			SchedldFcstdEarliestStartDate:  data.SchedldFcstdEarliestStartDate,
			LatestSchedldFcstdEndDate:      data.LatestSchedldFcstdEndDate,
			TentativeActualStartDate:       data.TentativeActualStartDate,
			TentativeActualEndDate:         data.TentativeActualEndDate,
			JointVenture:                   data.JointVenture,
			JointVentureEquityType:         data.JointVentureEquityType,
			JntVntrProjectType:             data.JntVntrProjectType,
			SubProject:                     data.SubProject,
			InvestmentReason:               data.InvestmentReason,
			InvestmentScale:                data.InvestmentScale,
			EnvironmentalInvestmentReason:  data.EnvironmentalInvestmentReason,
			RequestingCompanyCode:          data.RequestingCompanyCode,
			NetworkAssignmentType:          data.NetworkAssignmentType,
			CostObject:                     data.CostObject,
			WBSElementParentInternalID:     data.WBSElementParentInternalID,
			WBSElementChildInternalID:      data.WBSElementChildInternalID,
			LeftSiblingWBSElmntInternalID:  data.LeftSiblingWBSElmntInternalID,
			RightSiblingWBSElmntInternalID: data.RightSiblingWBSElmntInternalID,
			FreeDefinedAttribute01:         data.FreeDefinedAttribute01,
			FreeDefinedAttribute02:         data.FreeDefinedAttribute02,
			FreeDefinedAttribute03:         data.FreeDefinedAttribute03,
			FreeDefinedAttribute04:         data.FreeDefinedAttribute04,
			FreeDefinedQuantity1:           data.FreeDefinedQuantity1,
			FreeDefinedQuantity1Unit:       data.FreeDefinedQuantity1Unit,
			FreeDefinedQuantity2:           data.FreeDefinedQuantity2,
			FreeDefinedQuantity2Unit:       data.FreeDefinedQuantity2Unit,
			FreeDefinedAmount1:             data.FreeDefinedAmount1,
			FreeDefinedAmount1Currency:     data.FreeDefinedAmount1Currency,
			FreeDefinedAmount2:             data.FreeDefinedAmount2,
			FreeDefinedAmount2Currency:     data.FreeDefinedAmount2Currency,
			FreeDefinedDate1:               data.FreeDefinedDate1,
			FreeDefinedDate2:               data.FreeDefinedDate2,
			FreeDefinedIndicator1:          data.FreeDefinedIndicator1,
			FreeDefinedIndicator2:          data.FreeDefinedIndicator2,
		})
	}

	return toWBSElement, nil
}

func ConvertToWBSElement(raw []byte, l *logger.Logger) ([]WBSElement, error) {
	pm := &responses.WBSElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to WBSElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	wBSElement := make([]WBSElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		wBSElement = append(wBSElement, WBSElement{
			WBSElementInternalID:           data.WBSElementInternalID,
			WBSElementExternalID:           data.WBSElementExternalID,
			WBSElementShortID:              data.WBSElementShortID,
			WBSDescription:                 data.WBSDescription,
			WBSElementLangBsdDescription:   data.WBSElementLangBsdDescription,
			ResponsiblePerson:              data.ResponsiblePerson,
			ResponsiblePersonName:          data.ResponsiblePersonName,
			ApplicantCode:                  data.ApplicantCode,
			ApplicantName:                  data.ApplicantName,
			CompanyCode:                    data.CompanyCode,
			BusinessArea:                   data.BusinessArea,
			ControllingArea:                data.ControllingArea,
			FunctionalArea:                 data.FunctionalArea,
			ProfitCenter:                   data.ProfitCenter,
			ResponsibleCostCenter:          data.ResponsibleCostCenter,
			Plant:                          data.Plant,
			FreeDefinedTableFieldSemantic:  data.FreeDefinedTableFieldSemantic,
			FactoryCalendar:                data.FactoryCalendar,
			PriorityCode:                   data.PriorityCode,
			Currency:                       data.Currency,
			CostingSheet:                   data.CostingSheet,
			CostCenter:                     data.CostCenter,
			RequestingCostCenter:           data.RequestingCostCenter,
			ProjectInternalID:              data.ProjectInternalID,
			WBSElementIsBillingElement:     data.WBSElementIsBillingElement,
			InvestmentProfile:              data.InvestmentProfile,
			WBSIsStatisticalWBSElement:     data.WBSIsStatisticalWBSElement,
			WBSIsAccountAssignmentElement:  data.WBSIsAccountAssignmentElement,
			ProjectType:                    data.ProjectType,
			WBSElementIsPlanningElement:    data.WBSElementIsPlanningElement,
			WorkCenterLocation:             data.WorkCenterLocation,
			ResultAnalysisInternalID:       data.ResultAnalysisInternalID,
			TaxJurisdiction:                data.TaxJurisdiction,
			FunctionalLocation:             data.FunctionalLocation,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			WBSIsMarkedForIntegratedPlng:   data.WBSIsMarkedForIntegratedPlng,
			Equipment:                      data.Equipment,
			ProjectObjectChangeNumber:      data.ProjectObjectChangeNumber,
			WBSElementHierarchyLevel:       data.WBSElementHierarchyLevel,
			ReferenceElement:               data.ReferenceElement,
			ProjInterestCalcProfile:        data.ProjInterestCalcProfile,
			Language:                       data.Language,
			IsMarkedForDeletion:            data.IsMarkedForDeletion,
			BasicStartDate:                 data.BasicStartDate,
			ForecastedStartDate:            data.ForecastedStartDate,
			ActualStartDate:                data.ActualStartDate,
			BasicEndDate:                   data.BasicEndDate,
			ForecastedEndDate:              data.ForecastedEndDate,
			ActualEndDate:                  data.ActualEndDate,
			BasicDuration:                  data.BasicDuration,
			BasicDurationUnit:              data.BasicDurationUnit,
			ForecastedDuration:             data.ForecastedDuration,
			ForecastedDurationUnit:         data.ForecastedDurationUnit,
			ActualDuration:                 data.ActualDuration,
			ActualDurationUnit:             data.ActualDurationUnit,
			SchedldBasicEarliestStartDate:  data.SchedldBasicEarliestStartDate,
			ScheduledBasicLatestEndDate:    data.ScheduledBasicLatestEndDate,
			SchedldFcstdEarliestStartDate:  data.SchedldFcstdEarliestStartDate,
			LatestSchedldFcstdEndDate:      data.LatestSchedldFcstdEndDate,
			TentativeActualStartDate:       data.TentativeActualStartDate,
			TentativeActualEndDate:         data.TentativeActualEndDate,
			JointVenture:                   data.JointVenture,
			JointVentureEquityType:         data.JointVentureEquityType,
			JntVntrProjectType:             data.JntVntrProjectType,
			SubProject:                     data.SubProject,
			InvestmentReason:               data.InvestmentReason,
			InvestmentScale:                data.InvestmentScale,
			EnvironmentalInvestmentReason:  data.EnvironmentalInvestmentReason,
			RequestingCompanyCode:          data.RequestingCompanyCode,
			NetworkAssignmentType:          data.NetworkAssignmentType,
			CostObject:                     data.CostObject,
			WBSElementParentInternalID:     data.WBSElementParentInternalID,
			WBSElementChildInternalID:      data.WBSElementChildInternalID,
			LeftSiblingWBSElmntInternalID:  data.LeftSiblingWBSElmntInternalID,
			RightSiblingWBSElmntInternalID: data.RightSiblingWBSElmntInternalID,
			FreeDefinedAttribute01:         data.FreeDefinedAttribute01,
			FreeDefinedAttribute02:         data.FreeDefinedAttribute02,
			FreeDefinedAttribute03:         data.FreeDefinedAttribute03,
			FreeDefinedAttribute04:         data.FreeDefinedAttribute04,
			FreeDefinedQuantity1:           data.FreeDefinedQuantity1,
			FreeDefinedQuantity1Unit:       data.FreeDefinedQuantity1Unit,
			FreeDefinedQuantity2:           data.FreeDefinedQuantity2,
			FreeDefinedQuantity2Unit:       data.FreeDefinedQuantity2Unit,
			FreeDefinedAmount1:             data.FreeDefinedAmount1,
			FreeDefinedAmount1Currency:     data.FreeDefinedAmount1Currency,
			FreeDefinedAmount2:             data.FreeDefinedAmount2,
			FreeDefinedAmount2Currency:     data.FreeDefinedAmount2Currency,
			FreeDefinedDate1:               data.FreeDefinedDate1,
			FreeDefinedDate2:               data.FreeDefinedDate2,
			FreeDefinedIndicator1:          data.FreeDefinedIndicator1,
			FreeDefinedIndicator2:          data.FreeDefinedIndicator2,
		})
	}

	return wBSElement, nil
}
