const axios = require('axios')

if (process.argv.length != 3) {
    console.log("First argument should be an ICAO station.");
    process.exit();
}

axios
  .get(`https://api.paradox.ovh/metar?station=${process.argv[2]}`)
  .then(res => {
    const full_description =  res.data.details.descriptions.full_description;
    console.log(full_description);
  })
  .catch(error => {
    console.error("Err: "+error.response.data.msg);
  })

