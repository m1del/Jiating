const cheerio = require('cheerio');
const request = require('request');
const fs = require('fs');

//write stream
let writeStream = fs.createWriteStream("imageslink.txt", "utf8");
request("https://www.ufl.edu/", (err, response, html) =>{

    if(!err && response.statusCode === 200) {
        console.log("Request succeeded");

        //Define cheerio e.g. the $ Object
        var $ = cheerio.load(html);

        $("img").each((index, image) => {
            var img = $(image).attr('src');
            var baseUrl = 'https://www.ufl.edu/';
            var link = baseUrl + img;

            console.log("Getting image...");
            console.log(link);

            writeStream.write(link);
            writeStream.write('\n');
            writeStream.end();
        });

    } else {
        console.log("Request failed");
    }
});