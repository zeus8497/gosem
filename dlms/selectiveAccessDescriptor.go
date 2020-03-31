package cosem

import (
	"bytes"
	. "gosem/axdr"
	"time"
)

type accesSelector uint8

const (
	AccessSelectorRange accesSelector = 0x1
	AccessSelectorEntry accesSelector = 0x2
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s accesSelector) Value() uint8 {
	return uint8(s)
}

type SelectiveAccessDescriptor struct {
	AccessSelector  accesSelector
	AccessParameter DlmsData
}

func CreateSelectiveAccessDescriptor(as accesSelector, ap interface{}) *SelectiveAccessDescriptor {
	if as == AccessSelectorRange {
		// make sure AccessParameter is a [2]time.Time
		ranges := ap.([]time.Time)
		// selector range should be of:
		// structure { structure {classid, obis, attributeid, dataidx}, range-start, range-end, selected val }
		var classId DlmsData = *CreateAxdrLongUnsigned(8)
		var obisCode DlmsData = *CreateAxdrOctetString("0.0.1.0.0.255") // obis of clock
		var attributeId DlmsData = *CreateAxdrInteger(2)
		var dataIdx DlmsData = *CreateAxdrLongUnsigned(0)
		var rangeStart DlmsData = *CreateAxdrDateTime(ranges[0])
		var rangeEnd DlmsData = *CreateAxdrDateTime(ranges[1])
		var selectedValue DlmsData = *CreateAxdrArray([]*DlmsData{})

		var restrictingObject DlmsData = *CreateAxdrStructure([]*DlmsData{&classId, &obisCode, &attributeId, &dataIdx})
		var rangeDescriptor DlmsData = *CreateAxdrStructure([]*DlmsData{&restrictingObject, &rangeStart, &rangeEnd, &selectedValue})

		return &SelectiveAccessDescriptor{AccessSelector: as, AccessParameter: rangeDescriptor}
	} else {
		// make sure AccessParameter is a [2]uint32
		entries := ap.([]uint32)
		// selector enty should be of:
		// structure {fromEntry, toEntry, fromSelectedValue, toSelectedValue}
		var fromEntry DlmsData = *CreateAxdrDoubleLongUnsigned(entries[0])
		var toEntry DlmsData = *CreateAxdrDoubleLongUnsigned(entries[1])
		var fromSelectedValue DlmsData = *CreateAxdrLongUnsigned(0)
		var toSelectedValue DlmsData = *CreateAxdrLongUnsigned(0)

		var entryDescriptor DlmsData = *CreateAxdrStructure([]*DlmsData{&fromEntry, &toEntry, &fromSelectedValue, &toSelectedValue})
		return &SelectiveAccessDescriptor{AccessSelector: as, AccessParameter: entryDescriptor}
	}
}

func (s *SelectiveAccessDescriptor) Encode() []byte {
	var out bytes.Buffer
	out.WriteByte(byte(s.AccessSelector.Value()))
	out.Write(s.AccessParameter.Encode())
	return out.Bytes()
}

func DecodeSelectiveAccessDescriptor(src *[]byte) (out SelectiveAccessDescriptor, err error) {
	if (*src)[0] == AccessSelectorRange.Value() {
		out.AccessSelector = AccessSelectorRange
	} else {
		out.AccessSelector = AccessSelectorEntry
	}
	(*src) = (*src)[1:] // remove access-selector byte

	thisDataTag, e := CheckTag(src)
	if e != nil {
		err = e
		return
	}

	var axdrDecoder Decoder = *NewDecoder(thisDataTag)
	out.AccessParameter, err = axdrDecoder.Decode(src)
	if err != nil {
		return
	}

	return
}
