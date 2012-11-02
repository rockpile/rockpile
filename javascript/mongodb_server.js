#!/usr/bin/env node


var mongoose = require('mongoose');
db = mongoose.createConnection('localhost', 'test');

db.on('error', console.error.bind(console, 'connection error:'));
db.once('open', function() {
	console.log("open");
	var kittySchema = new mongoose.Schema({
		name: String
	});

	var Kitten = db.model('kitten', kittySchema);

	var silence = new Kitten({ name: 'Silence' });
	console.log(silence.name); // 'Silence'

	kittySchema.methods.speak = function () {
		var greeting = this.name
	? "Meow name is " + this.name
	: "I don't have a name"
	console.log(greeting);
	};

	var Kitten = db.model('Kitten', kittySchema);

	var fluffy = new Kitten({ name: 'fluffy' });
	fluffy.speak(); // "Meow name is fluffy"

	fluffy.save(function (err) {
		if (err) // TODO handle the error
		console.log('meow\n')
	});

	Kitten.find(function (err, kittens) {
		//if (err) // TODO handle err
		console.log(kittens)
	})
});


//Kitten.find({ name: /fluff/i }, callback)
