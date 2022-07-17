 - This project is contain to mak CRUD with RESTFul API using :
    ====================================================================
    1. Go Language for programming language
    2. MySql for Database Driver
    3. Gin for Go Framework
    4. Goose for Database Migration
   

 - Listen at port 3000 :
   ======================================================================
   - Run basic with command: "go run main.go"
      at config.go use this :
     - db, err := sql.Open("mysql", "admin:admin@/cake_store") 
      
   
   - Run with docker command: "docker-compose up"
     at config.go use this :
     - db, err := sql.Open("mysql", "admin:admin@tcp(godockerDB)/cake_store") "
      
   
 - Database Migration will create a table with this structure:
   ==========================================================================

   ========================================
 - 
   CREATE TABLE IF NOT EXISTS cake
   (
       id          int(100) auto_increment
           primary key,
       title       varchar(100) not null,
       description varchar(200) not null,
       rating      float(5)     not null,
       image       varchar(500) not null,
       created_at  datetime     not null,
       updated_at  datetime     not null
   );
   ========================================

 - Using this RESTFul API with Postman :
   ======================================================================
  1. List Of Cakes

    ===============================================================
     curl --location --request GET 'http://localhost:3000/cakes' \
     --header 'Access-Control-Allow-Origin: *' \
     --data-raw ''
    ===============================================================
     Description =>
     - Return a list of the cakes in JSON format, the cakes will be sorted by rating and alphabetically
     - Return Response with http code 200

  2. Detail Of Cakes

     ===============================================================
     curl --location --request GET 'http://localhost:3000/cakes/:id'
     ===============================================================
      Description =>
      - At Params, fill the Path Variable with "key = id", "value = {id value}"
        example :
        'http://localhost:3000/cakes/1'
      - Return the details of a cake in JSON Format
      - Return Response with http code 200

  4. Add New Cake

     ==============================================================
     curl --location --request POST 'http://localhost:3000/cakes' \
     --header 'Access-Control-Allow-Origin: *' \
     --header 'Content-Type: application/json' \
     --data-raw '{
         "title":{type string},
         "description":{type string},
         "rating":{type float},
         "image":{type string}
     }'
     ===============================================================
     Description =>
     - At Request Body, fill the raw in JSON
     example:
     {
        "title": "Brownies",         //cannot empty
        "description": "Black Cake", //cannot empty
        "rating": 6.1,               //cannot empty
        "image": "brownies.jpg"      //cannot empty
     }
     - Add a cake to the cakes list, the data will be sent as a JSON in the request body
     - Field "created_at" will set with time.Now()
     - Return Response with http code 200

  5. Update Cake

     =================================================================
     curl --location --request PATCH 'http://localhost:3000/cakes/:id' \
     --header 'Content-Type: application/json' \
     --data-raw '{
                 "title":{type string},
                 "description":{type string},
                 "rating":{type float},
                 "image":{type string}
     }'
     ==================================================================

     Description =>
     - At Params, fill the Path Variable with "key = id", "value = {id value}"
        example :
        'http://localhost:3000/cakes/1'
     - At Request Body, fill the raw in JSON
       example:
       {
          "title": "Brownies",         //cannot empty
          "description": "Black Cake", //cannot empty
          "rating": 6.1,               //cannot empty
          "image": "brownies.jpg"      //cannot empty
       }
     - Update a cake to the cakes list, the data will be sent as a JSON in the request body
     = Field "update_at" will set with time.Now()
     - Return Response with http code 200

  6. Delete Cake

     ==================================================================
     curl --location --request DELETE 'http://localhost:3000/cakes/:id'
     ==================================================================

     Description =>
     - At Params, fill the Path Variable with "key = id", "value = {id value}"
        example :
        'http://localhost:3000/cakes/1'
     - Delete a cake from database
     - Return Response with http code 200