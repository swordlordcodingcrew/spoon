package main
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
 **-----------------------------------------------------------------------------
 **
 ** This file is based on Frisby which is Copyright (c) 2015 Tony Worm. Frisby
 ** is licenced under the The MIT License (MIT)
 **
 **
-----------------------------------------------------------------------------*/
import (
	"fmt"
	"os"
	"swordlord.com/spoon/cmd"
	"swordlord.com/spoon/utils"
	"swordlord.com/spoon/frisby"
)

func main() {

	// Initialise env and params
	utils.InitConfig()

	frisby.Create("Test GET homepage").
		Get("http://127.0.0.1:8888/").
		Send().
		ExpectStatus(200)

	//.ExpectContent("The Go Programming Language")

	frisby.Global.PrintReport()

	// initialise the command structure
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)

		// yes, we deferred closing of the db, but that only works when ending with dignity
		os.Exit(1)
	}
}