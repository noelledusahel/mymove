package models

//
// Tab 1b: Service Areas
//
type StageDomesticServiceArea struct {
	BasePointCity     string `db:"base_point_city" csv:"base_point_city"`
	State             string `db:"state" csv:"state"`
	ServiceAreaNumber string `db:"service_area_number" csv:"service_area_number"`
	Zip3s             string `db:"zip3s" csv:"zip3s"`
}

type StageInternationalServiceArea struct {
	RateArea   string `db:"rate_area" csv:"rate_area"`
	RateAreaID string `db:"rate_area_id" csv:"rate_area_id"`
}

//
// Tab 2a: Domestic Linehaul Prices
//
type StageDomesticLinehaulPrice struct {
	ServiceAreaNumber string `db:"service_area_number" csv:"service_area_number"`
	OriginServiceArea string `db:"origin_service_area" csv:"origin_service_area"`
	ServicesSchedule  string `db:"services_schedule" csv:"services_schedule"`
	Season            string `db:"season" csv:"season"`
	WeightLower       string `db:"weight_lower" csv:"weight_lower"`
	WeightUpper       string `db:"weight_upper" csv:"weight_upper"`
	MilesLower        string `db:"miles_lower" csv:"miles_lower"`
	MilesUpper        string `db:"miles_upper" csv:"miles_upper"`
	EscalationNumber  string `db:"escalation_number" csv:"escalation_number"`
	Rate              string `db:"rate" csv:"rate"`
}

//
// Tab 2b: Domestic Service Area Prices
//
type StageDomesticServiceAreaPrice struct {
	ServiceAreaNumber                     string `db:"service_area_number" csv:"service_area_number"`
	ServiceAreaName                       string `db:"service_area_name" csv:"service_area_name"`
	ServicesSchedule                      string `db:"services_schedule" csv:"services_schedule"`
	SITPickupDeliverySchedule             string `db:"sit_pickup_delivery_schedule" csv:"sit_pickup_delivery_schedule"`
	Season                                string `db:"season" csv:"season"`
	ShorthaulPrice                        string `db:"shorthaul_price" csv:"shorthaul_price"`
	OriginDestinationPrice                string `db:"origin_destination_price" csv:"origin_destination_price"`
	OriginDestinationSITFirstDayWarehouse string `db:"origin_destination_sit_first_day_warehouse" csv:"origin_destination_sit_first_day_warehouse"`
	OriginDestinationSITAddlDays          string `db:"origin_destination_sit_addl_days" csv:"origin_destination_sit_addl_days"`
}

//
// Tab 2c: Other Domestic Prices
//
type StageDomesticOtherPackPrice struct {
	ServicesSchedule   string `db:"services_schedule" csv:"services_schedule"`
	ServiceProvided    string `db:"service_provided" csv:"service_provided"`
	NonPeakPricePerCwt string `db:"non_peak_price_per_cwt" csv:"non_peak_price_per_cwt"`
	PeakPricePerCwt    string `db:"peak_price_per_cwt" csv:"peak_price_per_cwt"`
}

type StageDomesticOtherSitPrice struct {
	SITPickupDeliverySchedule string `db:"sit_pickup_delivery_schedule" csv:"sit_pickup_delivery_schedule"`
	ServiceProvided           string `db:"service_provided" csv:"service_provided"`
	NonPeakPricePerCwt        string `db:"non_peak_price_per_cwt" csv:"non_peak_price_per_cwt"`
	PeakPricePerCwt           string `db:"peak_price_per_cwt" csv:"peak_price_per_cwt"`
}

//
// Tab 3a: OCONUS to OCONUS Prices
//
type StageOconusToOconusPrice struct {
	OriginIntlPriceAreaID      string `db:"origin_intl_price_area_id" csv:"origin_intl_price_area_id"`
	OriginIntlPriceArea        string `db:"origin_intl_price_area" csv:"origin_intl_price_area"`
	DestinationIntlPriceAreaID string `db:"destination_intl_price_area_id" csv:"destination_intl_price_area_id"`
	DestinationIntlPriceArea   string `db:"destination_intl_price_area" csv:"destination_intl_price_area"`
	Season                     string `db:"season" csv:"season"`
	HHGShippingLinehaulPrice   string `db:"hhg_shipping_linehaul_price" csv:"hhg_shipping_linehaul_price"`
	UBPrice                    string `db:"ub_price" csv:"ub_price"`
}

//
// Tab 3b: CONUS to OCONUS Prices
//
type StageConusToOconusPrice struct {
	OriginDomesticPriceAreaCode string `db:"origin_domestic_price_area_code" csv:"origin_domestic_price_area_code"`
	OriginDomesticPriceArea     string `db:"origin_domestic_price_area" csv:"origin_domestic_price_area"`
	DestinationIntlPriceAreaID  string `db:"destination_intl_price_area_id" csv:"destination_intl_price_area_id"`
	DestinationIntlPriceArea    string `db:"destination_intl_price_area" csv:"destination_intl_price_area"`
	Season                      string `db:"season" csv:"season"`
	HHGShippingLinehaulPrice    string `db:"hhg_shipping_linehaul_price" csv:"hhg_shipping_linehaul_price"`
	UBPrice                     string `db:"ub_price" csv:"ub_price"`
}

//
// Tab 3c: OCONUS to CONUS Prices
//
type StageOconusToConusPrice struct {
	OriginIntlPriceAreaID            string `db:"origin_intl_price_area_id" csv:"origin_intl_price_area_id"`
	OriginIntlPriceArea              string `db:"origin_intl_price_area" csv:"origin_intl_price_area"`
	DestinationDomesticPriceAreaCode string `db:"destination_domestic_price_area_area" csv:"destination_domestic_price_area_area"`
	DestinationDomesticPriceArea     string `db:"destination_domestic_price_area" csv:"destination_domestic_price_area"`
	Season                           string `db:"season" csv:"season"`
	HHGShippingLinehaulPrice         string `db:"hhg_shipping_linehaul_price" csv:"hhg_shipping_linehaul_price"`
	UBPrice                          string `db:"ub_price" csv:"ub_price"`
}

//
// Tab 3d: Other International Prices
//
type StageOtherIntlPrice struct {
	RateAreaCode                          string `db:"rate_area_code" csv:"rate_area_code"`
	RateAreaName                          string `db:"rate_area_name" csv:"rate_area_name"`
	HHGOriginPackPrice                    string `db:"hhg_origin_pack_price" csv:"hhg_origin_pack_price"`
	HHGDestinationUnPackPrice             string `db:"hhg_destination_unpack_price" csv:"hhg_destination_unpack_price"`
	UBOriginPackPrice                     string `db:"ub_origin_pack_price" csv:"ub_origin_pack_price"`
	UBDestinationUnPackPrice              string `db:"ub_destination_unpack_price" csv:"ub_destination_unpack_price"`
	OriginDestinationSITFirstDayWarehouse string `db:"origin_destination_sit_first_day_warehouse" csv:"origin_destination_sit_first_day_warehouse"`
	OriginDestinationSITAddlDays          string `db:"origin_destination_sit_addl_days" csv:"origin_destination_sit_addl_days"`
	SITLte50Miles                         string `db:"sit_lte_50_miles" csv:"sit_lte_50_miles"`
	SITGt50Miles                          string `db:"sit_gt_50_miles" csv:"sit_gt_50_miles"`
	Season                                string `db:"season" csv:"season"`
}

//
// Tab 3e: Non-Standard Location Prices
//
type StageNonStandardLocnPrice struct {
	OriginID        string `db:"origin_id" csv:"origin_id"`
	OriginArea      string `db:"origin_area" csv:"origin_area"`
	DestinationID   string `db:"destination_id" csv:"destination_id"`
	DestinationArea string `db:"destination_area" csv:"destination_area"`
	MoveType        string `db:"move_type" csv:"move_type"`
	Season          string `db:"season" csv:"season"`
	HHGPrice        string `db:"hhg_price" csv:"hhg_price"`
	UBPrice         string `db:"ub_price" csv:"ub_price"`
}

//
// Tab 4a: Management, Counseling, and Transition Prices
//
type StageShipmentManagementServicesPrice struct {
	ContractYear      string `db:"contract_year" csv:"contract_year"`
	PricePerTaskOrder string `db:"price_per_task_order" csv:"price_per_task_order"`
}

type StageCounselingServicesPrice struct {
	ContractYear      string `db:"contract_year" csv:"contract_year"`
	PricePerTaskOrder string `db:"price_per_task_order" csv:"price_per_task_order"`
}

type StageTransitionPrice struct {
	ContractYear      string `db:"contract_year" csv:"contract_year"`
	PricePerTaskOrder string `db:"price_total_cost" csv:"price_total_cost"`
}

//
// Tab 5a: Accessorial and Additional Prices
//
type StageDomesticMoveAccessorialPrices struct {
	ServicesSchedule string `db:"services_schedule" csv:"services_schedule"`
	ServiceProvided  string `db:"service_provided" csv:"service_provided"`
	PricePerUnit     string `db:"price_per_unit" csv:"price_per_unit"`
}

type StageInternationalMoveAccessorialPrices struct {
	Market          string `db:"market" csv:"market"`
	ServiceProvided string `db:"service_provided" csv:"service_provided"`
	PricePerUnit    string `db:"price_per_unit" csv:"price_per_unit"`
}

type StageDomesticInternationalAdditionalPrices struct {
	Market       string `db:"market" csv:"market"`
	ShipmentType string `db:"shipment_type" csv:"shipment_type"`
	Factor       string `db:"factor" csv:"factor"`
}

//
// Tab 5b: Price Escalation Discount
//
type StagePriceEscalationDiscount struct {
	ContractYear          string `db:"contract_year" csv:"contract_year"`
	ForecastingAdjustment string `db:"forecasting_adjustment" csv:"forecasting_adjustment"`
	Discount              string `db:"discount" csv:"discount"`
	PriceEscalation       string `db:"price_escalation" csv:"price_escalation"`
}
