# AWS Go Logger

### A simple logger for your lambda function cloudwatch logs in AWS.

Step 1. Build it using `go build main.go alg.go`  
Step 2. Set up your AWS credentials locally  
Step 3. Inspect your lambdas by running `agl <lambda-name>` where the <lambda-name> can also be partial, not necessarily the entire name.  
Step 4. See the cloudwatch logs in your terminal.
