/*
Copyright © 2021 Fotiadis Michail <fotiadis.michalis20@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion {bash | zsh | fish | powershell}",
	Short: "Output shell completion code for the specified shell",
	Long: `To load completions:

	Bash:
	
	  $ source <(ditctl completion bash)
	
	  # To load completions for each session, execute once:
	  # Linux:
	  $ ditctl completion bash > /etc/bash_completion.d/ditctl
	  # macOS:
	  $ ditctl completion bash > /usr/local/etc/bash_completion.d/ditctl
	
	Zsh:
	
	  # If shell completion is not already enabled in your environment,
	  # you will need to enable it.  You can execute the following once:
	
	  $ echo "autoload -U compinit; compinit" >> ~/.zshrc
	
	  # To load completions for each session, execute once:
	  $ echo "source <(ditctl completion bash)" >> ~/.zshrc
	
	  # You will need to start a new shell for this setup to take effect.
	
	fish:
	
	  $ ditctl completion fish | source
	
	  # To load completions for each session, execute once:
	  $ ditctl completion fish > ~/.config/fish/completions/ditctl.fish
	
	PowerShell:
	
	  PS> ditctl completion powershell | Out-String | Invoke-Expression
	
	  # To load completions for every new session, run:
	  PS> ditctl completion powershell > ditctl.ps1
	  # and source this file from your PowerShell profile.
	`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			rootCmd.GenZshCompletion(os.Stdout)
			fmt.Println("compdef _ditctl ditctl")
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			rootCmd.GenPowerShellCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
