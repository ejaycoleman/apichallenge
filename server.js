const express = require('express')
const request = require('request');
const app = express()
const port = 3000







app.get('/api/dwarves', (req, res) => {
	
	request('https://thedwarves.pusherplatform.io/api/dwarves', { json: true }, (err, res, body) => {
		if (err) { return console.log(err); }
		console.log(body.dwarves.map(lol => lol.name) );
	});
	

	res.send('Hello World!')
})

app.get('/api/dwarves/:name', (req, res) => {
	
	request('https://thedwarves.pusherplatform.io/api/dwarves', { json: true }, (err, res, body) => {
		if (err) { return console.log(err); }
		console.log(body.dwarves.filter(lol => lol.name === req.params.name) );
	});
	

	res.send('Hello World!')
})

app.listen(port, () => console.log(`Example app listening on port ${port}!`))