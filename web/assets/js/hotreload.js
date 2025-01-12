const eventSource = new EventSource('/sse');

eventSource.onmessage = (event) => {
    if (event.data === 'reload') {
        location.reload();
    }
};