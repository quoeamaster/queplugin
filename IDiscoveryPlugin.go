/*
 *  Copyright Project - quePlugin, Author - quoeamaster, (C) 2018
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package queplugin

// PS. a "new" method must be created for each DiscoveryPlugin implementor
// => New{PLUGIN_NAME} e.g. NewSimpleDiscoveryPlugin()

type DiscoveryPlugin interface {
    // method to ping for the given payload information (the payload could be a json document)
    Ping (payload string, options map[string]interface{}) (valid bool, jsonResponse string, returnValues map[string]interface{}, err error)

    // perform a master broker election; return the master's broker id, extra params and any error
    ElectMaster (payload string, params map[string]interface{}) (brokerId string, jsonResponse string, returnValues map[string]interface{}, err error)
}



