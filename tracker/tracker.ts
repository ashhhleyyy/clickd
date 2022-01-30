(() => {
    const loc = window.location;
    const path = loc.pathname + loc.search + loc.hash;
    fetch('%CLICKD_HOST%/track', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            type: 'visit',
            path,
            referer: document.referrer,
            width: window.innerWidth,
            height: window.innerHeight,
        }),
    });
})();
