package main

func main(){

}

// necessary todo 
    // setup database in one node 
	// run this program as background  just like systemctl start xxx.service
	
	// in db set two tables for users
	// 1. user table         ----> cointains user's name and password and last login time
	// 2. user's data table   ----> cointains user's data and time  such as resources that he currently use or used

	// one table for metrics it will be time series database
	// 1. metrics table  
	// 2. request table  ----> cointains user's request and resource type and resource name and resource value and time that request was made 


	// listen to port 8080
	// get request from client and authenticate user by checking user table
	// if user is authenticated then proceed to next step
	// if not then return error message to client

	// then see the user request and process it based  on the request , get the rsource type and resource name and resource value 
	// then save the data in request table and assign it as failed request

	// call the function that will check the resource type and resource name and resource value and check if it is available or not
	// if not available then in request table add reason for failure and assign it as failed request

	// if available then update the user's data table and update the resource value and assign it as success request
	// call appropriate function to create resource and assign it as success requestin database and add time and date to it

	// once the request is processed then return the response to client


	// one other way we can send it to user is we can set a boolean variable if all thing worked and resources created successfully
	// then set the boolean variable to true and user will watch the boolean variable and if it is true then it will show new resource created
	// if it is false then it will show error message that resource is not created and reason for failure




// Structure implementation

/* 
				1. make 1 interface for all resources
                 |
				 |----------- create resource
				 |----------- delete resource
				 |----------- update resource

*/