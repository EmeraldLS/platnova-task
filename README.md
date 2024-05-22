
# PDF Generator

A simple and efficient Golang application designed to parse JSON data and generate PDF documents. This tool is ideal for converting structured data stored in JSON format into professionally formatted PDF reports. The project leverages the power of the UniDoc library to handle PDF creation and manipulation, ensuring high-quality output.

## JSON Schema for PDF Statement Generation

The following JSON structure outlines the schema for generating PDF statements:

```json
{
    "statement": {
        "title": "string",
        "generated_date": "string",
        "bank_name": "string",
        "customer_name": "string",
        "customer_address": {
            "line1": "string",
            "line2": "string",
            "city": "string",
            "county": "string",
            "postcode": "string"
        },
        "balance_summary":[
            {
            "product": "string",
            "opening_balance": number,
            "money_out": number,
            "money_in": number,
            "closing_balance": number
            }
        ],
        "account_transactions": [
            {
                "date": "string",
                "description": "string",
                "money_out": number,
                "money_in": number,
                "balance": number
            }
        ],
        "iban_details": [
            {
                "iban": "string",
                "bic": "string",
                "note": "string"
            }
        ]
    }
}
```

## Usage

- #### Using Online Version
    - Visit the [PDF Generator](https://pdf-generator-uni.herokuapp.com/) web application.
    - Upload a JSON file containing the data to be converted into a PDF document.
    - Click the "Generate PDF" button to create the PDF statement.
    - Download the generated PDF file to your local machine.
- #### Using Local Version
    1. Local version is made of two aspects, first is, you can use the api which support upload of `.json` file and downloads the generated pdf file to your computer instantly. It requires some data in the `.env` file, there's a `.env.example` file you can follow through:
        - Clone the repository to your local machine.
        - Navigate to the project directory.
        - Run the following command to start the server:
            ```bash
            go run main.go -use_api
            ```
        - Open your browser and visit `http://localhost:8080`.
        - Upload a JSON file containing the data to be converted into a PDF document using the defined schema above.

    2. Second is, you can run the application on your local machine and generate the pdf file.

    To use the local version of the application, follow the steps below:

    - You need an API key to use the local version of the application. You can obtain an API key by signing up on the [UniDoc](https://unidoc.io/) website. 
    - Clone the repository to your local machine.
    - Navigate to the project directory.
    - Run the following command to generate a PDF statement:
        ```bash
        go run main.go -api_key=<your_api_key> -json_file=<path_to_json_file> 
        ```
    - Replace `<path_to_json_file>` with the path to the JSON file containing the data.
    - Replace `<path_to_pdf_file>` with the desired output path for the generated PDF document.

## Testing
- Run the following command to execute the unit tests:
    ```bash
    make test
    ```
    or
    
    ```bash
   go test -v ./...
    ```
## Project Structure
    main.go - The entry point of the application.
    static/ - Contains the static files for the web application.
    internal/
        generator/ - Contains the API client for interacting with the UniDoc API and the PDF generator.
        rest/ - Contains the REST API server implementation.
        types/ - Contains the data structures used in the application.
        utils/ - Contains utility functions used in the application.