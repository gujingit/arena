// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package evaluate

import (
	"github.com/spf13/cobra"
)

var (
	dataLong = `manage evaluate job.

Available Commands:
  model                Submit a evaluate job.
  list,ls              List the evaluate job.
  get                  Get evaluate job by name.
  delete,del           Delete evaluate job by name.
`
)

// NewEvaluateCommand manage evaluate job
func NewEvaluateCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "evaluate",
		Short: "Manage evaluate job. (deprecated)",
		Long:  dataLong,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewEvaluateModelCommand())
	command.AddCommand(NewEvaluateDeleteCommand())
	command.AddCommand(NewEvaluateListCommand())
	command.AddCommand(NewEvaluateGetCommand())

	return command
}
