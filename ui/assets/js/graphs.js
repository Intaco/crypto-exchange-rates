function drawGraphInitial() {
    var ctx = getChartContext();
    drawGraph(ctx);
    updateGraphData(0, 0);
}
function newDate(days) {
    return moment().add(days, 'd');
}

function getChartContext() {
    return document.getElementById("myChart").getContext('2d');
}

function convertToDataSets(response, interval){
    console.log(JSON.stringify(response));
}
function partitionData(data) {
    //TODO split by type into arrays
}


function drawGraph(ctx) {
    var config = {
        type: 'line',
        data: {
            labels: [newDate(-4), newDate(-3), newDate(2), newDate(3), newDate(4), newDate(5), newDate(6)],
            datasets: [{
                label: "My First dataset",
                fill: true, // fill underlying
                data: [1, 3, 4, 2, 1, 4, 2],
            }]
        },
        options: {
            scales: {
                xAxes: [{
                    type: 'time',
                    time: {
                        displayFormats: {
                            'millisecond': 'MMM DD',
                            'second': 'MMM DD',
                            'minute': 'MMM DD',
                            'hour': 'MMM DD',
                            'day': 'MMM DD',
                            'week': 'MMM DD',
                            'month': 'MMM DD',
                            'quarter': 'MMM DD',
                            'year': 'MMM DD',
                        }
                    }
                }],
            },
        }
    };

    var myChart = new Chart(ctx, config);
}
function updateGraphData(from, to) {
    const INTERVAL = {
        OTHER : 31,
        WEEK: 7,
        DAY : 24
    };

    fetch('./api/data', {
        method: "POST",
        body: `{
			    "from": "2012-11-01T22:08:41+00:00",
                "to": "2017-11-01T22:08:41+00:00",
                "duration": "DAY"

	    }`
    }).then(
            function (response) {
                if (response.status !== 200) {
                    console.log('Looks like there was a problem. Status Code: ' +
                        response.status);
                    return;
                }

                // Examine the text in the response
                response.json().then(function (data) {
                    convertToDataSets(data, INTERVAL.OTHER);
                });
            }
        )
        .catch(function (err) {
            console.log('Fetch Error :-S', err);
        });
}