// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 */
contract Supplier {

    struct Order_builder {
        string OrderNumber;
        string CommissionDate;      //委托时间
        string ProjectName;         //工程名称
        string ProjectLocation;     //工程地点
        string BuilderName;         //建造方名称
        string SupplierName;        //供应商名称
        string SuppervisorName;     //监理方名称
        string IPFSHASH;        //上传文件的IPFS地址
        string NameChargeEntrusted;  //委托单位负责人姓名
    }
    struct Order_production {
        string ProductionTime;      //生产时间
        string FactoryCertificate;  //预拌混凝土出厂合格证的IPFS地址
        string ConcreteProportioning;   //混凝土配合比通知单
        string RawMateriaQuality;       //原材料质量检测报告
        string NameChargeSupplier;      //供应单位负责人姓名
    }
    struct Order_logistics {
        string TransportsVolume;    //本车运输方量
        string CumulativeVolume;    //混凝土累计方量
        string ConcreteDeliveryNote; //混凝土送货单
        string DispatcherName;      //调度员姓名
    }
    struct Order_receiving {
        string ArrivalTime;         //抵达时间
        string MeasuredConcreteSlump;    //实测的混凝土坍落度
        string ConcreteEntryList;       //混凝土进场清单
        string SigningNameProject;      //施工单位签收人姓名
    }
    
    mapping(string => Order_builder) public Orders_builder;
    mapping(string => Order_production) public Orders_production;
    mapping(string => Order_logistics) public Orders_logistics;
    mapping(string => Order_receiving) public Orders_receiving;
    
    function submitOrder_builder(string memory OrderNumber, string memory ProjectLocation, string memory BuilderName,string memory SupplierName, string memory SuppervisorName,string memory IPFSHASH)public {
        
        Orders_builder[OrderNumber].SuppervisorName = SuppervisorName;
        Orders_builder[OrderNumber].ProjectLocation = ProjectLocation;
        Orders_builder[OrderNumber].BuilderName = BuilderName;
        Orders_builder[OrderNumber].SupplierName = SupplierName;
        Orders_builder[OrderNumber].IPFSHASH = IPFSHASH;
        
    }
    
    function submitOrder_production(string memory OrderNumber, string memory ProductionTime, string memory FactoryCertificate, string memory ConcreteProportioning, string memory RawMateriaQuality, string memory NameChargeSupplier) public{
        Orders_production[OrderNumber].ProductionTime = ProductionTime;
        Orders_production[OrderNumber].FactoryCertificate = FactoryCertificate;
        Orders_production[OrderNumber].ConcreteProportioning = ConcreteProportioning;
        Orders_production[OrderNumber].RawMateriaQuality = RawMateriaQuality;
        Orders_production[OrderNumber].NameChargeSupplier = NameChargeSupplier;
        
    }
    
    function submitOrder_logistics(string memory OrderNumber, string memory TransportsVolume, string memory CumulativeVolume, string memory ConcreteDeliveryNote, string memory DispatcherName)public {
        Orders_logistics[OrderNumber].TransportsVolume = TransportsVolume;
        Orders_logistics[OrderNumber].CumulativeVolume = CumulativeVolume;
        Orders_logistics[OrderNumber].ConcreteDeliveryNote = ConcreteDeliveryNote;
        Orders_logistics[OrderNumber].DispatcherName = DispatcherName;
    }
    
    function submitOrder_receiving(string memory OrderNumber, string memory ArrivalTime, string memory MeasuredConcreteSlump, string memory ConcreteEntryList, string memory SigningNameProject)public{
        Orders_receiving[OrderNumber].ArrivalTime = ArrivalTime;
        Orders_receiving[OrderNumber].MeasuredConcreteSlump = MeasuredConcreteSlump;
        Orders_receiving[OrderNumber].ConcreteEntryList = ConcreteEntryList;
        Orders_receiving[OrderNumber].SigningNameProject = SigningNameProject;
    }
    
    function queryOrder_builder(string memory OrderNumber) public view returns (Order_builder memory) {
        return Orders_builder[OrderNumber];
        
    }
    function queryOrder_production(string memory OrderNumber) public view returns (Order_production memory) {
        return Orders_production[OrderNumber];
        
    }
    function queryOrder_logistics(string memory OrderNumber) public view returns (Order_logistics memory) {
        return Orders_logistics[OrderNumber];
        
    }
    function queryOrder_receiving(string memory OrderNumber) public view returns (Order_receiving memory) {
        return Orders_receiving[OrderNumber];
        
    }

    /**
     * @dev Store value in variable
     * @param num value to store
     */
   

    /**
     * @dev Return value 
     * @return value of 'number'
     */
 
}