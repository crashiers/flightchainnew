import shim = require('fabric-shim');
import FlightStatusChaincode from './chaincode';

shim.start(new FlightStatusChaincode());
