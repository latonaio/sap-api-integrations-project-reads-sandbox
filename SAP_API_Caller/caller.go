package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-project-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetProject(projectInternalID, wBSElementInternalID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Project":
			func() {
				c.Project(projectInternalID)
				wg.Done()
			}()

		case "WBSElement":
			func() {
				c.WBSElement(wBSElementInternalID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Project(projectInternalID string) {
	projectData, err := c.callProjectSrvAPIRequirementProject("Project", projectInternalID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(projectData)

	wBSElementData, err := c.callToWBSElement(projectData[0].ToWBSElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(wBSElementData)

}

func (c *SAPAPICaller) callProjectSrvAPIRequirementProject(api, projectInternalID string) ([]sap_api_output_formatter.Project, error) {
	url := strings.Join([]string{c.baseURL, "API_PROJECT_V2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithProject(req, projectInternalID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProject(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToWBSElement(url string) ([]sap_api_output_formatter.ToWBSElement, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToWBSElement(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) WBSElement(wBSElementInternalID string) {
	data, err := c.callProjectSrvAPIRequirementWBSElement("WBSElement", wBSElementInternalID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callProjectSrvAPIRequirementWBSElement(api, wBSElementInternalID string) ([]sap_api_output_formatter.WBSElement, error) {
	url := strings.Join([]string{c.baseURL, "API_PROJECT_V2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithWBSElement(req, wBSElementInternalID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToWBSElement(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithProject(req *http.Request, projectInternalID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ProjectInternalID eq '%s'", projectInternalID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithWBSElement(req *http.Request, wBSElementInternalID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("WBSElementInternalID eq '%s'", wBSElementInternalID))
	req.URL.RawQuery = params.Encode()
}
