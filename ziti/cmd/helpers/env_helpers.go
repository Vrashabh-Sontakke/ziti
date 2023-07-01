/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package helpers

import (
	edge "github.com/openziti/edge/controller/config"
	"github.com/openziti/ziti/ziti/constants"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
	"time"
)

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	h := os.Getenv("USERPROFILE") // windows
	if h == "" {
		h = "."
	}
	return NormalizePath(h)
}

func WorkingDir() (string, error) {
	wd, err := os.Getwd()
	if wd == "" || err != nil {
		return "", err
	}

	return NormalizePath(wd), nil
}

func GetZitiHome() (string, error) {
	// Get path from env variable
	retVal := os.Getenv(constants.ZitiHomeVarName)

	if retVal == "" {
		// If not set, create a default path of the current working directory
		workingDir, err := WorkingDir()
		if err != nil {
			// If there is an error just use .
			workingDir = "."
		}

		err = os.Setenv(constants.ZitiHomeVarName, workingDir)
		if err != nil {
			return "", err
		}

		retVal = os.Getenv(constants.ZitiHomeVarName)
	}

	return NormalizePath(retVal), nil
}

func HostnameOrNetworkName() string {
	val := os.Getenv("ZITI_NETWORK_NAME")
	if val == "" {
		h, err := os.Hostname()
		if err != nil {
			return "localhost"
		}
		return h
	}
	return val
}

func defaultValue(val string) func() string {
	return func() string {
		return val
	}
}
func defaultIntValue(val int64) func() string {
	return func() string {
		return strconv.FormatInt(val, 10)
	}
}

func GetCtrlBindAddress() string {
	return getFromEnv(constants.CtrlBindAddressVarName, defaultValue(constants.DefaultCtrlBindAddress))
}

func GetCtrlAdvertisedAddress() string {
	return getFromEnv(constants.CtrlAdvertisedAddressVarName, HostnameOrNetworkName)
}

func GetCtrlAdvertisedPort() string {
	return getFromEnv(constants.CtrlAdvertisedPortVarName, defaultValue(constants.DefaultCtrlAdvertisedPort))
}

func GetCtrlEdgeBindAddress() string {
	return getFromEnv(constants.CtrlEdgeBindAddressVarName, defaultValue(constants.DefaultCtrlEdgeBindAddress))
}
func GetCtrlEdgeAdvertisedAddress() string {
	return getFromEnv(constants.CtrlEdgeAdvertisedAddressVarName, HostnameOrNetworkName)
}
func GetCtrlEdgeAdvertisedPort() string {
	return getFromEnv(constants.CtrlEdgeAdvertisedPortVarName, defaultValue(constants.DefaultCtrlEdgeAdvertisedPort))
}

func GetZitiEdgeRouterPort() string {
	return getFromEnv(constants.ZitiEdgeRouterPortVarName, defaultValue(constants.DefaultZitiEdgeRouterPort))
}

func GetZitiEdgeRouterListenerBindPort() string {
	return getFromEnv(constants.ZitiEdgeRouterListenerBindPortVarName, defaultValue(constants.DefaultZitiEdgeRouterListenerBindPort))
}

func GetCtrlEdgeIdentityEnrollmentDuration() (time.Duration, error) {
	retVal := getFromEnv(constants.CtrlEdgeIdentityEnrollmentDurationVarName, defaultIntValue(int64(edge.DefaultEdgeEnrollmentDuration.Minutes())))
	retValInt, err := strconv.Atoi(retVal)
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+constants.CtrlEdgeIdentityEnrollmentDurationVarDescription)
		if err != nil {
			return edge.DefaultEdgeEnrollmentDuration, err
		}
	}

	return time.Duration(retValInt) * time.Minute, nil
}

func GetCtrlEdgeRouterEnrollmentDuration() (time.Duration, error) {
	retVal := getFromEnv(constants.CtrlEdgeRouterEnrollmentDurationVarName, defaultIntValue(int64(edge.DefaultEdgeEnrollmentDuration.Minutes())))
	retValInt, err := strconv.Atoi(retVal)
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+constants.CtrlEdgeRouterEnrollmentDurationVarDescription)
		if err != nil {
			return edge.DefaultEdgeEnrollmentDuration, err
		}
	}

	return time.Duration(retValInt) * time.Minute, nil
}

type envVarNotFound func() string

func getFromEnv(envVarName string, onNotFound envVarNotFound) string {
	// Get path from env variable
	retVal := os.Getenv(envVarName)
	if retVal != "" {
		return retVal
	}
	return onNotFound()
}

// NormalizePath replaces windows \ with / which windows allows for
func NormalizePath(input string) string {
	return strings.ReplaceAll(input, "\\", "/")
}
