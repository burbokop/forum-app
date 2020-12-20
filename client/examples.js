// This file contains examples of scenarios implementation using

const channels = require('./interface/client');

const client = channels.Client('http://localhost:8080');

// Scenario 1: List Forums
client.listForums()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Forums:');
        list.forEach((c) => console.log(c));
    })
    .catch((e) => {
        console.log(`Problem listing forums: ${e.message}`);
    });

// Scenario 2: Add user.
client.addUser("robert", [ "music", "programming" ])
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Adding user responce:');
        for(const r of resp) {
            console.log(r);
        }
    })
    .catch((e) => {
        console.log(`Problem adding user: ${e.message}`);
    });
