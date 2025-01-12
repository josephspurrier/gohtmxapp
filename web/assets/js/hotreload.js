const eventSource = new EventSource('/reload');

eventSource.onmessage = (event) => {
    if (event.data === 'reload') {
        location.reload();
    }
};