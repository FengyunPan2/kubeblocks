/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package component

import (
	"os"
)

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func GetSQLChannelProc() (*os.Process, error) {
	// sqlChannel pid is usually 1
	sqlChannelPid := os.Getppid()
	sqlChannelProc, err := os.FindProcess(sqlChannelPid)
	if err != nil {
		return nil, err
	}

	return sqlChannelProc, nil
}
