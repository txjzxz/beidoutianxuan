package controller

import (
)

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName string
	Password  string
	Org       string
	IsAdmin   bool
}

type Status struct {
	Order   bool
	Produce bool
	Supply  bool
	Receive bool
	Pouring bool
}

type Check struct {
	IsQualified bool
	Reason string
}

type step1Data struct {
	Date                  string `json:"date"`
	ProjectName           string `json:"project_name"`
	ProjectLocation       string `json:"project_location"`
	ConstructionOrg       string `json:"construction_org"`
	SupplyOrg             string `json:"supply_org"`
	SupervisorOrg         string `json:"supervisor_org"`
	ScheduledTime         string `json:"scheduled_time"`
	PouringPosition       string `json:"pouring_position"`
	GradesOfConcrete      string `json:"grades_of_concrete"`
	ConcreteVolume        string `json:"concrete_volume"`
	ConcretePumpingMethod string `json:"concrete_pumping_method"`
	ConcreteSlump         string `json:"concrete_slump"`
	ConcretePerformance   string `json:"concrete_performance"`
	Hash                  string `json:"hash"`
	Principal             string `json:"principal"`
}

type step1Order struct {
	Id        string `json:"id"`
	OrderData *step1Data
}

type Step2Data struct {
	Date                  string `json:"date"`
	ScheduledTime         string `json:"scheduled_time"`
	ProjectName           string `json:"project_name"`
	ProjectLocation       string `json:"project_location"`
	ConstructionOrg       string `json:"construction_org"`
	PouringPosition       string `json:"pouring_position"`
	GradesOfConcrete      string `json:"grades_of_concrete"`
	ConcreteVolume        string `json:"concrete_volume"`
	ConcretePumpingMethod string `json:"concrete_pumping_method"`
	ConcreteSlump         string `json:"concrete_slump"`
	Hash1                 string `json:"hash_1"`
	Hash2                 string `json:"hash_2"`
	Hash3                 string `json:"hash_3"`
	Principal             string `json:"principal"`
}

type step2 struct {
	Id        string `json:"id"`
	OrderData Step2Data
}

type Step3Data struct {
	ProductionTime        string `json:"production_time"`
	ScheduledTime         string `json:"scheduled_time"`
	ProjectName           string `json:"project_name"`
	ProjectLocation       string `json:"project_location"`
	ConstructionOrg       string `json:"construction_org"`
	PouringPosition       string `json:"pouring_position"`
	GradesOfConcrete      string `json:"grades_of_concrete"`
	TheCarSupplyVolume    string `json:"the_car_supply_volume"`
	SumVolume             string `json:"sum_volume"`
	PlanVolume            string `json:"plan_volume"`
	ConcretePumpingMethod string `json:"concrete_pumping_method"`
	ConcreteSlump         string `json:"concrete_slump"`
	SupplyDistance        string `json:"supply_distance"`
	Hash                  string `json:"hash"`
	Yardman               string `json:"yardman"`
}

type step3 struct {
	Id        string `json:"id"`
	OrderData Step3Data
}

type step4Data struct {
	ProductionTime        string `json:"production_time"`
	ArriveTime            string `json:"arrive_time"`
	ProjectName           string `json:"project_name"`
	ProjectLocation       string `json:"project_location"`
	ConstructionOrg       string `json:"construction_org"`
	SupplyOrg             string `json:"supply_org"`
	SupervisorOrg         string `json:"supervisor_org"`
	PouringPosition       string `json:"pouring_position"`
	GradesOfConcrete      string `json:"grades_of_concrete"`
	ConcreteVolume        string `json:"concrete_volume"`
	ConcretePumpingMethod string `json:"concrete_pumping_method"`
	ConcreteActualSlump   string `json:"concrete_actual_slump"`
	Hash                  string `json:"hash"`
	IsQualified           string `json:"is_qualified"`
	ConstructionPrincipal string `json:"construction_principal"`
}

type step4 struct {
	ArriveId  string `json:"arrive_id"`
	OrderData step4Data
}

type step5Data struct {
	ConcreteDischargeTime string `json:"concrete_discharge_time"`
	Weather               string `json:"weather"`
	Temperature           string `json:"temperature"`
	ProjectName           string `json:"project_name"`
	ProjectLocation       string `json:"project_location"`
	ConstructionOrg       string `json:"construction_org"`
	SupplyOrg             string `json:"supply_org"`
	SupervisorOrg         string `json:"supervisor_org"`
	PouringPosition       string `json:"pouring_position"`
	GradesOfConcrete      string `json:"grades_of_concrete"`
	ConcreteVolume        string `json:"concrete_volume"`
	ConcretePumpingMethod string `json:"concrete_pumping_method"`
	Hash1                 string `json:"hash_1"`
	Hash2                 string `json:"hash_2"`
	ConstructionPrincipal string `json:"construction_principal"`
}

type step5 struct {
	Id        string `json:"id"`
	OrderData step5Data
}

type block struct {
	BCI      *common.BlockchainInfo
	Endorser string
	Status   int32
}

var users []User

func init() {

	org1 := User{LoginName: "Org1@gmail.com", Password: "123456", Org: "1", IsAdmin: true}
	org2 := User{LoginName: "Org2@gmail.com", Password: "123456", Org: "2", IsAdmin: true}
	org3 := User{LoginName: "Org3@gmail.com", Password: "123456", Org: "3", IsAdmin: true}
	org4 := User{LoginName: "Org4@gmail.com", Password: "123456", Org: "4", IsAdmin: true}
	org5 := User{LoginName: "Org5@gmail.com", Password: "123456", Org: "5", IsAdmin: true}
	org6 := User{LoginName: "Org6@gmail.com", Password: "123456", Org: "6", IsAdmin: true}

	users = append(users, org1)
	users = append(users, org2)
	users = append(users, org3)
	users = append(users, org4)
	users = append(users, org5)
	users = append(users, org6)

}

