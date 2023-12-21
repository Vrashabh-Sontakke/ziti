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

package edge

import (
	"fmt"
	"github.com/openziti/ziti/controller/db"
	"github.com/openziti/ziti/tunnel/entities"
	"github.com/openziti/ziti/ziti/cmd/common"
	cmdhelper "github.com/openziti/ziti/ziti/cmd/helpers"
	"github.com/spf13/cobra"
	"io"
	"net/url"
	"os"
)

const (
	optionEndpoint = "endpoint"
)

// SecureOptions the options for the secure command
type SecureOptions struct {
	common.CommonOptions

	Endpoint string
}

// newSecureCmd consolidates network configuration steps for securing a service.
func newSecureCmd(out io.Writer, errOut io.Writer) *cobra.Command {
	options := &SecureOptions{}

	cmd := &cobra.Command{
		Use:   "secure <service_name> <protocol>:<address>:<port>",
		Short: "creates a service, configs, and policies for a resource",
		Long:  "creates a service, configs, and policies for a resource",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := runSecure(options)
			cmdhelper.CheckErr(err)
		},
	}

	cmd.Flags().StringVar(&options.Endpoint, optionEndpoint, "", "the custom endpoint name for your service")
	options.AddCommonFlags(cmd)

	cmd.AddCommand(newShowConfigTypeAction(out, errOut))
	cmd.AddCommand(newShowConfigAction(out, errOut))
	return cmd
}

// runSecure implements the command to secure a resource
func runSecure(o *SecureOptions) (err error) {

	svcName := o.Args[0]
	address := o.Args[1]

	// Parse the url argument
	u, err := url.Parse(address)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	protocol := u.Scheme
	hostname := u.Hostname()
	port := u.Port()

	if protocol == "" {
		fmt.Println("Protocol is missing")
	}
	if hostname == "" {
		fmt.Println("Hostname is missing")
	}
	if port == "" {
		fmt.Println("Port is missing")
	}

	// Create a bind config
	bindCfgName := svcName + ".host.v1"
	jsonStr := fmt.Sprintf(`{"protocol":"%s", "address":"%s", "port":%s}`, protocol, hostname, port)

	cmd := newCreateConfigCmd(os.Stdout, os.Stderr)
	args := []string{bindCfgName, entities.HostConfigV1, jsonStr}
	cmd.SetArgs(args)

	// Run the command
	err = cmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Create a dial config
	endpoint := svcName + ".ziti"
	if o.Endpoint != "" {
		endpoint = o.Endpoint
	}
	dialCfgName := svcName + ".intercept.v1"
	jsonStr = fmt.Sprintf(`{"protocols":["%s"], "addresses":["%s"], "portRanges":[{"low":%s, "high":%s}]}`, protocol, endpoint, port, port)
	fmt.Printf("jsonStr: %s\n", jsonStr)

	cmd = newCreateConfigCmd(os.Stdout, os.Stderr)
	args = []string{dialCfgName, entities.InterceptV1, jsonStr}
	cmd.SetArgs(args)

	// Run the command
	err = cmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Create service
	cmd = newCreateServiceCmd(os.Stdout, os.Stderr)
	args = []string{svcName, "--configs", bindCfgName + "," + dialCfgName}
	cmd.SetArgs(args)
	fmt.Printf("Creating service with args: %v\n", args)

	// Run the command
	err = cmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Create service policies
	svcRole := "@" + svcName

	dialSvcPolName := svcName + ".dial"
	dialIdRole := "#" + svcName + ".clients"
	cmd = newCreateServicePolicyCmd(os.Stdout, os.Stderr)
	args = []string{dialSvcPolName, db.PolicyTypeDialName, "--service-roles", svcRole, "--identity-roles", dialIdRole}
	cmd.SetArgs(args)

	// Run the command
	err = cmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
	}

	bindSvcPolName := svcName + ".bind"
	bindIdRole := "#" + svcName + ".servers"
	cmd = newCreateServicePolicyCmd(os.Stdout, os.Stderr)
	args = []string{bindSvcPolName, db.PolicyTypeBindName, "--service-roles", svcRole, "--identity-roles", bindIdRole}
	cmd.SetArgs(args)

	// Run the command
	err = cmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
	}

	return
}
