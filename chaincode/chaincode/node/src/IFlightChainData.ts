import {AcrisFlight} from './acris-schema/AcrisFlight';

/**
 * This is the data model stored on the blockchain.
 * Each ACRIS flight entry/update is stored with the transaction Id and
 * updater Id so it can conveniently be accessed by the client apps.
 */
export interface IFlightChainData {
   
    flightData: AcrisFlight;
    flightKey: string;
    updaterId: string;
    txId: string;
    docType: string;
}

