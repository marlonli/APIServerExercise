## Dependencies
[gorilla/mux](github.com/gorilla/mux)
A powerful router for golang

Also, [Postman](https://www.getpostman.com/) is a great tool for API development. I used Postman for mock http request. 
## API Contract

### Persist Metadata
It will check whether the metadata format is valid (all fields are required) or not, and if yes, persist the metadata.
- Path: /v1/persist
- Method: POST
- Request Body: YAML
```yaml
title: Valid App 1
version: 0.0.1
maintainers:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
- name: secondmaintainer app1
  email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
 ### Interesting Title
 Some application content, and description
```
- Response: 
  - Status code _200_: The request body (metadata) is valid and has successfully persisted
    - Response Body: "Persist completed!"
  - Status code _400_: The request body (metadata) is invalid, probably missing some field
    - Response Body: "Invalid request!"

### Search Metadata
It will search for the metadata that matches the query string in any field and retrieve a list of metadata that matches.
- Path: /v1/search
- Method: POST
- Request Body: A query string
- Response:
  - Status code _200_: The query is valid 
    - Response Body: A empty body or a list of metadata in YAML
    ```yaml
    - title: Valid App 1
      version: 0.0.1
      maintainers:
      - name: firstmaintainer app1
        email: firstmaintainer@hotmail.com
      - name: secondmaintainer app1
        email: secondmaintainer@gmail.com
      company: Random Inc.
      website: https://website.com
      source: https://github.com/random/repo
      license: Apache-2.0
      description: |
       ### Interesting Title
       Some application content, and description
    - title: Valid App 2
      version: 1.0.1
      maintainers:
      - name: AppTwo Maintainer
        email: apptwo@hotmail.com
      company: Upbound Inc.
      website: https://upbound.io
      source: https://github.com/upbound/repo
      license: Apache-2.0
      description: |
       ### Why app 2 is the best
       Because it simply is...
    ``` 
  - Status code _400_: The query is invalid, 
    - Response Body: "Invalid request!"
  
### GetMetadata
- Path: /v1/metadata
- Method: GET
- Response: The list of metadata we have.
  
## Improvements
For future improvements:
 - Vague keyword searching: If the query string does not match any record, we can show results that close to the query.
 - "Not contain" searching: Search for records that not contain a specific string.  
 - Improve the search algorithm: Currently the search algorithm goes through each field of each record to see if there is a match. Probably we can use Trie to improve the searching. Also, we can cache the most frequent queries to give a quick response.