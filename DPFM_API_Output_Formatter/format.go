package dpfm_api_output_formatter

import (
	"data-platform-api-product-master-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	general := make([]General, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.General{}

		err := rows.Scan(
			&pm.FinInstCountry,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &general, err
		}

		data := pm
		general = append(general, General{
			FinInstCountry:      data.FinInstCountry,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &general, nil
	}

	return &general, nil
}
func ConvertToGeneralBusinessPartner(rows *sql.Rows) (*[]GeneralBusinessPartner, error) {
	defer rows.Close()
	generalBusinessPartner := make([]GeneralBusinessPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.GeneralBusinessPartner{}

		err := rows.Scan(
			&pm.FinInstCountry,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &generalBusinessPartner, err
		}

		data := pm
		generalBusinessPartner = append(generalBusinessPartner, GeneralBusinessPartner{
			FinInstCountry:      data.FinInstCountry,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &generalBusinessPartner, nil
	}

	return &generalBusinessPartner, nil
}

func ConvertToBranch(rows *sql.Rows) (*[]Branch, error) {
	defer rows.Close()
	branch := make([]Branch, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Branch{}

		err := rows.Scan(
			&pm.FinInstCountry,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &branch, err
		}

		data := pm
		branch = append(branch, Branch{
			FinInstCountry:      *&data.FinInstCountry,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &branch, nil
	}

	return &branch, nil
}

func ConvertToBranchBusinessPartner(rows *sql.Rows) (*[]BranchBusinessPartner, error) {
	defer rows.Close()
	branchBusinessPartner := make([]BranchBusinessPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.BranchBusinessPartner{}

		err := rows.Scan(
			&pm.FinInstCountry,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &branchBusinessPartner, err
		}

		data := pm
		branchBusinessPartner = append(branchBusinessPartner, BranchBusinessPartner{
			FinInstCountry:      data.FinInstCountry,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &branchBusinessPartner, nil
	}

	return &branchBusinessPartner, nil
}
