/*
  Copyright 2015 W. Max Lees

  This file is part of jarvisos.

  Jarvisos is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  Jarvisos is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with jarvisos.  If not, see <http://www.gnu.org/licenses/>.

  File: network.go
  Author: W. Max Lees <max.lees@gmail.com>
  Date: 06.12.2015
*/

package network

// Network Interface
// This interface is used by Jarvis NLP to actually run any
// NLP problems. It contains the actual Neural Network to use
// for determining meaning.
type Network interface {
	// initalize Function
	// Used to run any setup required for the network
	//
	// return bool
	// 		a boolean value based on the success of the
	// 		setup
	Initalize(bool) bool

	// setInput Function
	// Used to set the input string for the network to
	// process
	//
	// string parameter
	//		A string for the network to process
	SetInput(string)

	// destroy Function
	// Used to run any internal shutdown the network needs
	// before it is released
	//
	// return bool
	// 		A boolean value based on the success of the
	//		destruction
	Destroy() bool
}
