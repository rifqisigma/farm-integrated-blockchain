// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_distributorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_destinationId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestCollectorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProcessorId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_transportation\",\"type\":\"string\"}],\"name\":\"createDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_farmerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cropId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_regencyId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_basePrice\",\"type\":\"uint256\"}],\"name\":\"createHarvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collectorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_basePrice\",\"type\":\"uint256\"}],\"name\":\"createHarvestCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_processorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestCollectorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"createHarvestProcessor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sellerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_distributionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"createSellerBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"distributionIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"distributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"harvestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"harvestCollectorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"harvestProcessorId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"transportation\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllDistributionTuples\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllHarvestCollectorTuples\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllHarvestProcessorTuples\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllHarvestTuples\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllSellerBoxTuples\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"harvestCollectorIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"harvestCollectors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"harvestId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"harvestIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"harvestProcessorIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"harvestProcessors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"harvestCollectorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"harvestId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"harvests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cropId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"regencyId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sellerBoxIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sellerBoxes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657612465908161001c8239f35b600080fdfe6101c0604052600436101561001357600080fd5b60003560e01c8063184024a11461170e5780631a84ce9b146116535780631fde561c1461152c57806335a423b01461147d5780633d6140ce1461143f5780634487d3df1461134a5780634841fe1e14611224578063579255b3146111fb5780635b80e7d814610eb95780636910f26014610e9057806379360aa114610b36578063867c23f314610b0d5780638a8b1ece14610ad557806392bb06b314610a4a57806398728e9a146109a6578063abd9f6c414610939578063acad494914610902578063b7d5e2da1461074a578063bdec4ed7146105fd5763e59d49d4146100f957600080fd5b346105f857610180806003193601126105f85760c4356001600160401b0381116105f85761012b903690600401611fc2565b9060e4356001600160401b0381116105f85761014b903690600401611fc2565b610164356001600160401b0381116105f85761016b903690600401611fc2565b6040519161017883611bf1565b60043583526020830194602435865260443560408501526064356060850152608435608085015260a43560a085015260c084015260e0830152610104356101008301526101243561012083015261014435610140830152610160820152428282015260043560005260066020526040600020928151845551600184015560408101516002840155606081015160038401556080810151600484015560a0810151600584015560c08101519283516001600160401b0381116103a2576102406006830154611bb7565b601f81116105b1575b506020601f8211600114610542578192939495600092610537575b50508160011b916000199060031b1c19161760068201555b60e08201519283516001600160401b0381116103a25761029f6007840154611bb7565b601f81116104f0575b506020601f8211600114610481578192939495600092610476575b50508160011b916000199060031b1c19161760078301555b61010083015160088301556101208301516009830155610140830151600a830155600b820192610160810151938451946001600160401b0386116103a2576103238254611bb7565b601f811161042e575b50602090601f87116001146103c357958091600c96976000926103b8575b50508160011b916000199060031b1c19161790555b0151910155600754600160401b8110156103a2576103868160016103a09301600755611e99565b6004359082549060031b91821b91600019901b1916179055565b005b634e487b7160e01b600052604160045260246000fd5b01519050388061034a565b90601f198716918360005260206000209260005b8181106104165750916001939189600c999a94106103fd575b505050811b01905561035f565b015160001960f88460031b161c191690553880806103f0565b929360206001819287860151815501950193016103d7565b826000526020600020601f880160051c8101916020891061046c575b601f0160051c01905b818110610460575061032c565b60008155600101610453565b909150819061044a565b0151905038806102c3565b6007840160005260206000209060005b601f19841681106104d8575060019394959683601f198116106104bf575b505050811b0160078301556102db565b015160001960f88460031b161c191690553880806104af565b9091602060018192858b015181550193019101610491565b600784016000526020600020601f830160051c810160208410610530575b601f830160051c820181106105245750506102a8565b6000815560010161050e565b508061050e565b015190503880610264565b6006830160005260206000209060005b601f1984168110610599575060019394959683601f19811610610580575b505050811b01600682015561027c565b015160001960f88460031b161c19169055388080610570565b9091602060018192858b015181550193019101610552565b600683016000526020600020601f830160051c8101602084106105f1575b601f830160051c820181106105e5575050610249565b600081556001016105cf565b50806105cf565b600080fd5b346105f85760003660031901126105f857600380549061061c8261208d565b6106258361208d565b9161062f8461208d565b93610639816120bf565b610642826120bf565b61064b8361208d565b916106558461208d565b9361065f8161208d565b9560005b8281106106835750505061067f95969760405198899889611e07565b0390f35b8061069061074592611f8b565b905490841b1c60005260e08a8c6106cb846020600281526106b460406000206123bc565b946106c18387519261212e565b528401519261212e565b528d6106dc8460408401519261212e565b5260608101516106ec848961212e565b526106f7838861212e565b506080810151610707848a61212e565b52610712838961212e565b5060a0810151610722848b61212e565b5260c0810151610732848c61212e565b52015161073f828b61212e565b52612109565b610663565b346105f85760003660031901126105f8576005546107678161208d565b6107708261208d565b9061077a8361208d565b6107838461208d565b61078c856120bf565b610795866120bf565b61079e8761208d565b916107a88861208d565b936107b28961208d565b9560005b8a81106107cf57506040519950899861067f988a611d0c565b806107dc6108fd92611f54565b90549060031b1c6000526004602052610100604060002060096040519161080283611c29565b8054835260018101546020840152600281015460408401526003810154606084015261083060048201611c66565b608084015261084160058201611c66565b60a0840152600681015460c0840152600781015460e084015260088101548484015201546101208201528051610877848e61212e565b526020810151610887848f61212e565b526040810151610897848761212e565b5260608101516108a7848861212e565b5260808101516108b7848961212e565b526108c2838861212e565b5060a08101516108d2848a61212e565b526108dd838961212e565b5060c08101516108ed848b61212e565b5260e0810151610732848c61212e565b6107b6565b346105f85760203660031901126105f8576004356007548110156105f85761092b602091611e99565b90546040519160031b1c8152f35b346105f85760203660031901126105f85760043560005260086020526040600020805461067f60018301549260028101549061097760038201611c66565b61098360048301611c66565b6005830154906006840154926008600786015495015495604051998a998a611db4565b346105f857610a0e6109b736612018565b94604098979894919493929351966109ce88611c0d565b898852602088015260408701526060860152608085015260a084015260c083015260e0820152426101008201528260005260026020526040600020612142565b60035490600160401b8210156103a257610a318260016103a09401600355611f8b565b90919082549060031b91821b91600019901b1916179055565b346105f857610ab2610a5b36612018565b9460409897989491949392935196610a7288611c0d565b898852602088015260408701526060860152608085015260a084015260c083015260e0820152426101008201528260005260086020526040600020612142565b60095490600160401b8210156103a257610a318260016103a09401600955611f1d565b346105f85760203660031901126105f8576004356003548110156105f857610afe602091611f8b565b90549060031b1c604051908152f35b346105f85760203660031901126105f8576004356001548110156105f85761092b602091611ee6565b346105f857610120806003193601126105f8576084356001600160401b0381116105f857610b68903690600401611fc2565b60a4356001600160401b0381116105f857610b87903690600401611fc2565b60405191610b9483611c29565b600435835260208301916024358352604084016044358152606085019160643583526080860193845260a086015260c43560c086015260e43560e0860152610104356101008601524286860152600435600052600460205260406000209385518555516001850155516002840155516003830155519283516001600160401b0381116103a257610c276004840154611bb7565b601f8111610e49575b506020601f8211600114610ddc578192939495600092610dd1575b50508160011b916000199060031b1c19161760048301555b60a08301519283516001600160401b0381116103a257610c866005850154611bb7565b601f8111610d86575b506020601f8211600114610d1457819060099596600092610d09575b50508160011b916000199060031b1c19161760058501555b60c0810151600685015560e0810151600785015561010081015160088501550151910155600554600160401b8110156103a2576103868160016103a09301600555611f54565b015190508680610cab565b6005850160005260206000209560005b601f1984168110610d6e575095829160099697600194601f19811610610d55575b505050811b016005850155610cc3565b015160001960f88460031b161c19169055868080610d45565b82820151885560019097019660209283019201610d24565b600585016000526020600020601f830160051c81019160208410610dc7575b601f0160051c01905b818110610dbb5750610c8f565b60008155600101610dae565b9091508190610da5565b015190508580610c4b565b60048401600052602060002090600096601f198416975b888110610e3157508360019596979810610e18575b505050811b016004830155610c63565b015160001960f88460031b161c19169055858080610e08565b91926020600181928685015181550194019201610df3565b600484016000526020600020601f830160051c810160208410610e89575b601f830160051c82018110610e7d575050610c30565b60008155600101610e67565b5080610e67565b346105f85760203660031901126105f8576004356005548110156105f85761092b602091611f54565b346105f857610100806003193601126105f8576084356001600160401b0381116105f857610eeb903690600401611fc2565b60a4356001600160401b0381116105f857610f0a903690600401611fc2565b60405191610f1783611c0d565b600435835260208301916024358352604084016044358152606085019160643583526080860193845260a086015260c43560c086015260e43560e08601524286860152600435600052600060205260406000209385518555516001850155516002840155516003830155519283516001600160401b0381116103a257610fa06004840154611bb7565b601f81116111b0575b50602094601f821160011461114457948192939495600092611139575b50508160011b916000199060031b1c19161760048301555b600582019260a0810151938451946001600160401b0386116103a2576110048254611bb7565b601f81116110f1575b50602090601f8711600114611086579580916008969760009261107b575b50508160011b916000199060031b1c19161790555b60c0810151600685015560e081015160078501550151910155600154600160401b8110156103a2576103868160016103a09301600155611ee6565b01519050878061102b565b90601f198716918360005260206000209260005b8181106110d957509160019391896008999a94106110c0575b505050811b019055611040565b015160001960f88460031b161c191690558780806110b3565b9293602060018192878601518155019501930161109a565b826000526020600020601f880160051c8101916020891061112f575b601f0160051c01905b818110611123575061100d565b60008155600101611116565b909150819061110d565b015190508580610fc6565b601f198216956004850160005260206000209160005b8881106111985750836001959697981061117f575b505050811b016004830155610fde565b015160001960f88460031b161c1916905585808061116f565b9192602060018192868501518155019401920161115a565b600484016000526020600020601f830160051c810191602084106111f1575b601f0160051c01905b8181106111e55750610fa9565b600081556001016111d8565b90915081906111cf565b346105f85760203660031901126105f8576004356009548110156105f85761092b602091611f1d565b346105f85760003660031901126105f8576009546112418161208d565b61124a8261208d565b906112548361208d565b61125d846120bf565b611266856120bf565b61126f8661208d565b906112798761208d565b926112838861208d565b9460005b8981106112a057506040519850889761067f9789611e07565b806112ad61134592611f1d565b90549060031b1c60005260e0898b6112d2846020600881526106b460406000206123bc565b5260408101516112e2848761212e565b5260608101516112f2848861212e565b526112fd838761212e565b50608081015161130d848961212e565b52611318838861212e565b5060a0810151611328848a61212e565b5260c0810151611338848b61212e565b52015161073f828a61212e565b611287565b346105f85760203660031901126105f85760043560005260066020526040600020805460018201549160028101549060038101546004820154906005830154916006840161139790611c66565b6113a360078601611c66565b90600886015494600987015493600a88015495600b89016113c390611c66565b98600c0154996040519c8d9c8d5260208d015260408c015260608b015260808a015260a08901526101a08060c08a015288016113fe91611b1b565b87810360e089015261140f91611b1b565b9261010087015261012086015261014085015283810361016085015261143491611b1b565b906101808301520390f35b346105f85760203660031901126105f85760043560005260026020526040600020805461067f60018301549260028101549061097760038201611c66565b346105f85760203660031901126105f85760043560005260006020526040600020805460018201549160028101546003820154916115176114c060048301611c66565b6115096114cf60058501611c66565b9160068501549660086007870154960154966040519a8b9a8b5260208b015260408a015260608901526101208060808a0152880190611b1b565b9086820360a0880152611b1b565b9260c085015260e08401526101008301520390f35b346105f85760003660031901126105f8576001546115498161208d565b6115528261208d565b9061155c8361208d565b6115658461208d565b61156e856120bf565b611577866120bf565b6115808761208d565b9161158a8861208d565b936115948961208d565b9560005b8a81106115b157506040519950899861067f988a611d0c565b806115be61164e92611ee6565b90549060031b1c600052600060205261010060406000206008604051916115e483611c0d565b8054835260018101546020840152600281015460408401526003810154606084015261161260048201611c66565b608084015261162360058201611c66565b60a0840152600681015460c0840152600781015460e08401520154828201528051610877848e61212e565b611598565b346105f85760203660031901126105f857600435600052600460205260406000208054600182015491600281015460038201549161169360048201611c66565b926116f36116a360058401611c66565b6116e560068501549660078601549460096008880154970154976040519b8c9b8c5260208c015260408b015260608a01526101408060808b0152890190611b1b565b9087820360a0890152611b1b565b9360c086015260e08501526101008401526101208301520390f35b346105f85760003660031901126105f85760075461172b8161208d565b60a0526117378161208d565b6117408261208d565b60805261174c8261208d565b6117558361208d565b9161175f8461208d565b93611769816120bf565b60c052611775816120bf565b610140526117828161208d565b6101605261178f8161208d565b6101005261179c8161208d565b610180526117a9816120bf565b6101a0526000610120525b61012051818110156119f2576117c990611e99565b90549060031b1c6000526006602052600c60406000206117ee6040518060e052611bf1565b805460e051526001810154602060e05101526002810154604060e05101526003810154606060e05101526004810154608060e0510152600581015460a060e051015261183c60068201611c66565b60c060e051015261184f60078201611c66565b60e080510152600881015461010060e0510152600981015461012060e0510152600a81015461014060e0510152611888600b8201611c66565b61016060e0510152015461018060e051015260e051516118ad6101205160a05161212e565b52602060e05101516118c2610120518461212e565b52604060e05101516118d96101205160805161212e565b52606060e05101516118ee610120518561212e565b52608060e0510151611903610120518661212e565b5260a060e0510151611918610120518761212e565b5260c060e051015161192f6101205160c05161212e565b5261193f6101205160c05161212e565b5060e080510151611956610120516101405161212e565b52611967610120516101405161212e565b5061010060e0510151611980610120516101605161212e565b5261012060e0510151611999610120516101005161212e565b5261014060e05101516119b2610120516101805161212e565b5261016060e05101516119cb610120516101a05161212e565b526119dc610120516101a05161212e565b506119e961012051612109565b610120526117b4565b5050611a4361067f611ad5611ac3611ab1611a9f611a8e611a7e8b611a708c611a628d611a546040519e8f9e8f916101808352611a35610180840160a051611ae7565b908382036020850152611ae7565b906040818303910152608051611ae7565b8d810360608f015290611ae7565b908b820360808d0152611ae7565b9089820360a08b0152611ae7565b87810360c089015260c051611b5b565b86810360e088015261014051611b5b565b85810361010087015261016051611ae7565b84810361012086015261010051611ae7565b83810361014085015261018051611ae7565b8281036101608401526101a051611b5b565b90815180825260208080930193019160005b828110611b07575050505090565b835185529381019392810192600101611af9565b919082519283825260005b848110611b47575050826000602080949584010152601f8019910116010190565b602081830181015184830182015201611b26565b908082519081815260208091019281808460051b8301019501936000915b848310611b895750505050505090565b9091929394958480611ba7600193601f198682030187528a51611b1b565b9801930193019194939290611b79565b90600182811c92168015611be7575b6020831014611bd157565b634e487b7160e01b600052602260045260246000fd5b91607f1691611bc6565b6101a081019081106001600160401b038211176103a257604052565b61012081019081106001600160401b038211176103a257604052565b61014081019081106001600160401b038211176103a257604052565b90601f801991011681019081106001600160401b038211176103a257604052565b9060405191826000825492611c7a84611bb7565b908184526001948581169081600014611ce95750600114611ca6575b5050611ca492500383611c45565b565b9093915060005260209081600020936000915b818310611cd1575050611ca493508201013880611c96565b85548884018501529485019487945091830191611cb9565b915050611ca494506020925060ff191682840152151560051b8201013880611c96565b9794611d7890611d6a611d9496611d5c8c611db19e9c98611d4e611d8699611d40611da29f9a610120808752860190611ae7565b908482036020860152611ae7565b916040818403910152611ae7565b8c810360608e015290611ae7565b908a820360808c0152611b5b565b9088820360a08a0152611b5b565b9086820360c0880152611ae7565b9084820360e0860152611ae7565b91610100818403910152611ae7565b90565b95999897949193611df3936101009895611de593895260208901526040880152610120806060890152870190611b1b565b908582036080870152611b1b565b9660a084015260c083015260e08201520152565b969394611e61611e8b96611e53611e7d96611e458c611db19e9c98611e37611e6f99610100808552840190611ae7565b916020818403910152611ae7565b8c810360408e015290611ae7565b908a820360608c0152611b5b565b9088820360808a0152611b5b565b9086820360a0880152611ae7565b9084820360c0860152611ae7565b9160e0818403910152611ae7565b600754811015611ed05760076000527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880190600090565b634e487b7160e01b600052603260045260246000fd5b600154811015611ed05760016000527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60190600090565b600954811015611ed05760096000527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0190600090565b600554811015611ed05760056000527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00190600090565b600354811015611ed05760036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0190600090565b81601f820112156105f8578035906001600160401b0382116103a25760405192611ff6601f8401601f191660200185611c45565b828452602083830101116105f857816000926020809301838601378301015290565b6101006003198201126105f8576004359160243591604435916001600160401b036064358181116105f8578361205091600401611fc2565b926084359182116105f85761206791600401611fc2565b9060a4359060c4359060e43590565b6001600160401b0381116103a25760051b60200190565b9061209782612076565b6120a46040519182611c45565b82815280926120b5601f1991612076565b0190602036910137565b906120c982612076565b6120d66040519182611c45565b82815280926120e7601f1991612076565b019060005b8281106120f857505050565b8060606020809385010152016120ec565b60001981146121185760010190565b634e487b7160e01b600052601160045260246000fd5b8051821015611ed05760209160051b010190565b9080518255602090818101516001908185015560408201516002850155600384016060830151918251916001600160401b03928381116103a257806121878354611bb7565b95601f9687811161236a575b508890878311600114612307576000926122fc575b5050600019600383901b1c191690831b1790555b600486019260808501519586519384116103a2576121da8554611bb7565b8281116122b6575b508091841160011461224657509180809261010096959460089860009461223b575b50501b916000199060031b1c19161790555b60a0810151600585015560c0810151600685015560e081015160078501550151910155565b015192503880612204565b91939495601f1984168660005283600020936000905b82821061229f575050916008979593918561010098969410612286575b505050811b019055612216565b015160001960f88460031b161c19169055388080612279565b80888697829497870151815501960194019061225c565b85600052816000208380870160051c8201928488106122f3575b0160051c019084905b8281106122e75750506121e2565b600081550184906122d9565b925081926122d0565b0151905038806121a8565b90859350601f19831691856000528a6000209260005b8c828210612354575050841161233b575b505050811b0190556121bc565b015160001960f88460031b161c1916905538808061232e565b838501518655899790950194938401930161231d565b90915060008481528981208880860160051c8201938c87106123b3575b91889187969594930160051c01925b8381106123a557505050612193565b828155869550889101612396565b93508193612387565b906040516123c981611c0d565b610100600882948054845260018101546020850152600281015460408501526123f460038201611c66565b606085015261240560048201611c66565b6080850152600581015460a0850152600681015460c0850152600781015460e0850152015491015256fea264697066735822122052c97c799fa3a53399471c48307ff92090c3c39b372780250b2d3615b737d63a64736f6c63430008130033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DistributionIds is a free data retrieval call binding the contract method 0xacad4949.
//
// Solidity: function distributionIds(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) DistributionIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "distributionIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistributionIds is a free data retrieval call binding the contract method 0xacad4949.
//
// Solidity: function distributionIds(uint256 ) view returns(uint256)
func (_Contract *ContractSession) DistributionIds(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.DistributionIds(&_Contract.CallOpts, arg0)
}

// DistributionIds is a free data retrieval call binding the contract method 0xacad4949.
//
// Solidity: function distributionIds(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) DistributionIds(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.DistributionIds(&_Contract.CallOpts, arg0)
}

// Distributions is a free data retrieval call binding the contract method 0x4487d3df.
//
// Solidity: function distributions(uint256 ) view returns(uint256 id, uint256 distributorProfileId, uint256 destinationId, uint256 harvestId, uint256 harvestCollectorId, uint256 harvestProcessorId, string name, string desc, uint256 quantity, uint256 basePrice, uint256 price, string transportation, uint256 createdAt)
func (_Contract *ContractCaller) Distributions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                   *big.Int
	DistributorProfileId *big.Int
	DestinationId        *big.Int
	HarvestId            *big.Int
	HarvestCollectorId   *big.Int
	HarvestProcessorId   *big.Int
	Name                 string
	Desc                 string
	Quantity             *big.Int
	BasePrice            *big.Int
	Price                *big.Int
	Transportation       string
	CreatedAt            *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "distributions", arg0)

	outstruct := new(struct {
		Id                   *big.Int
		DistributorProfileId *big.Int
		DestinationId        *big.Int
		HarvestId            *big.Int
		HarvestCollectorId   *big.Int
		HarvestProcessorId   *big.Int
		Name                 string
		Desc                 string
		Quantity             *big.Int
		BasePrice            *big.Int
		Price                *big.Int
		Transportation       string
		CreatedAt            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DistributorProfileId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DestinationId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.HarvestId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.HarvestCollectorId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HarvestProcessorId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Desc = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.Quantity = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.BasePrice = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Transportation = *abi.ConvertType(out[11], new(string)).(*string)
	outstruct.CreatedAt = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Distributions is a free data retrieval call binding the contract method 0x4487d3df.
//
// Solidity: function distributions(uint256 ) view returns(uint256 id, uint256 distributorProfileId, uint256 destinationId, uint256 harvestId, uint256 harvestCollectorId, uint256 harvestProcessorId, string name, string desc, uint256 quantity, uint256 basePrice, uint256 price, string transportation, uint256 createdAt)
func (_Contract *ContractSession) Distributions(arg0 *big.Int) (struct {
	Id                   *big.Int
	DistributorProfileId *big.Int
	DestinationId        *big.Int
	HarvestId            *big.Int
	HarvestCollectorId   *big.Int
	HarvestProcessorId   *big.Int
	Name                 string
	Desc                 string
	Quantity             *big.Int
	BasePrice            *big.Int
	Price                *big.Int
	Transportation       string
	CreatedAt            *big.Int
}, error) {
	return _Contract.Contract.Distributions(&_Contract.CallOpts, arg0)
}

// Distributions is a free data retrieval call binding the contract method 0x4487d3df.
//
// Solidity: function distributions(uint256 ) view returns(uint256 id, uint256 distributorProfileId, uint256 destinationId, uint256 harvestId, uint256 harvestCollectorId, uint256 harvestProcessorId, string name, string desc, uint256 quantity, uint256 basePrice, uint256 price, string transportation, uint256 createdAt)
func (_Contract *ContractCallerSession) Distributions(arg0 *big.Int) (struct {
	Id                   *big.Int
	DistributorProfileId *big.Int
	DestinationId        *big.Int
	HarvestId            *big.Int
	HarvestCollectorId   *big.Int
	HarvestProcessorId   *big.Int
	Name                 string
	Desc                 string
	Quantity             *big.Int
	BasePrice            *big.Int
	Price                *big.Int
	Transportation       string
	CreatedAt            *big.Int
}, error) {
	return _Contract.Contract.Distributions(&_Contract.CallOpts, arg0)
}

// GetAllDistributionTuples is a free data retrieval call binding the contract method 0x184024a1.
//
// Solidity: function getAllDistributionTuples() view returns(uint256[], uint256[], uint256[], uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[], string[])
func (_Contract *ContractCaller) GetAllDistributionTuples(opts *bind.CallOpts) ([]*big.Int, []*big.Int, []*big.Int, []*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, []string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAllDistributionTuples")

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]string), *new([]string), *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	out3 := *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	out4 := *abi.ConvertType(out[4], new([]*big.Int)).(*[]*big.Int)
	out5 := *abi.ConvertType(out[5], new([]*big.Int)).(*[]*big.Int)
	out6 := *abi.ConvertType(out[6], new([]string)).(*[]string)
	out7 := *abi.ConvertType(out[7], new([]string)).(*[]string)
	out8 := *abi.ConvertType(out[8], new([]*big.Int)).(*[]*big.Int)
	out9 := *abi.ConvertType(out[9], new([]*big.Int)).(*[]*big.Int)
	out10 := *abi.ConvertType(out[10], new([]*big.Int)).(*[]*big.Int)
	out11 := *abi.ConvertType(out[11], new([]string)).(*[]string)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, out11, err

}

// GetAllDistributionTuples is a free data retrieval call binding the contract method 0x184024a1.
//
// Solidity: function getAllDistributionTuples() view returns(uint256[], uint256[], uint256[], uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[], string[])
func (_Contract *ContractSession) GetAllDistributionTuples() ([]*big.Int, []*big.Int, []*big.Int, []*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, []string, error) {
	return _Contract.Contract.GetAllDistributionTuples(&_Contract.CallOpts)
}

// GetAllDistributionTuples is a free data retrieval call binding the contract method 0x184024a1.
//
// Solidity: function getAllDistributionTuples() view returns(uint256[], uint256[], uint256[], uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[], string[])
func (_Contract *ContractCallerSession) GetAllDistributionTuples() ([]*big.Int, []*big.Int, []*big.Int, []*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, []string, error) {
	return _Contract.Contract.GetAllDistributionTuples(&_Contract.CallOpts)
}

// GetAllHarvestCollectorTuples is a free data retrieval call binding the contract method 0xbdec4ed7.
//
// Solidity: function getAllHarvestCollectorTuples() view returns(uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractCaller) GetAllHarvestCollectorTuples(opts *bind.CallOpts) ([]*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAllHarvestCollectorTuples")

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]string), *new([]string), *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	out3 := *abi.ConvertType(out[3], new([]string)).(*[]string)
	out4 := *abi.ConvertType(out[4], new([]string)).(*[]string)
	out5 := *abi.ConvertType(out[5], new([]*big.Int)).(*[]*big.Int)
	out6 := *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	out7 := *abi.ConvertType(out[7], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetAllHarvestCollectorTuples is a free data retrieval call binding the contract method 0xbdec4ed7.
//
// Solidity: function getAllHarvestCollectorTuples() view returns(uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractSession) GetAllHarvestCollectorTuples() ([]*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	return _Contract.Contract.GetAllHarvestCollectorTuples(&_Contract.CallOpts)
}

// GetAllHarvestCollectorTuples is a free data retrieval call binding the contract method 0xbdec4ed7.
//
// Solidity: function getAllHarvestCollectorTuples() view returns(uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractCallerSession) GetAllHarvestCollectorTuples() ([]*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	return _Contract.Contract.GetAllHarvestCollectorTuples(&_Contract.CallOpts)
}

// GetAllHarvestProcessorTuples is a free data retrieval call binding the contract method 0xb7d5e2da.
//
// Solidity: function getAllHarvestProcessorTuples() view returns(uint256[], uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractCaller) GetAllHarvestProcessorTuples(opts *bind.CallOpts) ([]*big.Int, []*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAllHarvestProcessorTuples")

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]string), *new([]string), *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	out3 := *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	out4 := *abi.ConvertType(out[4], new([]string)).(*[]string)
	out5 := *abi.ConvertType(out[5], new([]string)).(*[]string)
	out6 := *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	out7 := *abi.ConvertType(out[7], new([]*big.Int)).(*[]*big.Int)
	out8 := *abi.ConvertType(out[8], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAllHarvestProcessorTuples is a free data retrieval call binding the contract method 0xb7d5e2da.
//
// Solidity: function getAllHarvestProcessorTuples() view returns(uint256[], uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractSession) GetAllHarvestProcessorTuples() ([]*big.Int, []*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	return _Contract.Contract.GetAllHarvestProcessorTuples(&_Contract.CallOpts)
}

// GetAllHarvestProcessorTuples is a free data retrieval call binding the contract method 0xb7d5e2da.
//
// Solidity: function getAllHarvestProcessorTuples() view returns(uint256[], uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractCallerSession) GetAllHarvestProcessorTuples() ([]*big.Int, []*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	return _Contract.Contract.GetAllHarvestProcessorTuples(&_Contract.CallOpts)
}

// GetAllHarvestTuples is a free data retrieval call binding the contract method 0x1fde561c.
//
// Solidity: function getAllHarvestTuples() view returns(uint256[], uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractCaller) GetAllHarvestTuples(opts *bind.CallOpts) ([]*big.Int, []*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAllHarvestTuples")

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]string), *new([]string), *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	out3 := *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	out4 := *abi.ConvertType(out[4], new([]string)).(*[]string)
	out5 := *abi.ConvertType(out[5], new([]string)).(*[]string)
	out6 := *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	out7 := *abi.ConvertType(out[7], new([]*big.Int)).(*[]*big.Int)
	out8 := *abi.ConvertType(out[8], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAllHarvestTuples is a free data retrieval call binding the contract method 0x1fde561c.
//
// Solidity: function getAllHarvestTuples() view returns(uint256[], uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractSession) GetAllHarvestTuples() ([]*big.Int, []*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	return _Contract.Contract.GetAllHarvestTuples(&_Contract.CallOpts)
}

// GetAllHarvestTuples is a free data retrieval call binding the contract method 0x1fde561c.
//
// Solidity: function getAllHarvestTuples() view returns(uint256[], uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractCallerSession) GetAllHarvestTuples() ([]*big.Int, []*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	return _Contract.Contract.GetAllHarvestTuples(&_Contract.CallOpts)
}

// GetAllSellerBoxTuples is a free data retrieval call binding the contract method 0x4841fe1e.
//
// Solidity: function getAllSellerBoxTuples() view returns(uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractCaller) GetAllSellerBoxTuples(opts *bind.CallOpts) ([]*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAllSellerBoxTuples")

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), *new([]string), *new([]string), *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	out3 := *abi.ConvertType(out[3], new([]string)).(*[]string)
	out4 := *abi.ConvertType(out[4], new([]string)).(*[]string)
	out5 := *abi.ConvertType(out[5], new([]*big.Int)).(*[]*big.Int)
	out6 := *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	out7 := *abi.ConvertType(out[7], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetAllSellerBoxTuples is a free data retrieval call binding the contract method 0x4841fe1e.
//
// Solidity: function getAllSellerBoxTuples() view returns(uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractSession) GetAllSellerBoxTuples() ([]*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	return _Contract.Contract.GetAllSellerBoxTuples(&_Contract.CallOpts)
}

// GetAllSellerBoxTuples is a free data retrieval call binding the contract method 0x4841fe1e.
//
// Solidity: function getAllSellerBoxTuples() view returns(uint256[], uint256[], uint256[], string[], string[], uint256[], uint256[], uint256[])
func (_Contract *ContractCallerSession) GetAllSellerBoxTuples() ([]*big.Int, []*big.Int, []*big.Int, []string, []string, []*big.Int, []*big.Int, []*big.Int, error) {
	return _Contract.Contract.GetAllSellerBoxTuples(&_Contract.CallOpts)
}

// HarvestCollectorIds is a free data retrieval call binding the contract method 0x8a8b1ece.
//
// Solidity: function harvestCollectorIds(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) HarvestCollectorIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "harvestCollectorIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestCollectorIds is a free data retrieval call binding the contract method 0x8a8b1ece.
//
// Solidity: function harvestCollectorIds(uint256 ) view returns(uint256)
func (_Contract *ContractSession) HarvestCollectorIds(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.HarvestCollectorIds(&_Contract.CallOpts, arg0)
}

// HarvestCollectorIds is a free data retrieval call binding the contract method 0x8a8b1ece.
//
// Solidity: function harvestCollectorIds(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) HarvestCollectorIds(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.HarvestCollectorIds(&_Contract.CallOpts, arg0)
}

// HarvestCollectors is a free data retrieval call binding the contract method 0x3d6140ce.
//
// Solidity: function harvestCollectors(uint256 ) view returns(uint256 id, uint256 collectorProfileId, uint256 harvestId, string name, string desc, uint256 quantity, uint256 price, uint256 basePrice, uint256 createdAt)
func (_Contract *ContractCaller) HarvestCollectors(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                 *big.Int
	CollectorProfileId *big.Int
	HarvestId          *big.Int
	Name               string
	Desc               string
	Quantity           *big.Int
	Price              *big.Int
	BasePrice          *big.Int
	CreatedAt          *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "harvestCollectors", arg0)

	outstruct := new(struct {
		Id                 *big.Int
		CollectorProfileId *big.Int
		HarvestId          *big.Int
		Name               string
		Desc               string
		Quantity           *big.Int
		Price              *big.Int
		BasePrice          *big.Int
		CreatedAt          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CollectorProfileId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HarvestId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Desc = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Quantity = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.BasePrice = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// HarvestCollectors is a free data retrieval call binding the contract method 0x3d6140ce.
//
// Solidity: function harvestCollectors(uint256 ) view returns(uint256 id, uint256 collectorProfileId, uint256 harvestId, string name, string desc, uint256 quantity, uint256 price, uint256 basePrice, uint256 createdAt)
func (_Contract *ContractSession) HarvestCollectors(arg0 *big.Int) (struct {
	Id                 *big.Int
	CollectorProfileId *big.Int
	HarvestId          *big.Int
	Name               string
	Desc               string
	Quantity           *big.Int
	Price              *big.Int
	BasePrice          *big.Int
	CreatedAt          *big.Int
}, error) {
	return _Contract.Contract.HarvestCollectors(&_Contract.CallOpts, arg0)
}

// HarvestCollectors is a free data retrieval call binding the contract method 0x3d6140ce.
//
// Solidity: function harvestCollectors(uint256 ) view returns(uint256 id, uint256 collectorProfileId, uint256 harvestId, string name, string desc, uint256 quantity, uint256 price, uint256 basePrice, uint256 createdAt)
func (_Contract *ContractCallerSession) HarvestCollectors(arg0 *big.Int) (struct {
	Id                 *big.Int
	CollectorProfileId *big.Int
	HarvestId          *big.Int
	Name               string
	Desc               string
	Quantity           *big.Int
	Price              *big.Int
	BasePrice          *big.Int
	CreatedAt          *big.Int
}, error) {
	return _Contract.Contract.HarvestCollectors(&_Contract.CallOpts, arg0)
}

// HarvestIds is a free data retrieval call binding the contract method 0x867c23f3.
//
// Solidity: function harvestIds(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) HarvestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "harvestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestIds is a free data retrieval call binding the contract method 0x867c23f3.
//
// Solidity: function harvestIds(uint256 ) view returns(uint256)
func (_Contract *ContractSession) HarvestIds(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.HarvestIds(&_Contract.CallOpts, arg0)
}

// HarvestIds is a free data retrieval call binding the contract method 0x867c23f3.
//
// Solidity: function harvestIds(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) HarvestIds(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.HarvestIds(&_Contract.CallOpts, arg0)
}

// HarvestProcessorIds is a free data retrieval call binding the contract method 0x6910f260.
//
// Solidity: function harvestProcessorIds(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) HarvestProcessorIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "harvestProcessorIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProcessorIds is a free data retrieval call binding the contract method 0x6910f260.
//
// Solidity: function harvestProcessorIds(uint256 ) view returns(uint256)
func (_Contract *ContractSession) HarvestProcessorIds(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.HarvestProcessorIds(&_Contract.CallOpts, arg0)
}

// HarvestProcessorIds is a free data retrieval call binding the contract method 0x6910f260.
//
// Solidity: function harvestProcessorIds(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) HarvestProcessorIds(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.HarvestProcessorIds(&_Contract.CallOpts, arg0)
}

// HarvestProcessors is a free data retrieval call binding the contract method 0x1a84ce9b.
//
// Solidity: function harvestProcessors(uint256 ) view returns(uint256 id, uint256 processorProfileId, uint256 harvestCollectorId, uint256 harvestId, string name, string desc, uint256 quantity, uint256 basePrice, uint256 price, uint256 createdAt)
func (_Contract *ContractCaller) HarvestProcessors(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                 *big.Int
	ProcessorProfileId *big.Int
	HarvestCollectorId *big.Int
	HarvestId          *big.Int
	Name               string
	Desc               string
	Quantity           *big.Int
	BasePrice          *big.Int
	Price              *big.Int
	CreatedAt          *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "harvestProcessors", arg0)

	outstruct := new(struct {
		Id                 *big.Int
		ProcessorProfileId *big.Int
		HarvestCollectorId *big.Int
		HarvestId          *big.Int
		Name               string
		Desc               string
		Quantity           *big.Int
		BasePrice          *big.Int
		Price              *big.Int
		CreatedAt          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ProcessorProfileId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HarvestCollectorId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.HarvestId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Desc = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Quantity = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.BasePrice = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// HarvestProcessors is a free data retrieval call binding the contract method 0x1a84ce9b.
//
// Solidity: function harvestProcessors(uint256 ) view returns(uint256 id, uint256 processorProfileId, uint256 harvestCollectorId, uint256 harvestId, string name, string desc, uint256 quantity, uint256 basePrice, uint256 price, uint256 createdAt)
func (_Contract *ContractSession) HarvestProcessors(arg0 *big.Int) (struct {
	Id                 *big.Int
	ProcessorProfileId *big.Int
	HarvestCollectorId *big.Int
	HarvestId          *big.Int
	Name               string
	Desc               string
	Quantity           *big.Int
	BasePrice          *big.Int
	Price              *big.Int
	CreatedAt          *big.Int
}, error) {
	return _Contract.Contract.HarvestProcessors(&_Contract.CallOpts, arg0)
}

// HarvestProcessors is a free data retrieval call binding the contract method 0x1a84ce9b.
//
// Solidity: function harvestProcessors(uint256 ) view returns(uint256 id, uint256 processorProfileId, uint256 harvestCollectorId, uint256 harvestId, string name, string desc, uint256 quantity, uint256 basePrice, uint256 price, uint256 createdAt)
func (_Contract *ContractCallerSession) HarvestProcessors(arg0 *big.Int) (struct {
	Id                 *big.Int
	ProcessorProfileId *big.Int
	HarvestCollectorId *big.Int
	HarvestId          *big.Int
	Name               string
	Desc               string
	Quantity           *big.Int
	BasePrice          *big.Int
	Price              *big.Int
	CreatedAt          *big.Int
}, error) {
	return _Contract.Contract.HarvestProcessors(&_Contract.CallOpts, arg0)
}

// Harvests is a free data retrieval call binding the contract method 0x35a423b0.
//
// Solidity: function harvests(uint256 ) view returns(uint256 id, uint256 farmerProfileId, uint256 cropId, uint256 regencyId, string name, string description, uint256 quantity, uint256 basePrice, uint256 createdAt)
func (_Contract *ContractCaller) Harvests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              *big.Int
	FarmerProfileId *big.Int
	CropId          *big.Int
	RegencyId       *big.Int
	Name            string
	Description     string
	Quantity        *big.Int
	BasePrice       *big.Int
	CreatedAt       *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "harvests", arg0)

	outstruct := new(struct {
		Id              *big.Int
		FarmerProfileId *big.Int
		CropId          *big.Int
		RegencyId       *big.Int
		Name            string
		Description     string
		Quantity        *big.Int
		BasePrice       *big.Int
		CreatedAt       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FarmerProfileId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CropId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RegencyId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Quantity = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.BasePrice = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Harvests is a free data retrieval call binding the contract method 0x35a423b0.
//
// Solidity: function harvests(uint256 ) view returns(uint256 id, uint256 farmerProfileId, uint256 cropId, uint256 regencyId, string name, string description, uint256 quantity, uint256 basePrice, uint256 createdAt)
func (_Contract *ContractSession) Harvests(arg0 *big.Int) (struct {
	Id              *big.Int
	FarmerProfileId *big.Int
	CropId          *big.Int
	RegencyId       *big.Int
	Name            string
	Description     string
	Quantity        *big.Int
	BasePrice       *big.Int
	CreatedAt       *big.Int
}, error) {
	return _Contract.Contract.Harvests(&_Contract.CallOpts, arg0)
}

// Harvests is a free data retrieval call binding the contract method 0x35a423b0.
//
// Solidity: function harvests(uint256 ) view returns(uint256 id, uint256 farmerProfileId, uint256 cropId, uint256 regencyId, string name, string description, uint256 quantity, uint256 basePrice, uint256 createdAt)
func (_Contract *ContractCallerSession) Harvests(arg0 *big.Int) (struct {
	Id              *big.Int
	FarmerProfileId *big.Int
	CropId          *big.Int
	RegencyId       *big.Int
	Name            string
	Description     string
	Quantity        *big.Int
	BasePrice       *big.Int
	CreatedAt       *big.Int
}, error) {
	return _Contract.Contract.Harvests(&_Contract.CallOpts, arg0)
}

// SellerBoxIds is a free data retrieval call binding the contract method 0x579255b3.
//
// Solidity: function sellerBoxIds(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) SellerBoxIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "sellerBoxIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellerBoxIds is a free data retrieval call binding the contract method 0x579255b3.
//
// Solidity: function sellerBoxIds(uint256 ) view returns(uint256)
func (_Contract *ContractSession) SellerBoxIds(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SellerBoxIds(&_Contract.CallOpts, arg0)
}

// SellerBoxIds is a free data retrieval call binding the contract method 0x579255b3.
//
// Solidity: function sellerBoxIds(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) SellerBoxIds(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SellerBoxIds(&_Contract.CallOpts, arg0)
}

// SellerBoxes is a free data retrieval call binding the contract method 0xabd9f6c4.
//
// Solidity: function sellerBoxes(uint256 ) view returns(uint256 id, uint256 sellerProfileId, uint256 distributionId, string name, string desc, uint256 quantity, uint256 basePrice, uint256 price, uint256 createdAt)
func (_Contract *ContractCaller) SellerBoxes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              *big.Int
	SellerProfileId *big.Int
	DistributionId  *big.Int
	Name            string
	Desc            string
	Quantity        *big.Int
	BasePrice       *big.Int
	Price           *big.Int
	CreatedAt       *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "sellerBoxes", arg0)

	outstruct := new(struct {
		Id              *big.Int
		SellerProfileId *big.Int
		DistributionId  *big.Int
		Name            string
		Desc            string
		Quantity        *big.Int
		BasePrice       *big.Int
		Price           *big.Int
		CreatedAt       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SellerProfileId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DistributionId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Desc = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Quantity = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.BasePrice = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SellerBoxes is a free data retrieval call binding the contract method 0xabd9f6c4.
//
// Solidity: function sellerBoxes(uint256 ) view returns(uint256 id, uint256 sellerProfileId, uint256 distributionId, string name, string desc, uint256 quantity, uint256 basePrice, uint256 price, uint256 createdAt)
func (_Contract *ContractSession) SellerBoxes(arg0 *big.Int) (struct {
	Id              *big.Int
	SellerProfileId *big.Int
	DistributionId  *big.Int
	Name            string
	Desc            string
	Quantity        *big.Int
	BasePrice       *big.Int
	Price           *big.Int
	CreatedAt       *big.Int
}, error) {
	return _Contract.Contract.SellerBoxes(&_Contract.CallOpts, arg0)
}

// SellerBoxes is a free data retrieval call binding the contract method 0xabd9f6c4.
//
// Solidity: function sellerBoxes(uint256 ) view returns(uint256 id, uint256 sellerProfileId, uint256 distributionId, string name, string desc, uint256 quantity, uint256 basePrice, uint256 price, uint256 createdAt)
func (_Contract *ContractCallerSession) SellerBoxes(arg0 *big.Int) (struct {
	Id              *big.Int
	SellerProfileId *big.Int
	DistributionId  *big.Int
	Name            string
	Desc            string
	Quantity        *big.Int
	BasePrice       *big.Int
	Price           *big.Int
	CreatedAt       *big.Int
}, error) {
	return _Contract.Contract.SellerBoxes(&_Contract.CallOpts, arg0)
}

// CreateDistribution is a paid mutator transaction binding the contract method 0xe59d49d4.
//
// Solidity: function createDistribution(uint256 _id, uint256 _distributorProfileId, uint256 _destinationId, uint256 _harvestId, uint256 _harvestCollectorId, uint256 _harvestProcessorId, string _name, string _desc, uint256 _quantity, uint256 _basePrice, uint256 _price, string _transportation) returns()
func (_Contract *ContractTransactor) CreateDistribution(opts *bind.TransactOpts, _id *big.Int, _distributorProfileId *big.Int, _destinationId *big.Int, _harvestId *big.Int, _harvestCollectorId *big.Int, _harvestProcessorId *big.Int, _name string, _desc string, _quantity *big.Int, _basePrice *big.Int, _price *big.Int, _transportation string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createDistribution", _id, _distributorProfileId, _destinationId, _harvestId, _harvestCollectorId, _harvestProcessorId, _name, _desc, _quantity, _basePrice, _price, _transportation)
}

// CreateDistribution is a paid mutator transaction binding the contract method 0xe59d49d4.
//
// Solidity: function createDistribution(uint256 _id, uint256 _distributorProfileId, uint256 _destinationId, uint256 _harvestId, uint256 _harvestCollectorId, uint256 _harvestProcessorId, string _name, string _desc, uint256 _quantity, uint256 _basePrice, uint256 _price, string _transportation) returns()
func (_Contract *ContractSession) CreateDistribution(_id *big.Int, _distributorProfileId *big.Int, _destinationId *big.Int, _harvestId *big.Int, _harvestCollectorId *big.Int, _harvestProcessorId *big.Int, _name string, _desc string, _quantity *big.Int, _basePrice *big.Int, _price *big.Int, _transportation string) (*types.Transaction, error) {
	return _Contract.Contract.CreateDistribution(&_Contract.TransactOpts, _id, _distributorProfileId, _destinationId, _harvestId, _harvestCollectorId, _harvestProcessorId, _name, _desc, _quantity, _basePrice, _price, _transportation)
}

// CreateDistribution is a paid mutator transaction binding the contract method 0xe59d49d4.
//
// Solidity: function createDistribution(uint256 _id, uint256 _distributorProfileId, uint256 _destinationId, uint256 _harvestId, uint256 _harvestCollectorId, uint256 _harvestProcessorId, string _name, string _desc, uint256 _quantity, uint256 _basePrice, uint256 _price, string _transportation) returns()
func (_Contract *ContractTransactorSession) CreateDistribution(_id *big.Int, _distributorProfileId *big.Int, _destinationId *big.Int, _harvestId *big.Int, _harvestCollectorId *big.Int, _harvestProcessorId *big.Int, _name string, _desc string, _quantity *big.Int, _basePrice *big.Int, _price *big.Int, _transportation string) (*types.Transaction, error) {
	return _Contract.Contract.CreateDistribution(&_Contract.TransactOpts, _id, _distributorProfileId, _destinationId, _harvestId, _harvestCollectorId, _harvestProcessorId, _name, _desc, _quantity, _basePrice, _price, _transportation)
}

// CreateHarvest is a paid mutator transaction binding the contract method 0x5b80e7d8.
//
// Solidity: function createHarvest(uint256 _id, uint256 _farmerProfileId, uint256 _cropId, uint256 _regencyId, string _name, string _description, uint256 _quantity, uint256 _basePrice) returns()
func (_Contract *ContractTransactor) CreateHarvest(opts *bind.TransactOpts, _id *big.Int, _farmerProfileId *big.Int, _cropId *big.Int, _regencyId *big.Int, _name string, _description string, _quantity *big.Int, _basePrice *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createHarvest", _id, _farmerProfileId, _cropId, _regencyId, _name, _description, _quantity, _basePrice)
}

// CreateHarvest is a paid mutator transaction binding the contract method 0x5b80e7d8.
//
// Solidity: function createHarvest(uint256 _id, uint256 _farmerProfileId, uint256 _cropId, uint256 _regencyId, string _name, string _description, uint256 _quantity, uint256 _basePrice) returns()
func (_Contract *ContractSession) CreateHarvest(_id *big.Int, _farmerProfileId *big.Int, _cropId *big.Int, _regencyId *big.Int, _name string, _description string, _quantity *big.Int, _basePrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateHarvest(&_Contract.TransactOpts, _id, _farmerProfileId, _cropId, _regencyId, _name, _description, _quantity, _basePrice)
}

// CreateHarvest is a paid mutator transaction binding the contract method 0x5b80e7d8.
//
// Solidity: function createHarvest(uint256 _id, uint256 _farmerProfileId, uint256 _cropId, uint256 _regencyId, string _name, string _description, uint256 _quantity, uint256 _basePrice) returns()
func (_Contract *ContractTransactorSession) CreateHarvest(_id *big.Int, _farmerProfileId *big.Int, _cropId *big.Int, _regencyId *big.Int, _name string, _description string, _quantity *big.Int, _basePrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateHarvest(&_Contract.TransactOpts, _id, _farmerProfileId, _cropId, _regencyId, _name, _description, _quantity, _basePrice)
}

// CreateHarvestCollector is a paid mutator transaction binding the contract method 0x98728e9a.
//
// Solidity: function createHarvestCollector(uint256 _id, uint256 _collectorProfileId, uint256 _harvestId, string _name, string _desc, uint256 _quantity, uint256 _price, uint256 _basePrice) returns()
func (_Contract *ContractTransactor) CreateHarvestCollector(opts *bind.TransactOpts, _id *big.Int, _collectorProfileId *big.Int, _harvestId *big.Int, _name string, _desc string, _quantity *big.Int, _price *big.Int, _basePrice *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createHarvestCollector", _id, _collectorProfileId, _harvestId, _name, _desc, _quantity, _price, _basePrice)
}

// CreateHarvestCollector is a paid mutator transaction binding the contract method 0x98728e9a.
//
// Solidity: function createHarvestCollector(uint256 _id, uint256 _collectorProfileId, uint256 _harvestId, string _name, string _desc, uint256 _quantity, uint256 _price, uint256 _basePrice) returns()
func (_Contract *ContractSession) CreateHarvestCollector(_id *big.Int, _collectorProfileId *big.Int, _harvestId *big.Int, _name string, _desc string, _quantity *big.Int, _price *big.Int, _basePrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateHarvestCollector(&_Contract.TransactOpts, _id, _collectorProfileId, _harvestId, _name, _desc, _quantity, _price, _basePrice)
}

// CreateHarvestCollector is a paid mutator transaction binding the contract method 0x98728e9a.
//
// Solidity: function createHarvestCollector(uint256 _id, uint256 _collectorProfileId, uint256 _harvestId, string _name, string _desc, uint256 _quantity, uint256 _price, uint256 _basePrice) returns()
func (_Contract *ContractTransactorSession) CreateHarvestCollector(_id *big.Int, _collectorProfileId *big.Int, _harvestId *big.Int, _name string, _desc string, _quantity *big.Int, _price *big.Int, _basePrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateHarvestCollector(&_Contract.TransactOpts, _id, _collectorProfileId, _harvestId, _name, _desc, _quantity, _price, _basePrice)
}

// CreateHarvestProcessor is a paid mutator transaction binding the contract method 0x79360aa1.
//
// Solidity: function createHarvestProcessor(uint256 _id, uint256 _processorProfileId, uint256 _harvestCollectorId, uint256 _harvestId, string _name, string _desc, uint256 _quantity, uint256 _basePrice, uint256 _price) returns()
func (_Contract *ContractTransactor) CreateHarvestProcessor(opts *bind.TransactOpts, _id *big.Int, _processorProfileId *big.Int, _harvestCollectorId *big.Int, _harvestId *big.Int, _name string, _desc string, _quantity *big.Int, _basePrice *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createHarvestProcessor", _id, _processorProfileId, _harvestCollectorId, _harvestId, _name, _desc, _quantity, _basePrice, _price)
}

// CreateHarvestProcessor is a paid mutator transaction binding the contract method 0x79360aa1.
//
// Solidity: function createHarvestProcessor(uint256 _id, uint256 _processorProfileId, uint256 _harvestCollectorId, uint256 _harvestId, string _name, string _desc, uint256 _quantity, uint256 _basePrice, uint256 _price) returns()
func (_Contract *ContractSession) CreateHarvestProcessor(_id *big.Int, _processorProfileId *big.Int, _harvestCollectorId *big.Int, _harvestId *big.Int, _name string, _desc string, _quantity *big.Int, _basePrice *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateHarvestProcessor(&_Contract.TransactOpts, _id, _processorProfileId, _harvestCollectorId, _harvestId, _name, _desc, _quantity, _basePrice, _price)
}

// CreateHarvestProcessor is a paid mutator transaction binding the contract method 0x79360aa1.
//
// Solidity: function createHarvestProcessor(uint256 _id, uint256 _processorProfileId, uint256 _harvestCollectorId, uint256 _harvestId, string _name, string _desc, uint256 _quantity, uint256 _basePrice, uint256 _price) returns()
func (_Contract *ContractTransactorSession) CreateHarvestProcessor(_id *big.Int, _processorProfileId *big.Int, _harvestCollectorId *big.Int, _harvestId *big.Int, _name string, _desc string, _quantity *big.Int, _basePrice *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateHarvestProcessor(&_Contract.TransactOpts, _id, _processorProfileId, _harvestCollectorId, _harvestId, _name, _desc, _quantity, _basePrice, _price)
}

// CreateSellerBox is a paid mutator transaction binding the contract method 0x92bb06b3.
//
// Solidity: function createSellerBox(uint256 _id, uint256 _sellerProfileId, uint256 _distributionId, string _name, string _desc, uint256 _quantity, uint256 _basePrice, uint256 _price) returns()
func (_Contract *ContractTransactor) CreateSellerBox(opts *bind.TransactOpts, _id *big.Int, _sellerProfileId *big.Int, _distributionId *big.Int, _name string, _desc string, _quantity *big.Int, _basePrice *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createSellerBox", _id, _sellerProfileId, _distributionId, _name, _desc, _quantity, _basePrice, _price)
}

// CreateSellerBox is a paid mutator transaction binding the contract method 0x92bb06b3.
//
// Solidity: function createSellerBox(uint256 _id, uint256 _sellerProfileId, uint256 _distributionId, string _name, string _desc, uint256 _quantity, uint256 _basePrice, uint256 _price) returns()
func (_Contract *ContractSession) CreateSellerBox(_id *big.Int, _sellerProfileId *big.Int, _distributionId *big.Int, _name string, _desc string, _quantity *big.Int, _basePrice *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateSellerBox(&_Contract.TransactOpts, _id, _sellerProfileId, _distributionId, _name, _desc, _quantity, _basePrice, _price)
}

// CreateSellerBox is a paid mutator transaction binding the contract method 0x92bb06b3.
//
// Solidity: function createSellerBox(uint256 _id, uint256 _sellerProfileId, uint256 _distributionId, string _name, string _desc, uint256 _quantity, uint256 _basePrice, uint256 _price) returns()
func (_Contract *ContractTransactorSession) CreateSellerBox(_id *big.Int, _sellerProfileId *big.Int, _distributionId *big.Int, _name string, _desc string, _quantity *big.Int, _basePrice *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateSellerBox(&_Contract.TransactOpts, _id, _sellerProfileId, _distributionId, _name, _desc, _quantity, _basePrice, _price)
}
