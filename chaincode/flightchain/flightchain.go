package main
import (
   "encoding/json"
   "fmt"
   "bytes"
   "acris"
   "util"
   "strconv"
   "github.com/hyperledger/fabric/core/chaincode/shim"
     pb "github.com/hyperledger/fabric/protos/peer"
    "github.com/hyperledger/fabric/core/chaincode/lib/cid"
)


type FlightChainData struct{

  FlightData   acris.AcrisFlight

  FlightKey    string
    
  UpdaterId    string
    
  TxId         string
    
  DocType      string

}

type FlightChaincode struct {
}


func (c *FlightChaincode ) Init(stub shim.ChaincodeStubInterface) pb.Response {

    fmt.Printf("initialization done!!!")
    fmt.Printf("initialization done!!!")
    return shim.Success(nil)
  

  
}



func (c *FlightChaincode) getFlight(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	flightBytes, _ := APIstub.GetState(args[0])
	if flightBytes == nil {
		return shim.Error("Could not locate flight key!getFlight - cant find any data for " +args[0])
	}
	return shim.Success(flightBytes)
}

func (c *FlightChaincode) getFlightHistory(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) < 1 {
            return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    flightKey := args[0]
    resultsIterator, err := APIstub.GetHistoryForKey(flightKey)
    if err != nil {
            return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    var buffer bytes.Buffer

    buffer.WriteString("[")
    bArrayMemberAlreadyWritten := false
    for resultsIterator.HasNext() {
            response, err := resultsIterator.Next()
            if err != nil {
                    return shim.Error(err.Error())
            }
            // Add a comma before array members, suppress it for the first array member
            if bArrayMemberAlreadyWritten == true {
                    buffer.WriteString(",")
            }
            buffer.WriteString("{")
            buffer.WriteString(string(response.Value))

            buffer.WriteString("}")
            bArrayMemberAlreadyWritten = true
    }
    buffer.WriteString("]")
    fmt.Printf("- History returning:\n%s\n", buffer.String())
    return shim.Success(buffer.Bytes())
}


func (c *FlightChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

 fc, args := stub.GetFunctionAndParameters()
 if fc == "createFlight" {
    return c.createFlight(stub, args)
 }
 
  if fc == "getFlight" {
    return c.getFlight(stub, args)
 }
 
 if fc == "getFlightHistory" {
    return c.getFlightHistory(stub, args)
 }
 return shim.Error("Called function is not defined in the chaincode ")
}

func (c *FlightChaincode) createFlight(stub shim.ChaincodeStubInterface, args []string) pb.Response {

if len(args) != 1 {

		return shim.Error("Incorrect number of arguments. Expecting 1")

 }
value := c.getCustomAttribute(stub)
fmt.Println("IATACODe from cert"+ value)
var arg string
arg = args[0]
_, err := strconv.Unquote(arg )
fmt.Print("final")
if err != nil {
 fmt.Println(err)
}
fmt.Print("final") 	
var val []byte = []byte(arg)
var flight acris.AcrisFlight
err = json.Unmarshal([]byte(val), &flight)
if err != nil {
        fmt.Println(string(args[0]))
	return shim.Error("No ACRIS flightdata passed in as " + args[0])
  }

  validErr := util.VerifyValidACRIS(flight);
  if validErr != nil {
      fmt.Println(validErr)
	  return shim.Error(validErr.Error())
   }
   
 checkErr := util.VerifyAbleToCreateOrModifyFlight(flight.OperatingAirline.IataCode, flight);
 if checkErr != nil {
      fmt.Println(checkErr)
	  return shim.Error(checkErr.Error())
   }
   
 flightKey := util.GenerateUniqueKey(flight)
 
 flightBytes, _ := stub.GetState(flightKey)
 var flightChainData FlightChainData
 var txID  string
 var buffer bytes.Buffer
 if flightBytes == nil {

   txID = stub.GetTxID()
   flightChainData = FlightChainData{FlightData:flight,FlightKey:flightKey,UpdaterId:flight.OperatingAirline.IataCode,TxId:txID,DocType:"flight"}
   flightDataAsBytes, _ := json.Marshal(flightChainData)
   stub.PutState(flightKey, flightDataAsBytes)
   buffer.WriteString("New flight status created with key " +flightKey) 

 }else{

   args[0]= arg 
   args[1] = flightKey
   c.updateFlight(stub, args)

 }
 
return shim.Success(buffer.Bytes())

}

func (c *FlightChaincode) updateFlight(stub shim.ChaincodeStubInterface, args []string) pb.Response  {
       
	   if (len(args) != 2) {
		   return shim.Error("Incorrect number of arguments. Expecting 2 (flightKey & new ACRIS flight data)")
        }

       value := c.getCustomAttribute(stub)
       fmt.Println("IATACODe from cert"+ value)
       var flightKey string = args[0]
       var arg string
       arg = args[0]
       _, err := strconv.Unquote(arg)
      
       if err != nil {
          fmt.Println(err)
       }
	   
       var val []byte = []byte(arg)
	   var flightDelta acris.AcrisFlight
	   err = json.Unmarshal([]byte(val), &flightDelta)
       if err != nil {
           fmt.Println(string(args[0]))
	       return shim.Error("No ACRIS flightdata passed in as " + args[0])
       }

        existingFlightBytes, _ := stub.GetState(flightKey)
	   
        if(existingFlightBytes == nil) {
		   return shim.Error("A flight with this flight key " +flightKey +" does not yet exist. It must be created first")
	    }
		
        var existingFlight FlightChainData
        json.Unmarshal(existingFlightBytes,existingFlight)
		jb, errMarshal := json.Marshal(existingFlight.FlightData)
		
	    if errMarshal != nil {
		  fmt.Println("Marshal error existingFlight.FlightData:", err)
	    }
		
	    err = json.Unmarshal(jb, &flightDelta)
		
	    if err != nil {
		 fmt.Println("Unmarshal &flightDelta:", err)
	    }
		existingFlight.FlightData = flightDelta   
        validErr := util.VerifyValidACRIS(existingFlight.FlightData);
        if validErr != nil {
            fmt.Println(validErr)
	        return shim.Error(validErr.Error())
         }
         mergedFlightKey := util.GenerateUniqueKey(flightDelta);
        if mergedFlightKey != flightKey {
          
             return shim.Error("You cannot change data that will modify the flight key originDate, departureAirport, operatingAirline.iataCode or flightNumber.trackNumber" )
        }

        var buffer bytes.Buffer
        existingFlight.UpdaterId = flightDelta.OperatingAirline.IataCode;
        existingFlight.TxId = stub.GetTxID();
        existingFlight.FlightData = flightDelta;
        flightDataAsBytes, _ := json.Marshal(existingFlight)
        stub.PutState(mergedFlightKey, flightDataAsBytes)
		buffer.WriteString("Flight status updated with key " +mergedFlightKey) 
		return shim.Success(buffer.Bytes())
}


func (c *FlightChaincode) getCustomAttribute(stub shim.ChaincodeStubInterface) string {
	
        var value string
	var found bool
	var err error

	value, found, err = cid.GetAttributeValue(stub, "iata-code")
	if err != nil {
                fmt.Print(found)
		fmt.Printf( "Error getting MSP identity: %s\n", err.Error())
		return ""
	}

	return value
}

func main() {

 err := shim.Start(new(FlightChaincode))
if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

