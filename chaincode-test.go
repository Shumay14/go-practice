// ReadAsset reads the information from collection
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, assetID string) (*Asset, error) {

	}


// ReadAssetPrivateDetails reads the asset private details in organization specific collection
func (s *SmartContract)	


//


func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, ) error {

	// Get new asset from transient map
	transientMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return fmt.Errorf("error getting transient: %v", err)
	}

	// Asset properties are private, therefore they get passed in transient field, instead of func args
	transientAssetJSON, ok := 


	type assetTransientInput struct {
		Type 		   string
		ID			   string	
		Colour		   string
		Size		   int
		AppraisedValue int

	}
}