// (c) 2022 by flonatel GmbH & Co. KG / Andreas Florath
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This file is part of rqmgtool.
//
// rqmgtool is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// at your option any later version.
//
// rqmgtool is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with rqmgtool.  If not, see <https://www.gnu.org/licenses/>.

package data

type ReqStatusType int64

const (
	RSTATEUnknown        ReqStatusType = iota
	RSTATEOpen
	RSTATEClosed
	RSTATEAssigned
)

var reqStatusTypeToString = map[ReqStatusType]string{
	RSTATEUnknown:     "+++UNKNOWN_REQ_STATUS_TYPE+++",
	RSTATEOpen:         "open",
	RSTATEClosed:       "closed",
	RSTATEAssigned:     "assigned",
}

var reqStatusTypeToID = map[string]ReqStatusType{
	"open":            RSTATEOpen,
	"closed":          RSTATEClosed,
	"assigned":        RSTATEAssigned,
}

func (s ReqStatusType) String() string {
	return reqStatusTypeToString[s]
}

// MarshalYAML marshals the enum as a quoted YAML string
//func (s ReqSubType) MarshalYAML() ([]byte, error) {
//	buffer := bytes.NewBufferString(`"`)
//	buffer.WriteString(toString[s])
//	buffer.WriteString(`"`)
//	return buffer.Bytes(), nil
//}

// UnmarshalYAML unmashals a quoted YAML string to the enum value
func (s *ReqStatusType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var j string
	err := unmarshal(&j)
	if err != nil {
		return err
	}
	*s = reqStatusTypeToID[j]
	return nil
}

// Local Variables:
// tab-width: 4
// End:
