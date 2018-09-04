package core

import "github.com/ninjadotorg/uncommons/common/hexutil"

// GenesisContractControlInput is ...
var GenesisContractControlInput, _ = hexutil.Decode("0x608060405234801561001057600080fd5b50600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563018054600160a060020a031916736435192907ef452744560be39fce68835ee6d7e017905561092c806100736000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312c443f381146100a85780631785f53c146100cf57806324d7806c14610104578063268a66d81461012557806366c083b71461016457806367e13ee01461018557806370480275146101c95780639e23c209146101ea578063c6efaed71461020b578063eff25cd61461022f575b600080fd5b3480156100b457600080fd5b506100bd610268565b60408051918252519081900360200190f35b3480156100db57600080fd5b506100f0600160a060020a036004351661026f565b604080519115158252519081900360200190f35b34801561011057600080fd5b506100f0600160a060020a036004351661034e565b34801561013157600080fd5b50610149600160a060020a03600435166024356103ac565b60408051921515835260208301919091528051918290030190f35b34801561017057600080fd5b506100bd600160a060020a03600435166104a5565b34801561019157600080fd5b506101a6600160a060020a03600435166104e4565b60408051600160a060020a03909316835260208301919091528051918290030190f35b3480156101d557600080fd5b50610149600160a060020a0360043516610537565b3480156101f657600080fd5b506100f0600160a060020a03600435166105f4565b34801561021757600080fd5b50610149600160a060020a03600435166024356106b7565b34801561023b57600080fd5b5061024760043561075b565b60408051928352600160a060020a0390911660208301528051918290030190f35b6001545b90565b600080600061027d3361034e565b156103425761028b84610790565b9150915081801561029e57506000546001105b15610342575b600054600019018110156103255760008054600183019081106102c357fe5b60009182526020822001548154600160a060020a039091169190839081106102e757fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556001016102a4565b6000805490610338906000198301610854565b5060019250610347565b600092505b5050919050565b6000805b6000548110156103a15782600160a060020a031660008281548110151561037557fe5b600091825260209091200154600160a060020a0316141561039957600191506103a6565b600101610352565b600091505b50919050565b6000806000806103bb3361034e565b15610492576103c9866107f5565b9150915081156103df576001819350935061049c565b60408051808201909152858152600160a060020a0387811660208301908152600180548082018255600082815294517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf660029092029182015591517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf7909201805473ffffffffffffffffffffffffffffffffffffffff19169290931691909117909155905490945060001901925061049c565b6000935060001992505b50509250929050565b60008060006104b3846107f5565b9150915081156103425760018054829081106104cb57fe5b9060005260206000209060020201600001549250610347565b6000806000806104f3856107f5565b915091508115610528578460018281548110151561050d57fe5b90600052602060002090600202016000015493509350610530565b849350600092505b5050915091565b6000806000806004600080549050111561055957600093506000199250610530565b6105623361034e565b156105e65761057085610790565b9150915081156105865760018193509350610530565b60008054600180820183559180527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388161790559350839250610530565b600093506000199250610530565b60008060006106023361034e565b1561034257610610846107f5565b915091508115610342575b600154600019018110156106a4576001805482820190811061063957fe5b906000526020600020906002020160018281548110151561065657fe5b600091825260209091208254600290920201908155600191820154908201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790550161061b565b600180549061033890600019830161087d565b6000806000806106c63361034e565b15610492576106d4866107f5565b9150915081156104925760408051808201909152858152600160a060020a0387166020820152600180548390811061070857fe5b60009182526020918290208351600292909202019081559101516001918201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055935091508161049c565b600180548290811061076957fe5b600091825260209091206002909102018054600190910154909150600160a060020a031682565b600080805b6000548110156107e75783600160a060020a03166000828154811015156107b857fe5b600091825260209091200154600160a060020a031614156107df57600181925092506107ef565b600101610795565b600092508291505b50915091565b600080805b6001548110156107e75783600160a060020a031660018281548110151561081d57fe5b6000918252602090912060016002909202010154600160a060020a0316141561084c57600181925092506107ef565b6001016107fa565b815481835581811115610878576000838152602090206108789181019083016108a9565b505050565b8154818355818111156108785760020281600202836000526020600020918201910161087891906108c7565b61026c91905b808211156108c357600081556001016108af565b5090565b61026c91905b808211156108c3576000815560018101805473ffffffffffffffffffffffffffffffffffffffff191690556002016108cd5600a165627a7a72305820a70da5f33bea728cbcb2bd5f8dc20cef1aea00a5a92078c300a72fe44d3231600029")

// GenesisContractControlCode is ...
var GenesisContractControlCode, _ = hexutil.Decode("0x6080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312c443f381146100a85780631785f53c146100cf57806324d7806c14610104578063268a66d81461012557806366c083b71461016457806367e13ee01461018557806370480275146101c95780639e23c209146101ea578063c6efaed71461020b578063eff25cd61461022f575b600080fd5b3480156100b457600080fd5b506100bd610268565b60408051918252519081900360200190f35b3480156100db57600080fd5b506100f0600160a060020a036004351661026f565b604080519115158252519081900360200190f35b34801561011057600080fd5b506100f0600160a060020a036004351661034e565b34801561013157600080fd5b50610149600160a060020a03600435166024356103ac565b60408051921515835260208301919091528051918290030190f35b34801561017057600080fd5b506100bd600160a060020a03600435166104a5565b34801561019157600080fd5b506101a6600160a060020a03600435166104e4565b60408051600160a060020a03909316835260208301919091528051918290030190f35b3480156101d557600080fd5b50610149600160a060020a0360043516610537565b3480156101f657600080fd5b506100f0600160a060020a03600435166105f4565b34801561021757600080fd5b50610149600160a060020a03600435166024356106b7565b34801561023b57600080fd5b5061024760043561075b565b60408051928352600160a060020a0390911660208301528051918290030190f35b6001545b90565b600080600061027d3361034e565b156103425761028b84610790565b9150915081801561029e57506000546001105b15610342575b600054600019018110156103255760008054600183019081106102c357fe5b60009182526020822001548154600160a060020a039091169190839081106102e757fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556001016102a4565b6000805490610338906000198301610854565b5060019250610347565b600092505b5050919050565b6000805b6000548110156103a15782600160a060020a031660008281548110151561037557fe5b600091825260209091200154600160a060020a0316141561039957600191506103a6565b600101610352565b600091505b50919050565b6000806000806103bb3361034e565b15610492576103c9866107f5565b9150915081156103df576001819350935061049c565b60408051808201909152858152600160a060020a0387811660208301908152600180548082018255600082815294517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf660029092029182015591517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf7909201805473ffffffffffffffffffffffffffffffffffffffff19169290931691909117909155905490945060001901925061049c565b6000935060001992505b50509250929050565b60008060006104b3846107f5565b9150915081156103425760018054829081106104cb57fe5b9060005260206000209060020201600001549250610347565b6000806000806104f3856107f5565b915091508115610528578460018281548110151561050d57fe5b90600052602060002090600202016000015493509350610530565b849350600092505b5050915091565b6000806000806004600080549050111561055957600093506000199250610530565b6105623361034e565b156105e65761057085610790565b9150915081156105865760018193509350610530565b60008054600180820183559180527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388161790559350839250610530565b600093506000199250610530565b60008060006106023361034e565b1561034257610610846107f5565b915091508115610342575b600154600019018110156106a4576001805482820190811061063957fe5b906000526020600020906002020160018281548110151561065657fe5b600091825260209091208254600290920201908155600191820154908201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790550161061b565b600180549061033890600019830161087d565b6000806000806106c63361034e565b15610492576106d4866107f5565b9150915081156104925760408051808201909152858152600160a060020a0387166020820152600180548390811061070857fe5b60009182526020918290208351600292909202019081559101516001918201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055935091508161049c565b600180548290811061076957fe5b600091825260209091206002909102018054600190910154909150600160a060020a031682565b600080805b6000548110156107e75783600160a060020a03166000828154811015156107b857fe5b600091825260209091200154600160a060020a031614156107df57600181925092506107ef565b600101610795565b600092508291505b50915091565b600080805b6001548110156107e75783600160a060020a031660018281548110151561081d57fe5b6000918252602090912060016002909202010154600160a060020a0316141561084c57600181925092506107ef565b6001016107fa565b815481835581811115610878576000838152602090206108789181019083016108a9565b505050565b8154818355818111156108785760020281600202836000526020600020918201910161087891906108c7565b61026c91905b808211156108c357600081556001016108af565b5090565b61026c91905b808211156108c3576000815560018101805473ffffffffffffffffffffffffffffffffffffffff191690556002016108cd5600a165627a7a72305820a70da5f33bea728cbcb2bd5f8dc20cef1aea00a5a92078c300a72fe44d3231600029")

// GenesisContractControlJSON is ...
var GenesisContractControlJSON = `{"nonce":"0x0","gasPrice":"0x3b9aca00","gas":"0xb08a8","to":null,"value":"0x0","input":"0x608060405234801561001057600080fd5b50600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563018054600160a060020a031916736435192907ef452744560be39fce68835ee6d7e017905561092c806100736000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312c443f381146100a85780631785f53c146100cf57806324d7806c14610104578063268a66d81461012557806366c083b71461016457806367e13ee01461018557806370480275146101c95780639e23c209146101ea578063c6efaed71461020b578063eff25cd61461022f575b600080fd5b3480156100b457600080fd5b506100bd610268565b60408051918252519081900360200190f35b3480156100db57600080fd5b506100f0600160a060020a036004351661026f565b604080519115158252519081900360200190f35b34801561011057600080fd5b506100f0600160a060020a036004351661034e565b34801561013157600080fd5b50610149600160a060020a03600435166024356103ac565b60408051921515835260208301919091528051918290030190f35b34801561017057600080fd5b506100bd600160a060020a03600435166104a5565b34801561019157600080fd5b506101a6600160a060020a03600435166104e4565b60408051600160a060020a03909316835260208301919091528051918290030190f35b3480156101d557600080fd5b50610149600160a060020a0360043516610537565b3480156101f657600080fd5b506100f0600160a060020a03600435166105f4565b34801561021757600080fd5b50610149600160a060020a03600435166024356106b7565b34801561023b57600080fd5b5061024760043561075b565b60408051928352600160a060020a0390911660208301528051918290030190f35b6001545b90565b600080600061027d3361034e565b156103425761028b84610790565b9150915081801561029e57506000546001105b15610342575b600054600019018110156103255760008054600183019081106102c357fe5b60009182526020822001548154600160a060020a039091169190839081106102e757fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556001016102a4565b6000805490610338906000198301610854565b5060019250610347565b600092505b5050919050565b6000805b6000548110156103a15782600160a060020a031660008281548110151561037557fe5b600091825260209091200154600160a060020a0316141561039957600191506103a6565b600101610352565b600091505b50919050565b6000806000806103bb3361034e565b15610492576103c9866107f5565b9150915081156103df576001819350935061049c565b60408051808201909152858152600160a060020a0387811660208301908152600180548082018255600082815294517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf660029092029182015591517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf7909201805473ffffffffffffffffffffffffffffffffffffffff19169290931691909117909155905490945060001901925061049c565b6000935060001992505b50509250929050565b60008060006104b3846107f5565b9150915081156103425760018054829081106104cb57fe5b9060005260206000209060020201600001549250610347565b6000806000806104f3856107f5565b915091508115610528578460018281548110151561050d57fe5b90600052602060002090600202016000015493509350610530565b849350600092505b5050915091565b6000806000806004600080549050111561055957600093506000199250610530565b6105623361034e565b156105e65761057085610790565b9150915081156105865760018193509350610530565b60008054600180820183559180527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388161790559350839250610530565b600093506000199250610530565b60008060006106023361034e565b1561034257610610846107f5565b915091508115610342575b600154600019018110156106a4576001805482820190811061063957fe5b906000526020600020906002020160018281548110151561065657fe5b600091825260209091208254600290920201908155600191820154908201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790550161061b565b600180549061033890600019830161087d565b6000806000806106c63361034e565b15610492576106d4866107f5565b9150915081156104925760408051808201909152858152600160a060020a0387166020820152600180548390811061070857fe5b60009182526020918290208351600292909202019081559101516001918201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055935091508161049c565b600180548290811061076957fe5b600091825260209091206002909102018054600190910154909150600160a060020a031682565b600080805b6000548110156107e75783600160a060020a03166000828154811015156107b857fe5b600091825260209091200154600160a060020a031614156107df57600181925092506107ef565b600101610795565b600092508291505b50915091565b600080805b6001548110156107e75783600160a060020a031660018281548110151561081d57fe5b6000918252602090912060016002909202010154600160a060020a0316141561084c57600181925092506107ef565b6001016107fa565b815481835581811115610878576000838152602090206108789181019083016108a9565b505050565b8154818355818111156108785760020281600202836000526020600020918201910161087891906108c7565b61026c91905b808211156108c357600081556001016108af565b5090565b61026c91905b808211156108c3576000815560018101805473ffffffffffffffffffffffffffffffffffffffff191690556002016108cd5600a165627a7a72305820a70da5f33bea728cbcb2bd5f8dc20cef1aea00a5a92078c300a72fe44d3231600029","v":"0x26","r":"0xcb0c0d18cd4d9810266d19d7208f91b031dbd28030549ffef6dd8940927999c2","s":"0x483dddf8aa29211cbedb724ab75bbe1ac4e5989e36cf318dbcc15862adecd590","hash":"0xee0fd7d6a89cd033c574217f1b3c74a25a7175d9064fcf0af1804dd0896b69f9"}`
