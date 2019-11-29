package acris

import (
	"encoding/json"
	
)

type Iot struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Context json.RawMessage `json:"context"`
}
type OperatingAirline struct {
	IataCode string    `json:"iataCode,required,validate:presence,min=2,max=2"`
	IcaoCode string    `json:"icaoCode"`
	Name     string    `json:"name"`
}


type AircraftType struct {
	IcaoCode     string    `json:"icaoCode"`
	ModelName    string    `json:"modelName,omitempty"`
	Registration string    `json:"registration,omitempty"`
}

type FlightNumber struct {
	AirlineCode      string `json:"airlineCode,required,validate:presence,min=2,max=2"`
	TrackNumber      string `json:"trackNumber,required"`
	DepartureAirport string `json:"departureAirport"`
}

type CheckinInfo struct {
	CheckinLocation     string    `json:"checkinLocation,omitempty"`
	CheckInBeginTime    string    `json:"checkInBeginTime,omitempty"`
	CheckInEndTime      string    `json:"CheckInEndTime,omitempty"`
	AdditionalInfo      string    `json:"additionalInfo,omitempty"`
	Any map[string]interface{}    `json:"-,omitempty"` 
}

type  BoardingTime struct {
     
        BookingClass string  `json:"bookingClass,omitempty"` 
        Time         string   `json:"time,omitempty"` 
        Any map[string]interface{}    `json:"-,omitempty"` 
    }

type FlightDepartureInformation struct {

    Scheduled  string  `json:"scheduled"` 
    
    Estimated  string  `json:"estimated"` 
    
    Actual    string  `json:"actual"` 
    
    Terminal  string  `json:"terminal"` 
    
    Gate     string  `json:"gate"` 
   
    CheckinInfo *CheckinInfo  `json:"checkinInfo,omitempty"` 
    
    BoardingTime *[]BoardingTime `json:"boardingTime,omitempty"` 
   
    Any map[string]interface{}    `json:"-,omitempty"` 


}

type FlightArrivalInformation struct {

     Scheduled  string  `json:"scheduled"` 
    
    Estimated  string  `json:"estimated"` 
    
    Actual    string  `json:"actual"` 
    
    Terminal  string  `json:"terminal"` 
    
    Gate     string  `json:"gate"` 
   
    TransferInformation string  `json:"transferInformation "` 
    BaggageClaim *BaggageClaim  `json:"baggageClaim,omitempty"` 

}

type  BaggageClaim struct {
     
        Carousel string  `json:"carousel,omitempty"` 
     
        ExpectedTimeOnCarousel  string `json:"expectedTimeOnCarousel,omitempty"`

        Any map[string]interface{}    `json:"-,omitempty"` 
    }

type Via struct{
   
    ViaAirport string  `json:"viaAirport,omitempty"` 
    Departure *FlightDepartureInformation `json:"departure,omitempty"`
    Arrival   *FlightArrivalInformation   `json:"arrival,omitempty"`
}



type AcrisFlight struct{
   
   OperatingAirline *OperatingAirline `json:"operatingAirline,omitempty"`
   
   AircraftType *AircraftType `json:"aircraftType,omitempty"`

   FlightNumber *FlightNumber `json:"flightNumber,omitempty"`
    
   CodeShares *[]FlightNumber `json:"codeShares,omitempty"`
   
   DepartureAirport string  `json:"departureAirport,omitempty"` 
   
   ArrivalAirport string   `json:"ArrivalAirport,omitempty"` 
   
   OriginDate  string   `json:"originDate,omitempty"`
   
   Departure *FlightDepartureInformation `json:"departure ,omitempty"`

   Arrival *FlightArrivalInformation `json:"arrival,omitempty"`
    
   FlightStatus string   `json:"flightStatus ,omitempty"`
    
   Via *[]Via `json:"via,omitempty"`

}
