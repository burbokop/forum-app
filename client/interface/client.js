const http = require('../common/http');

const Client = (baseUrl) => {
    const client = http.Client(baseUrl);
    return {
        listForums: () => client.get('/forums'),
        addUser: (user, interests) => client.post('/add_user', { name: user, interests: Array.isArray(interests) ? interests : [ interests ] })
    }

};

module.exports = { Client };
