window.setTimeout(function() {
    $(".moment").fadeTo(500, 0).slideUp(500, function(){
        $(this).remove(); 
    });
}, 5000);

var dataURL = "http://localhost:8080/api/index";

fetch(dataURL)
    .then(function (response) {
        return response.json();
    })
    .then(function (data) {
        console.log(data);
        // get water data 
        var waterData = data.status.water;
        console.log(waterData);

        // get wind data
        var windData = data.status.wind;
        console.log(windData);

        if (waterData < 5){
            waterMessage = `Ketinggian air = ${waterData} meter, status aman`;
        }else if ( waterData <= 8 ) {
            waterMessage = `Ketinggian air = ${waterData} meter, status siaga`;
        }else if ( waterData > 8 ) {
            waterMessage = `Ketinggian air = ${waterData} meter, status bahaya`;
        }

        if (windData < 5){
            windMessage = `Kecepatan angin = ${windData}m/s, status aman`;
        }else if ( windData <= 8 ) {
            windMessage = `Kecepatan angin = ${windData}m/s, status siaga`;
        }else if ( windData > 8 ) {
            windMessage = `Kecepatan angin = ${windData}m/s, status bahaya`;
        }

        document.querySelector(".waterData").innerHTML = waterMessage;
        document.querySelector(".windData").innerHTML = windMessage;
    });


