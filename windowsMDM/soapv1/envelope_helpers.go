package soap

import (

	"encoding/xml"
	"errors"
)

func CheckFault(soapResponse []byte) error {
	xmlEnvelope := ResponseEnvelope{}

	err := xml.Unmarshal(soapResponse, &xmlEnvelope)
	if err != nil {
		return err
	}

	fault := xmlEnvelope.ResponseBodyBody.Fault
	if fault.XMLName.Local == "Fault" {
		sFault := fault.Code + " | " + fault.String + " | " + fault.Actor + " | " + fault.Detail
		return errors.New(sFault)
	}

	return nil
}
