// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SupplyChain {

    // ==================== HARVEST ====================
    struct Harvest {
        uint256 id;
        uint256 farmerProfileId;
        uint256 cropId;
        uint256 regencyId;
        string name;
        string description;
        uint256 quantity;
        uint256 basePrice;
        uint256 createdAt;
    }
    mapping(uint256 => Harvest) public harvests;
    uint256[] public harvestIds;

    function createHarvest(
        uint256 _id,
        uint256 _farmerProfileId,
        uint256 _cropId,
        uint256 _regencyId,
        string memory _name,
        string memory _description,
        uint256 _quantity,
        uint256 _basePrice
    ) public {
        Harvest memory h = Harvest(_id, _farmerProfileId, _cropId, _regencyId, _name, _description, _quantity, _basePrice, block.timestamp);
        harvests[_id] = h;
        harvestIds.push(_id);
    }

    function getAllHarvestTuples() public view returns (
        uint256[] memory, uint256[] memory, uint256[] memory, uint256[] memory,
        string[] memory, string[] memory, uint256[] memory, uint256[] memory, uint256[] memory
    ) {
        uint len = harvestIds.length;
        uint256[] memory ids = new uint256[](len);
        uint256[] memory farmerIds = new uint256[](len);
        uint256[] memory cropIds = new uint256[](len);
        uint256[] memory regencyIds = new uint256[](len);
        string[] memory names = new string[](len);
        string[] memory descs = new string[](len);
        uint256[] memory quantities = new uint256[](len);
        uint256[] memory basePrices = new uint256[](len);
        uint256[] memory createdAts = new uint256[](len);

        for (uint i = 0; i < len; i++) {
            Harvest memory h = harvests[harvestIds[i]];
            ids[i] = h.id;
            farmerIds[i] = h.farmerProfileId;
            cropIds[i] = h.cropId;
            regencyIds[i] = h.regencyId;
            names[i] = h.name;
            descs[i] = h.description;
            quantities[i] = h.quantity;
            basePrices[i] = h.basePrice;
            createdAts[i] = h.createdAt;
        }
        return (ids, farmerIds, cropIds, regencyIds, names, descs, quantities, basePrices, createdAts);
    }

    // ==================== HARVEST COLLECTOR ====================
    struct HarvestCollector {
        uint256 id;
        uint256 collectorProfileId;
        uint256 harvestId;
        string name;
        string desc;
        uint256 quantity;
        uint256 price;
        uint256 basePrice;
        uint256 createdAt;
    }
    mapping(uint256 => HarvestCollector) public harvestCollectors;
    uint256[] public harvestCollectorIds;

    function createHarvestCollector(
        uint256 _id,
        uint256 _collectorProfileId,
        uint256 _harvestId,
        string memory _name,
        string memory _desc,
        uint256 _quantity,
        uint256 _price,
        uint256 _basePrice
    ) public {
        HarvestCollector memory hc = HarvestCollector(_id, _collectorProfileId, _harvestId, _name, _desc, _quantity, _price, _basePrice, block.timestamp);
        harvestCollectors[_id] = hc;
        harvestCollectorIds.push(_id);
    }

    function getAllHarvestCollectorTuples() public view returns (
        uint256[] memory, uint256[] memory, uint256[] memory,
        string[] memory, string[] memory,
        uint256[] memory, uint256[] memory, uint256[] memory
    ) {
        uint len = harvestCollectorIds.length;
        uint256[] memory ids = new uint256[](len);
        uint256[] memory collectorIDs = new uint256[](len);
        uint256[] memory harvestIDs = new uint256[](len);
        string[] memory names = new string[](len);
        string[] memory descs = new string[](len);
        uint256[] memory quantities = new uint256[](len);
        uint256[] memory prices = new uint256[](len);
        uint256[] memory basePrices = new uint256[](len);

        for (uint i = 0; i < len; i++) {
            HarvestCollector memory hc = harvestCollectors[harvestCollectorIds[i]];
            ids[i] = hc.id;
            collectorIDs[i] = hc.collectorProfileId;
            harvestIDs[i] = hc.harvestId;
            names[i] = hc.name;
            descs[i] = hc.desc;
            quantities[i] = hc.quantity;
            prices[i] = hc.price;
            basePrices[i] = hc.basePrice;
        }
        return (ids, collectorIDs, harvestIDs, names, descs, quantities, prices, basePrices);
    }

    // ==================== HARVEST PROCESSOR ====================
    struct HarvestProcessor {
        uint256 id;
        uint256 processorProfileId;
        uint256 harvestCollectorId;
        uint256 harvestId;
        string name;
        string desc;
        uint256 quantity;
        uint256 basePrice;
        uint256 price;
        uint256 createdAt;
    }
    mapping(uint256 => HarvestProcessor) public harvestProcessors;
    uint256[] public harvestProcessorIds;

    function createHarvestProcessor(
        uint256 _id,
        uint256 _processorProfileId,
        uint256 _harvestCollectorId,
        uint256 _harvestId,
        string memory _name,
        string memory _desc,
        uint256 _quantity,
        uint256 _basePrice,
        uint256 _price
    ) public {
        HarvestProcessor memory hp = HarvestProcessor(_id, _processorProfileId, _harvestCollectorId, _harvestId, _name, _desc, _quantity, _basePrice, _price, block.timestamp);
        harvestProcessors[_id] = hp;
        harvestProcessorIds.push(_id);
    }

    function getAllHarvestProcessorTuples() public view returns (
        uint256[] memory, uint256[] memory, uint256[] memory, uint256[] memory,
        string[] memory, string[] memory,
        uint256[] memory, uint256[] memory, uint256[] memory
    ) {
        uint len = harvestProcessorIds.length;
        uint256[] memory ids = new uint256[](len);
        uint256[] memory processorIDs = new uint256[](len);
        uint256[] memory collectorIDs = new uint256[](len);
        uint256[] memory harvestIDs = new uint256[](len);
        string[] memory names = new string[](len);
        string[] memory descs = new string[](len);
        uint256[] memory quantities = new uint256[](len);
        uint256[] memory basePrices = new uint256[](len);
        uint256[] memory prices = new uint256[](len);

        for (uint i = 0; i < len; i++) {
            HarvestProcessor memory hp = harvestProcessors[harvestProcessorIds[i]];
            ids[i] = hp.id;
            processorIDs[i] = hp.processorProfileId;
            collectorIDs[i] = hp.harvestCollectorId;
            harvestIDs[i] = hp.harvestId;
            names[i] = hp.name;
            descs[i] = hp.desc;
            quantities[i] = hp.quantity;
            basePrices[i] = hp.basePrice;
            prices[i] = hp.price;
        }
        return (ids, processorIDs, collectorIDs, harvestIDs, names, descs, quantities, basePrices, prices);
    }

    // ==================== DISTRIBUTION ====================
    struct Distribution {
        uint256 id;
        uint256 distributorProfileId;
        uint256 destinationId;
        uint256 harvestId;
        uint256 harvestCollectorId;
        uint256 harvestProcessorId;
        string name;
        string desc;
        uint256 quantity;
        uint256 basePrice;
        uint256 price;
        string transportation;
        uint256 createdAt;
    }
    mapping(uint256 => Distribution) public distributions;
    uint256[] public distributionIds;

    function createDistribution(
        uint256 _id,
        uint256 _distributorProfileId,
        uint256 _destinationId,
        uint256 _harvestId,
        uint256 _harvestCollectorId,
        uint256 _harvestProcessorId,
        string memory _name,
        string memory _desc,
        uint256 _quantity,
        uint256 _basePrice,
        uint256 _price,
        string memory _transportation
    ) public {
        Distribution memory d = Distribution(_id,_distributorProfileId,_destinationId,_harvestId,_harvestCollectorId,_harvestProcessorId,_name,_desc,_quantity,_basePrice,_price,_transportation,block.timestamp);
        distributions[_id] = d;
        distributionIds.push(_id);
    }

    function getAllDistributionTuples() public view returns (
        uint256[] memory, uint256[] memory, uint256[] memory, uint256[] memory,
        uint256[] memory, uint256[] memory,
        string[] memory, string[] memory,
        uint256[] memory, uint256[] memory, uint256[] memory,
        string[] memory
    ) {
        uint len = distributionIds.length;
        uint256[] memory ids = new uint256[](len);
        uint256[] memory distributorIDs = new uint256[](len);
        uint256[] memory destinationIDs = new uint256[](len);
        uint256[] memory harvestIDs = new uint256[](len);
        uint256[] memory collectorIDs = new uint256[](len);
        uint256[] memory processorIDs = new uint256[](len);
        string[] memory names = new string[](len);
        string[] memory descs = new string[](len);
        uint256[] memory quantities = new uint256[](len);
        uint256[] memory basePrices = new uint256[](len);
        uint256[] memory prices = new uint256[](len);
        string[] memory transportations = new string[](len);

        for (uint i = 0; i < len; i++) {
            Distribution memory d = distributions[distributionIds[i]];
            ids[i] = d.id;
            distributorIDs[i] = d.distributorProfileId;
            destinationIDs[i] = d.destinationId;
            harvestIDs[i] = d.harvestId;
            collectorIDs[i] = d.harvestCollectorId;
            processorIDs[i] = d.harvestProcessorId;
            names[i] = d.name;
            descs[i] = d.desc;
            quantities[i] = d.quantity;
            basePrices[i] = d.basePrice;
            prices[i] = d.price;
            transportations[i] = d.transportation;
        }
        return (ids, distributorIDs, destinationIDs, harvestIDs, collectorIDs, processorIDs, names, descs, quantities, basePrices, prices, transportations);
    }

    // ==================== SELLER BOX ====================
    struct SellerBox {
        uint256 id;
        uint256 sellerProfileId;
        uint256 distributionId;
        string name;
        string desc;
        uint256 quantity;
        uint256 basePrice;
        uint256 price;
        uint256 createdAt;
    }
    mapping(uint256 => SellerBox) public sellerBoxes;
    uint256[] public sellerBoxIds;

    function createSellerBox(
        uint256 _id,
        uint256 _sellerProfileId,
        uint256 _distributionId,
        string memory _name,
        string memory _desc,
        uint256 _quantity,
        uint256 _basePrice,
        uint256 _price
    ) public {
        SellerBox memory s = SellerBox(_id,_sellerProfileId,_distributionId,_name,_desc,_quantity,_basePrice,_price,block.timestamp);
        sellerBoxes[_id] = s;
        sellerBoxIds.push(_id);
    }

    function getAllSellerBoxTuples() public view returns (
        uint256[] memory, uint256[] memory, uint256[] memory,
        string[] memory, string[] memory,
        uint256[] memory, uint256[] memory, uint256[] memory
    ) {
        uint len = sellerBoxIds.length;
        uint256[] memory ids = new uint256[](len);
        uint256[] memory sellerIDs = new uint256[](len);
        uint256[] memory distributionIDs = new uint256[](len);
        string[] memory names = new string[](len);
        string[] memory descs = new string[](len);
        uint256[] memory quantities = new uint256[](len);
        uint256[] memory basePrices = new uint256[](len);
        uint256[] memory prices = new uint256[](len);

        for (uint i = 0; i < len; i++) {
            SellerBox memory s = sellerBoxes[sellerBoxIds[i]];
            ids[i] = s.id;
            sellerIDs[i] = s.sellerProfileId;
            distributionIDs[i] = s.distributionId;
            names[i] = s.name;
            descs[i] = s.desc;
            quantities[i] = s.quantity;
            basePrices[i] = s.basePrice;
            prices[i] = s.price;
        }
        return (ids, sellerIDs, distributionIDs, names, descs, quantities, basePrices, prices);
    }

}
