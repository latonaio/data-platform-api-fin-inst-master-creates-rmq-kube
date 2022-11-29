package dpfm_api_input_reader

import (
	"data-platform-api-fin-inst-master-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.FinInstMaster
	return &requests.General{
		FinInstCountry:      data.FinInstCountry,
		FinInstCode:         data.FinInstCode,
		FinInstName:         data.FinInstName,
		FinInstFullName:     data.FinInstFullName,
		AddressID:           data.AddressID,
		SWIFTCode:           data.SWIFTCode,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}
