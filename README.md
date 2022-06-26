# sap-api-integrations-project-reads
sap-api-integrations-project-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で プロジェクトデータを取得するマイクロサービスです。    
sap-api-integrations-project-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-project-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_PROJECT_V2_0001/overview

## 動作環境  
sap-api-integrations-project-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-project-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-project-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PROJECT_V2_0001/overview
* APIサービス名(=baseURL): API_PROJECT_V2

## 本レポジトリ に 含まれる API名
sap-api-integrations-project-reads には、次の API をコールするためのリソースが含まれています。  

* Project（SAP プロジェクト - プロジェクト）※プロジェクトの詳細データを取得するために、ToWBSElement、と合わせて利用されます。  
* WBSElement（SAP プロジェクト - WBS要素）
* ToWBSElement（SAP プロジェクト - WBS要素）

## API への 値入力条件 の 初期値
sap-api-integrations-project-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ProjectInternalID,
* inoutSDC.WBSElement.WBSElementInternalID

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Project" が指定されています。

```
	"api_schema": "sap.s4.beh.salesorder.v1.SalesOrder.Created.v1",
	"accepter": ["Project"],
	"sales_order": "1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.salesorder.v1.SalesOrder.Created.v1",
	"accepter": ["All"],
	"sales_order": "1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetProject(projectInternalID, WBSElementInternalID string, accepter []string) {
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
				c.WBSElement(WBSElementInternalID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

```
## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP プロジェクト の プロジェクトデータ が取得された結果の JSON の例です。  
以下の項目のうち、"XXXXX" ～ "XXXXX" は、/SAP_API_Output_Formatter/type.go 内 の Type Project {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-project-reads/SAP_API_Caller/caller.go#L50",
	"function": "sap-api-integrations-project-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": "[{XXXXXXXXXXXXXXXXXXXXXXXXXXXXX}]",
	"time": "2021-12-11T15:33:00.054455+09:00"
}
```
