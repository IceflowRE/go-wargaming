// Auto generated file!

package wargaming

// WgnService Wargaming.NET service.
type WgnService service

// WotService World Of Tanks service.
type WotService service

// WotbService World Of Tanks Blitz service.
type WotbService service

// WotxService World Of Tanks Console service.
type WotxService service

// WowpService World Of Warplanes service.
type WowpService service

// WowsService World Of Warships service.
type WowsService service

type wgServices struct {
	// Wargaming.NET
	Wgn *WgnService
	// World Of Tanks
	Wot *WotService
	// World Of Tanks Blitz
	Wotb *WotbService
	// World Of Tanks Console
	Wotx *WotxService
	// World Of Warplanes
	Wowp *WowpService
	// World Of Warships
	Wows *WowsService
}

func newWgServices(common *service) wgServices {
	return wgServices{
		Wgn:  (*WgnService)(common),
		Wot:  (*WotService)(common),
		Wotb: (*WotbService)(common),
		Wotx: (*WotxService)(common),
		Wowp: (*WowpService)(common),
		Wows: (*WowsService)(common),
	}
}
