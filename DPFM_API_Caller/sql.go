package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-fin-inst-master-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-fin-inst-master-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *[]dpfm_api_output_formatter.General
	var generalBusinessPartner *[]dpfm_api_output_formatter.GeneralBusinessPartner
	var branch *[]dpfm_api_output_formatter.Branch
	var branchBusinessPartner *[]dpfm_api_output_formatter.BranchBusinessPartner
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = c.General(mtx, input, output, errs, log)
			}()
		case "GeneralBusinessPartner":
			func() {
				generalBusinessPartner = c.GeneralBusinessPartner(mtx, input, output, errs, log)
			}()
		case "Branch":
			func() {
				branch = c.Branch(mtx, input, output, errs, log)
			}()
		case "BranchBusinessPartner":
			func() {
				branchBusinessPartner = c.BranchBusinessPartner(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:                general,
		GeneralBusinessPartner: generalBusinessPartner,
		Branch:                 branch,
		BranchBusinessPartner:  branchBusinessPartner,
	}

	return data
}

func (c *DPFMAPICaller) General(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) []dpfm_api_output_formatter.General {
	FinInstCountry := input.General.FinInstCountry
	FinInstCode := input.General.FinInstCode
	FinInstName := input.General.FinInstName
	FinInstFullName := input.General.FinInstFullName
	AddressID := input.General.AddressID
	SWIFTCode := input.General.SWIFTCode

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_fin_inst_master_general_data
		WHERE (
			FinInstCountry,
		    FinInstCode,
		    FinInstName,
		    FinInstFullName,
		    AddressID,
		    SWIFTCode,
		) = (?, ?, ?, ?, ?, ?, ?);`,
		FinInstCountry,
		FinInstCode,
		FinInstName,
		FinInstFullName,
		AddressID,
		SWIFTCode,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToGeneral(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) GeneralBusinessPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.GeneralBusinessPartner {
	finInstCountry := input.General.FinInstCountry
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_fin_inst_master_nutritional_info_data
		WHERE Product = ?
		ORDER BY Product DESC, FinInstCountry DESC;`, finInstCountry,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToNutritionalInfo(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Branch(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Calories {
	product := input.General.Product
	bp := input.BusinessPartnerID

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_calories_data
		WHERE Product = ? AND BusinessPartner = ?
		ORDER BY Product DESC, BusinessPartner DESC;`, product, *bp,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToCalories(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) BranchBusinessPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BranchBusinessPartner {
	var args []interface{}
	finInstCountry := input.General.FinInstCountry
	finInstCode := input.General.FinInstCode[0]

	cnt := 0
	for _, v := range bPPlant {
		args = append(args, product, v.BusinessPartner, v.Plant)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` )
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToBPPlant(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
