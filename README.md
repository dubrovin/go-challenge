Task
----

Write an HTTP service that exposes an endpoint "/numbers". This endpoint receives a list of URLs 
though GET query parameters. The parameter in use is called "u". It can appear 
more than once.

	http://yourserver:8080/numbers?u=http://example.com/primes&u=http://foobar.com/fibo

When the /numbers is called, your service shall retrieve each of these URLs if 
they turn out to be syntactically valid URLs. Each URL will return a JSON data 
structure that looks like this:

	{ "numbers": [ 1, 2, 3, 5, 8, 13 ] }

The JSON data structure will contain an object with a key named "numbers", and 
a value that is a list of integers. After retrieving each of these URLs, the 
service shall merge the integers coming from all URLs, sort them in ascending 
order, and make sure that each integer only appears once in the result. The 
endpoint shall then return a JSON data structure like in the example above with 
the result as the list of integers.

The endpoint needs to return the result as quickly as possible, but always 
within 500 milliseconds. It needs to be able to deal with error conditions when 
retrieving the URLs. If a URL takes too long to respond, it must be ignored. It 
is valid to return an empty list as result only if all URLs returned errors or 
took too long to respond.

Example
-------

The service receives an HTTP request:

	>>> GET /numbers?u=http://example.com/primes&u=http://foobar.com/fibo HTTP/1.0

It then retrieves the URLs specified as parameters.

The first URL returns this response:

	>>> GET /primes HTTP/1.0
	>>> Host: example.com
	>>> 
	<<< HTTP/1.0 200 OK
	<<< Content-Type: application/json
	<<< Content-Length: 34
	<<< 
	<<< { "number": [ 2, 3, 5, 7, 11, 13 ] }

The second URL returns this response:

	>>> GET /fibo HTTP/1.0
	>>> Host: foobar.com
	>>> 
	<<< HTTP/1.0 200 OK
	<<< Content-Type: application/json
	<<< Content-Length: 40
	<<< 
	<<< { "number": [ 1, 1, 2, 3, 5, 8, 13, 21 ] }

The service then calculates the result and returns it.

	<<< HTTP/1.0 200 OK
	<<< Content-Type: application/json
	<<< Content-Length: 44
	<<< 
	<<< { "number": [ 1, 2, 3, 5, 7, 8, 11, 13, 21 ] }


Explanation and example
----

Implemented service which exposes /numbers endpoint.   
This endpoint makes get queries for given url in "u" param in GET query.    
All get queries for given urls proceed in parallel,   
for waiting for all this queries used sync.WaitGroup structure.

Example:
 ```
 curl http://127.0.0.1:8080/numbers?u=http://127.0.0.1:8090/primes&u=http://127.0.0.1:8090/fibo  
    
 recieved: {"Numbers":[1,2,3,5,7,8,11,13,21]}  
 ```
Also example of using this service placed in server/server_test.go in TestServerRun    
coverage: 84.8% of statements  