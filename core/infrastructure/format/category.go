/**
 * Copyright 2015 @ z3q.net.
 * name : category
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package format

import (
	"strconv"
	"strings"
)

// 获取栏目编号字符串
func IdArrJoinStr(ids []int) string {
	var strIds []string = make([]string, len(ids))
	for i, v := range ids {
		strIds[i] = strconv.Itoa(v)
	}
	return strings.Join(strIds, ",")
}

func IdArrJoinStr32(ids []int32) string {
	var strIds []string = make([]string, len(ids))
	for i, v := range ids {
		strIds[i] = strconv.Itoa(int(v))
	}
	return strings.Join(strIds, ",")
}

func IdArrJoinStr64(ids []int64) string {
	var strIds []string = make([]string, len(ids))
	for i, v := range ids {
		strIds[i] = strconv.Itoa(int(v))
	}
	return strings.Join(strIds, ",")
}
