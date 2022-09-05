import logging
import base64
import json

def lambda_handler(event, context):

    logger = logging.getLogger()
    logger.setLevel(logging.INFO)

    for record in event['Records']:
        #Kinesis data is base64 encoded so decode here
        payload=base64.b64decode(record["kinesis"]["data"]).decode()

        if payload == "InvalidData":
            print(f'Error: {payload}')
            raise Exception('fail')

        print("Decoded payload: " + payload)

    return {
        'statusCode': 200,
        'body': json.dumps('Hello from Lambda!')
    }
