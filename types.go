package main

type Event struct {
	name        string
	description string
	link        string
	location    string
	startTime   string
	endTime     string
}

func (e Event) ToArgSlice() []string {
	return []string{e.name, e.description, e.link, e.location, e.startTime, e.endTime}
}

// gno.land events past > present
var events = []Event{
	{
		name:        "Berlin Blockchain Week Buckle Up and Build with Cosmos",
		description: "",
		link:        "https://www.youtube.com/watch?v=hCLErPgnavI",
		location:    "Berlin, Germany",
		startTime:   "2022-09-11T09:00:00+02:00",
		endTime:     "2022-09-18T18:00:00+02:00",
	},
	{
		name:        "Cosmoverse",
		description: "",
		link:        "https://www.youtube.com/watch?v=6s1zG7hgxMk",
		location:    "Medellin, Colombia",
		startTime:   "2022-09-26T09:00:00-05:00",
		endTime:     "2022-09-28T18:00:00-05:00",
	},
	{
		name:        "Web Summit Buckle Up and Build with Cosmos",
		description: "",
		link:        "",
		location:    "Lisbon, Portugal",
		startTime:   "2022-11-01T09:00:00+01:00",
		endTime:     "2022-11-04T18:00:00+01:00",
	},
	{
		name:        "Istanbul Blockchain Week",
		description: "",
		link:        "https://www.youtube.com/watch?v=JX0gdWT0Cg4",
		location:    "Istanbul, Turkey",
		startTime:   "2022-11-14T10:00:00+03:00",
		endTime:     "2022-11-17T18:00:00+03:00",
	},
	{
		name:        "EthDenver 2023",
		description: "Side Event: Discover gno.land",
		link:        "https://www.youtube.com/watch?v=IJ0xel8lr4c",
		location:    "Denver, US",
		startTime:   "2023-02-24T10:00:00-06:00",
		endTime:     "2023-03-05T10:00:00-06:00",
	},
	{
		name:        "Game Developer Conference",
		description: "Side Event: Web3 Gaming Apps Powered by Gno",
		link:        "", // links for GDC & ETHDenver are the same on gno.land, by mistake
		location:    "San Francisco, US",
		startTime:   "2023-03-23T10:00:00-07:00",
		endTime:     "2023-03-23T18:00:00-07:00",
	},
	{
		name:        "BUIDL Asia",
		description: "Proof of Contribution in gno.land",
		link:        "https://www.buidl.asia/",
		location:    "Seoul, South Korea",
		startTime:   "2023-06-06T10:00:00+09:00",
		endTime:     "2023-06-07T18:00:00+09:00",
	},
	{
		name:        "ETH Seoul",
		description: "The Evolution of Smart Contracts: A Journey into gno.land",
		link:        "https://2023.ethseoul.org/",
		location:    "Seoul, South Korea",
		startTime:   "2023-06-02T10:00:00+09:00",
		endTime:     "2023-06-04T18:00:00+09:00",
	},
	{
		name:        "EthCC",
		description: "Come Meet Us at our Booth",
		link:        "https://ethcc.io/",
		location:    "Paris, France",
		startTime:   "2023-07-17T10:00:00+02:00",
		endTime:     "2023-07-20T18:00:00+02:00",
	},
	{
		name:        "Nebular Summit gno.land for Developers",
		description: "",
		link:        "https://www.nebular.builders/",
		location:    "Paris, France",
		startTime:   "2023-07-24T10:00:00+02:00",
		endTime:     "2023-07-25T18:00:00+02:00",
	},
	{
		name:        "GopherCon EU 2024",
		description: "Come Meet Us at our Booth",
		link:        "https://gophercon.eu/",
		location:    "Berlin, Germany",
		startTime:   "2023-07-26T10:00:00+02:00",
		endTime:     "2023-07-29T18:00:00+02:00",
	},
	{
		name:        "GopherCon US 2024",
		description: "Come Meet Us at our Booth",
		link:        "https://www.gophercon.com/",
		location:    "San Diego, US",
		startTime:   "2023-09-26T10:00:00-07:00",
		endTime:     "2023-09-29T18:00:00-07:00",
	},
	{
		name:        "Go to Gno Seoul",
		description: "Join the workshop!",
		link:        "https://medium.com/onbloc/go-to-gno-recap-intro-to-the-gno-stack-with-memeland-284a43d7f620",
		location:    "Seoul, South Korea",
		startTime:   "2024-03-23T10:00:00+09:00",
		endTime:     "2024-03-23T18:00:00+09:00",
	},
	{
		name:        "Intro to Gno Tokyo",
		description: "Join the meetup!",
		link:        "https://gno.land/r/gnoland/blog:p/gno-tokyo",
		location:    "Shinjuku City, Tokyo, Japan",
		startTime:   "2024-04-11T18:30:00+09:00",
		endTime:     "2024-04-11T22:00:00+09:00",
	},
	{
		name:        "Gno @ Golang Serbia",
		description: "Join the meetup!",
		link:        "https://gno.land/r/gnoland/blog:p/gnomes-in-serbia",
		location:    "Belgrade, Serbia",
		startTime:   "2024-05-23T18:00:00+02:00",
		endTime:     "2024-05-23T22:00:00+02:00",
	},
	{
		name:        "Nebular Summit 2024",
		description: "Join our workshop!",
		link:        "https://nebular.builders/",
		location:    "Brussels, Belgium",
		startTime:   "2024-07-12T10:00:00+02:00",
		endTime:     "2024-07-13T18:00:00+02:00",
	},
	{
		name:        "GopherCon US 2024",
		description: "Come Meet Us at our Booth",
		link:        "https://www.gophercon.com/",
		location:    "Chicago, US",
		startTime:   "2024-07-07T10:00:00-06:00",
		endTime:     "2024-07-10T18:00:00-06:00",
	},
	{
		name:        "GopherCon EU 2024",
		description: "Come Meet Us at our Booth",
		link:        "https://www.gophercon.com/",
		location:    "Berlin, Germany",
		startTime:   "2024-06-17T10:00:00+02:00",
		endTime:     "2024-06-20T10:00:00+02:00",
	},
	{
		name:        "Web3 Kamp 2024",
		description: "Workshop: \"Exploring Web3 Ecosystems - Building a dapp in Go\"",
		link:        "https://web3kamp.org/",
		location:    "Petnica, Serbia",
		startTime:   "2024-08-01T10:00:00+02:00",
		endTime:     "2024-08-09T17:00:00+02:00",
	},
	{
		name:        "Web3 Summit 2024",
		description: "Side Event: BUIDL with Cosmos Tooling",
		link:        "https://www.eventbrite.com/e/buidl-with-cosmos-tooling-tickets-981775686507?aff=oddtdtcreator",
		location:    "Berlin, Germany",
		startTime:   "2024-08-20T17:00:00+02:00",
		endTime:     "2024-08-20T22:00:00+02:00",
	},
}
