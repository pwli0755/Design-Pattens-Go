package main

func main() {
	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	windows := &windows{}
	windowsAdapter := &windowsAdapter{windowMachine: windows}
	client.insertLightningConnectorIntoComputer(windowsAdapter)
}
