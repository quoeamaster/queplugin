package queplugin

type DiscoveryPlugin interface {
    // method to ping for the given url / broker
    Ping (url string, options map[string]interface{}) (valid bool, err error)

    // perform a master broker election; return the master's broker id
    ElectMaster (params map[string]interface{}) (brokerId string, err error)

}

