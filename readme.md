## UDM-API

usage example:
```
// create client
client, _ := udm_api.CreateClient(
    "apiuser",             // user
    "hunter2",             // pass
    "https://192.168.1.1", // host (unifi console IP)
    3,                     // request timeout
    true,                  // skip ssl verification
)

// get devices from "default" site
devices, _ := network.ClientController_Active(*client, "default")

// print device names
for _, d := range devices {
    fmt.Println(d.DisplayName)
}
```
