package ghcimport

import (
	"fmt"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
)

type GHCRateEngineImporter struct {
	Logger       Logger
	ContractCode string
	ContractName string
	// TODO: add reference maps here as needed for dependencies between tables
	contractID                   uuid.UUID
	serviceAreaToIDMap           map[string]uuid.UUID
	domesticRateAreaToIDMap      map[string]uuid.UUID
	internationalRateAreaToIDMap map[string]uuid.UUID
	serviceToIDMap               map[string]uuid.UUID
	contractYearToIDMap          map[string]uuid.UUID
}

func (gre *GHCRateEngineImporter) runImports(dbTx *pop.Connection) error {
	// Reference tables
	err := gre.importREContract(dbTx) // Also populates gre.contractID
	if err != nil {
		return fmt.Errorf("failed to import re_contract: %w", err)
	}

	err = gre.importREContractYears(dbTx) // Populates gre.contractYearToIDMap
	if err != nil {
		return fmt.Errorf("failed to import re_contract_years: %w", err)
	}

	err = gre.importREDomesticServiceArea(dbTx) // Also populates gre.serviceAreaToIDMap
	if err != nil {
		return fmt.Errorf("failed to import re_domestic_service_area: %w", err)
	}

	err = gre.importRERateArea(dbTx) // Also populates gre.domesticRateAreaToIDMap and gre.internationalRateAreaToIDMap
	if err != nil {
		return fmt.Errorf("failed to import re_rate_area: %w", err)
	}

	err = gre.loadServiceMap(dbTx) // Populates gre.serviceToIDMap
	if err != nil {
		return fmt.Errorf("failed to load service map: %w", err)
	}

	// Non-reference tables
	err = gre.importREDomesticLinehaulPrices(dbTx)
	if err != nil {
		return fmt.Errorf("failed to import re_domestic_linehaul_prices: %w", err)
	}

	err = gre.importREDomesticServiceAreaPrices(dbTx)
	if err != nil {
		return fmt.Errorf("failed to import re_domestic_service_area_prices: %w", err)
	}

	err = gre.importREDomesticOtherPrices(dbTx)
	if err != nil {
		return fmt.Errorf("failed to import re_domestic_other_prices: %w", err)
	}

	err = gre.importREInternationalPrices(dbTx)
	if err != nil {
		return fmt.Errorf("failed to import re_intl_prices: %w", err)
	}

	err = gre.importREInternationalOtherPrices(dbTx)
	if err != nil {
		return fmt.Errorf("failed to import re_intl_other_prices: %w", err)
	}

	err = gre.importRETaskOrderFees(dbTx)
	if err != nil {
		return fmt.Errorf("failed to import re_task_order_fees: %w", err)
	}

	err = gre.importREDomesticAccessorialPrices(dbTx)
	if err != nil {
		return fmt.Errorf("failed to import re_domestic_accessorial_prices: %w", err)
	}

	err = gre.importREIntlAccessorialPrices(dbTx)
	if err != nil {
		return fmt.Errorf("failed to import re_intl_accessorial_prices: %w", err)
	}

	err = gre.importREShipmentTypePrices(dbTx)
	if err != nil {
		return fmt.Errorf("failed to import re_shipment_type_prices: %w", err)
	}
	return nil
}

func (gre *GHCRateEngineImporter) Import(db *pop.Connection) error {
	err := db.Transaction(func(connection *pop.Connection) error {
		dbTxError := gre.runImports(connection)
		return dbTxError
	})
	if err != nil {
		return fmt.Errorf("transaction failed during GHC Rate Engine Import(): %w", err)
	}
	return nil
}
