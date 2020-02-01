# How to run the GHC Pricing Import and Verify Data

To support loading GHC Pricing data you can use the `bin/ghc-pricing-parser` to do so.

## Running the parser

You will need to build the parser first.

```
make bin/ghc-pricing-parser
```

Once built you can run the command. This command will take some time, the sample data xlsx used below takes 5-6 minutes to complete.

```
ghc-pricing-parser --filename pkg/parser/pricing/fixtures/pricing_template_2019-09-19_fake-data.xlsx --contract-code=UNIQUECODE --contract-name="Unique Name"
```

Once complete move on to the next section to verify the import

## Verifying the data

The script will output the summary of the staging tables and the rate engine tables that were used for the import. The summary will include the total number of rows inserted as well as the first and last rows. For verifying a correct import it is recommended to select a random couple of rows from each rate engine table to verify correct parsing.

make sure all table counts match spreadsheet
verify first and last row matches
select a random row to verify

## Useful Command Options

You can run the parser with the `--help` flag to see all possible options. Below is a selection of the most commonly needed flags

* `--filename string` **Required**
  * Filename including path of the XLSX to parse for Rate Engine GHC import
* `--contract-code string` **Required**
  * Contract code to use for this import
* `--contract-name string`
  * Contract name to use for this import
* `--display`
  * Display output of parsed info
* `--save-csv`
  * Save output of xlsx sheets to CSV file
* `--verify`
  * Default is true, if false skip sheet format verification (default true) this will verify that the xlsx looks as we expect it too, this does not validate data.
* `--re-import`
  * Run GHC Rate Engine Import (default true)
* `--use-temp-tables`
  * Default is true, if false stage tables are NOT temp tables (default true)
* `--drop`
  * Default is false, if true stage tables will be dropped if they exist this is useful in conjunction with turning `--use-temp-tables` off
* `--db-env string`
  * database environment: container, test, development (default "development")
* `--db-name string`
  * Database Name (default "dev_db")
* `--db-host string`
  * Database Hostname (default "localhost")
* `--db-port int`
  * Database Port (default 5432)
* `--db-user string`
  * Database Username (default "postgres")
* `--db-password string`
  * Database Password
