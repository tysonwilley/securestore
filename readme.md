# SecureStore

SecureStore is a "proof of concept" storage system that will adhere to HIPAA & ePHI security specs and requirements. 

The API is written in Go and interfaces with a MariaDB instance that stores encrypted blobs of schemaless field:value pairs submitted through a RESTFUL endpoint. 
