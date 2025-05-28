export namespace types {
	
	export class ConnectResult {
	    IsConnected: boolean;
	    Type: string;
	    ID: string;
	    Data: any;
	
	    static createFrom(source: any = {}) {
	        return new ConnectResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.IsConnected = source["IsConnected"];
	        this.Type = source["Type"];
	        this.ID = source["ID"];
	        this.Data = source["Data"];
	    }
	}

}

