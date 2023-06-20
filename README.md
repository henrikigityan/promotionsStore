# promotionsStore
PRomotion Store is an app writteng in golang for dealing csv uploads, saving data from it, and providing opportunity to read that data by id.

## Running the app
Install go and docker on server. 
Pull the main branch from git hub repo. 
Run mogodb latest image. 
```bash
docker pull mongo:latest
docker run -d -p 27017:27017 –name=mongo-test mongo:latest
```
Then run the program 
```bash
go run "test/promotionsStore"
```

## Usage
Programs has 2 apis 
Post /promotions/csv csv name in body "file"
Get /promotions/:id id is path param


## Additional info about later development opportunities
This is the MVP version of the app. It has two REST endpoints one for uploading csv and the other one for getting promotions by id. Program is written in go and uses mongoDb. 

Performance with large csv files: When the program is running on my local for csv with 200k rows endpoint all reading and importing into 
db is being done for 3-4 seconds, and for csv with 10m records is being done for 3-4 minutes. On larger servers which would have more ram 
and better cpu it will obviously be able to use much more thread efficiently and would be much faster. For better performance there can be 
used db sharding mechanism, so that when the program will get csv it will read it concurrently and using hashing function (most simple one 
would be to use the first character of id) decide in which db shard to write that part of info. Which will make writing faster, and will 
help to also make GET operations much faster and have proper caching. 

The next proper step that comes to my mind would be using 2 different db collections, and keeping one of them empty, 
so that when a new csv would be uploaded it would write info in that one, and set the application to use it, and only after that 
delete information from the other one. These tecniche would help not to have app latencies.

Unfortunately I didn’t have enough time to run stress tests and don’t have info how the app would perform with high load. 
Main steps for app performance during high load  would be, at first, separate GET endpoint to another microservice. 
Then use the sharding mechanism I told about above, it would help to set a load balancer which would choose to which shard  
send requests and with that mechanism out caches would be concrete, of course saving big caches would not be possible and best 
cache mechanism for our situation would be least recently used (LRU) approach. Meanwhile it will be meaningful to use rateLimit-ing 
in our load balancer to prevent our app from DOS attacks.

For deploying, scaling the way that comes to my mind would be to run db and app in docker and deploy it in k8s. We would be able 
to set CI/CD (for example with Jenkins). Which for CI will at first start unit and integration tests as soon as we push codes to github. 
After that it will build the application. Later CD will be triggered and will deploy the built app to k8s using predefined deployment 
and config files. 
For monitoring we could also deploy a Middleware agent which will help us to visualize app load and performance. 
We could also set it to notify us about some alerts.
 
