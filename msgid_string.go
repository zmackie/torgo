// Code generated by "stringer -type=msgID"; DO NOT EDIT.

package main

import "fmt"

const _msgID_name = "KPALIVECHOKEUNCHOKEINTERSTUNINTERSTHAVEBITFLDREQPIECECNCL"

var _msgID_index = [...]uint8{0, 7, 12, 19, 26, 35, 39, 45, 48, 53, 57}

func (i msgID) String() string {
	i -= -1
	if i < 0 || i >= msgID(len(_msgID_index)-1) {
		return fmt.Sprintf("msgID(%d)", i+-1)
	}
	return _msgID_name[_msgID_index[i]:_msgID_index[i+1]]
}
