function drawGraphInitial() {
    updateGraphData();
}

function getChartContext() {
    return document.getElementById("myChart").getContext('2d');
}


let Currencies = function (obj) {
    this.BTC = obj.BTC;
    this.ETH = obj.ETH;
    this.LTC = obj.LTC;
    this.XMR = obj.XMR;
    this.ETC = obj.ETC;
    this.DASH = obj.DASH;
    this.MAID = obj.MAID;
    this.REP = obj.REP;
    this.XEM = obj.XEM;
};

let XY = function(x, y) {
    this.x = x;
    this.y = y;
};
function commaSeparateNumber(val){
    while (/(\d+)(\d{3})/.test(val.toString())){
        val = val.toString().replace(/(\d+)(\d{3})/, '$1'+','+'$2');
    }
    return val;
}

function drawGraph(curr) {
    let currencies = new Currencies(curr);
    let ctx = getChartContext();
    console.log(JSON.stringify(curr));
    console.log([currencies.BTC.map(e => new Date(e.Moment))].toString());
    console.log([currencies.BTC.map(e => e.RUR)].toString());
    let config = {
        type: 'line',
        data: {
            labels: [currencies.BTC.map(e => new Date(e.Moment))],
            datasets: [{
                label: "BTC",
                data: currencies.BTC.map(e => new XY(new Date(e.Moment), e.RUR)),
                borderColor: "#3e95cd",
                fill: false
            }, {
                label: "ETH",
                data: currencies.ETH.map(e => new XY(new Date(e.Moment), e.RUR)),
                borderColor: "#8e5ea2",
                fill: false
            }, {
                label: "LTC",
                data: currencies.LTC.map(e => new XY(new Date(e.Moment), e.RUR)),
                borderColor: "#3cba9f",
                fill: false
            }, {
                label: "XMR",
                data: currencies.XMR.map(e => new XY(new Date(e.Moment), e.RUR)),
                borderColor: "#e8c3b9",
                fill: false
            }, {
                label: "ETC",
                data: currencies.ETC.map(e => new XY(new Date(e.Moment), e.RUR)),
                borderColor: "#c45850",
                fill: false
            }, {
                label: "DASH",
                data: currencies.DASH.map(e => new XY(new Date(e.Moment), e.RUR)),
                borderColor: "#54cd45",
                fill: false
            }, {
                label: "MAID",
                data: currencies.MAID.map(e => new XY(new Date(e.Moment), e.RUR)),
                borderColor: "#c76acd",
                fill: false
            }, {
                label: "REP",
                data: currencies.REP.map(e => new XY(new Date(e.Moment), e.RUR)),
                borderColor: "#000000",
                fill: false
            }, {
                label: "XEM",
                data: currencies.XEM.map(e => new XY(new Date(e.Moment), e.RUR)),
                borderColor: "#b8cd19",
                fill: false
            }]
        },
        options: {
            scales: {
                xAxes: [{
                    scaleLabel: {
                        display: true,
                        labelString: "Time",
                        fontColor: "red"
                    },
                    type: 'time',
                    time: {
                        displayFormats: {
                            hour: 'hh:mm'
                        }
                    }
                }],
                yAxes: [{
                    ticks: {
                        userCallback: function (value) {
                            return commaSeparateNumber(value);
                        }
                    },
                    scaleLabel: {
                        display: true,
                        labelString: "Currency price",
                        fontColor: "green"
                    }
                }]
            }
        }
    };

    let myChart = new Chart(ctx, config);
}

function updateGraphData() {
    const INTERVAL = {
        MONTH: "MONTH",
        WEEK: "WEEK",
        DAY: "DAY"
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
                drawGraph(data);
            });
        }
    )
        .catch(function (err) {
            console.log('Fetch Error :-S', err);
        });
}