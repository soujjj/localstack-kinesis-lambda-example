exports.handler = function(event, context) {
    console.log(`event = ${JSON.stringify(event)}`);
    console.log(`context = ${JSON.stringify(context)}`);

    return `Hello World!`;
};
