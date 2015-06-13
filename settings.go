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

   File: jarvis-nlp.go
   Author: W. Max Lees <max.lees@gmail.com>
   Date: 06.12.2015
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Settings settings

type settings struct {
	NewNetwork bool
}

func (s *settings) SaveSettings() error {
	// Open the configuration file
	fmt.Println("Saving Config File")
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	ioutil.WriteFile("config.json", data, os.FileMode(0666))
	if err != nil {
		return err
	}

	return nil
}

func (s *settings) LoadSettings() error {
	// Open the configuration file
	fmt.Println("Reading Config File")
	fileData, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, s)
	if err != nil {
		return err
	}

	return nil
}
