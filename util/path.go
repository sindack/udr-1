// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

//+build !debug

package util

import (
	"github.com/free5gc/path_util"
)

var (
	UdrLogPath           = path_util.Free5gcPath("free5gc/udrsslkey.log")
	UdrPemPath           = path_util.Free5gcPath("free5gc/support/TLS/udr.pem")
	UdrKeyPath           = path_util.Free5gcPath("free5gc/support/TLS/udr.key")
	DefaultUdrConfigPath = path_util.Free5gcPath("free5gc/config/udrcfg.yaml")
)
