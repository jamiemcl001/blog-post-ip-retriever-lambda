# IP Retriever Lambda

This is just a basic lambda function which can be uploaded to AWS Lambda and run. It will attempt to get the public facing IP of the lambda and return it as a response.

## Building Instructions

To build and package the app for use in AWS Lambda just run the shell script in the `scripts` folder.

```
bash scripts/build.sh
```

This will build the app and add it to the `app.zip` folder. It builds it in the correct format (for Linux based architecture) which can be uploaded via the Lambda console.