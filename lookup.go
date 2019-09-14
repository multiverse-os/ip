package ip

import (
	"fmt"
	"image/color"
	"net"
	"os"

	globe "github.com/multiverse-os/ip/globe"

	stun "github.com/ccding/go-stun/stun"
	maxmind "github.com/oschwald/maxminddb-golang"
)

func DrawConnection(remote string) {
	localCity := LookupSelf()
	fmt.Println("==================================")
	fmt.Println("Latitude:", localCity.Location.Latitude)
	fmt.Println("Longitude:", localCity.Location.Longitude)
	fmt.Println("==================================")

	remoteCity := Lookup(remote)
	fmt.Println("==================================")
	fmt.Println("Latitude:", remoteCity.Location.Latitude)
	fmt.Println("Longitude:", remoteCity.Location.Longitude)
	fmt.Println("==================================")

	g := globe.New()
	g.DrawGraticule(10.0)
	g.DrawLandBoundaries()
	g.DrawLine(
		localCity.Location.Latitude, localCity.Location.Longitude,
		remoteCity.Location.Latitude, remoteCity.Location.Longitude,
		// TODO: Color, size, and other line attributes will indicate PORT,
		// PROTOCOL, BANDWIDTH USED, etc. The idea is that one can get a lot of
		// information about their current network connections at a glance and
		// importantly can see major differences after watching it a while
		globe.Color(color.NRGBA{255, 0, 0, 255}),
	)
	// TODO: Can use this to slowly rotate the globe to show all the maps. Maybe
	// even support dragging it around
	g.CenterOn(50.244440, -37.207949)
	g.SavePNG("globe.png", 400)

}

func LookupSelf() City {
	client := stun.NewClient()
	nat, host, err := client.Discover()
	if err != nil {
		fmt.Println("[error] failed to make stun request:", err)
		os.Exit(1)
	}

	fmt.Println("NAT Type:", nat)
	if host != nil {
		fmt.Println("External IP Family:", host.Family())
		fmt.Println("External IP:", host.IP())
		fmt.Println("External Port:", host.Port())
	}
	return Lookup(host.IP())
}

func Lookup(address string) City {
	fmt.Println("Loading database...")

	db, err := maxmind.Open("./db/city.mmdb")
	if err != nil {
		fmt.Println("[error] failed to load ip database:", err)
		os.Exit(1)
	}
	defer db.Close()

	ip := net.ParseIP(address)

	var city City

	err = db.Lookup(ip, &city)
	if err != nil {
		fmt.Println("[error] failed to lookup ip:", err)
		os.Exit(1)
	}
	fmt.Println(city.Location)
	return city

}
