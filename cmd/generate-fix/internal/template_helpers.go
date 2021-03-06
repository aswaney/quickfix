package internal

import (
	"fmt"

	"github.com/quickfixgo/quickfix/datadictionary"
)

func checkFieldTimeRequired(f *datadictionary.FieldDef) (required bool, err error) {
	var globalType *datadictionary.FieldType
	if globalType, err = getGlobalFieldType(f); err != nil {
		return
	}

	var t string
	if t, err = quickfixType(globalType); err != nil {
		return
	}

	var vt string
	if vt, err = quickfixValueType(t); err != nil {
		return
	}

	if vt == "time.Time" {
		required = true
		return
	}

	for _, groupField := range f.Fields {
		if required, err = checkFieldTimeRequired(groupField); required || err != nil {
			return
		}
	}

	return
}

func collectExtraImports(m *datadictionary.MessageDef) (imports []string, err error) {
	//NOTE: the time package is the only extra import considered here
	var timeRequired bool
	for _, f := range m.Fields {
		if timeRequired, err = checkFieldTimeRequired(f); timeRequired {
			imports = []string{"time"}
			break
		} else if err != nil {
			return
		}
	}

	return
}

func quickfixValueType(quickfixType string) (goType string, err error) {
	switch quickfixType {
	case "FIXString":
		goType = "string"
	case "FIXBoolean":
		goType = "bool"
	case "FIXInt":
		goType = "int"
	case "FIXUTCTimestamp":
		goType = "time.Time"
	case "FIXFloat":
		goType = "float64"
	default:
		err = fmt.Errorf("Unknown QuickFIX Type: %v", quickfixType)
	}

	return
}

func quickfixType(field *datadictionary.FieldType) (quickfixType string, err error) {
	switch field.Type {
	case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
		fallthrough
	case "MULTIPLECHARVALUE":
		fallthrough
	case "CHAR":
		fallthrough
	case "CURRENCY":
		fallthrough
	case "DATA":
		fallthrough
	case "MONTHYEAR":
		fallthrough
	case "LOCALMKTDATE":
		fallthrough
	case "TIME":
		fallthrough
	case "DATE":
		fallthrough
	case "EXCHANGE":
		fallthrough
	case "LANGUAGE":
		fallthrough
	case "XMLDATA":
		fallthrough
	case "COUNTRY":
		fallthrough
	case "UTCTIMEONLY":
		fallthrough
	case "UTCDATE":
		fallthrough
	case "UTCDATEONLY":
		fallthrough
	case "TZTIMEONLY":
		fallthrough
	case "TZTIMESTAMP":
		fallthrough
	case "STRING":
		quickfixType = "FIXString"

	case "BOOLEAN":
		quickfixType = "FIXBoolean"

	case "LENGTH":
		fallthrough
	case "DAYOFMONTH":
		fallthrough
	case "NUMINGROUP":
		fallthrough
	case "SEQNUM":
		fallthrough
	case "INT":
		quickfixType = "FIXInt"

	case "UTCTIMESTAMP":
		quickfixType = "FIXUTCTimestamp"

	case "QTY":
		fallthrough
	case "QUANTITY":
		fallthrough
	case "AMT":
		fallthrough
	case "PRICE":
		fallthrough
	case "PRICEOFFSET":
		fallthrough
	case "PERCENTAGE":
		fallthrough
	case "FLOAT":
		quickfixType = "FIXFloat"

	default:
		err = fmt.Errorf("Unknown type '%v' for tag '%v'\n", field.Type, field.Tag())
	}

	return
}

func requiredFields(m *datadictionary.MessageDef) (required []*datadictionary.FieldDef) {
	for _, part := range m.RequiredParts() {
		if part.Required() {
			switch pType := part.(type) {
			case *datadictionary.FieldDef:
				if !pType.IsGroup() {
					required = append(required, pType)
				}
			case *datadictionary.Component:
				for _, f := range pType.RequiredFields() {
					if !f.IsGroup() {
						required = append(required, f)
					}
				}
			}
		}
	}

	return
}

func routerBeginString(spec *datadictionary.DataDictionary) (routerBeginString string) {
	switch {
	case spec.FIXType == "FIXT":
		routerBeginString = "FIXT.1.1"
	case spec.Major != 5 && spec.ServicePack == 0:
		routerBeginString = fmt.Sprintf("FIX.%v.%v", spec.Major, spec.Minor)
		//ApplVerID enums
	case spec.Major == 2:
		routerBeginString = "0"
	case spec.Major == 3:
		routerBeginString = "1"
	case spec.Major == 4 && spec.Minor == 0:
		routerBeginString = "2"
	case spec.Major == 4 && spec.Minor == 1:
		routerBeginString = "3"
	case spec.Major == 4 && spec.Minor == 2:
		routerBeginString = "4"
	case spec.Major == 4 && spec.Minor == 3:
		routerBeginString = "5"
	case spec.Major == 4 && spec.Minor == 4:
		routerBeginString = "6"
	case spec.Major == 5 && spec.Minor == 0 && spec.ServicePack == 0:
		routerBeginString = "7"
	case spec.Major == 5 && spec.Minor == 0 && spec.ServicePack == 1:
		routerBeginString = "8"
	case spec.Major == 5 && spec.Minor == 0 && spec.ServicePack == 2:
		routerBeginString = "9"
	}

	return
}
