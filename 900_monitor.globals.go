package main

/* ----------------------------------- */
type serviceStruct struct {
	LogicalName         string
	URL                 string
	LastCheckResultText string
}
type listOfServices []serviceStruct

type ServiceStatStruc struct {
	CreateRangeTabName string
	ResultCreateRange  string
	SchemaCreateRange  string
	ShowStateTabName   string
	ResultShowRange    string
	SchemaShowRange    string
	GetNextTabName     string
	ResultNextTabName  string
	SchemaNextRange    string
}

type ServiceStatTable struct {
	ServicesStatInfo ServiceStatStruc
}

var viewpagedata ServiceStatTable

/* ----------------------------------- */

/* ----------------------------------- */
type subscriberStruct struct {
	SubscrName     string
	SubscrEmail    string
	SubscrDateFrom string
	//SubscrDateTimeFrom string
	SubscrDateTo string
	//SubscrDateTimeTo   string
	Index      int
	delFlagSet bool
}
type listOfSubscribers []subscriberStruct

// SubscriberInfoTable ...
type SubscriberInfoTable struct {
	SubscribersInfo listOfSubscribers
}

var subscrviewpagedata SubscriberInfoTable

/* ----------------------------------- */

/* ============Konstanten ============================================= */

/* ============Types ============================================= */

type nameValStruct struct {
	name string
	val  string
}
