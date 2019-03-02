pragma solidity ^ 0.4.18;

contract EcommerceStore {
    enum ProductStatus {
        Open,
        Sold,
        Unsold
    }
    enum ProductCondition {
        New,
        Used
    }
    uint public productIndex;
    mapping(address => mapping(uint => Product))stores;
    mapping(uint => address)productIdInStore;

    struct Product {
        uint id; //产品id
        string name; //产品名字
        string category; //分类
        string imageLink; //图片hash
        string descLink; //图片描述信息的hash
        uint auctionStartTime; //开始竞标的时间
        uint auctionEndTime; // 竞标结束时间
        uint startPrice; // 拍卖价格
        address highestBidder; // 赢家的钱包地址
        uint highestBid; // 赢家竞标的价格
        uint secondHighestBid; // 第二高的这个人的地址
        uint totalBids; // 一共有多少人参与竞标
        ProductStatus status; //状态
        ProductCondition condition; // 新、旧
    }

    function EcommerceStore() public { //电子钱包
        productIndex = 0;
    }

    /*  添加产品到区块链*/
    function addProductToStore(string _name, string _category, string _imageLink, string _descLink, uint _auctionStartTime, uint _auctionEndTime, uint _startPrice, uint _productCondition) public {
        require(_auctionStartTime < _auctionEndTime);//断言 false中断
        productIndex += 1;
        Product memory product = Product(productIndex, _name, _category, _imageLink, _descLink, _auctionStartTime, _auctionEndTime, _startPrice, 0, 0, 0, 0, ProductStatus.Open, ProductCondition(_productCondition));
        stores[msg.sender][productIndex] = product;
        productIdInStore[productIndex] = msg.sender;
    }
    /* 通过产品ID读取产品信息 */
    function getProduct(uint _productId) view public returns (uint, string, string, string, string, uint, uint, uint, ProductStatus, ProductCondition) {
        Product memory product = stores[productIdInStore[_productId]][_productId];
        return (product.id, product.name, product.category, product.imageLink, product.descLink, product.auctionStartTime, product.auctionEndTime, product.startPrice, product.status, product.condition);
    }
}
