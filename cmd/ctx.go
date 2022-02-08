/*
Copyright © 2022 Mohammed Abu-Taleb - mohatb@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// ctxCmd represents the ctx command
var ctxCmd = &cobra.Command{
	Use:   "ctx",
	Short: "switch between contexts (clusters) on kubectl faster",
	Long: `switch between contexts (clusters) on kubectl faster
	The script is not tied to kgo, it works by calling https://github.com/ahmetb/kubectx (Special thanks to the creators).`,
	Run: func(cmd *cobra.Command, args []string) {

		kubectx()

	},
}

func init() {
	rootCmd.AddCommand(ctxCmd)
}

func kubectx() {

	cmd := exec.Command("bash", "-c", "scripts/kubectx")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking
}
