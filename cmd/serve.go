// Copyright © 2016 John Morrice <john@functorama.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"net"

	"github.com/spf13/cobra"
	"github.com/johnny-morrice/gear/lib"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run gear server",
	Long: `Run gear server until terminated.  Example:

	$ gear serve --bind 10.0.1.1 --port 7777`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()

		bind, binderr := flags.GetIP("bind")
		port, porterr := flags.GetUint("port")

		err := binderr
		if err == nil {
			err = porterr
		}

		if err != nil {
			die(err)
		}

		err = lib.Serve(bind, port, nil)

		if err != nil {
			die(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	flags := serveCmd.PersistentFlags()
	flags.IP("bind", net.IP([]byte{127,0,0,1}), "Bind interface")
	flags.Uint("port", 6666, "Server port")
}
