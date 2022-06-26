package responses

type Project struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ProjectInternalID             string `json:"ProjectInternalID"`
			ProjectExternalID             string `json:"ProjectExternalID"`
			ProjectDescription            string `json:"ProjectDescription"`
			ProjectLangBsdDescription     string `json:"ProjectLangBsdDescription"`
			ProjectProfileCode            string `json:"ProjectProfileCode"`
			CompanyCode                   string `json:"CompanyCode"`
			ControllingArea               string `json:"ControllingArea"`
			FunctionalArea                string `json:"FunctionalArea"`
			ProfitCenter                  string `json:"ProfitCenter"`
			PlannedStartDate              string `json:"PlannedStartDate"`
			PlannedEndDate                string `json:"PlannedEndDate"`
			WorkCenterLocation            string `json:"WorkCenterLocation"`
			ResponsiblePerson             string `json:"ResponsiblePerson"`
			ResponsiblePersonName         string `json:"ResponsiblePersonName"`
			ApplicantCode                 string `json:"ApplicantCode"`
			ApplicantName                 string `json:"ApplicantName"`
			CreationDate                  string `json:"CreationDate"`
			LastChangeDate                string `json:"LastChangeDate"`
			BasicDatesLastScheduledDate   string `json:"BasicDatesLastScheduledDate"`
			FcstdDatesLastScheduledDate   string `json:"FcstdDatesLastScheduledDate"`
			FactoryCalendar               string `json:"FactoryCalendar"`
			SchedulingDurationUnit        string `json:"SchedulingDurationUnit"`
			ForecastedStartDate           string `json:"ForecastedStartDate"`
			ForecastedEndDate             string `json:"ForecastedEndDate"`
			BusinessArea                  string `json:"BusinessArea"`
			Plant                         string `json:"Plant"`
			Currency                      string `json:"Currency"`
			BudgetProfile                 string `json:"BudgetProfile"`
			PlanningProfile               string `json:"PlanningProfile"`
			InvestmentProfile             string `json:"InvestmentProfile"`
			ProjInterestCalcProfile       string `json:"ProjInterestCalcProfile"`
			ResultAnalysisInternalID      string `json:"ResultAnalysisInternalID"`
			NetworkProfile                string `json:"NetworkProfile"`
			WBSSchedulingProfile          string `json:"WBSSchedulingProfile"`
			PlanningMethForProjBasicDate  string `json:"PlanningMethForProjBasicDate"`
			PlanningMethForProjFcstdDate  string `json:"PlanningMethForProjFcstdDate"`
			NetworkAssignmentType         string `json:"NetworkAssignmentType"`
			WBSIsStatisticalWBSElement    bool   `json:"WBSIsStatisticalWBSElement"`
			WBSIsMarkedForIntegratedPlng  bool   `json:"WBSIsMarkedForIntegratedPlng"`
			ProjectHasOwnStock            bool   `json:"ProjectHasOwnStock"`
			InventorySpecialStockValnType string `json:"InventorySpecialStockValnType"`
			WBSIsMarkedForAutomReqmtGrpg  bool   `json:"WBSIsMarkedForAutomReqmtGrpg"`
			SalesOrganization             string `json:"SalesOrganization"`
			DistributionChannel           string `json:"DistributionChannel"`
			Language                      string `json:"Language"`
			WBSElementExternalID          string `json:"WBSElementExternalID"`
			Division                      string `json:"Division"`
			JointVenture                  string `json:"JointVenture"`
			JointVentureEquityType        string `json:"JointVentureEquityType"`
			JointVentureObjectType        string `json:"JointVentureObjectType"`
			StatusProfile                 string `json:"StatusProfile"`
			WBSStatusProfile              string `json:"WBSStatusProfile"`
			SimulationProfile             string `json:"SimulationProfile"`
			SchedulingScenario            string `json:"SchedulingScenario"`
			DistributionProfile           string `json:"DistributionProfile"`
			PartnerDeterminationProcedure string `json:"PartnerDeterminationProcedure"`
			FreeDefinedTableFieldSemantic string `json:"FreeDefinedTableFieldSemantic"`
			FreeDefinedAttribute01        string `json:"FreeDefinedAttribute01"`
			FreeDefinedAttribute02        string `json:"FreeDefinedAttribute02"`
			FreeDefinedAttribute03        string `json:"FreeDefinedAttribute03"`
			FreeDefinedAttribute04        string `json:"FreeDefinedAttribute04"`
			FreeDefinedQuantity1          string `json:"FreeDefinedQuantity1"`
			FreeDefinedQuantity1Unit      string `json:"FreeDefinedQuantity1Unit"`
			FreeDefinedQuantity2          string `json:"FreeDefinedQuantity2"`
			FreeDefinedQuantity2Unit      string `json:"FreeDefinedQuantity2Unit"`
			FreeDefinedAmount1            string `json:"FreeDefinedAmount1"`
			FreeDefinedAmount1Currency    string `json:"FreeDefinedAmount1Currency"`
			FreeDefinedAmount2            string `json:"FreeDefinedAmount2"`
			FreeDefinedAmount2Currency    string `json:"FreeDefinedAmount2Currency"`
			FreeDefinedDate1              string `json:"FreeDefinedDate1"`
			FreeDefinedDate2              string `json:"FreeDefinedDate2"`
			FreeDefinedIndicator1         bool   `json:"FreeDefinedIndicator1"`
			FreeDefinedIndicator2         bool   `json:"FreeDefinedIndicator2"`
			ToWBSElement                  struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_WBSElement"`
		} `json:"results"`
	} `json:"d"`
}
