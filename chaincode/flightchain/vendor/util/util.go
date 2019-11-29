package util

import "acris"
import "fmt"
import "log"
import "errors"
import "strings"

func GenerateUniqueKey(flight acris.AcrisFlight) string {
 
  var flightNum string
  var flightKey string

  flightNum = flight.FlightNumber.TrackNumber;

      
  for (len(flightNum) < 4){

      flightNum = "0" + flightNum;
 }

  fmt.Println(flightNum )
  flightKey= flight.OriginDate + flight.DepartureAirport + flight.OperatingAirline.IataCode + flightNum
  log.Print("flight Key ***" , flightNum )
  return flightKey;


}


func VerifyValidACRIS(flight acris.AcrisFlight) error {

        var msg string

        if (acris.AcrisFlight{} == flight  || flight.OperatingAirline == nil || flight.OperatingAirline.IataCode == "" || len(flight.OperatingAirline.IataCode) != 2) {
            msg = "Invalid flight data, there is no valid flight.operatingAirline.iataCode set."
            log.Print(msg, flight.OperatingAirline)
            return  errors.New(msg)
        }

       if ( acris.AcrisFlight{} == flight  || flight.DepartureAirport == "" || len(flight.DepartureAirport) != 3) {
             msg = "Invalid flight data, there is no valid flight.departureAirport set."
             log.Print(msg, flight)
             return  errors.New(msg)
        }

      if ( acris.AcrisFlight{} == flight  || flight.ArrivalAirport == "" || len(flight.ArrivalAirport) != 3) {
            msg = "Invalid flight data, there is no valid flight.arrivalAirport set."
            log.Print(msg, flight)
            return  errors.New(msg)
        }
        if ( acris.AcrisFlight{} == flight  || flight.FlightNumber == nil || flight.FlightNumber.TrackNumber == "" || len(flight.FlightNumber.TrackNumber) != 4) {
            msg = "Invalid flight data, there is no valid 4 digit flight.flightNumber.trackNumber set."
            log.Print(msg, flight);
            return  errors.New(msg)
        }
        if (acris.AcrisFlight{} == flight  || flight.OriginDate == "" ) {
            msg = "Invalid flight data, there is no valid flight.originDate set (e.g. 2018-09-13)."
            log.Print(msg, flight)
            return  errors.New(msg)
        }
       
      
		return nil
    }


 func VerifyAbleToCreateOrModifyFlight(iata_code string, flight acris.AcrisFlight) error {
        
	var msg string
	var operatingAirlne string 
        var departureAirport string
        var arrivalAirport string

		
        if (iata_code ==""  || len(iata_code) > 3) {
            msg = "Invalid iata-code" + iata_code;
            log.Print(msg);
            return  errors.New(msg)
        }

        if (IsAirline(iata_code)) {
            operatingAirlne = GetOperatingAirline(flight);
            if (strings.ToUpper(operatingAirlne) != strings.ToUpper(iata_code)) {
                 msg = "Operating airline " + operatingAirlne + " does not match certificate iata-code " + iata_code;
                 log.Print(msg);
                 return  errors.New(msg)
            }
        } else {
            departureAirport = GetDepartureAirport(flight);
            arrivalAirport   = GetArrivalAirport(flight);
           if( strings.ToUpper(iata_code) !=  strings.ToUpper(departureAirport) && strings.ToUpper(iata_code) !=  strings.ToUpper(arrivalAirport)) {
                msg = "The iata airport code "+ iata_code + " does not match the departure " +
                            "airport "+departureAirport+ " or the arrival airport "+arrivalAirport;
                log.Print(msg);
                return  errors.New(msg)
            }
            
        }
		return nil

    }




func IsAirline(iata_code string) bool {
    

    return len(iata_code) == 2   

 }

    

func GetOperatingAirline(flight  acris.AcrisFlight) string {
        

  return flight.OperatingAirline.IataCode
    
 }

    

func GetDepartureAirport(flight  acris.AcrisFlight) string  {

       
    return flight.DepartureAirport    
}

    

func GetArrivalAirport(flight acris.AcrisFlight) string {
        

 return flight.ArrivalAirport
    

}


