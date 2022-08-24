exports.handler = function(event, context) {
    console.log(`event = ${JSON.stringify(event)}`);
    console.log(`context = ${JSON.stringify(context)}`);

    event.Records.forEach((record) => {
        var payload = Buffer.from(record.kinesis.data, 'base64').toString();
        console.log('record.kinesis.data: ', record.kinesis.data);
        console.log('payload: ', payload);
    });

    return `Hello World!`;
};
