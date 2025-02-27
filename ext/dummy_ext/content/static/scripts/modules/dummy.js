document.addEventListener('DOMContentLoaded', function() {
    fetch('/api/dummy')
        .then(response => response.json())
        .then(data => {
            const node = d3.select('#dummy-js-replaced');
            node.text(data.data.message)
        });
});
