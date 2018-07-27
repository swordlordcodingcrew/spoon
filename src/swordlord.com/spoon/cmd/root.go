package cmd
/*-----------------------------------------------------------------------------
 **
 ** - Spoon -
 **
 ** The tasting spoon for the Fennel project. CalDAV and CardDAV testing toolkit.
 **
 ** Copyright 2018 by SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 ** This program is free software; you can redistribute it and/or modify it
 ** under the terms of the GNU Affero General Public License as published by the
 ** Free Software Foundation, either version 3 of the License, or (at your option)
 ** any later version.
 **
 ** This program is distributed in the hope that it will be useful, but WITHOUT
 ** ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 ** FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
 ** for more details.
 **
 ** You should have received a copy of the GNU Affero General Public License
 ** along with this program. If not, see <http://www.gnu.org/licenses/>.
 **
 **-----------------------------------------------------------------------------
 **
 ** Original Authors:
 ** LordEidi@swordlord.com
 ** LordCuttlery@swordlord.com
 **
-----------------------------------------------------------------------------*/
import (
	"github.com/spf13/cobra"
)

// var cfgFile string // see init() for details

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "spoon",
	Short: "The tasting spoon for the Fennel project. CalDAV and CardDAV testing toolkit.",
	Long: `The tasting spoon for the Fennel project. CalDAV and CardDAV testing toolkit.

This tool helps you in testing the CalDAV and CardDAV functionality of your Fennel instance
(or another server of your liking). This is mostly a development helper, no production tool.`,
}

func init() {
	// following lines just for reference.

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gohjasmincli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


