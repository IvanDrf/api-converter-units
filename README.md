# REST API for conversion  units of measurement

## Methods(CRUD)
- POST
  ```shell
  curl -X POST -H "Content-Type: application/json" -d '{"type":"_type_", "units":"_units_", "value":_value_, "converted_units":"_new_"}' http://localhost:8080 
  ```
  - Request
    ```json
    {
      "type": "length",
      "units": "m",
      "value": 100000,
      "converted_units": "km"
    }

    ```
 
    
  - Response
    ```json
    {
      "id": 10,
      "type": "length",
      "units": "m",
      "value": 100000,
      "converted_units": "km",
      "converted_value": 100
    }
    ```
  
  
- GET
  ```shell
  curl -X GET http://localhost:8080
  ```

  - Response
    ```json
    [
    {
        "id": 9,
        "type": "length",
        "units": "m",
        "value": 100,
        "converted_units": "cm",
        "converted_value": 10000
    },
    {
        "id": 8,
        "type": "length",
        "units": "m",
        "value": 100000,
        "converted_units": "cm",
        "converted_value": 10000000
    },
    {
        "id": 10,
        "type": "length",
        "units": "m",
        "value": 100000,
        "converted_units": "km",
        "converted_value": 100
    }
    ]
    ```

- PATCH
  ```shell
  curl -X PATCH -H "Content-Type: application/json" -d '{"type":"_type_", "units":"_units_", "value":_value_, "converted_units":"_new_"}' "http://localhost:8080/conversions/_id_"
  ```

  - Request
    ```json
    {
      "type": "length",
      "units": "m",
      "value": 100000,
      "converted_units": "cm"
    }
    ```
  - Response
    ```json
    {
      "id": 8,
      "type": "length",
      "units": "m",
      "value": 100000,
      "converted_units": "cm",
      "converted_value": 10000000
    }
    ```
    
- DELETE
  ```shell
  curl -X DELETE http://localhost:8080/conversions/_id_
  ```
***

## Supported units of measurement
- Length
  - millimeter
  - centimetre
  - meter
  - kilometer
  - inch
  - feet
  - yard
  - mile
    
- Temperature
  - celsius
  - fahrenheit
  - kelvin
    
- Time
  - second
  - minute
  - hour
  - day
    
- Weight
  - milligram
  - gram
  - kilogram
  - pound
  - ounce

***

## Technology used
- Golang
- Web-framework Echo 
- PostgreSQL

    
