var nodes = [];
var addresses = [];
let protocol = 'http://'

Vue.component('node', {
    props: ["data"],
    template: `
        <div class="card mb-4 shadow-sm">
            <h5>{{ node.Name }}</h5>
        </div>`,
    watch: {
        data: {
            handler: function(newValue) {
                console.log(newValue)
            },
            deep: true
        }
    }
});

//vue.js setup
var app;
function vueSetup() {
    console.log(nodes)
    app = new Vue({
        el: '#nodes',
        data: {
            test: "bla",
            nodes: nodes
        }
    });
}

// Helpers
function updatePeer(name) {
    $.getJSON(protocol + addresses[name] + "/status", function (data) {
        nodes[data.Name] = data;
    });
}

function updateAll() {
    for (key in addresses) {
        updatePeer(key)
    }
}

// Initialize all nodes by fetching "our" own status and populating all data
currentHost = $(location).attr('host');
$.getJSON(protocol + currentHost + "/status", function (data) {
    addresses[data.Name] = currentHost;
    nodes[data.Name] = data;
    data.Peers.forEach(function (e) {
        addresses[e.Name] = e.Address;
        updatePeer(e.Name)
    });
    vueSetup();
});
setInterval(updateAll, 2000);
