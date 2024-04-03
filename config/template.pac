const hostConfig = {{.HostConfig}};
function hostMatch(host, rule) {
    if (rule.startsWith('*.')) {
        return host.endsWith(rule.slice(2));
    }
    return host === rule;
}
function FindProxyForURL(url, host) {
    const { hosts, rules } = hostConfig;
    let proxy = [];
    const hostRule = rules.find((item) => hostMatch(host, item[0]));
    if (hostRule) {
        const rules2 = [];
        hostRule[1].forEach((hostId) => {
            const rule = hosts.find((host2) => host2.id === hostId);
            if (rule) {
                rules2.push(rule);
            }
        });
        proxy = rules2.map((item) => `${item.type || 'PROXY'} ${item.host}:${item.port}`);
    }
    proxy.push('DIRECT');
    return proxy.join(';');
}
